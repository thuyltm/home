package node

import (
	"bytes"
	"fmt"
	"home/pacman/network/binary"
	. "home/pacman/network/protocolv2"
	"io"
	"math/rand/v2"
	"net"
	"time"

	"github.com/sirupsen/logrus"
)

type Node struct {
	Network      string
	NetworkMagic Magic
	UserAgent    string
}

func New(network, userAgent string) (*Node, error) {
	networkMagic, ok := Networks[network]
	if !ok {
		return nil, fmt.Errorf("unsupported network %s", network)
	}
	return &Node{
		Network:      network,
		NetworkMagic: networkMagic,
		UserAgent:    userAgent,
	}, nil
}

func (no Node) Run(nodeAddr string) error {
	peerAddr, err := ParseNodeAddr(nodeAddr)
	if err != nil {
		return err
	}
	version := MsgVersion{
		Version:   Version,
		Services:  SrvNodeNetwork,
		Timestamp: time.Now().UTC().Unix(),
		AddrRecv: VersionNetAddr{
			Services: SrvNodeNetwork,
			IP:       peerAddr.IP,
			Port:     peerAddr.Port,
		},
		AddrFrom: VersionNetAddr{
			Services: SrvNodeNetwork,
			IP:       NewIPv4(127, 0, 0, 1),
			Port:     9334,
		},
		Nonce:       nonce(),
		UserAgent:   NewUserAgent(no.UserAgent),
		StartHeight: -1,
		Relay:       true,
	}
	msg, err := NewMessage("version", no.Network, version)
	if err != nil {
		logrus.Fatalln(err)
	}
	msgSerialized, err := binary.Marshal(msg)
	if err != nil {
		logrus.Fatalln(err)
	}
	conn, err := net.Dial("tcp", nodeAddr)
	if err != nil {
		logrus.Fatalln(err)
	}
	defer conn.Close()
	_, err = conn.Write(msgSerialized)
	if err != nil {
		logrus.Fatalln(err)
	}
	tmp := make([]byte, MsgHeaderLength)
Loop:
	for {
		n, err := conn.Read(tmp)
		if err != nil {
			if err != io.EOF {
				return err
			}
			break Loop
		}
		logrus.Debugf("received header: %x", tmp[:n])
		var msgHeader MessageHeader
		if err := binary.NewDecoder(bytes.NewReader(tmp[:n])).Decode(&msgHeader); err != nil {
			logrus.Error(err)
			continue
		}
		logrus.Debugf("received message: %s", msgHeader.Command)
		switch msgHeader.CommandString() {
		case "version":
			if err := no.handleVersion(&msgHeader, conn); err != nil {
				logrus.Errorf("failed to handle 'version': %+v", err)
				continue
			}
		}
	}
	return nil
}

func nonce() uint64 {
	return rand.Uint64()
}
