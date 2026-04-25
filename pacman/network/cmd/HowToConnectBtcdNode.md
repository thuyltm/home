1. Serialize the message
```golang
msgSerialized, err := msg.Serialize()
```
2. Connect to the local btcd node
```golang
conn, err := net.Dial("tcp", "127.0.0.1:9333")
defer conn.Close()
```
3. Send message by writing the serialized message to the TCP connection
```golang
_, err = conn.Write(msgSerialized)
```
4. Wait for any response and print it out
```golang
tmp := make([]byte, 256)
for {
    n, err := conn.Read(tmp)
    if err != nil {
        if err != io.EOF {
            logrus.Fatalln(err)
        }
        return
    }
    logrus.Infof("received: %x", tmp[:n])
}
```
One important thing to keep in mind about TCP connections is that they;re streams. TCP messages don't carry information about their sizes. This forces use to use a buffer when reading from a TCP connection.