# API Reference

This document provides detailed information about all available methods and types in the Juice WRLD API Wrapper (Go).

## Client

### New

```go
func New() *Client
```

Creates a new API client with default settings.

**Returns:**
- `*Client` - A new client instance

**Example:**
```go
client := jw.New()
defer client.CloseIdleConnections()
```

### CloseIdleConnections

```go
func (c *Client) CloseIdleConnections()
```

Closes idle HTTP connections. Should be called when the client is no longer needed.

## Core Information Methods

### GetAPIOverview

```go
func (c *Client) GetAPIOverview(ctx context.Context) (map[string]interface{}, error)
```

Gets API and wrapper version information.

**Parameters:**
- `ctx` - Context for cancellation and timeouts

**Returns:**
- `map[string]interface{}` - API overview data including versions
- `error` - Any error that occurred

**Example:**
```go
overview, err := client.GetAPIOverview(ctx)
if err != nil {
    log.Fatal(err)
}
fmt.Printf("API Version: %v\n", overview["api_version"])
```

### GetArtists

```go
func (c *Client) GetArtists(ctx context.Context) ([]Artist, error)
```

Gets all available artists.

**Parameters:**
- `ctx` - Context for cancellation and timeouts

**Returns:**
- `[]Artist` - Slice of artist information
- `error` - Any error that occurred

**Example:**
```go
artists, err := client.GetArtists(ctx)
if err != nil {
    log.Fatal(err)
}
for _, artist := range artists {
    fmt.Printf("%s: %d songs\n", artist.Name, artist.SongCount)
}
```

### GetStats

```go
func (c *Client) GetStats(ctx context.Context) (Stats, error)
```

Gets API statistics.

**Parameters:**
- `ctx` - Context for cancellation and timeouts

**Returns:**
- `Stats` - Statistics data
- `error` - Any error that occurred

## Albums & Songs Methods

### GetAlbum

```go
func (c *Client) GetAlbum(ctx context.Context, albumID int) (Album, error)
```

Gets album details by ID.

**Parameters:**
- `ctx` - Context for cancellation and timeouts
- `albumID` - The album ID to retrieve

**Returns:**
- `Album` - Album information including songs
- `error` - Any error that occurred

**Example:**
```go
album, err := client.GetAlbum(ctx, 1)
if err != nil {
    log.Fatal(err)
}
fmt.Printf("Album: %s (%d songs)\n", album.Name, album.SongCount)
```

### GetSongs

```go
func (c *Client) GetSongs(ctx context.Context, page, limit int) (PaginatedSongsResponse, error)
```

Gets paginated songs list.

**Parameters:**
- `ctx` - Context for cancellation and timeouts
- `page` - Page number (1-based)
- `limit` - Number of songs per page

**Returns:**
- `PaginatedSongsResponse` - Paginated songs data
- `error` - Any error that occurred

**Example:**
```go
songs, err := client.GetSongs(ctx, 1, 10)
if err != nil {
    log.Fatal(err)
}
fmt.Printf("Found %d songs (page %d of %d)\n", 
    len(songs.Results), songs.CurrentPage, songs.TotalPages)
```

### GetJuiceWRLDSongs

```go
func (c *Client) GetJuiceWRLDSongs(ctx context.Context, page, limit int) (PaginatedSongsResponse, error)
```

Gets Juice WRLD songs specifically.

**Parameters:**
- `ctx` - Context for cancellation and timeouts
- `page` - Page number (1-based)
- `limit` - Number of songs per page

**Returns:**
- `PaginatedSongsResponse` - Paginated Juice WRLD songs data
- `error` - Any error that occurred

### SearchSongs

```go
func (c *Client) SearchSongs(ctx context.Context, query string, page, limit int) (PaginatedSongsResponse, error)
```

Searches songs by query.

**Parameters:**
- `ctx` - Context for cancellation and timeouts
- `query` - Search query string
- `page` - Page number (1-based)
- `limit` - Number of songs per page

**Returns:**
- `PaginatedSongsResponse` - Search results
- `error` - Any error that occurred

**Example:**
```go
results, err := client.SearchSongs(ctx, "lucid dreams", 1, 10)
if err != nil {
    log.Fatal(err)
}
for _, song := range results.Results {
    fmt.Printf("Found: %s by %s\n", song.Title, song.Artist)
}
```

### GetSongsByCategory

```go
func (c *Client) GetSongsByCategory(ctx context.Context, category string, page, limit int) (PaginatedSongsResponse, error)
```

Gets songs by category.

**Parameters:**
- `ctx` - Context for cancellation and timeouts
- `category` - Category name
- `page` - Page number (1-based)
- `limit` - Number of songs per page

**Returns:**
- `PaginatedSongsResponse` - Songs in the specified category
- `error` - Any error that occurred

## Eras & Categories Methods

### GetEras

```go
func (c *Client) GetEras(ctx context.Context) ([]Era, error)
```

Gets all available eras.

**Parameters:**
- `ctx` - Context for cancellation and timeouts

**Returns:**
- `[]Era` - Slice of era information
- `error` - Any error that occurred

### GetCategories

```go
func (c *Client) GetCategories(ctx context.Context) ([]map[string]interface{}, error)
```

Gets all song categories.

**Parameters:**
- `ctx` - Context for cancellation and timeouts

**Returns:**
- `[]map[string]interface{}` - Slice of category information
- `error` - Any error that occurred

## File Operations Methods

### BrowseFiles

```go
func (c *Client) BrowseFiles(ctx context.Context, path string, search *string) (DirectoryInfo, error)
```

Browses the file system.

**Parameters:**
- `ctx` - Context for cancellation and timeouts
- `path` - Directory path to browse (empty for root)
- `search` - Optional search term (nil for no search)

**Returns:**
- `DirectoryInfo` - Directory contents and metadata
- `error` - Any error that occurred

**Example:**
```go
dir, err := client.BrowseFiles(ctx, "Compilation", nil)
if err != nil {
    log.Fatal(err)
}
fmt.Printf("Found %d files in %s\n", dir.TotalFiles, dir.CurrentPath)
```

### GetFileInfo

```go
func (c *Client) GetFileInfo(ctx context.Context, filePath string) (FileInfo, error)
```

Gets file information.

**Parameters:**
- `ctx` - Context for cancellation and timeouts
- `filePath` - Path to the file

**Returns:**
- `FileInfo` - File metadata
- `error` - Any error that occurred

**Example:**
```go
info, err := client.GetFileInfo(ctx, "path/to/file.mp3")
if err != nil {
    log.Fatal(err)
}
fmt.Printf("File: %s, Size: %s\n", info.Name, info.SizeHuman)
```

### StreamAudioFile

```go
func (c *Client) StreamAudioFile(ctx context.Context, filePath string) (map[string]interface{}, error)
```

Gets streaming URL for audio file.

**Parameters:**
- `ctx` - Context for cancellation and timeouts
- `filePath` - Path to the audio file

**Returns:**
- `map[string]interface{}` - Streaming information including URL
- `error` - Any error that occurred

**Example:**
```go
stream, err := client.StreamAudioFile(ctx, "path/to/song.mp3")
if err != nil {
    log.Fatal(err)
}
fmt.Printf("Stream URL: %v\n", stream["stream_url"])
```

### DownloadFile

```go
func (c *Client) DownloadFile(ctx context.Context, filePath string) ([]byte, error)
```

Downloads file as bytes.

**Parameters:**
- `ctx` - Context for cancellation and timeouts
- `filePath` - Path to the file to download

**Returns:**
- `[]byte` - File contents
- `error` - Any error that occurred

**Example:**
```go
data, err := client.DownloadFile(ctx, "path/to/song.mp3")
if err != nil {
    log.Fatal(err)
}
fmt.Printf("Downloaded %d bytes\n", len(data))
```

### DownloadFileTo

```go
func (c *Client) DownloadFileTo(ctx context.Context, filePath, savePath string) (string, error)
```

Downloads file to disk.

**Parameters:**
- `ctx` - Context for cancellation and timeouts
- `filePath` - Path to the file to download
- `savePath` - Local path to save the file

**Returns:**
- `string` - Path where file was saved
- `error` - Any error that occurred

**Example:**
```go
savedPath, err := client.DownloadFileTo(ctx, "path/to/song.mp3", "local_song.mp3")
if err != nil {
    log.Fatal(err)
}
fmt.Printf("File saved to: %s\n", savedPath)
```

### GetCoverArt

```go
func (c *Client) GetCoverArt(ctx context.Context, filePath string) ([]byte, error)
```

Extracts cover art from file.

**Parameters:**
- `ctx` - Context for cancellation and timeouts
- `filePath` - Path to the audio file

**Returns:**
- `[]byte` - Cover art image data
- `error` - Any error that occurred

**Example:**
```go
coverArt, err := client.GetCoverArt(ctx, "path/to/song.mp3")
if err != nil {
    log.Fatal(err)
}
fmt.Printf("Cover art size: %d bytes\n", len(coverArt))
```

## ZIP Operations Methods

### CreateZip

```go
func (c *Client) CreateZip(ctx context.Context, filePaths []string) ([]byte, error)
```

Creates ZIP archive.

**Parameters:**
- `ctx` - Context for cancellation and timeouts
- `filePaths` - Slice of file paths to include in ZIP

**Returns:**
- `[]byte` - ZIP file contents
- `error` - Any error that occurred

**Example:**
```go
filePaths := []string{"path1", "path2", "path3"}
zipData, err := client.CreateZip(ctx, filePaths)
if err != nil {
    log.Fatal(err)
}
fmt.Printf("Created ZIP: %d bytes\n", len(zipData))
```

### StartZipJob

```go
func (c *Client) StartZipJob(ctx context.Context, filePaths []string) (map[string]interface{}, error)
```

Starts ZIP creation job (for large files).

**Parameters:**
- `ctx` - Context for cancellation and timeouts
- `filePaths` - Slice of file paths to include in ZIP

**Returns:**
- `map[string]interface{}` - Job information including job ID
- `error` - Any error that occurred

### GetZipJobStatus

```go
func (c *Client) GetZipJobStatus(ctx context.Context, jobID string) (map[string]interface{}, error)
```

Checks ZIP job status.

**Parameters:**
- `ctx` - Context for cancellation and timeouts
- `jobID` - Job ID returned by StartZipJob

**Returns:**
- `map[string]interface{}` - Job status information
- `error` - Any error that occurred

### CancelZipJob

```go
func (c *Client) CancelZipJob(ctx context.Context, jobID string) (map[string]interface{}, error)
```

Cancels ZIP job.

**Parameters:**
- `ctx` - Context for cancellation and timeouts
- `jobID` - Job ID to cancel

**Returns:**
- `map[string]interface{}` - Cancellation result
- `error` - Any error that occurred

## Player Methods

### PlayJuiceWRLDSong

```go
func (c *Client) PlayJuiceWRLDSong(ctx context.Context, songID int) (map[string]interface{}, error)
```

Gets playable URL for song.

**Parameters:**
- `ctx` - Context for cancellation and timeouts
- `songID` - Song ID to play

**Returns:**
- `map[string]interface{}` - Play information including URL
- `error` - Any error that occurred

## Data Types

### Artist

```go
type Artist struct {
    ID        int    `json:"id"`
    Name      string `json:"name"`
    SongCount int    `json:"song_count"`
}
```

### Album

```go
type Album struct {
    ID          int         `json:"id"`
    Name        string      `json:"name"`
    ReleaseDate FlexibleTime `json:"release_date"`
    SongCount   int         `json:"song_count"`
    Songs       []Song      `json:"songs"`
}
```

### Song

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

### FileInfo

```go
type FileInfo struct {
    Name      string       `json:"name"`
    Path      string       `json:"path"`
    Size      int64        `json:"size"`
    SizeHuman string       `json:"size_human"`
    Type      string       `json:"type"`
    Modified  FlexibleTime `json:"modified"`
}
```

### DirectoryInfo

```go
type DirectoryInfo struct {
    CurrentPath      string     `json:"current_path"`
    TotalFiles       int        `json:"total_files"`
    TotalDirectories int        `json:"total_directories"`
    Items            []FileItem `json:"items"`
}
```

### FileItem

```go
type FileItem struct {
    Name string `json:"name"`
    Type string `json:"type"`
}
```

### Era

```go
type Era struct {
    ID   int    `json:"id"`
    Name string `json:"name"`
}
```

### Stats

```go
type Stats struct {
    TotalSongs    int `json:"total_songs"`
    TotalArtists  int `json:"total_artists"`
    TotalAlbums   int `json:"total_albums"`
    TotalEras     int `json:"total_eras"`
}
```

### PaginatedSongsResponse

```go
type PaginatedSongsResponse struct {
    Results     []Song `json:"results"`
    CurrentPage int    `json:"current_page"`
    TotalPages  int    `json:"total_pages"`
    TotalCount  int    `json:"total_count"`
}
```

### FlexibleTime

```go
type FlexibleTime struct {
    time.Time
}
```

Custom time type that can parse various date formats from the API.

## Error Types

### APIError

```go
type APIError struct {
    StatusCode int
    Message    string
}
```

General API errors.

### RateLimitError

```go
type RateLimitError struct {
    Message string
}
```

Rate limiting errors.

### NotFoundError

```go
type NotFoundError struct {
    Message string
}
```

Resource not found errors.

### AuthenticationError

```go
type AuthenticationError struct {
    Message string
}
```

Authentication errors.

### ValidationError

```go
type ValidationError struct {
    Message string
}
```

Input validation errors.

## Error Handling Example

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
