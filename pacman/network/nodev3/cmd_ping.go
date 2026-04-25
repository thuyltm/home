package nodev3

import (
	"home/pacman/network/binary"
	. "home/pacman/network/protocolv3"
	"io"
)

func (n Node) handlePing(header *MessageHeader, conn io.ReadWriter) error {
	var ping MsgPing
	lr := io.LimitReader(conn, int64(header.Length))
	if err := binary.NewDecoder(lr).Decode(&ping); err != nil {
		return err
	}
	pong, err := NewPongMsg(n.Network, ping.Nonce)
	if err != nil {
		return err
	}
	msg, err := binary.Marshal(pong)
	if err != nil {
		return nil
	}
	if _, err := conn.Write(msg); err != nil {
		return err
	}
	return nil
}
