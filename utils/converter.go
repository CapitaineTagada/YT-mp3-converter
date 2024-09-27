package utils

import (
	"bytes"
	"fmt"
	"os/exec"
)

func ConvertToMp3(inputFile, outputFile string) error {
	// FFmpeg command to convert audio to MP3
	cmd := exec.Command("ffmpeg",
		"-i", inputFile,
		"-acodec", "libmp3lame",
		"-b:a", "128k",
		"-sample_fmt", "s16p",
		outputFile)

	// Capture the output
	var out bytes.Buffer
	var stderr bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &stderr

	// Run the FFmpeg command
	err := cmd.Run()
	if err != nil {
		// Return an error with details if the command fails
		return fmt.Errorf("FFmpeg error: %v\nDetails: %s", err, stderr.String())
	}

	return nil
}
