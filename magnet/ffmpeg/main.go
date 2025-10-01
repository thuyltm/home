package main

import (
	"log"
	"os/exec"
)

func main() {
	cmd := exec.Command("ffplay", "magnet/video/sample.mp4")
	err := cmd.Run()
	if err != nil {
		log.Fatalf("ffplay command fail: %v", err)
	}

	cmd = exec.Command("ffmpeg", "-i", "magnet/video/sample.mp4", "-i",
		"magnet/video/beach.mp3", "-c:v", "copy", "-c:a", "aac",
		"-map", "0:v", "-map", "1:a", "magnet/video/result/output.mp4")
	err = cmd.Run()
	if err != nil {
		log.Fatalf("ffmpeq combine fail: %v", err)
	}

}
