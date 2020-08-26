// access to the solid state of the smart contract
package stateapi

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/iotaledger/goshimmer/dapps/valuetransfers/packages/address"
	"github.com/iotaledger/wasp/packages/kv"
	"github.com/iotaledger/wasp/packages/state"
	"github.com/iotaledger/wasp/plugins/webapi/misc"
	"github.com/labstack/echo"
)

type ValueType string

const (
	ValueTypeBytes       = ValueType("bytes")
	ValueTypeInt64       = ValueType("int64")
	ValueTypeArray       = ValueType("array")
	ValueTypeDict        = ValueType("dict")
	ValueTypeDictElement = ValueType("dict-elem")
)

type KeyQuery struct {
	Key    []byte
	Type   ValueType
	Params json.RawMessage // one of DictQueryParams, ArrayQueryParams, ...
}

type DictQueryParams struct {
	Limit uint32
}

type DictElementQueryParams struct {
	Key []byte
}

type ArrayQueryParams struct {
	From uint16
	To   uint16
}

type QueryRequest struct {
	Address string
	Query   []*KeyQuery
}

type QueryResult struct {
	Key   []byte
	Type  ValueType
	Value json.RawMessage // one of DictResult, ArrayResult, ...
}

type KeyValuePair struct {
	Key   []byte
	Value []byte
}

type Int64Result struct {
	Value int64
}

type DictResult struct {
	Len     uint32
	Entries []KeyValuePair
}

type DictElementResult struct {
	Value []byte
}

type ArrayResult struct {
	Len    uint16
	Values [][]byte
}

type QueryResponse struct {
	Results []*QueryResult
	Error   string
}

func NewQueryRequest(address *address.Address) *QueryRequest {
	return &QueryRequest{Address: address.String()}
}

func (q *QueryRequest) AddBytes(key kv.Key) {
	q.Query = append(q.Query, &KeyQuery{
		Key:    []byte(key),
		Type:   ValueTypeBytes,
		Params: nil,
	})
}

func (q *QueryRequest) AddInt64(key kv.Key) {
	q.Query = append(q.Query, &KeyQuery{
		Key:    []byte(key),
		Type:   ValueTypeInt64,
		Params: nil,
	})
}

func (q *QueryRequest) AddArray(key kv.Key, from uint16, to uint16) {
	p := &ArrayQueryParams{From: from, To: to}
	params, _ := json.Marshal(p)
	q.Query = append(q.Query, &KeyQuery{
		Key:    []byte(key),
		Type:   ValueTypeArray,
		Params: json.RawMessage(params),
	})
}

func (q *QueryRequest) AddDictionary(key kv.Key, limit uint32) {
	p := &DictQueryParams{Limit: limit}
	params, _ := json.Marshal(p)
	q.Query = append(q.Query, &KeyQuery{
		Key:    []byte(key),
		Type:   ValueTypeDict,
		Params: json.RawMessage(params),
	})
}

func (q *QueryRequest) AddDictionaryElement(dictKey kv.Key, elemKey []byte) {
	p := &DictElementQueryParams{Key: elemKey}
	params, _ := json.Marshal(p)
	q.Query = append(q.Query, &KeyQuery{
		Key:    []byte(dictKey),
		Type:   ValueTypeDictElement,
		Params: json.RawMessage(params),
	})
}

func (r *QueryResult) MustBytes() ([]byte, bool, error) {
	var b []byte
	err := json.Unmarshal(r.Value, &b)
	if err != nil {
		return nil, false, err
	}
	return b, b != nil, nil
}

func (r *QueryResult) MustInt64() (int64, bool, error) {
	var ir *Int64Result
	err := json.Unmarshal(r.Value, &ir)
	if err != nil {
		return 0, false, err
	}
	if ir == nil {
		return 0, false, nil
	}
	return ir.Value, true, nil
}

func (r *QueryResult) MustArrayResult() *ArrayResult {
	var ar ArrayResult
	err := json.Unmarshal(r.Value, &ar)
	if err != nil {
		panic(err)
	}
	return &ar
}

func (r *QueryResult) MustDictionaryResult() *DictResult {
	var dr DictResult
	err := json.Unmarshal(r.Value, &dr)
	if err != nil {
		panic(err)
	}
	return &dr
}

func (r *QueryResult) MustDictionaryElementResult() *DictElementResult {
	var dr DictElementResult
	err := json.Unmarshal(r.Value, &dr)
	if err != nil {
		panic(err)
	}
	return &dr
}

func HandlerQueryState(c echo.Context) error {
	var req QueryRequest

	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, &QueryResponse{Error: err.Error()})
	}
	addr, err := address.FromBase58(req.Address)
	if err != nil {
		return c.JSON(http.StatusBadRequest, &QueryResponse{Error: err.Error()})
	}
	// TODO serialize access to solid state
	state, _, exist, err := state.LoadSolidState(&addr)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, &QueryResponse{Error: err.Error()})
	}
	if !exist {
		return c.JSON(http.StatusNotFound, &QueryResponse{
			Error: fmt.Sprintf("State not found with address %s", addr),
		})
	}
	ret := &QueryResponse{
		Results: make([]*QueryResult, 0),
	}
	vars := state.Variables()
	for _, q := range req.Query {
		value, err := processQuery(q, vars)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, &QueryResponse{Error: err.Error()})
		}
		b, err := json.Marshal(value)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, &QueryResponse{Error: err.Error()})
		}
		ret.Results = append(ret.Results, &QueryResult{
			Key:   q.Key,
			Type:  q.Type,
			Value: json.RawMessage(b),
		})
	}

	return misc.OkJson(c, ret)
}

func processQuery(q *KeyQuery, vars kv.BufferedKVStore) (interface{}, error) {
	key := kv.Key(q.Key)
	switch q.Type {
	case ValueTypeBytes:
		value, err := vars.Get(key)
		if err != nil {
			return nil, err
		}
		return value, nil

	case ValueTypeInt64:
		value, err := vars.Get(key)
		if err != nil || value == nil {
			return nil, err
		}
		n, err := kv.DecodeInt64(value)
		if err != nil {
			return nil, err
		}
		return Int64Result{Value: n}, nil

	case ValueTypeArray:
		var params ArrayQueryParams
		err := json.Unmarshal(q.Params, &params)
		if err != nil {
			return nil, err
		}

		arr, err := vars.Codec().GetArray(key)
		if err != nil {
			return nil, err
		}

		size := arr.Len()
		values := make([][]byte, 0)
		for i := params.From; i < size && i < params.To; i++ {
			v, err := arr.GetAt(i)
			if err != nil {
				return nil, err
			}
			values = append(values, v)
		}
		return ArrayResult{Len: size, Values: values}, nil

	case ValueTypeDict:
		var params DictQueryParams
		err := json.Unmarshal(q.Params, &params)
		if err != nil {
			return nil, err
		}

		dict, err := vars.Codec().GetDictionary(key)
		if err != nil {
			return nil, err
		}

		entries := make([]KeyValuePair, 0)
		err = dict.Iterate(func(elemKey []byte, value []byte) bool {
			entries = append(entries, KeyValuePair{Key: elemKey, Value: value})
			return len(entries) < int(params.Limit)
		})
		if err != nil {
			return nil, err
		}
		return DictResult{Len: dict.Len(), Entries: entries}, nil

	case ValueTypeDictElement:
		var params DictElementQueryParams
		err := json.Unmarshal(q.Params, &params)
		if err != nil {
			return nil, err
		}

		dict, err := vars.Codec().GetDictionary(key)
		if err != nil {
			return nil, err
		}

		v, err := dict.GetAt(params.Key)
		if err != nil {
			return nil, err
		}
		return DictElementResult{Value: v}, nil
	}

	return nil, fmt.Errorf("No handler for type %s", q.Type)
}
