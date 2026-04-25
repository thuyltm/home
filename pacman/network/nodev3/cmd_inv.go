package nodev3

import (
	"home/pacman/network/binary"
	. "home/pacman/network/protocolv3"
	"io"
)

func (no Node) handleInv(header *MessageHeader, conn io.ReadWriter) error {
	var inv MsgInv
	lr := io.LimitReader(conn, int64(header.Length))
	if err := binary.NewDecoder(lr).Decode(&inv); err != nil {
		return err
	}
	var getData MsgGetData
	getData.Inventory = inv.Inventory
	getData.Count = inv.Count
	getDataMsg, err := NewMessage("getdata", no.Network, getData)
	if err != nil {
		return err
	}
	msg, err := binary.Marshal(getDataMsg)
	if err != nil {
		return err
	}
	if _, err := conn.Write(msg); err != nil {
		return err
	}
	return nil
}
