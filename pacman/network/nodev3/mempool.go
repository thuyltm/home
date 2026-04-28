package nodev3

import (
	"encoding/hex"
	. "home/pacman/network/protocolv3"

	"github.com/sirupsen/logrus"
)

type Mempool struct {
	NewBlockCh chan MsgBlock
	NewTxCh    chan MsgTx
	txs        map[string]*MsgTx
}

func NewMempool() *Mempool {
	return &Mempool{
		NewBlockCh: make(chan MsgBlock),
		NewTxCh:    make(chan MsgTx),
		txs:        make(map[string]*MsgTx),
	}
}

func (m Mempool) Run() {
	for {
		select {
		case tx := <-m.NewTxCh:
			hash, err := tx.Hash()
			if err != nil {
				logrus.Errorf("failed to calculate tx hash: %+v", err)
				break
			}
			txid := hex.EncodeToString(hash)
			m.txs[txid] = &tx
		case block := <-m.NewBlockCh:
			for _, tx := range block.Txs {
				hash, err := tx.Hash()
				if err != nil {
					logrus.Errorf("failed to calculate tx hash: %+v", err)
					break
				}
				txid := hex.EncodeToString(hash)
				delete(m.txs, txid)
			}
		}
	}
}
