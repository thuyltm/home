# Serialize
1. For serialization, we'll implement the **non-streaming approach** because, before starting serializing, we already have all the data we need
2. Marshal/Unmarshal functions are the well-known non-streaming method 
3. Serialization is type-based, not message-based
binray/marshaler.go
```go
func Marshal(v interface{}) ([]byte, error) {
    switch val := v.(type) {
        case uint8, int32, uint32, int64, uint64, bool:
            ...
        case uint16:
            ...
        case [magicAndChecksumLength]byte:
            ...
        case [commandLength]byte:
            ...
        case []byte:
            ...
        case string:
            ...
        case Marshaler:
            ...
        default:
            if reflect.ValueOf(v).Kind() == reflect.Struct {
                ...
            }
    }
}
```
Now, we can use binary.Marshal to serialize messages:
```go
msg, err := protocol.NewMessage("version", network, version)
msgSerialized, err := binary.Marshal(msg)
```

# Deserialize
1. For deserialization, we'll implemtn the streaming approach because the node will be reading serialized message from a TCP connection, which is a stream
2. The streaming implementation consits of Decode/Encode methods to decode raw data from input stream and write encoded data to output stream

Specifically, we read raw bytes from a TCP connection and convert them to message structure

    1. splitting it into pieces (one piece per struct field)
    2. decode
    3. assign decoded values to struct fields
    This implies that we must have a full message before deserializing it. But TPC connections are stream, we will receive a long sequence of bytes. But we need to get something complete before starting deserialization

This structure is a message header that contains meta information about the message
```go
type Message struct {
    Magic    [magicLength]byte
	Command  [commandLength]byte
	Length   uint32
	Checksum [checksumLength]byte
	Payload  []byte
}
```
After our node has received, deserialized, and validate a message header, it knows:
    
    1. What network this message is for
    2. What command it is
    3. What's the length of the message payload
3. binary.Decode decode message payload
```go
type Decoder struct {
	r io.Reader
}
```
io.Reader reads and decodes data from a stream and do
    
    1. Takes a pointer to a value of any supported type
    2. Checks the type of the pointer
    3. Reads proper number of bytes from the stream
    4. Properly decodes the bytes
    5. Saves the decoded value at the passed pointer
```go
func (d Decoder) Decode(v any) error {
	switch val := v.(type) {
	case *bool:
		...
	case *int32:
		...
	case *int64:
		...
	case *uint8:
		...
	case *uint16:
		...
	case *uint32:
		...
	case *uint64:
		...
	case *[magicAndChecksumLength]byte:
		...
	case *[hashLength]byte:
		...
	case *[commandLength]byte:
		...
	case Unmarshaler:
		...
	default:
		if reflect.ValueOf(v).Kind() == reflect.Ptr &&
			reflect.ValueOf(v).Elem().Kind() == reflect.Struct {
			if err := d.decodeStruct(v); err != nil {
				return err
			}
			break
		}

		return fmt.Errorf("unsupported type %s", reflect.TypeOf(v).String())
	}
	return nil
}
```
4. Deseralize 'version' message
We split Message structure into two structures: MessageHeader and Payload
```go
type MessageHeader struct {
	Magic    [magicLength]byte
	Command  [commandLength]byte
	Length   uint32
	Checksum [checksumLength]byte
}

type Message struct {
	MessageHeader
	Payload []byte
}
```
If we want to read a message, 
1. we have to read its header first
2. After the header was decoded, we must validate it. We validate magic and command
3. If the header is correct, we can start handling the command
```go
tmp := make([]byte, MsgHeaderLength)
Loop:
	for {
		n, err := conn.Read(tmp)
		if err != nil {
			...
			break Loop
		}		
		var msgHeader MessageHeader
		if err := binary.NewDecoder(bytes.NewReader(tmp[:n])).Decode(&msgHeader); err != nil {
			logrus.Error(err)
			continue
		}
		
		switch msgHeader.CommandString() {
		case "version":
			if err := no.handleVersion(&msgHeader, conn); err != nil {
				logrus.Errorf(err)
				continue
			}
		}
	}
	return nil
```
