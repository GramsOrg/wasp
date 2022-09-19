// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

package test

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/iotaledger/wasp/contracts/wasm/testwasmlib/go/testwasmlib"
	"github.com/iotaledger/wasp/packages/wasmvm/wasmhost"
	"github.com/iotaledger/wasp/packages/wasmvm/wasmlib/go/wasmlib/wasmtypes"
	"github.com/iotaledger/wasp/packages/wasmvm/wasmsolo"
)

func TestStringMapOfStringArrayClear(t *testing.T) {
	ctx := setupTest(t)

	as := testwasmlib.ScFuncs.StringMapOfStringArrayAppend(ctx)
	as.Params.Name().SetValue("bands")
	as.Params.Value().SetValue("Simple Minds")
	as.Func.Post()
	require.NoError(t, ctx.Err)

	as = testwasmlib.ScFuncs.StringMapOfStringArrayAppend(ctx)
	as.Params.Name().SetValue("bands")
	as.Params.Value().SetValue("Dire Straits")
	as.Func.Post()
	require.NoError(t, ctx.Err)

	as = testwasmlib.ScFuncs.StringMapOfStringArrayAppend(ctx)
	as.Params.Name().SetValue("bands")
	as.Params.Value().SetValue("ELO")
	as.Func.Post()
	require.NoError(t, ctx.Err)

	al := testwasmlib.ScFuncs.StringMapOfStringArrayLength(ctx)
	al.Params.Name().SetValue("bands")
	al.Func.Call()
	require.NoError(t, ctx.Err)
	length := al.Results.Length()
	require.True(t, length.Exists())
	require.EqualValues(t, 3, length.Value())

	av := testwasmlib.ScFuncs.StringMapOfStringArrayValue(ctx)
	av.Params.Name().SetValue("bands")
	av.Params.Index().SetValue(0)
	av.Func.Call()
	require.NoError(t, ctx.Err)
	value := av.Results.Value()
	require.True(t, value.Exists())
	require.EqualValues(t, "Simple Minds", value.Value())

	av = testwasmlib.ScFuncs.StringMapOfStringArrayValue(ctx)
	av.Params.Name().SetValue("bands")
	av.Params.Index().SetValue(1)
	av.Func.Call()
	require.NoError(t, ctx.Err)
	value = av.Results.Value()
	require.True(t, value.Exists())
	require.EqualValues(t, "Dire Straits", value.Value())

	av = testwasmlib.ScFuncs.StringMapOfStringArrayValue(ctx)
	av.Params.Name().SetValue("bands")
	av.Params.Index().SetValue(2)
	av.Func.Call()
	require.NoError(t, ctx.Err)
	value = av.Results.Value()
	require.True(t, value.Exists())
	require.EqualValues(t, "ELO", value.Value())

	ac := testwasmlib.ScFuncs.StringMapOfStringArrayClear(ctx)
	ac.Params.Name().SetValue("bands")
	ac.Func.Post()
	require.NoError(t, ctx.Err)

	al = testwasmlib.ScFuncs.StringMapOfStringArrayLength(ctx)
	al.Params.Name().SetValue("bands")
	al.Func.Call()
	require.NoError(t, ctx.Err)
	length = al.Results.Length()
	require.True(t, length.Exists())
	require.EqualValues(t, 0, length.Value())

	av = testwasmlib.ScFuncs.StringMapOfStringArrayValue(ctx)
	av.Params.Name().SetValue("bands")
	av.Params.Index().SetValue(0)
	av.Func.Call()
	require.Error(t, ctx.Err)
}

func genTestAddress(ctx *wasmsolo.SoloContext, num int) []wasmtypes.ScAddress {
	addrs := make([]wasmtypes.ScAddress, num)
	for i := 0; i < num; i++ {
		_, addr := ctx.Chain.Env.NewKeyPair()
		addrs[i] = wasmhost.WasmConvertor{}.ScAddress(addr)
	}

	return addrs
}

func TestAddressMapOfAddressArrayClear(t *testing.T) {
	ctx := setupTest(t)
	mapNames, mapVals := genTestAddress(ctx, 2), genTestAddress(ctx, 3)

	as := testwasmlib.ScFuncs.AddressMapOfAddressArrayAppend(ctx)
	as.Params.NameAddr().SetValue(mapNames[0])
	as.Params.ValueAddr().SetValue(mapVals[0])
	as.Func.Post()
	require.NoError(t, ctx.Err)

	as = testwasmlib.ScFuncs.AddressMapOfAddressArrayAppend(ctx)
	as.Params.NameAddr().SetValue(mapNames[0])
	as.Params.ValueAddr().SetValue(mapVals[1])
	as.Func.Post()
	require.NoError(t, ctx.Err)

	as = testwasmlib.ScFuncs.AddressMapOfAddressArrayAppend(ctx)
	as.Params.NameAddr().SetValue(mapNames[1])
	as.Params.ValueAddr().SetValue(mapVals[2])
	as.Func.Post()
	require.NoError(t, ctx.Err)

	al := testwasmlib.ScFuncs.AddressMapOfAddressArrayLength(ctx)
	al.Params.NameAddr().SetValue(mapNames[0])
	al.Func.Call()
	require.NoError(t, ctx.Err)
	length := al.Results.Length()
	require.True(t, length.Exists())
	require.EqualValues(t, 2, length.Value())

	al = testwasmlib.ScFuncs.AddressMapOfAddressArrayLength(ctx)
	al.Params.NameAddr().SetValue(mapNames[1])
	al.Func.Call()
	require.NoError(t, ctx.Err)
	length = al.Results.Length()
	require.True(t, length.Exists())
	require.EqualValues(t, 1, length.Value())

	av := testwasmlib.ScFuncs.AddressMapOfAddressArrayValue(ctx)
	av.Params.NameAddr().SetValue(mapNames[0])
	av.Params.Index().SetValue(0)
	av.Func.Call()
	require.NoError(t, ctx.Err)
	value := av.Results.ValueAddr()
	require.True(t, value.Exists())
	require.EqualValues(t, mapVals[0], value.Value())

	av = testwasmlib.ScFuncs.AddressMapOfAddressArrayValue(ctx)
	av.Params.NameAddr().SetValue(mapNames[0])
	av.Params.Index().SetValue(1)
	av.Func.Call()
	require.NoError(t, ctx.Err)
	value = av.Results.ValueAddr()
	require.True(t, value.Exists())
	require.EqualValues(t, mapVals[1], value.Value())

	av = testwasmlib.ScFuncs.AddressMapOfAddressArrayValue(ctx)
	av.Params.NameAddr().SetValue(mapNames[1])
	av.Params.Index().SetValue(0)
	av.Func.Call()
	require.NoError(t, ctx.Err)
	value = av.Results.ValueAddr()
	require.True(t, value.Exists())
	require.EqualValues(t, mapVals[2], value.Value())

	ac := testwasmlib.ScFuncs.AddressMapOfAddressArrayClear(ctx)
	ac.Params.NameAddr().SetValue(mapNames[0])
	ac.Func.Post()
	require.NoError(t, ctx.Err)

	al = testwasmlib.ScFuncs.AddressMapOfAddressArrayLength(ctx)
	al.Params.NameAddr().SetValue(mapNames[0])
	al.Func.Call()
	require.NoError(t, ctx.Err)
	length = al.Results.Length()
	require.True(t, length.Exists())
	require.EqualValues(t, 0, length.Value())

	av = testwasmlib.ScFuncs.AddressMapOfAddressArrayValue(ctx)
	av.Params.NameAddr().SetValue(mapNames[0])
	av.Params.Index().SetValue(0)
	av.Func.Call()
	require.Error(t, ctx.Err)
}

func TestStringMapOfStringArraySet(t *testing.T) {
	ctx := setupTest(t)

	ap := testwasmlib.ScFuncs.StringMapOfStringArrayAppend(ctx)
	ap.Params.Name().SetValue("bands")
	ap.Params.Value().SetValue("Simple Minds")
	ap.Func.Post()
	require.NoError(t, ctx.Err)

	ap = testwasmlib.ScFuncs.StringMapOfStringArrayAppend(ctx)
	ap.Params.Name().SetValue("bands")
	ap.Params.Value().SetValue("Dire Straits")
	ap.Func.Post()
	require.NoError(t, ctx.Err)

	al := testwasmlib.ScFuncs.StringMapOfStringArrayLength(ctx)
	al.Params.Name().SetValue("bands")
	al.Func.Call()
	require.NoError(t, ctx.Err)
	length := al.Results.Length()
	require.True(t, length.Exists())
	require.EqualValues(t, 2, length.Value())

	av := testwasmlib.ScFuncs.StringMapOfStringArrayValue(ctx)
	av.Params.Name().SetValue("bands")
	av.Params.Index().SetValue(0)
	av.Func.Call()
	require.NoError(t, ctx.Err)
	value := av.Results.Value()
	require.True(t, value.Exists())
	require.EqualValues(t, "Simple Minds", value.Value())

	av = testwasmlib.ScFuncs.StringMapOfStringArrayValue(ctx)
	av.Params.Name().SetValue("bands")
	av.Params.Index().SetValue(1)
	av.Func.Call()
	require.NoError(t, ctx.Err)
	value = av.Results.Value()
	require.True(t, value.Exists())
	require.EqualValues(t, "Dire Straits", value.Value())

	as := testwasmlib.ScFuncs.StringMapOfStringArraySet(ctx)
	as.Params.Name().SetValue("bands")
	as.Params.Index().SetValue(0)
	as.Params.Value().SetValue("Collage")
	as.Func.Post()
	require.NoError(t, ctx.Err)

	al = testwasmlib.ScFuncs.StringMapOfStringArrayLength(ctx)
	al.Params.Name().SetValue("bands")
	al.Func.Call()
	require.NoError(t, ctx.Err)
	length = al.Results.Length()
	require.True(t, length.Exists())
	require.EqualValues(t, 2, length.Value())

	av = testwasmlib.ScFuncs.StringMapOfStringArrayValue(ctx)
	av.Params.Name().SetValue("bands")
	av.Params.Index().SetValue(0)
	av.Func.Call()
	require.NoError(t, ctx.Err)
	value = av.Results.Value()
	require.True(t, value.Exists())
	require.EqualValues(t, "Collage", value.Value())

	av = testwasmlib.ScFuncs.StringMapOfStringArrayValue(ctx)
	av.Params.Name().SetValue("bands")
	av.Params.Index().SetValue(1)
	av.Func.Call()
	require.NoError(t, ctx.Err)
	value = av.Results.Value()
	require.True(t, value.Exists())
	require.EqualValues(t, "Dire Straits", value.Value())
}

func TestAddressMapOfAddressArraySet(t *testing.T) {
	ctx := setupTest(t)
	mapNames, mapVals := genTestAddress(ctx, 2), genTestAddress(ctx, 4)

	aap := testwasmlib.ScFuncs.AddressMapOfAddressArrayAppend(ctx)
	aap.Params.NameAddr().SetValue(mapNames[0])
	aap.Params.ValueAddr().SetValue(mapVals[0])
	aap.Func.Post()
	require.NoError(t, ctx.Err)

	aap = testwasmlib.ScFuncs.AddressMapOfAddressArrayAppend(ctx)
	aap.Params.NameAddr().SetValue(mapNames[0])
	aap.Params.ValueAddr().SetValue(mapVals[1])
	aap.Func.Post()
	require.NoError(t, ctx.Err)

	aap = testwasmlib.ScFuncs.AddressMapOfAddressArrayAppend(ctx)
	aap.Params.NameAddr().SetValue(mapNames[1])
	aap.Params.ValueAddr().SetValue(mapVals[2])
	aap.Func.Post()
	require.NoError(t, ctx.Err)

	aal := testwasmlib.ScFuncs.AddressMapOfAddressArrayLength(ctx)
	aal.Params.NameAddr().SetValue(mapNames[0])
	aal.Func.Call()
	require.NoError(t, ctx.Err)
	length := aal.Results.Length()
	require.True(t, length.Exists())
	require.EqualValues(t, 2, length.Value())

	aav := testwasmlib.ScFuncs.AddressMapOfAddressArrayValue(ctx)
	aav.Params.NameAddr().SetValue(mapNames[0])
	aav.Params.Index().SetValue(0)
	aav.Func.Call()
	require.NoError(t, ctx.Err)
	value := aav.Results.ValueAddr()
	require.True(t, value.Exists())
	require.EqualValues(t, mapVals[0], value.Value())

	aas := testwasmlib.ScFuncs.AddressMapOfAddressArraySet(ctx)
	aas.Params.NameAddr().SetValue(mapNames[0])
	aas.Params.Index().SetValue(0)
	aas.Params.ValueAddr().SetValue(mapVals[3])
	aas.Func.Post()
	require.NoError(t, ctx.Err)

	aal = testwasmlib.ScFuncs.AddressMapOfAddressArrayLength(ctx)
	aal.Params.NameAddr().SetValue(mapNames[0])
	aal.Func.Call()
	require.NoError(t, ctx.Err)
	length = aal.Results.Length()
	require.True(t, length.Exists())
	require.EqualValues(t, 2, length.Value())

	aav = testwasmlib.ScFuncs.AddressMapOfAddressArrayValue(ctx)
	aav.Params.NameAddr().SetValue(mapNames[0])
	aav.Params.Index().SetValue(0)
	aav.Func.Call()
	require.NoError(t, ctx.Err)
	value = aav.Results.ValueAddr()
	require.True(t, value.Exists())
	require.EqualValues(t, mapVals[3], value.Value())

	aav = testwasmlib.ScFuncs.AddressMapOfAddressArrayValue(ctx)
	aav.Params.NameAddr().SetValue(mapNames[0])
	aav.Params.Index().SetValue(1)
	aav.Func.Call()
	require.NoError(t, ctx.Err)
	value = aav.Results.ValueAddr()
	require.True(t, value.Exists())
	require.EqualValues(t, mapVals[1], value.Value())

	aav = testwasmlib.ScFuncs.AddressMapOfAddressArrayValue(ctx)
	aav.Params.NameAddr().SetValue(mapNames[1])
	aav.Params.Index().SetValue(0)
	aav.Func.Call()
	require.NoError(t, ctx.Err)
	value = aav.Results.ValueAddr()
	require.True(t, value.Exists())
	require.EqualValues(t, mapVals[2], value.Value())
}

func TestInvalidIndexInGetStringMapOfStringArrayElt(t *testing.T) {
	ctx := setupTest(t)

	as := testwasmlib.ScFuncs.StringMapOfStringArrayAppend(ctx)
	as.Params.Name().SetValue("bands")
	as.Params.Value().SetValue("Simple Minds")
	as.Func.Post()
	require.NoError(t, ctx.Err)

	as = testwasmlib.ScFuncs.StringMapOfStringArrayAppend(ctx)
	as.Params.Name().SetValue("bands")
	as.Params.Value().SetValue("Dire Straits")
	as.Func.Post()
	require.NoError(t, ctx.Err)

	as = testwasmlib.ScFuncs.StringMapOfStringArrayAppend(ctx)
	as.Params.Name().SetValue("bands")
	as.Params.Value().SetValue("ELO")
	as.Func.Post()
	require.NoError(t, ctx.Err)

	al := testwasmlib.ScFuncs.StringMapOfStringArrayLength(ctx)
	al.Params.Name().SetValue("bands")
	al.Func.Call()
	require.NoError(t, ctx.Err)
	length := al.Results.Length()
	require.True(t, length.Exists())
	require.EqualValues(t, 3, length.Value())

	av := testwasmlib.ScFuncs.StringMapOfStringArrayValue(ctx)
	av.Params.Name().SetValue("bands")
	av.Params.Index().SetValue(100)
	av.Func.Call()
	require.Contains(t, ctx.Err.Error(), "invalid index")
}

func TestArrayOfStringArrayAppend(t *testing.T) {
	ctx := setupTest(t)

	aap := testwasmlib.ScFuncs.ArrayOfStringArrayAppend(ctx)
	aap.Params.Index().SetValue(0)
	aap.Params.Value().AppendString().SetValue("support")
	aap.Params.Value().AppendString().SetValue("freedom")
	aap.Func.Post()
	require.NoError(t, ctx.Err)

	aap = testwasmlib.ScFuncs.ArrayOfStringArrayAppend(ctx)
	aap.Params.Index().SetValue(1)
	aap.Params.Value().AppendString().SetValue("hail")
	aap.Params.Value().AppendString().SetValue("life")
	aap.Func.Post()
	require.NoError(t, ctx.Err)

	aav := testwasmlib.ScFuncs.ArrayOfStringArrayValue(ctx)
	aav.Params.Index0().SetValue(0)
	aav.Params.Index1().SetValue(0)
	aav.Func.Call()
	require.NoError(t, ctx.Err)
	require.EqualValues(t, "support", aav.Results.Value().Value())

	aav = testwasmlib.ScFuncs.ArrayOfStringArrayValue(ctx)
	aav.Params.Index0().SetValue(0)
	aav.Params.Index1().SetValue(1)
	aav.Func.Call()
	require.NoError(t, ctx.Err)
	require.EqualValues(t, "freedom", aav.Results.Value().Value())

	aav = testwasmlib.ScFuncs.ArrayOfStringArrayValue(ctx)
	aav.Params.Index0().SetValue(1)
	aav.Params.Index1().SetValue(0)
	aav.Func.Call()
	require.NoError(t, ctx.Err)
	require.EqualValues(t, "hail", aav.Results.Value().Value())

	aav = testwasmlib.ScFuncs.ArrayOfStringArrayValue(ctx)
	aav.Params.Index0().SetValue(1)
	aav.Params.Index1().SetValue(1)
	aav.Func.Call()
	require.NoError(t, ctx.Err)
	require.EqualValues(t, "life", aav.Results.Value().Value())
}

func TestArrayOfStringArrayClear(t *testing.T) {
	ctx := setupTest(t)

	aap := testwasmlib.ScFuncs.ArrayOfStringArrayAppend(ctx)
	aap.Params.Index().SetValue(0)
	aap.Params.Value().AppendString().SetValue("support")
	aap.Params.Value().AppendString().SetValue("freedom")
	aap.Func.Post()
	require.NoError(t, ctx.Err)

	aap = testwasmlib.ScFuncs.ArrayOfStringArrayAppend(ctx)
	aap.Params.Index().SetValue(1)
	aap.Params.Value().AppendString().SetValue("hail")
	aap.Params.Value().AppendString().SetValue("life")
	aap.Func.Post()
	require.NoError(t, ctx.Err)

	aav := testwasmlib.ScFuncs.ArrayOfStringArrayValue(ctx)
	aav.Params.Index0().SetValue(0)
	aav.Params.Index1().SetValue(0)
	aav.Func.Call()
	require.NoError(t, ctx.Err)
	require.EqualValues(t, "support", aav.Results.Value().Value())

	aav = testwasmlib.ScFuncs.ArrayOfStringArrayValue(ctx)
	aav.Params.Index0().SetValue(0)
	aav.Params.Index1().SetValue(1)
	aav.Func.Call()
	require.NoError(t, ctx.Err)
	require.EqualValues(t, "freedom", aav.Results.Value().Value())

	aav = testwasmlib.ScFuncs.ArrayOfStringArrayValue(ctx)
	aav.Params.Index0().SetValue(1)
	aav.Params.Index1().SetValue(0)
	aav.Func.Call()
	require.NoError(t, ctx.Err)
	require.EqualValues(t, "hail", aav.Results.Value().Value())

	aav = testwasmlib.ScFuncs.ArrayOfStringArrayValue(ctx)
	aav.Params.Index0().SetValue(1)
	aav.Params.Index1().SetValue(1)
	aav.Func.Call()
	require.NoError(t, ctx.Err)
	require.EqualValues(t, "life", aav.Results.Value().Value())

	ac := testwasmlib.ScFuncs.ArrayOfStringArrayClear(ctx)
	ac.Func.Post()
	require.NoError(t, ctx.Err)

	aav = testwasmlib.ScFuncs.ArrayOfStringArrayValue(ctx)
	aav.Params.Index0().SetValue(1)
	aav.Params.Index1().SetValue(0)
	aav.Func.Call()
	require.Error(t, ctx.Err)
}

func TestArrayOfAddressArrayClear(t *testing.T) {
	ctx := setupTest(t)
	mapVals := genTestAddress(ctx, 4)

	aap := testwasmlib.ScFuncs.ArrayOfAddressArrayAppend(ctx)
	aap.Params.Index().SetValue(0)
	aap.Params.ValueAddr().AppendAddress().SetValue(mapVals[0])
	aap.Params.ValueAddr().AppendAddress().SetValue(mapVals[1])
	aap.Func.Post()
	require.NoError(t, ctx.Err)

	aap = testwasmlib.ScFuncs.ArrayOfAddressArrayAppend(ctx)
	aap.Params.Index().SetValue(1)
	aap.Params.ValueAddr().AppendAddress().SetValue(mapVals[2])
	aap.Params.ValueAddr().AppendAddress().SetValue(mapVals[3])
	aap.Func.Post()
	require.NoError(t, ctx.Err)

	aav := testwasmlib.ScFuncs.ArrayOfAddressArrayValue(ctx)
	aav.Params.Index0().SetValue(0)
	aav.Params.Index1().SetValue(0)
	aav.Func.Call()
	require.NoError(t, ctx.Err)
	require.EqualValues(t, mapVals[0], aav.Results.ValueAddr().Value())

	aav = testwasmlib.ScFuncs.ArrayOfAddressArrayValue(ctx)
	aav.Params.Index0().SetValue(0)
	aav.Params.Index1().SetValue(1)
	aav.Func.Call()
	require.NoError(t, ctx.Err)
	require.EqualValues(t, mapVals[1], aav.Results.ValueAddr().Value())

	aav = testwasmlib.ScFuncs.ArrayOfAddressArrayValue(ctx)
	aav.Params.Index0().SetValue(1)
	aav.Params.Index1().SetValue(0)
	aav.Func.Call()
	require.NoError(t, ctx.Err)
	require.EqualValues(t, mapVals[2], aav.Results.ValueAddr().Value())

	aav = testwasmlib.ScFuncs.ArrayOfAddressArrayValue(ctx)
	aav.Params.Index0().SetValue(1)
	aav.Params.Index1().SetValue(1)
	aav.Func.Call()
	require.NoError(t, ctx.Err)
	require.EqualValues(t, mapVals[3], aav.Results.ValueAddr().Value())

	ac := testwasmlib.ScFuncs.ArrayOfAddressArrayClear(ctx)
	ac.Func.Post()
	require.NoError(t, ctx.Err)

	aav = testwasmlib.ScFuncs.ArrayOfAddressArrayValue(ctx)
	aav.Params.Index0().SetValue(1)
	aav.Params.Index1().SetValue(0)
	aav.Func.Call()
	require.Error(t, ctx.Err)
}

func TestArrayOfStringArraySet(t *testing.T) {
	ctx := setupTest(t)

	aap := testwasmlib.ScFuncs.ArrayOfStringArrayAppend(ctx)
	aap.Params.Index().SetValue(0)
	aap.Params.Value().AppendString().SetValue("support")
	aap.Params.Value().AppendString().SetValue("freedom")
	aap.Func.Post()
	require.NoError(t, ctx.Err)

	aap = testwasmlib.ScFuncs.ArrayOfStringArrayAppend(ctx)
	aap.Params.Index().SetValue(1)
	aap.Params.Value().AppendString().SetValue("hail")
	aap.Params.Value().AppendString().SetValue("life")
	aap.Func.Post()
	require.NoError(t, ctx.Err)

	aav := testwasmlib.ScFuncs.ArrayOfStringArrayValue(ctx)
	aav.Params.Index0().SetValue(0)
	aav.Params.Index1().SetValue(0)
	aav.Func.Call()
	require.NoError(t, ctx.Err)
	require.EqualValues(t, "support", aav.Results.Value().Value())

	aav = testwasmlib.ScFuncs.ArrayOfStringArrayValue(ctx)
	aav.Params.Index0().SetValue(0)
	aav.Params.Index1().SetValue(1)
	aav.Func.Call()
	require.NoError(t, ctx.Err)
	require.EqualValues(t, "freedom", aav.Results.Value().Value())

	aav = testwasmlib.ScFuncs.ArrayOfStringArrayValue(ctx)
	aav.Params.Index0().SetValue(1)
	aav.Params.Index1().SetValue(0)
	aav.Func.Call()
	require.NoError(t, ctx.Err)
	require.EqualValues(t, "hail", aav.Results.Value().Value())

	aav = testwasmlib.ScFuncs.ArrayOfStringArrayValue(ctx)
	aav.Params.Index0().SetValue(1)
	aav.Params.Index1().SetValue(1)
	aav.Func.Call()
	require.NoError(t, ctx.Err)
	require.EqualValues(t, "life", aav.Results.Value().Value())

	aas := testwasmlib.ScFuncs.ArrayOfStringArraySet(ctx)
	aas.Params.Index0().SetValue(1)
	aas.Params.Index1().SetValue(1)
	aas.Params.Value().SetValue("moon")
	aas.Func.Post()
	require.NoError(t, ctx.Err)

	aav = testwasmlib.ScFuncs.ArrayOfStringArrayValue(ctx)
	aav.Params.Index0().SetValue(1)
	aav.Params.Index1().SetValue(1)
	aav.Func.Call()
	require.EqualValues(t, "moon", aav.Results.Value().Value())
}

func TestArrayOfAddressArraySet(t *testing.T) {
	ctx := setupTest(t)
	mapVals := genTestAddress(ctx, 4)

	aap := testwasmlib.ScFuncs.ArrayOfAddressArrayAppend(ctx)
	aap.Params.Index().SetValue(0)
	aap.Params.ValueAddr().AppendAddress().SetValue(mapVals[0])
	aap.Params.ValueAddr().AppendAddress().SetValue(mapVals[1])
	aap.Func.Post()
	require.NoError(t, ctx.Err)

	aap = testwasmlib.ScFuncs.ArrayOfAddressArrayAppend(ctx)
	aap.Params.Index().SetValue(1)
	aap.Params.ValueAddr().AppendAddress().SetValue(mapVals[2])
	aap.Func.Post()
	require.NoError(t, ctx.Err)

	aav := testwasmlib.ScFuncs.ArrayOfAddressArrayValue(ctx)
	aav.Params.Index0().SetValue(0)
	aav.Params.Index1().SetValue(0)
	aav.Func.Call()
	require.NoError(t, ctx.Err)
	require.EqualValues(t, mapVals[0], aav.Results.ValueAddr().Value())

	aav = testwasmlib.ScFuncs.ArrayOfAddressArrayValue(ctx)
	aav.Params.Index0().SetValue(0)
	aav.Params.Index1().SetValue(1)
	aav.Func.Call()
	require.NoError(t, ctx.Err)
	require.EqualValues(t, mapVals[1], aav.Results.ValueAddr().Value())

	aav = testwasmlib.ScFuncs.ArrayOfAddressArrayValue(ctx)
	aav.Params.Index0().SetValue(1)
	aav.Params.Index1().SetValue(0)
	aav.Func.Call()
	require.NoError(t, ctx.Err)
	require.EqualValues(t, mapVals[2], aav.Results.ValueAddr().Value())

	aas := testwasmlib.ScFuncs.ArrayOfAddressArraySet(ctx)
	aas.Params.Index0().SetValue(0)
	aas.Params.Index1().SetValue(1)
	aas.Params.ValueAddr().SetValue(mapVals[3])
	aas.Func.Post()
	require.NoError(t, ctx.Err)

	aav = testwasmlib.ScFuncs.ArrayOfAddressArrayValue(ctx)
	aav.Params.Index0().SetValue(0)
	aav.Params.Index1().SetValue(1)
	aav.Func.Call()
	require.EqualValues(t, mapVals[3], aav.Results.ValueAddr().Value())
}

func TestStringMapOfStringMapClear(t *testing.T) {
	// test reproduces a problem that needs fixing
	t.SkipNow()

	ctx := setupTest(t)

	as := testwasmlib.ScFuncs.StringMapOfStringMapSet(ctx)
	as.Params.Name().SetValue("albums")
	as.Params.Key().SetValue("Simple Minds")
	as.Params.Value().SetValue("New Gold Dream")
	as.Func.Post()
	require.NoError(t, ctx.Err)

	as = testwasmlib.ScFuncs.StringMapOfStringMapSet(ctx)
	as.Params.Name().SetValue("albums")
	as.Params.Key().SetValue("Dire Straits")
	as.Params.Value().SetValue("Calling Elvis")
	as.Func.Post()
	require.NoError(t, ctx.Err)

	as = testwasmlib.ScFuncs.StringMapOfStringMapSet(ctx)
	as.Params.Name().SetValue("albums")
	as.Params.Key().SetValue("ELO")
	as.Params.Value().SetValue("Mr. Blue Sky")
	as.Func.Post()
	require.NoError(t, ctx.Err)

	av := testwasmlib.ScFuncs.StringMapOfStringMapValue(ctx)
	av.Params.Name().SetValue("albums")
	av.Params.Key().SetValue("Dire Straits")
	av.Func.Call()
	require.NoError(t, ctx.Err)
	value := av.Results.Value()
	require.True(t, value.Exists())
	require.EqualValues(t, "Calling Elvis", value.Value())

	ac := testwasmlib.ScFuncs.StringMapOfStringMapClear(ctx)
	ac.Params.Name().SetValue("albums")
	ac.Func.Post()
	require.NoError(t, ctx.Err)

	av = testwasmlib.ScFuncs.StringMapOfStringMapValue(ctx)
	av.Params.Name().SetValue("albums")
	av.Params.Key().SetValue("Dire Straits")
	av.Func.Call()
	require.NoError(t, ctx.Err)
	require.EqualValues(t, "", value.Value())
}

func TestAddressMapOfAddressMapClear(t *testing.T) {
	// test reproduces a problem that needs fixing
	t.SkipNow()

	ctx := setupTest(t)

	mapNames, mapKeys, mapVals := genTestAddress(ctx, 2), genTestAddress(ctx, 4), genTestAddress(ctx, 4)

	as := testwasmlib.ScFuncs.AddressMapOfAddressMapSet(ctx)
	as.Params.NameAddr().SetValue(mapNames[0])
	as.Params.KeyAddr().SetValue(mapKeys[0])
	as.Params.ValueAddr().SetValue(mapVals[0])
	as.Func.Post()
	require.NoError(t, ctx.Err)

	as = testwasmlib.ScFuncs.AddressMapOfAddressMapSet(ctx)
	as.Params.NameAddr().SetValue(mapNames[0])
	as.Params.KeyAddr().SetValue(mapKeys[1])
	as.Params.ValueAddr().SetValue(mapVals[1])
	as.Func.Post()
	require.NoError(t, ctx.Err)

	as = testwasmlib.ScFuncs.AddressMapOfAddressMapSet(ctx)
	as.Params.NameAddr().SetValue(mapNames[1])
	as.Params.KeyAddr().SetValue(mapKeys[2])
	as.Params.ValueAddr().SetValue(mapVals[2])
	as.Func.Post()
	require.NoError(t, ctx.Err)

	av := testwasmlib.ScFuncs.AddressMapOfAddressMapValue(ctx)
	av.Params.NameAddr().SetValue(mapNames[0])
	av.Params.KeyAddr().SetValue(mapKeys[0])
	av.Func.Call()
	require.NoError(t, ctx.Err)
	value := av.Results.ValueAddr()
	require.True(t, value.Exists())
	require.EqualValues(t, mapVals[0], value.Value())

	ac := testwasmlib.ScFuncs.AddressMapOfAddressMapClear(ctx)
	ac.Params.NameAddr().SetValue(mapNames[0])
	ac.Func.Post()
	require.NoError(t, ctx.Err)

	av = testwasmlib.ScFuncs.AddressMapOfAddressMapValue(ctx)
	av.Params.NameAddr().SetValue(mapNames[0])
	av.Params.KeyAddr().SetValue(mapKeys[0])
	av.Func.Call()
	require.NoError(t, ctx.Err)
	require.EqualValues(t, "", value.Value())
}

func TestAddressMapOfAddressMapSet(t *testing.T) {
	ctx := setupTest(t)

	mapNames, mapKeys, mapVals := genTestAddress(ctx, 2), genTestAddress(ctx, 4), genTestAddress(ctx, 4)

	as := testwasmlib.ScFuncs.AddressMapOfAddressMapSet(ctx)
	as.Params.NameAddr().SetValue(mapNames[0])
	as.Params.KeyAddr().SetValue(mapKeys[0])
	as.Params.ValueAddr().SetValue(mapVals[0])
	as.Func.Post()
	require.NoError(t, ctx.Err)

	as = testwasmlib.ScFuncs.AddressMapOfAddressMapSet(ctx)
	as.Params.NameAddr().SetValue(mapNames[0])
	as.Params.KeyAddr().SetValue(mapKeys[1])
	as.Params.ValueAddr().SetValue(mapVals[1])
	as.Func.Post()
	require.NoError(t, ctx.Err)

	as = testwasmlib.ScFuncs.AddressMapOfAddressMapSet(ctx)
	as.Params.NameAddr().SetValue(mapNames[1])
	as.Params.KeyAddr().SetValue(mapKeys[2])
	as.Params.ValueAddr().SetValue(mapVals[2])
	as.Func.Post()
	require.NoError(t, ctx.Err)

	av := testwasmlib.ScFuncs.AddressMapOfAddressMapValue(ctx)
	av.Params.NameAddr().SetValue(mapNames[0])
	av.Params.KeyAddr().SetValue(mapKeys[0])
	av.Func.Call()
	require.NoError(t, ctx.Err)
	value := av.Results.ValueAddr()
	require.True(t, value.Exists())
	require.EqualValues(t, mapVals[0], value.Value())

	av = testwasmlib.ScFuncs.AddressMapOfAddressMapValue(ctx)
	av.Params.NameAddr().SetValue(mapNames[0])
	av.Params.KeyAddr().SetValue(mapKeys[1])
	av.Func.Call()
	require.NoError(t, ctx.Err)
	value = av.Results.ValueAddr()
	require.True(t, value.Exists())
	require.EqualValues(t, mapVals[1], value.Value())

	av = testwasmlib.ScFuncs.AddressMapOfAddressMapValue(ctx)
	av.Params.NameAddr().SetValue(mapNames[1])
	av.Params.KeyAddr().SetValue(mapKeys[2])
	av.Func.Call()
	require.NoError(t, ctx.Err)
	value = av.Results.ValueAddr()
	require.True(t, value.Exists())
	require.EqualValues(t, mapVals[2], value.Value())

	as = testwasmlib.ScFuncs.AddressMapOfAddressMapSet(ctx)
	as.Params.NameAddr().SetValue(mapNames[1])
	as.Params.KeyAddr().SetValue(mapKeys[2])
	as.Params.ValueAddr().SetValue(mapVals[3])
	as.Func.Post()
	require.NoError(t, ctx.Err)

	av = testwasmlib.ScFuncs.AddressMapOfAddressMapValue(ctx)
	av.Params.NameAddr().SetValue(mapNames[1])
	av.Params.KeyAddr().SetValue(mapKeys[2])
	av.Func.Call()
	require.NoError(t, ctx.Err)
	value = av.Results.ValueAddr()
	require.True(t, value.Exists())
	require.EqualValues(t, mapVals[3], value.Value())

	as = testwasmlib.ScFuncs.AddressMapOfAddressMapSet(ctx)
	as.Params.NameAddr().SetValue(mapNames[0])
	as.Params.KeyAddr().SetValue(mapKeys[1])
	as.Params.ValueAddr().SetValue(mapVals[3])
	as.Func.Post()
	require.NoError(t, ctx.Err)

	av = testwasmlib.ScFuncs.AddressMapOfAddressMapValue(ctx)
	av.Params.NameAddr().SetValue(mapNames[0])
	av.Params.KeyAddr().SetValue(mapKeys[1])
	av.Func.Call()
	require.NoError(t, ctx.Err)
	value = av.Results.ValueAddr()
	require.True(t, value.Exists())
	require.EqualValues(t, mapVals[3], value.Value())
}

func TestStringMapOfStringMapSet(t *testing.T) {
	ctx := setupTest(t)

	as := testwasmlib.ScFuncs.StringMapOfStringMapSet(ctx)
	as.Params.Name().SetValue("albums")
	as.Params.Key().SetValue("Simple Minds")
	as.Params.Value().SetValue("New Gold Dream")
	as.Func.Post()
	require.NoError(t, ctx.Err)

	as = testwasmlib.ScFuncs.StringMapOfStringMapSet(ctx)
	as.Params.Name().SetValue("albums")
	as.Params.Key().SetValue("Dire Straits")
	as.Params.Value().SetValue("Calling Elvis")
	as.Func.Post()
	require.NoError(t, ctx.Err)

	as = testwasmlib.ScFuncs.StringMapOfStringMapSet(ctx)
	as.Params.Name().SetValue("albums")
	as.Params.Key().SetValue("ELO")
	as.Params.Value().SetValue("Mr. Blue Sky")
	as.Func.Post()
	require.NoError(t, ctx.Err)

	av := testwasmlib.ScFuncs.StringMapOfStringMapValue(ctx)
	av.Params.Name().SetValue("albums")
	av.Params.Key().SetValue("Dire Straits")
	av.Func.Call()
	require.NoError(t, ctx.Err)
	value := av.Results.Value()
	require.True(t, value.Exists())
	require.EqualValues(t, "Calling Elvis", value.Value())

	av = testwasmlib.ScFuncs.StringMapOfStringMapValue(ctx)
	av.Params.Name().SetValue("albums")
	av.Params.Key().SetValue("Simple Minds")
	av.Func.Call()
	require.NoError(t, ctx.Err)
	require.True(t, value.Exists())
	require.EqualValues(t, "New Gold Dream", av.Results.Value().Value())

	av = testwasmlib.ScFuncs.StringMapOfStringMapValue(ctx)
	av.Params.Name().SetValue("albums")
	av.Params.Key().SetValue("ELO")
	av.Func.Call()
	require.NoError(t, ctx.Err)
	require.True(t, value.Exists())
	require.EqualValues(t, "Mr. Blue Sky", av.Results.Value().Value())

	as = testwasmlib.ScFuncs.StringMapOfStringMapSet(ctx)
	as.Params.Name().SetValue("albums")
	as.Params.Key().SetValue("Simple Minds")
	as.Params.Value().SetValue("Life in a Day")
	as.Func.Post()
	require.NoError(t, ctx.Err)

	av = testwasmlib.ScFuncs.StringMapOfStringMapValue(ctx)
	av.Params.Name().SetValue("albums")
	av.Params.Key().SetValue("Simple Minds")
	av.Func.Call()
	require.NoError(t, ctx.Err)
	require.True(t, value.Exists())
	require.EqualValues(t, "Life in a Day", av.Results.Value().Value())
}

func TestArrayOfStringMapClear(t *testing.T) {
	ctx := setupTest(t)

	ams := testwasmlib.ScFuncs.ArrayOfStringMapSet(ctx)
	ams.Params.Index().SetValue(0)
	ams.Params.Key().SetValue("Simple Minds")
	ams.Params.Value().SetValue("New Gold Dream")
	ams.Func.Post()
	require.NoError(t, ctx.Err)

	ams = testwasmlib.ScFuncs.ArrayOfStringMapSet(ctx)
	ams.Params.Index().SetValue(0)
	ams.Params.Key().SetValue("Dire Straits")
	ams.Params.Value().SetValue("Calling Elvis")
	ams.Func.Post()
	require.NoError(t, ctx.Err)

	ams = testwasmlib.ScFuncs.ArrayOfStringMapSet(ctx)
	ams.Params.Index().SetValue(1)
	ams.Params.Key().SetValue("ELO")
	ams.Params.Value().SetValue("Mr. Blue Sky")
	ams.Func.Post()
	require.NoError(t, ctx.Err)

	amv := testwasmlib.ScFuncs.ArrayOfStringMapValue(ctx)
	amv.Params.Index().SetValue(0)
	amv.Params.Key().SetValue("Simple Minds")
	amv.Func.Call()
	require.NoError(t, ctx.Err)
	value := amv.Results.Value()
	require.True(t, value.Exists())
	require.EqualValues(t, "New Gold Dream", value.Value())

	amv = testwasmlib.ScFuncs.ArrayOfStringMapValue(ctx)
	amv.Params.Index().SetValue(0)
	amv.Params.Key().SetValue("Dire Straits")
	amv.Func.Call()
	require.NoError(t, ctx.Err)
	value = amv.Results.Value()
	require.True(t, value.Exists())
	require.EqualValues(t, "Calling Elvis", value.Value())

	amv = testwasmlib.ScFuncs.ArrayOfStringMapValue(ctx)
	amv.Params.Index().SetValue(1)
	amv.Params.Key().SetValue("ELO")
	amv.Func.Call()
	require.NoError(t, ctx.Err)
	value = amv.Results.Value()
	require.True(t, value.Exists())
	require.EqualValues(t, "Mr. Blue Sky", value.Value())

	amc := testwasmlib.ScFuncs.ArrayOfStringMapClear(ctx)
	amc.Func.Post()
	require.NoError(t, ctx.Err)

	amv = testwasmlib.ScFuncs.ArrayOfStringMapValue(ctx)
	amv.Params.Index().SetValue(1)
	amv.Params.Key().SetValue("ELO")
	amv.Func.Call()
	require.Error(t, ctx.Err)

	amv = testwasmlib.ScFuncs.ArrayOfStringMapValue(ctx)
	amv.Params.Index().SetValue(0)
	amv.Params.Key().SetValue("Simple Minds")
	amv.Func.Call()
	require.Error(t, ctx.Err)
}
