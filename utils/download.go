package utils

import (
	"errors"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/kkdai/youtube/v2" //YT API package
)

func DownloadAndConvert(input, outputDir string) error {
	// Extract the video ID
	fmt.Println("Extracting video information...")
	videoID, err := ExtractVideoID(input)
	if err != nil {
		return fmt.Errorf("failed to extract video ID: %w", err)
	}

	// Initialize YouTube API client
	client := youtube.Client{}

	// Fetch video metadata
	fmt.Println("Fetching video metadata...")
	video, err := client.GetVideo(videoID)
	if err != nil {
		return fmt.Errorf("failed to fetch video metadata: %w", err)
	}

	// Get audio formats available in the video
	formats := video.Formats.WithAudioChannels()
	if len(formats) == 0 {
		return fmt.Errorf("no audio formats found")
	}

	// Select the first available format with audio
	format := formats[0]

	// Clean the video title to use it as a valid file name
	cleanTitle := CleanFileName(video.Title)
	downloadedFile := filepath.Join(outputDir, cleanTitle+".tmp")
	mp3File := filepath.Join(outputDir, cleanTitle+".mp3")

	// Download the audio
	fmt.Println("Downloading audio stream...")
	stream, _, err := client.GetStream(video, &format)
	if err != nil {
		return fmt.Errorf("failed to get stream: %w", err)
	}
	defer stream.Close()

	// Create a file to store the downloaded audio
	out, err := os.Create(downloadedFile)
	if err != nil {
		return fmt.Errorf("failed to create output file: %w", err)
	}
	defer out.Close()

	// Copy the stream data to the file
	_, err = io.Copy(out, stream)
	if err != nil {
		return fmt.Errorf("failed to download audio: %w", err)
	}
	out.Close()

	// Convert the downloaded file to MP3 format
	fmt.Println("Converting to MP3...")
	err = ConvertToMp3(downloadedFile, mp3File)
	if err != nil {
		return fmt.Errorf("failed to convert to MP3: %w", err)
	}

	// Remove the temporary downloaded file
	err = os.Remove(downloadedFile)
	if err != nil {
		fmt.Printf("Warning: Could not remove temporary file: %v\n", err)
	}

	fmt.Printf("Successfully downloaded and converted to MP3: %s\n", mp3File)
	return nil
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
