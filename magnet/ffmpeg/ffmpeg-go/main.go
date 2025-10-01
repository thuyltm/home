package main

import (
	"bytes"
	"fmt"
	"image"
	"image/jpeg"
	"io"
	"log"
	"os"

	ffmpeg "github.com/u2takey/ffmpeg-go"
)

func concatAndOverlay() {
	split := ffmpeg.Input("magnet/video/sample.mp4").VFlip().Split()
	split0, split1 := split.Get("0"), split.Get("1")
	overlayFile := ffmpeg.Input("magnet/video/earth.mp4").Crop(10, 10, 158, 112)
	err := ffmpeg.Concat([]*ffmpeg.Stream{
		split0.Trim(ffmpeg.KwArgs{"start_frame": 10, "end_frame": 20}),
		split1.Trim(ffmpeg.KwArgs{"start_frame": 30, "end_frame": 40}),
	}).
		Overlay(overlayFile.HFlip(), "").
		DrawBox(50, 50, 120, 120, "red", 5).
		Output("magnet/video/result/concat.mp4").
		OverWriteOutput().
		Run()
	if err != nil {
		log.Fatal(err)
	}
}

func transcode() {
	err := ffmpeg.Input("magnet/video/sample.mp4").
		Output("magnet/video/result/transcode.mp4", ffmpeg.KwArgs{"c:v": "libx265"}).
		OverWriteOutput().ErrorToStdOut().Run()
	if err != nil {
		log.Fatal(err)
	}
}

func watermark() {
	overlay := ffmpeg.Input("magnet/images/capture_0.jpg").Filter("scale", ffmpeg.Args{"64:-1"})
	err := ffmpeg.Filter(
		[]*ffmpeg.Stream{
			ffmpeg.Input("magnet/video/sample.mp4"),
			overlay,
		}, "overlay", ffmpeg.Args{"10:10"}, ffmpeg.KwArgs{"enable": "gte(t,1)"}).
		Output("magnet/video/result/watermark.mp4").OverWriteOutput().ErrorToStdOut().Run()
	if err != nil {
		log.Fatal(err)
	}
}

func gif() {
	err := ffmpeg.Input("magnet/video/sample.mp4", ffmpeg.KwArgs{"ss": "1"}).
		Output("magnet/video/result/gif.gif",
			ffmpeg.KwArgs{"s": "320x240", "pix_fmt": "rgb24", "t": "3", "r": "3"}).
		OverWriteOutput().ErrorToStdOut().Run()
	if err != nil {
		log.Fatal(err)
	}
}

func ExampleReadFrameAsJpeg(inFileName string, frameNum int) io.Reader {
	buf := bytes.NewBuffer(nil)
	err := ffmpeg.Input(inFileName).
		Filter("select", ffmpeg.Args{fmt.Sprintf("gte(n,%d)", frameNum)}).
		Output("pipe:", ffmpeg.KwArgs{"vframes": 1, "format": "jpeg_pipe", "vcodec": "mjpeg"}).
		WithOutput(buf, os.Stdout).
		Run()
	if err != nil {
		panic(err)
	}
	return buf
}

func readFrameAsJpeg() {
	reader := ExampleReadFrameAsJpeg("magnet/video/sample.mp4", 5)
	img, _, err := image.Decode(reader)
	if err != nil {
		log.Fatal(err)
	}
	jpegFile, err := os.Create("magnet/video/result/frame.jpeg")
	if err != nil {
		log.Fatal(err)
	}
	defer jpegFile.Close()
	if err := jpeg.Encode(jpegFile, img, nil); err != nil {
		log.Fatal(err)
	}
}

func mergeOutputs() {
	input := ffmpeg.Input("magnet/video/sample.mp4").Split()
	out1 := input.Get("0").Filter("scale", ffmpeg.Args{"1920:-1"}).
		Output("magnet/video/result/1920.mp4", ffmpeg.KwArgs{"b:v": "5000k"})
	out2 := input.Get("1").Filter("scale", ffmpeg.Args{"1280:-1"}).
		Output("magnet/video/result/1280.mp4", ffmpeg.KwArgs{"b:v": "2800k"})
	err := ffmpeg.MergeOutputs(out1, out2).OverWriteOutput().ErrorToStdOut().Run()
	if err != nil {
		log.Fatal(err)
	}
}

func getVideoInfo(fileName string) {
	data, err := ffmpeg.Probe(fileName)
	if err != nil {
		panic(err)
	}
	log.Println("got video info", data)
}

func main() {
	//concatAndOverlay()
	//transcode()
	//watermark()
	//gif()
	//readFrameAsJpeg()
	//mergeOutputs()
	//ExampleShowProgress("magnet/video/result/1280.mp4", "magnet/video/result/1920.mp4")
	//RunExampleStream("magnet/video/sample.mp4", "magnet/video/result/output.mp4")
	getVideoInfo("magnet/video/sample.mp4")
}
