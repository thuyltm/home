package protocolv3

import (
	"home/pacman/network/binary"
	"io"
)

const (
	DataObjectError = iota
	DataObjectTx
	DataObjectBlock
	DataObjectFilterBlock
	DataObjectCmpctBlock
)

type MsgInv struct {
	Count     uint8
	Inventory []InvVector
}

func (inv *MsgInv) UnmarshalBinary(r io.Reader) error {
	d := binary.NewDecoder(r)
	if err := d.Decode(&inv.Count); err != nil {
		return err
	}
	for i := uint8(0); i < inv.Count; i++ {
		var v InvVector
		if err := d.Decode(&v); err != nil {
			return err
		}
		inv.Inventory = append(inv.Inventory, v)
	}
	return nil
}

type InvVector struct {
	Type uint32
	Hash [32]byte
}
