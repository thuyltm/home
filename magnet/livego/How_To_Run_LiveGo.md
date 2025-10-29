## Prepare
### Generate server.key, server.crt files
* Generate the private key
```sh
openssl genrsa -out server.key 2048
```
* Generate a certificate signing request
```sh
openssl req -new -key server.key -out server.csr
```
* Generate the self-signed certifacte
```sh
openssl x509 -req -days 365 -in server.csr -signkey server.key -out server.crt
```
### Get *appName* in config/live.yaml
```sh
server:
- appname: live
```
### Get file info
```sh
ffprobe -v error -show_format -show_streams -print_format json magnet/video/demo.flv
```

## Run
* GET channelKey
```sh
curl http://localhost:8090/control/get?room=movie
```



* Upstream
```sh
ffmpeg -re -i magnet/video/demo.flv -c copy -f flv rtmps://localhost:1935/live/{channelkey}
```
* Downstream

    RTMP:rtmp://localhost:1935/live/movie
    
    FLV:http://127.0.0.1:7001/live/movie.flv
    
    HLS:http://127.0.0.1:7002/live/movie.m3u8

```sh
ffplay rtmps://localhost:1935/live/movie

ffplay https://127.0.0.1:7002/live/movie.m3u8

ffplay http://127.0.0.1:7001/live/movie.flv
```
