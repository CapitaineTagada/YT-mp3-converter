# YouTube to MP3 Converter
[![Made with Go](https://img.shields.io/badge/Go-1-blue?logo=go&logoColor=white)](https://golang.org "Go to Go homepage")  
This is a command-line tool written in Go that downloads the audio from a YouTube video and converts it to an MP3 file. The program uses the [kkdai/youtube](https://github.com/kkdai/youtube) package to interact with YouTube, and [FFmpeg](https://ffmpeg.org/) to handle the conversion from video/audio formats to MP3.

## Features

- Downloads audio streams from YouTube videos.
- Converts downloaded audio streams to MP3 format using FFmpeg.
- Allows users to specify a YouTube URL or video ID.
- Allows users to specify the output directory for the MP3 file.
- Automatically handles invalid characters in filenames.

## Requirements

- Go (1.18 or later)
- [FFmpeg](https://ffmpeg.org/download.html) installed on your system and available in your PATH.
- Correct internet connection to fetch video streams from YouTube.

## Installation

1. Clone the repository:

   ```bash
   git clone https://github.com/CapitaineTagada/YT-mp3-converter.git
   cd YT-mp3-converter
   ```

2. Install dependencies using Go modules:

   ```bash
   go mod tidy
   ```

3. Install [FFmpeg](https://ffmpeg.org/download.html) if it's not already installed:

   - On Ubuntu :

     ```bash
     sudo apt update
     sudo apt install ffmpeg
     ```

   - On Windows:

     Download the FFmpeg executable and make sure to add it to your system's PATH.

4. Build the program:

   ```bash
   go build -o YT-mp3-converter
   ```

## Usage

### Command-Line Options

- `-url` (required): The YouTube video URL or video ID.
- `-output` (optional): The output directory for the MP3 file. If not specified, the current directory (`.`) is used by default.

### Example

To download and convert a YouTube video to MP3:

```bash
./YT-mp3-converter -url "https://www.youtube.com/watch?v=dQw4w9WgXcQ" -output "./audios"
```

This will:
1. Extract the audio stream from the given YouTube video.
2. Convert the audio to MP3 format.
3. Save the MP3 file in the `./audios` directory.

## File Naming

- The program automatically cleans the video title to create a valid file name for the MP3 file.
- Invalid characters like `<>:"/\|?*` are removed from the file name.
- Non-ASCII characters are removed to avoid issues with file systems that do not support them.
- If the title is too long (over 200 characters), it is truncated to a maximum of 200 characters.

## Error Handling

If any errors occur during the download or conversion process, the program will display a detailed error message. Common issues include:
- Failure to connect to YouTube.
- Invalid or missing YouTube video ID.
- Issues during conversion with FFmpeg.

## Dependencies

- [kkdai/youtube](https://github.com/kkdai/youtube) — Go library for accessing YouTube video data and streams.
- [FFmpeg](https://ffmpeg.org/) — A multimedia framework used for converting the downloaded audio stream to MP3.

## License

This project is licensed under the GPL 3.0 License.

