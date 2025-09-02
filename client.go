package juicewrld

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"time"
)

const goWrapperVersion = "1.0.0"

type Client struct {
	BaseURL    string
	HTTPClient *http.Client
	userAgent  string
	timeout    time.Duration
}

func New(baseURL string) *Client {
	if baseURL == "" {
		baseURL = "https://juicewrldapi.com"
	}
	return &Client{
		BaseURL: baseURL,
		HTTPClient: &http.Client{
			Timeout: 30 * time.Second,
		},
		userAgent: "JuiceWRLD-API-Wrapper-Go/" + goWrapperVersion,
		timeout:   30 * time.Second,
	}
}

func (c *Client) CloseIdleConnections() {
	if c.HTTPClient != nil {
		c.HTTPClient.CloseIdleConnections()
	}
}

func (c *Client) do(ctx context.Context, method, path string, query url.Values, body interface{}, out interface{}) error {
	u, err := url.Parse(c.BaseURL)
	if err != nil {
		return err
	}
	u.Path = u.ResolveReference(&url.URL{Path: path}).Path
	if query != nil {
		u.RawQuery = query.Encode()
	}

	var reqBody io.Reader
	if body != nil {
		buf, err := json.Marshal(body)
		if err != nil {
			return err
		}
		reqBody = bytes.NewBuffer(buf)
	}

	req, err := http.NewRequestWithContext(ctx, method, u.String(), reqBody)
	if err != nil {
		return err
	}
	req.Header.Set("User-Agent", c.userAgent)
	req.Header.Set("Accept", "application/json")
	if body != nil {
		req.Header.Set("Content-Type", "application/json")
	}

	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusTooManyRequests {
		b, _ := io.ReadAll(resp.Body)
		return &RateLimitError{APIError{StatusCode: resp.StatusCode, Message: string(b)}}
	}
	if resp.StatusCode == http.StatusNotFound {
		b, _ := io.ReadAll(resp.Body)
		return &NotFoundError{APIError{StatusCode: resp.StatusCode, Message: string(b)}}
	}
	if resp.StatusCode == http.StatusUnauthorized {
		b, _ := io.ReadAll(resp.Body)
		return &AuthenticationError{APIError{StatusCode: resp.StatusCode, Message: string(b)}}
	}
	if resp.StatusCode >= 400 {
		b, _ := io.ReadAll(resp.Body)
		return &APIError{StatusCode: resp.StatusCode, Message: string(b)}
	}

	if out == nil {
		io.Copy(io.Discard, resp.Body)
		return nil
	}
	dec := json.NewDecoder(resp.Body)
	return dec.Decode(out)
}

func (c *Client) get(ctx context.Context, path string, query url.Values, out interface{}) error {
	return c.do(ctx, http.MethodGet, path, query, nil, out)
}

func (c *Client) post(ctx context.Context, path string, body interface{}, out interface{}) error {
	return c.do(ctx, http.MethodPost, path, nil, body, out)
}

func (c *Client) GetAPIOverview(ctx context.Context) (map[string]interface{}, error) {
	var endpoints interface{}
	if err := c.get(ctx, "/juicewrld/", nil, &endpoints); err != nil {
		return nil, err
	}
	return map[string]interface{}{
		"endpoints":   endpoints,
		"title":       "Juice WRLD API",
		"description": "Comprehensive API for Juice WRLD discography and content",
		"version":     goWrapperVersion,
	}, nil
}

func (c *Client) GetArtists(ctx context.Context) ([]Artist, error) {
	var raw struct {
		Results []Artist `json:"results"`
	}
	if err := c.get(ctx, "/juicewrld/artists/", nil, &raw); err != nil {
		return nil, err
	}
	return raw.Results, nil
}

func (c *Client) GetArtist(ctx context.Context, artistID int) (Artist, error) {
	var out Artist
	err := c.get(ctx, fmt.Sprintf("/juicewrld/artists/%d/", artistID), nil, &out)
	return out, err
}

func (c *Client) GetAlbums(ctx context.Context) ([]Album, error) {
	var raw struct {
		Results []Album `json:"results"`
	}
	if err := c.get(ctx, "/juicewrld/albums/", nil, &raw); err != nil {
		return nil, err
	}
	return raw.Results, nil
}

func (c *Client) GetAlbum(ctx context.Context, albumID int) (Album, error) {
	var out Album
	err := c.get(ctx, fmt.Sprintf("/juicewrld/albums/%d/", albumID), nil, &out)
	return out, err
}

func (c *Client) GetSongs(ctx context.Context, page int, category, era, search *string, pageSize int) (PaginatedSongsResponse, error) {
	q := url.Values{}
	if page > 0 {
		q.Set("page", fmt.Sprintf("%d", page))
	}
	if pageSize > 0 {
		q.Set("page_size", fmt.Sprintf("%d", pageSize))
	}
	if category != nil && *category != "" {
		q.Set("category", *category)
	}
	if era != nil && *era != "" {
		q.Set("era", *era)
	}
	if search != nil && *search != "" {
		q.Set("search", *search)
	}

	var raw map[string]interface{}
	if err := c.get(ctx, "/juicewrld/songs/", q, &raw); err != nil {
		return PaginatedSongsResponse{}, err
	}

	if _, ok := raw["results"]; !ok {
		b, _ := json.Marshal(raw)
		return PaginatedSongsResponse{}, errors.New(string(b))
	}

	var out PaginatedSongsResponse
	buf, _ := json.Marshal(raw)
	if err := json.Unmarshal(buf, &out); err != nil {
		return PaginatedSongsResponse{}, err
	}
	return out, nil
}

func (c *Client) GetSong(ctx context.Context, songID int) (Song, error) {
	var out Song
	err := c.get(ctx, fmt.Sprintf("/juicewrld/songs/%d/", songID), nil, &out)
	return out, err
}

func (c *Client) GetEras(ctx context.Context) ([]Era, error) {
	var raw struct {
		Results []Era `json:"results"`
	}
	if err := c.get(ctx, "/juicewrld/eras/", nil, &raw); err != nil {
		return nil, err
	}
	return raw.Results, nil
}

func (c *Client) GetEra(ctx context.Context, eraID int) (Era, error) {
	var out Era
	err := c.get(ctx, fmt.Sprintf("/juicewrld/eras/%d/", eraID), nil, &out)
	return out, err
}

func (c *Client) GetStats(ctx context.Context) (Stats, error) {
	var out Stats
	err := c.get(ctx, "/juicewrld/stats/", nil, &out)
	return out, err
}

func (c *Client) GetCategories(ctx context.Context) ([]map[string]interface{}, error) {
	var out struct {
		Categories []map[string]interface{} `json:"categories"`
	}
	if err := c.get(ctx, "/juicewrld/categories/", nil, &out); err != nil {
		return nil, err
	}
	return out.Categories, nil
}

func (c *Client) GetJuiceWRLDSongs(ctx context.Context, page, pageSize int) (map[string]interface{}, error) {
	q := url.Values{}
	if page > 0 {
		q.Set("page", fmt.Sprintf("%d", page))
	}
	if pageSize > 0 {
		q.Set("page_size", fmt.Sprintf("%d", pageSize))
	}
	var out map[string]interface{}
	err := c.get(ctx, "/juicewrld/player/songs/", q, &out)
	return out, err
}

func (c *Client) GetJuiceWRLDSong(ctx context.Context, songID int) (map[string]interface{}, error) {
	var out map[string]interface{}
	err := c.get(ctx, fmt.Sprintf("/juicewrld/player/songs/%d/", songID), nil, &out)
	return out, err
}

func (c *Client) PlayJuiceWRLDSong(ctx context.Context, songID int) (map[string]interface{}, error) {
	songData, err := c.GetJuiceWRLDSong(ctx, songID)
	if err != nil {
		return nil, err
	}
	fileAny, ok := songData["file"]
	if !ok {
		return map[string]interface{}{"error": "Song file information not found", "song_id": songID, "status": "no_file_info"}, nil
	}
	fileStr, _ := fileAny.(string)
	if fileStr == "" {
		return map[string]interface{}{"error": "Invalid file URL format", "song_id": songID, "status": "invalid_url"}, nil
	}
	idx := -1
	for i := 0; i+6 <= len(fileStr); i++ {
		if fileStr[i:i+7] == "/media/" {
			idx = i
			break
		}
	}
	if idx == -1 {
		return map[string]interface{}{"error": "Invalid file URL format", "song_id": songID, "status": "invalid_url"}, nil
	}
	filePath := fileStr[idx+7:]

	possiblePaths := []string{
		fmt.Sprintf("Compilation/1. Released Discography/%v/%v.mp3", songData["album"], songData["title"]),
		fmt.Sprintf("Compilation/2. Unreleased Discography/%v.mp3", songData["title"]),
		fmt.Sprintf("Snippets/%v/%v.mp4", songData["title"], songData["title"]),
		fmt.Sprintf("Session Edits/%v.mp3", songData["title"]),
	}

	for _, p := range possiblePaths {
		streamURL := fmt.Sprintf("%s/juicewrld/files/download/?path=%s", c.BaseURL, url.QueryEscape(p))
		req, _ := http.NewRequestWithContext(ctx, http.MethodGet, streamURL, nil)
		req.Header.Set("Range", "bytes=0-0")
		resp, err := c.HTTPClient.Do(req)
		if err == nil && (resp.StatusCode == http.StatusOK || resp.StatusCode == http.StatusPartialContent) {
			ct := resp.Header.Get("content-type")
			resp.Body.Close()
			return map[string]interface{}{
				"status":       "success",
				"song_id":      songID,
				"stream_url":   streamURL,
				"file_path":    p,
				"content_type": ct,
			}, nil
		}
		if resp != nil {
			resp.Body.Close()
		}
	}

	streamURL := fmt.Sprintf("%s/juicewrld/files/download/?path=%s", c.BaseURL, url.QueryEscape(filePath))
	return map[string]interface{}{
		"status":     "file_not_found_but_url_provided",
		"song_id":    songID,
		"stream_url": streamURL,
		"file_path":  filePath,
		"note":       "File may not exist at this path, but streaming URL is provided",
	}, nil
}

func (c *Client) StreamAudioFile(ctx context.Context, filePath string) (map[string]interface{}, error) {
	streamURL := fmt.Sprintf("%s/juicewrld/files/download/?path=%s", c.BaseURL, url.QueryEscape(filePath))
	req, _ := http.NewRequestWithContext(ctx, http.MethodGet, streamURL, nil)
	req.Header.Set("Range", "bytes=0-0")
	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return map[string]interface{}{"error": fmt.Sprintf("Request failed: %v", err), "file_path": filePath, "status": "request_error"}, nil
	}
	defer resp.Body.Close()
	if resp.StatusCode == http.StatusOK || resp.StatusCode == http.StatusPartialContent {
		supportsRange := resp.Header.Get("accept-ranges")
		return map[string]interface{}{
			"status":         "success",
			"stream_url":     streamURL,
			"file_path":      filePath,
			"content_type":   resp.Header.Get("content-type"),
			"content_length": resp.Header.Get("content-length"),
			"supports_range": supportsRange != "" && supportsRange != "none",
		}, nil
	}
	if resp.StatusCode == http.StatusNotFound {
		return map[string]interface{}{"error": "Audio file not found", "file_path": filePath, "status": "file_not_found"}, nil
	}
	return map[string]interface{}{"error": fmt.Sprintf("HTTP %d", resp.StatusCode), "file_path": filePath, "status": "http_error"}, nil
}

func (c *Client) BrowseFiles(ctx context.Context, path string, search *string) (DirectoryInfo, error) {
	q := url.Values{}
	if path != "" {
		q.Set("path", path)
	}
	if search != nil && *search != "" {
		q.Set("search", *search)
	}
	var out DirectoryInfo
	err := c.get(ctx, "/juicewrld/files/browse/", q, &out)
	return out, err
}

func (c *Client) GetFileInfo(ctx context.Context, filePath string) (FileInfo, error) {
	q := url.Values{"path": {filePath}}
	var out FileInfo
	err := c.get(ctx, "/juicewrld/files/info/", q, &out)
	return out, err
}

func (c *Client) DownloadFile(ctx context.Context, filePath string) ([]byte, error) {
	u := fmt.Sprintf("%s/juicewrld/files/download/?path=%s", c.BaseURL, url.QueryEscape(filePath))
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, u, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("User-Agent", c.userAgent)
	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode >= 400 {
		b, _ := io.ReadAll(resp.Body)
		return nil, &APIError{StatusCode: resp.StatusCode, Message: string(b)}
	}
	return io.ReadAll(resp.Body)
}

func (c *Client) DownloadFileTo(ctx context.Context, filePath, savePath string) (string, error) {
	data, err := c.DownloadFile(ctx, filePath)
	if err != nil {
		return "", err
	}
	if err := writeFileAtomic(savePath, data); err != nil {
		return "", err
	}
	return savePath, nil
}

func writeFileAtomic(path string, data []byte) error {
	tmp := path + ".tmp"
	if err := os.WriteFile(tmp, data, 0o644); err != nil {
		return err
	}
	return os.Rename(tmp, path)
}

func (c *Client) GetCoverArt(ctx context.Context, filePath string) ([]byte, error) {
	u := fmt.Sprintf("%s/juicewrld/files/cover-art/?path=%s", c.BaseURL, url.QueryEscape(filePath))
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, u, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("User-Agent", c.userAgent)
	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode >= 400 {
		b, _ := io.ReadAll(resp.Body)
		return nil, &APIError{StatusCode: resp.StatusCode, Message: string(b)}
	}
	return io.ReadAll(resp.Body)
}

func (c *Client) CreateZip(ctx context.Context, filePaths []string) ([]byte, error) {
	u := fmt.Sprintf("%s/juicewrld/files/zip-selection/", c.BaseURL)
	reqBody := map[string]interface{}{"paths": filePaths}
	buf, err := json.Marshal(reqBody)
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, u, bytes.NewReader(buf))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")
	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode >= 400 {
		b, _ := io.ReadAll(resp.Body)
		return nil, &APIError{StatusCode: resp.StatusCode, Message: string(b)}
	}
	return io.ReadAll(resp.Body)
}

func (c *Client) StartZipJob(ctx context.Context, filePaths []string) (string, error) {
	var out struct {
		JobID string `json:"job_id"`
	}
	err := c.post(ctx, "/juicewrld/start-zip-job/", map[string]interface{}{"paths": filePaths}, &out)
	if err != nil {
		return "", err
	}
	return out.JobID, nil
}

func (c *Client) GetZipJobStatus(ctx context.Context, jobID string) (map[string]interface{}, error) {
	var out map[string]interface{}
	err := c.get(ctx, fmt.Sprintf("/juicewrld/zip-job-status/%s/", url.PathEscape(jobID)), nil, &out)
	return out, err
}

func (c *Client) CancelZipJob(ctx context.Context, jobID string) (bool, error) {
	var out map[string]interface{}
	err := c.post(ctx, fmt.Sprintf("/juicewrld/cancel-zip-job/%s/", url.PathEscape(jobID)), nil, &out)
	if err != nil {
		return false, err
	}
	return true, nil
}

func (c *Client) SearchSongs(ctx context.Context, query string, category *string, year *int, tags []string, limit int, offset int) (SearchResult, error) {
	page := 1
	if limit > 0 {
		page = (offset / limit) + 1
	}
	q := url.Values{
		"search":    {query},
		"page_size": {fmt.Sprintf("%d", limit)},
		"page":      {fmt.Sprintf("%d", page)},
	}
	if category != nil && *category != "" {
		q.Set("category", *category)
	}
	if year != nil && *year > 0 {
		q.Set("year", fmt.Sprintf("%d", *year))
	}
	if len(tags) > 0 {
		joined := ""
		for i, t := range tags {
			if i > 0 {
				joined += ","
			}
			joined += t
		}
		q.Set("tags", joined)
	}

	var raw PaginatedSongsResponse
	if err := c.get(ctx, "/juicewrld/songs/", q, &raw); err != nil {
		return SearchResult{}, err
	}
	res := SearchResult{
		Songs:     raw.Results,
		Total:     raw.Count,
		QueryTime: "0ms",
	}
	if category != nil {
		res.Category = category
	}
	return res, nil
}

func (c *Client) GetSongsByCategory(ctx context.Context, category string, page, pageSize int) (PaginatedSongsResponse, error) {
	return c.GetSongs(ctx, page, &category, nil, nil, pageSize)
}
