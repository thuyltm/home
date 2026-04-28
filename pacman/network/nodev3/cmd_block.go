package nodev3

import (
	"fmt"
	"home/pacman/network/binary"
	. "home/pacman/network/protocolv3"
	"io"

	"github.com/sirupsen/logrus"
)

func (no Node) handleBlock(header *MessageHeader, conn io.ReadWriter) error {
	var block MsgBlock
	lr := io.LimitReader(conn, int64(header.Length))
	if err := binary.NewDecoder(lr).Decode(&block); err != nil {
		return err
	}
	hash, err := block.Hash()
	if err != nil {
		return fmt.Errorf("block.Hash: %+v", err)
	}
	logrus.Debugf("block: %x", hash)
	if err := block.Verify(); err != nil {
		return fmt.Errorf("rejected invalid block %x", hash)
	}
	no.mempool.NewBlockCh <- block
	return nil
}
