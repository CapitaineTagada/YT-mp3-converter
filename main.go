package main

import (
	"YT-mp3-converter/utils"
	"flag"
	"fmt"
	"os"
)

func main() {
	// Define command-line flags
	urlFlag := flag.String("url", "", "YouTube video URL or ID")
	outputDirFlag := flag.String("output", ".", "Output directory for MP3 files")
	flag.Parse()

	// Validate input
	if *urlFlag == "" {
		fmt.Println("Please provide a YouTube URL or video ID using the -url flag")
		flag.PrintDefaults()
		os.Exit(1)
	}

	// Ensure output directory exists or create it if it doesn't
	err := os.MkdirAll(*outputDirFlag, os.ModePerm)
	if err != nil {
		fmt.Printf("Error creating output directory: %v\n", err)
		os.Exit(1)
	}

	// Call the download and conversion function
	err = utils.DownloadAndConvert(*urlFlag, *outputDirFlag)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(1)
	}

	fmt.Println("Process completed yipeeee!")
}
