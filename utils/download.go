package utils

import (
	"errors"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"strings"
)

func DownloadAndConvert(input, outputDir string) error {
	// Extract the video ID
	fmt.Println("Extracting video information...")
	videoID, err := ExtractVideoID(input)
	if err != nil {
		return fmt.Errorf("failed to extract video ID: %w", err)
	}

	// Temp name of the file
	tmpFile := filepath.Join(outputDir, videoID+".tmp")
	mp3File := filepath.Join(outputDir, videoID+".mp3")

	// Download only the audio with yt-dlp in m4a format (without conversion)
	fmt.Println("Downloading audio stream with yt-dlp...")
	cmd := exec.Command("yt-dlp",
		"--no-playlist",
		"-f", "bestaudio",
		"-o", tmpFile,
		input,
	)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("yt-dlp download error: %w", err)
	}

	// Convert to mp3 with ffmpeg
	fmt.Println("Converting to MP3...")
	err = ConvertToMp3(tmpFile, mp3File)
	if err != nil {
		return fmt.Errorf("failed to convert to MP3: %w", err)
	}

	// Remove temp file
	if err := os.Remove(tmpFile); err != nil {
		fmt.Printf("Warning: Could not remove temporary file: %v\n", err)
	}

	fmt.Printf("Successfully downloaded and converted to MP3: %s\n", mp3File)
	return nil
}

func CleanYouTubeURL(url string) string {
	// Si on trouve un "&list=", on garde juste la partie avant
	if strings.Contains(url, "&list=") {
		parts := strings.Split(url, "&list=")
		url = parts[0]
	}
	// Supprime aussi "&start_radio=" et tout autre paramètre résiduel
	if strings.Contains(url, "&") {
		parts := strings.Split(url, "&")
		url = parts[0]
	}
	return url
}

func ExtractVideoID(input string) (string, error) {
	// Regular expression to match YouTube video URLs
	videoIDRegex := regexp.MustCompile(`((?:https?:)?\/\/)?((?:www|m)\.)?((?:youtube\.com|youtu.be))(\/(?:[\w\-]+\?v=|embed\/|v\/)?)([\w\-]+)(\S+)?`)
	matches := videoIDRegex.FindStringSubmatch(input)

	// Check if the regex match contains a valid video ID
	if len(matches) < 6 {
		return "", errors.New("invalid YouTube URL or ID")
	}

	// Return the extracted video ID
	return matches[5], nil
}

func CleanFileName(fileName string) string {
	// Replace spaces with underscores
	fileName = strings.ReplaceAll(fileName, " ", "_")

	// Remove or replace invalid characters
	fileName = regexp.MustCompile(`[<>:"/\\|?*]`).ReplaceAllString(fileName, "")

	// Remove any non-ASCII characters
	fileName = regexp.MustCompile(`[^\x00-\x7F]`).ReplaceAllString(fileName, "")

	// Trim to a reasonable length if needed
	if len(fileName) > 200 {
		fileName = fileName[:200]
	}

	// Ensure the filename is not empty
	if fileName == "" {
		fileName = "unnamed_audio"
	}

	return fileName
}
