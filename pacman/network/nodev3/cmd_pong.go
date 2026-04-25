package nodev3

import (
	"home/pacman/network/binary"
	. "home/pacman/network/protocolv3"
	"io"
)

func (n Node) handlePong(header *MessageHeader, conn io.ReadWriter) error {
	var pong MsgPing
	lr := io.LimitReader(conn, int64(header.Length))
	if err := binary.NewDecoder(lr).Decode(&pong); err != nil {
		return err
	}
	n.PongCh <- pong.Nonce
	return nil
}
