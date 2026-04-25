package protocolv3

import (
	"bytes"
	"home/pacman/network/binary"
)

type MsgGetData struct {
	Count     uint8
	Inventory []InvVector
}

func (gd MsgGetData) MarshalBinary() ([]byte, error) {
	buf := bytes.NewBuffer([]byte{})

	b, err := binary.Marshal(gd.Count)
	if err != nil {
		return nil, err
	}
	if _, err := buf.Write(b); err != nil {
		return nil, err
	}
	for _, i := range gd.Inventory {
		b, err := binary.Marshal(i)
		if err != nil {
			return nil, err
		}
		if _, err := buf.Write(b); err != nil {
			return nil, err
		}
	}
	return buf.Bytes(), nil
}
