* Convert an input media file to a different format, by re-encoding media streams:
```sh
ffmpeg -i input.avi output.mp4
```
* Set the video __bitrate__ of the output file to 64 kbit/s
```sh
ffmpeg -i input.avi -b:v 64k -bufsize 64k output.mp4
```
* Force the __frame rate __of the output file to 24 fps
```sh
ffmpeg 0 input.avi -r 24 output.mp4
```
* Force the frame rate of the input file (valid for raw formats only) to 1 fps and the frame rate of the output file to 24 fps
```sh
ffmpeg -r 1 -i input.m2v -r 24 output.mp4
```
* The simplest pipeline in *ffmpeg* is single-stream *streamcopy*, that is copying one input elementary stream's packets without decoding, filtering, or encoding them
```sh
ffmpeg -i INPUT.mkv -map 0:1 -c copy OUTPUT.mp4
```
* Combine streams from two input files into a single output
```sh
ffmpeg -i INPUT0.mkv -i INPUT1.acc -map 0:0 -map 1:0 -c copy OUTPUT.mp4
```
* Split multiple streams from a single input into multiple outputs:
```sh
ffmpeg -i INPUT.mkv -map 0:0 -c copy OUTPUT0.mp4 -map 0:1 -c copy OUTPUT1.mp4
```
* Pipeline that reads an input file with one audio and one video stream, transcodes the video and copies the audio into a single output file
```sh
ffmpeg -i INPUT.mkv -map 0:v -map 0:a -c:v libx264 -c:a copy OUTPUT.mp4
```
* create "loopback" decoders that decode the output from some encoder and allow it to be fed back to complex filtergraphs. 

*-dec* directive takes as a paramter the index of the output stream that should be decoded
```sh
ffmpeg -i INPUT
    -map 0:v:0 -c:v libx264 -crf 45 -f null -
    -threads 3 -dec 0:0
    -filter_complex '[0:v][dec:0]hstack[stack]'
    -map '[stack]' -c:v ffv1 OUTPUT
```
reads an input video and
    * (line 2) encodes it with *libx264* at low quality
    * (line 3) decodes this encoded stream using 3 threads
    * (line 4) places decoded video side by side with the original input video
    * (line 5) combined video is then losslessly encoded and written to OUTPUT
* Example
```sh
input file 'A.avi'
      stream 0: video 640x360
      stream 1: audio 2 channels

input file 'B.mp4'
      stream 0: video 1920x1080
      stream 1: audio 2 channels
      stream 2: subtitles (text)
      stream 3: audio 5.1 channels
      stream 4: subtitles (text)

input file 'C.mkv'
      stream 0: video 1280x720
      stream 1: audio 2 channels
      stream 2: subtitles (image)
```
* automaic stream selection
```sh
ffmpeg -i A.avi -i B.mp4 out1.mkv out2.wav -map 1:a -c:a copy out3.mov
```
__out1.mkv__ is a Matroska container file and accepts video, audio and subtitle streams, so ffmpeg will try to select one of each type. For video, it will select *stream 0* from B.mp4, which has *highest resolution* among all the input video streams. For audio, it will select *stream 3* from B.mp4, since it has *the greatest number of channels*. For subtitles, it will select *stream 2* from B.mp4, which is the *first subtitle stream* from among A.avi and B.mp4

__out2.wav__ accepts only audio streams, so only *stream 3* from B.mp4 is selected

__out3.mov__, since a *-map* option is set, no automatic stream selection will occure. The *-map 1:a* option will select all audio streams from the second input B.mp4. codec option for audio stream has been set to *copy*, so no decoding-filtering-encoding operations will occur
```sh
ffmpeg -i A.avi -i B.mp4 -i C.mkv -filter_comple "[1:v]hue=s=0,split=2[outv1][outv2];overlay;aresample" \
-map '[outv1]' -an out1.mp4 \
                  out2.mkv \
-map '[outv2]' -map 1:a:0 out3.mkv
```
The video stream from B.mp3 is sent to __the hue filter__, whose output is cloned once using __the split filter__, and both outputs __labelled__. The a copy each is mapped to the first and third output files

__The overlay filter__, requiring two video inputs, uses the first two unused video streams. The overlay output isn't labelled, so it is sent to the first output file out1.mp4, regardless of the presence

__The aresample filter__ is sent the first unused audio stream, that of A.avi. Since this filter output is also unlabelled, it too is mapped to the first output file out1.mp4

the video, audo and subtitle streams mapped to __out2.mkv are entirely determined by automatic__ stream selection

out3.mkv consists of the cloned video output from the hue filter and the first audio stream from B.mp4

```sh
stream_type[:additional_stream_specifier]
```
stream_type is one of following: ’v’ or ’V’ for video, ’a’ for audio, ’s’ for subtitle, ’d’ for data, and ’t’ for attachments. ’v’ matches all video streams, ’V’ only matches video streams which are not attached pictures, video thumbnails or cover arts.

* overlay an image over video
```sh
ffmpeg -i video.mkv -i image.png -filter_complex '[0:v][1:v]overlay[out]' -map '[out]' out.mkv
```



