Messages serialization and deserialization is very important. Without proper (de)serialization it won't be possible to build communication between nodes. 

Golang comes with encoding/gob library that allows to serialize and deserialize Golang structs. This is very Golang way of serialization, other languages don't support it

We use JSON, YAML, etc. to serialize the message, but Bitcoin nodes won't understand because the Bitcoin specification describes a different way

The Bitcoin protocol uses a very simple serialization mechanic: Concatenate byte presentation of every field in the order

-------------------------------------------------------------------------------------------------------
| BYTES(Msg.Magic) + BYTES(Msg.Command) + BYTES(Msg.Length) + BYTES(Msg.Checksum) + BYTES(Msg.Payload)|
-------------------------------------------------------------------------------------------------------
```golang

type Message struct {
	Magic    [magicLength]byte
	Command  [commandLength]byte
	Length   uint32
	Checksum [checksumLength]byte
	Payload  []byte
}

func (m Message) Serialize() ([]byte, error) {
	var buf bytes.Buffer
	if _, err := buf.Write(m.Magic[:]); err != nil {
		return nil, err
	}

	if _, err := buf.Write(m.Command[:]); err != nil {
		return nil, err
	}

	if err := binary.Write(&buf, binary.LittleEndian, m.Length); err != nil {
		return nil, err
	}

	if _, err := buf.Write(m.Checksum[:]); err != nil {
		return nil, err
	}

	if _, err := buf.Write(m.Payload); err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}
```