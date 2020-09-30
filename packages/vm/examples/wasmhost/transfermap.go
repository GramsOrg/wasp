package wasmhost

import (
	"github.com/iotaledger/goshimmer/dapps/valuetransfers/packages/address"
	"github.com/iotaledger/goshimmer/dapps/valuetransfers/packages/balance"
	"github.com/mr-tron/base58"
)

type TransferMap struct {
	MapObject
	address string
	amount  int64
	color   string
}

func NewTransferMap(vm *wasmVMPocProcessor) HostObject {
	return &TransferMap{MapObject: MapObject{vm: vm, name: "Transfer"}}
}

func (o *TransferMap) Send() {
	o.vm.Trace("TRANSFER a%d c'%16s' a'%16s'", o.amount, o.color, o.address)
	addr, err := address.FromBase58(o.address)
	if err != nil {
		o.vm.ctx.Panic("Invalid base58 address")
	}

	// when no color specified default is ColorIOTA
	color := balance.ColorIOTA
	if o.color != "iota" {
		bytes, err := base58.Decode(o.color)
		if err != nil || len(bytes) != balance.ColorLength {
			o.vm.ctx.Panic("Invalid base58 color")
		}
		copy(color[:], bytes)
	}

	if !o.vm.ctx.AccessSCAccount().MoveTokens(&addr, &color, o.amount) {
		o.vm.ctx.Panic("Failed to move tokens")
	}
}

func (o *TransferMap) SetInt(keyId int32, value int64) {
	switch keyId {
	case KeyLength:
		// clear transfer, tracker will still know about it
		// so maybe move it to an allocation pool for reuse
		o.address = ""
		o.color = ""
		o.amount = 0
	case KeyAmount:
		o.amount = value
	default:
		o.MapObject.SetInt(keyId, value)
	}
}

func (o *TransferMap) SetString(keyId int32, value string) {
	switch keyId {
	case KeyAddress:
		o.address = value
	case KeyColor:
		o.color = value
	default:
		o.MapObject.SetString(keyId, value)
	}
}
