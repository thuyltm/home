package main

import (
	"fmt"
	"image"
	"image/jpeg"
	"os"

	vidio "github.com/AlexEidt/Vidio"
)

func writeTheFirstAndLastFramesOfVdieoAsImages() {
	video, _ := vidio.NewVideo("magnet/video/sample.mp4")
	frames, _ := video.ReadFrames(0, video.Frames()-1)
	for index, frame := range frames {
		f, _ := os.Create(fmt.Sprintf("magnet/video/result/%d.jpg", index))
		jpeg.Encode(f, frame, nil)
		f.Close()
	}
}

func writeAllFramesOfVideoAsImages() {
	video, _ := vidio.NewVideo("magnet/video/sample.mp4")
	img := image.NewRGBA(image.Rect(0, 0, video.Width(), video.Height()))
	video.SetFrameBuffer(img.Pix)
	frame := 0
	for video.Read() {
		f, _ := os.Create(fmt.Sprintf("magnet/video/result/frame_%d.jpg", frame))
		jpeg.Encode(f, img, nil)
		f.Close()
		frame++
	}
}

func createGifFromSeriesOfPngFiles() {
	w, h, _, _ := vidio.Read("magnet/images/capture_0.jpg")
	options := vidio.Options{FPS: 1, Loop: 0, Delay: 1000}
	gif, _ := vidio.NewVideoWriter("magnet/images/result/output.gif", w, h, &options)
	defer gif.Close()
	for i := 1; i <= 10; i++ {
		_, _, img, _ := vidio.Read(fmt.Sprintf("magnet/video/result/gif_%d.png", i))
		gif.Write(img)
	}
}

func read10FramesOfWebcamStream() {
	webcam, _ := vidio.NewCamera(0)
	defer webcam.Close()
	options := vidio.Options{FPS: webcam.FPS()}
	writer, _ := vidio.NewVideoWriter("magnet/video/result/webcam.mp4", webcam.Width(), webcam.Height(), &options)
	defer writer.Close()
	count := 0
	for webcam.Read() && count < 10 {
		writer.Write(webcam.FrameBuffer())
		count++
	}
}

func streamInputVideoToAnotherVideo() {
	video, _ := vidio.NewVideo("magnet/video/sample.mp4")
	options := vidio.Options{
		FPS:     video.FPS(),
		Bitrate: video.Bitrate(),
	}
	if video.HasStreams() {
		options.StreamFile = video.FileName()
	}
	writer, _ := vidio.NewVideoWriter("magnet/video/result/stream.mp4", video.Width(), video.Height(), &options)
	defer writer.Close()
	for video.Read() {
		writer.Write(video.FrameBuffer())
	}
}

func main() {
	writeTheFirstAndLastFramesOfVdieoAsImages()
	writeAllFramesOfVideoAsImages()
	createGifFromSeriesOfPngFiles()
	read10FramesOfWebcamStream()
	streamInputVideoToAnotherVideo()
}
