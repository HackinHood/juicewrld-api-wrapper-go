# Juice WRLD API Wrapper (Go)

[![Go 1.22+](https://img.shields.io/badge/go-1.22+-blue.svg)](https://golang.org/dl/)
[![Go Report Card](https://goreportcard.com/badge/github.com/hackinhood/juicewrld-api-wrapper-go)](https://goreportcard.com/report/github.com/hackinhood/juicewrld-api-wrapper-go)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)

A comprehensive Go wrapper for the Juice WRLD API, providing easy access to Juice WRLD's complete discography including released tracks, unreleased songs, recording sessions, and unsurfaced content.

## Features

- **Full API Coverage**: Access to all Juice WRLD API endpoints
- **Type Safety**: Strong typing with Go structs and interfaces
- **Error Handling**: Comprehensive error types and handling
- **Context Support**: Full context.Context integration for timeouts and cancellation
- **Search & Filtering**: Advanced song search across all categories
- **File Operations**: Browse, download, and manage audio files
- **Audio Streaming**: Modern streaming support with range requests
- **ZIP Operations**: Create and manage ZIP archives
- **Zero Dependencies**: Uses only Go standard library


## Requirements

- **Go**: 1.22+ (for generics support)
- **Dependencies**: None (uses only Go standard library)

## Installation

### Quick Install (Recommended)
```bash
go get github.com/hackinhood/juicewrld-api-wrapper-go
```

### Alternative Installation Methods

**From GitHub Repository (Latest Version)**
```bash
go get github.com/hackinhood/juicewrld-api-wrapper-go@latest
```

**From Source**
```bash
git clone https://github.com/hackinhood/juicewrld-api-wrapper-go.git
cd juicewrld-api-wrapper/go
go mod tidy
```

## Quick Start

```go
package main

import (
    "context"
    "fmt"
    "log"
    
    jw "github.com/hackinhood/juicewrld-api-wrapper-go"
)

func main() {
    // Create a new client
    client := jw.New()
    defer client.CloseIdleConnections()
    
    ctx := context.Background()
    
    // Get API overview
    overview, err := client.GetAPIOverview(ctx)
    if err != nil {
        log.Fatal(err)
    }
    
    fmt.Printf("API Version: %s\n", overview.APIVersion)
    fmt.Printf("Go Wrapper Version: %s\n", overview.GoWrapperVersion)
    
    // Get all artists
    artists, err := client.GetArtists(ctx)
    if err != nil {
        log.Fatal(err)
    }
    
    fmt.Printf("Found %d artists\n", len(artists))
    for _, artist := range artists {
        fmt.Printf("- %s (%d songs)\n", artist.Name, artist.SongCount)
    }
}
```

## API Reference

### Client Methods

#### Core Information
- `GetAPIOverview(ctx)` - Get API and wrapper version information
- `GetArtists(ctx)` - Get all available artists
- `GetStats(ctx)` - Get API statistics

#### Albums & Songs
- `GetAlbum(ctx, albumID)` - Get album details by ID
- `GetSongs(ctx, page, limit)` - Get paginated songs list
- `GetJuiceWRLDSongs(ctx, page, limit)` - Get Juice WRLD songs specifically
- `SearchSongs(ctx, query, page, limit)` - Search songs by query
- `GetSongsByCategory(ctx, category, page, limit)` - Get songs by category

#### Eras & Categories
- `GetEras(ctx)` - Get all available eras
- `GetCategories(ctx)` - Get all song categories

#### File Operations
- `BrowseFiles(ctx, path, search)` - Browse file system
- `GetFileInfo(ctx, filePath)` - Get file information
- `StreamAudioFile(ctx, filePath)` - Get streaming URL for audio file
- `DownloadFile(ctx, filePath)` - Download file as bytes
- `DownloadFileTo(ctx, filePath, savePath)` - Download file to disk
- `GetCoverArt(ctx, filePath)` - Extract cover art from file

#### ZIP Operations
- `CreateZip(ctx, filePaths)` - Create ZIP archive
- `StartZipJob(ctx, filePaths)` - Start ZIP creation job
- `GetZipJobStatus(ctx, jobID)` - Check ZIP job status
- `CancelZipJob(ctx, jobID)` - Cancel ZIP job

#### Player
- `PlayJuiceWRLDSong(ctx, songID)` - Get playable URL for song

### Data Models

#### Artist
```go
type Artist struct {
    ID        int    `json:"id"`
    Name      string `json:"name"`
    SongCount int    `json:"song_count"`
}
```

#### Album
```go
type Album struct {
    ID          int         `json:"id"`
    Name        string      `json:"name"`
    ReleaseDate FlexibleTime `json:"release_date"`
    SongCount   int         `json:"song_count"`
    Songs       []Song      `json:"songs"`
}
```

#### Song
```go
type Song struct {
    ID          int         `json:"id"`
    Title       string      `json:"title"`
    Artist      string      `json:"artist"`
    Album       string      `json:"album"`
    Duration    int         `json:"duration"`
    PublicID    interface{} `json:"public_id"`
    Category    string      `json:"category"`
    Era         string      `json:"era"`
    ReleaseDate FlexibleTime `json:"release_date"`
}
```

#### FileInfo
```go
type FileInfo struct {
    Name      string `json:"name"`
    Path      string `json:"path"`
    Size      int64  `json:"size"`
    SizeHuman string `json:"size_human"`
    Type      string `json:"type"`
    Modified  FlexibleTime `json:"modified"`
}
```

### Error Types

The wrapper provides specific error types for different scenarios:

- `APIError` - General API errors
- `RateLimitError` - Rate limiting errors
- `NotFoundError` - Resource not found
- `AuthenticationError` - Authentication issues
- `ValidationError` - Input validation errors

```go
if err != nil {
    switch e := err.(type) {
    case *jw.RateLimitError:
        fmt.Printf("Rate limited: %s\n", e.Message)
    case *jw.NotFoundError:
        fmt.Printf("Not found: %s\n", e.Message)
    case *jw.APIError:
        fmt.Printf("API error %d: %s\n", e.StatusCode, e.Message)
    default:
        fmt.Printf("Error: %v\n", err)
    }
}
```


## Examples

### Basic Usage
```go
package main

import (
    "context"
    "fmt"
    "log"
    
    jw "github.com/hackinhood/juicewrld-api-wrapper-go"
)

func main() {
    client := jw.New()
    defer client.CloseIdleConnections()
    
    ctx := context.Background()
    
    // Get API overview
    overview, err := client.GetAPIOverview(ctx)
    if err != nil {
        log.Fatal(err)
    }
    fmt.Printf("API Version: %s\n", overview.APIVersion)
}
```

### Search and Download
```go
// Search for songs
results, err := client.SearchSongs(ctx, "lucid dreams", 1, 10)
if err != nil {
    log.Fatal(err)
}

for _, song := range results.Results {
    fmt.Printf("Found: %s by %s\n", song.Title, song.Artist)
}

// Download a file
data, err := client.DownloadFile(ctx, "path/to/song.mp3")
if err != nil {
    log.Fatal(err)
}
fmt.Printf("Downloaded %d bytes\n", len(data))

// Save to file
err = client.DownloadFileTo(ctx, "path/to/song.mp3", "local_song.mp3")
if err != nil {
    log.Fatal(err)
}
fmt.Println("File saved to local_song.mp3")
```

### File Operations
```go
// Browse files
dir, err := client.BrowseFiles(ctx, "Compilation", nil)
if err != nil {
    log.Fatal(err)
}

fmt.Printf("Found %d files in %s\n", dir.TotalFiles, dir.CurrentPath)
for _, item := range dir.Items {
    fmt.Printf("- %s (%s)\n", item.Name, item.Type)
}

// Get file info
info, err := client.GetFileInfo(ctx, "path/to/file.mp3")
if err != nil {
    log.Fatal(err)
}
fmt.Printf("File: %s, Size: %s, Type: %s\n", info.Name, info.SizeHuman, info.Type)

// Stream audio
stream, err := client.StreamAudioFile(ctx, "path/to/file.mp3")
if err != nil {
    log.Fatal(err)
}
fmt.Printf("Stream URL: %s\n", stream["stream_url"])
```

### ZIP Operations
```go
// Create ZIP
filePaths := []string{"path1", "path2", "path3"}
zipData, err := client.CreateZip(ctx, filePaths)
if err != nil {
    log.Fatal(err)
}
fmt.Printf("Created ZIP: %d bytes\n", len(zipData))

// Start ZIP job (for large files)
job, err := client.StartZipJob(ctx, filePaths)
if err != nil {
    log.Fatal(err)
}
fmt.Printf("Started ZIP job: %s\n", job["job_id"])

// Check job status
status, err := client.GetZipJobStatus(ctx, job["job_id"].(string))
if err != nil {
    log.Fatal(err)
}
fmt.Printf("Job status: %s\n", status["status"])
```

## Error Handling

The wrapper provides comprehensive error handling with specific error types:

```go
result, err := client.GetAlbum(ctx, 999)
if err != nil {
    switch e := err.(type) {
    case *jw.NotFoundError:
        fmt.Println("Album not found")
    case *jw.RateLimitError:
        fmt.Printf("Rate limited: %s\n", e.Message)
    case *jw.APIError:
        fmt.Printf("API error %d: %s\n", e.StatusCode, e.Message)
    default:
        fmt.Printf("Unexpected error: %v\n", err)
    }
    return
}
```

## Context Usage

All methods accept a `context.Context` for cancellation and timeouts:

```go
// With timeout
ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
defer cancel()

// With cancellation
ctx, cancel := context.WithCancel(context.Background())
defer cancel()

// Use context
result, err := client.GetArtists(ctx)
```

## Performance

- **Zero Dependencies**: Uses only Go standard library
- **Connection Pooling**: Automatic HTTP connection reuse
- **Context Support**: Full cancellation and timeout support
- **Memory Efficient**: Streaming support for large files
- **Type Safe**: Compile-time type checking

## Contributing

1. Fork the repository
2. Create a feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## Related Projects

- [Python Wrapper](https://github.com/hackinhood/juicewrld-api-wrapper) - Original Python implementation
- [Juice WRLD API](https://juicewrldapi.com) - Official API documentation

## Support

- **Issues**: [GitHub Issues](https://github.com/hackinhood/juicewrld-api-wrapper-go/issues)
- **Discussions**: [GitHub Discussions](https://github.com/hackinhood/juicewrld-api-wrapper-go/discussions)
- **API Documentation**: [juicewrldapi.com](https://juicewrldapi.com)

---

**Made with ❤️ for the Juice WRLD community**

