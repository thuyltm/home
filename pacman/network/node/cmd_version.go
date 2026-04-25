package node

import (
	"io"

	"home/pacman/network/binary"
	. "home/pacman/network/protocolv2"

	"github.com/sirupsen/logrus"
)

func (n Node) handleVersion(header *MessageHeader, conn io.ReadWriter) error {
	var version MsgVersion
	lr := io.LimitReader(conn, int64(header.Length))
	if err := binary.NewDecoder(lr).Decode(&version); err != nil {
		return err
	}
	logrus.Infof("VERSION: %+v", version.UserAgent.String)
	verack, err := NewVerackMsg(n.Network)
	if err != nil {
		return err
	}
	msg, err := binary.Marshal(verack)
	if err != nil {
		return err
	}
	if _, err := conn.Write(msg); err != nil {
		return err
	}

	return nil
}
