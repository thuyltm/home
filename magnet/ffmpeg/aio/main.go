package main

import (
	"fmt"

	"github.com/AlexEidt/aio"
)

func streamWavToMp3() {
	audio, _ := aio.NewAudio("magnet/video/example.wav", nil)
	options := aio.Options{
		SampleRate: audio.SampleRate(),
		Channels:   audio.Channels(),
		Bitrate:    audio.Bitrate(),
		Format:     audio.Format(),
	}
	writer, _ := aio.NewAudioWriter("magnet/video/result/stream.mp3", &options)
	defer writer.Close()
	for audio.Read() {
		writer.Write(audio.Buffer())
	}
}

func capture10SecondsOfAudioFromMicrophone() {
	micOptions := aio.Options{Format: "s16", Channels: 2, SampleRate: 44100}
	mic, _ := aio.NewMicrophone(0, &micOptions)
	defer mic.Close()
	writerOptions := aio.Options{
		SampleRate: mic.SampleRate(),
		Channels:   mic.Channels(),
		Format:     mic.Format(),
	}
	writer, _ := aio.NewAudioWriter("magnet/video/result/capture.wav", &writerOptions)
	seconds := 0
	for mic.Read() && seconds < 10 {
		writer.Write(mic.Buffer())
		seconds++
	}
}

func stream() {
	streams, _ := aio.NewAudioStreams("magnet/video/sample.mp4", nil)
	for _, stream := range streams {
		player, _ := aio.NewPlayer(stream.Channels(), stream.SampleRate(), stream.Format())
		for stream.Read() {
			player.Play(stream.Buffer())
		}
		player.Close()
	}
}
func play() {
	audio, _ := aio.NewAudio("magnet/video/sample.mp4", nil)
	player, _ := aio.NewPlayer(audio.Channels(), audio.SampleRate(), audio.Format())
	defer player.Close()
	for audio.Read() {
		player.Play(audio.Buffer())
	}
}
func sample() {
	audio, _ := aio.NewAudio("magnet/video/example.wav", nil)
	for audio.Read() {
		samples := audio.Samples().([]int16)
		for i := range samples {
			fmt.Println(samples[i])
		}
	}
}
func combine() {
	audio, _ := aio.NewAudio("magnet/video/example.wav", nil)
	options := aio.Options{
		SampleRate: audio.SampleRate(),
		Channels:   audio.Channels(),
		Bitrate:    audio.Bitrate(),
		Format:     audio.Format(),
		Codec:      "aac",
		StreamFile: "magnet/video/beach.mp3",
	}
	writer, _ := aio.NewAudioWriter("magnet/video/result/combine.mp4", &options)
	defer writer.Close()
	for audio.Read() {
		writer.Write(audio.Buffer)
	}
}
func playMicrophone() {
	mic, _ := aio.NewMicrophone(0, nil)
	defer mic.Close()
	player, _ := aio.NewPlayer(mic.Channels(), mic.SampleRate(), mic.Format())
	defer player.Close()
	for mic.Read() {
		player.Play(mic.Buffer)
	}
}
func main() {
	streamWavToMp3()
	capture10SecondsOfAudioFromMicrophone()
	stream()
	play()
	sample()
	combine()
	playMicrophone()
}
