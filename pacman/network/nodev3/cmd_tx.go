package nodev3

import (
	"home/pacman/network/binary"
	. "home/pacman/network/protocolv3"
	"io"

	"github.com/sirupsen/logrus"
)

func (no Node) handleTx(header *MessageHeader, conn io.ReadWriter) error {
	var tx MsgTx
	lr := io.LimitReader(conn, int64(header.Length))
	if err := binary.NewDecoder(lr).Decode(&tx); err != nil {
		return err
	}
	logrus.Debugf("transaction: %+v", tx)
	return nil
}
