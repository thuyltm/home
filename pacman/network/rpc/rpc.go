package rpc

import (
	"fmt"
	. "home/pacman/network/protocolv3"
)

type Node interface {
	Mempool() map[string]*MsgTx
}

type RPC struct {
	node Node
}

type MempoolArgs interface{}

type MempoolReply string

func (r RPC) GetMempool(args *MempoolArgs, reply *MempoolReply) error {
	txs := r.node.Mempool()
	*reply = MempoolReply(formatMempoolReply(txs))
	return nil
}

func formatMempoolReply(txs map[string]*MsgTx) string {
	var result string
	for k := range txs {
		result += fmt.Sprintf("%s\n", k)
	}
	result += fmt.Sprintf("Total %d transactions", len(txs))
	return result
}
