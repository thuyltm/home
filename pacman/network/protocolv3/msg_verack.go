package protocolv3

func NewVerackMsg(network string) (*Message, error) {
	msg, err := NewMessage("verack", network, []byte{})
	if err != nil {
		return nil, err
	}
	return msg, nil
}
