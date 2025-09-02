package juicewrld

import (
	"encoding/json"
	"strings"
	"time"
)

type FlexibleTime struct {
	time.Time
}

func (ft *FlexibleTime) UnmarshalJSON(data []byte) error {
	str := strings.Trim(string(data), `"`)
	if str == "null" || str == "" {
		return nil
	}

	formats := []string{
		"2006-01-02T15:04:05Z07:00",
		"2006-01-02T15:04:05.999999Z07:00",
		"2006-01-02T15:04:05.999999",
		"2006-01-02T15:04:05",
		"2006-01-02",
		time.RFC3339,
		time.RFC3339Nano,
	}

	for _, format := range formats {
		if t, err := time.Parse(format, str); err == nil {
			ft.Time = t
			return nil
		}
	}

	return nil
}

func (ft FlexibleTime) MarshalJSON() ([]byte, error) {
	return json.Marshal(ft.Time.Format(time.RFC3339))
}

type Artist struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Bio  string `json:"bio"`
}

type Album struct {
	ID          int          `json:"id"`
	Title       string       `json:"title"`
	Type        string       `json:"type"`
	Artist      Artist       `json:"artist"`
	ReleaseDate FlexibleTime `json:"release_date"`
	Description string       `json:"description"`
}

type Era struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	TimeFrame   string `json:"time_frame"`
}

type Song struct {
	ID                    int         `json:"id"`
	Name                  string      `json:"name"`
	OriginalKey           string      `json:"original_key"`
	Category              string      `json:"category"`
	Era                   Era         `json:"era"`
	TrackTitles           []string    `json:"track_titles"`
	CreditedArtists       string      `json:"credited_artists"`
	Producers             string      `json:"producers"`
	Engineers             string      `json:"engineers"`
	AdditionalInformation string      `json:"additional_information"`
	FileNames             string      `json:"file_names"`
	Instrumentals         string      `json:"instrumentals"`
	RecordingLocations    string      `json:"recording_locations"`
	RecordDates           string      `json:"record_dates"`
	PreviewDate           string      `json:"preview_date"`
	ReleaseDate           string      `json:"release_date"`
	Dates                 string      `json:"dates"`
	Length                string      `json:"length"`
	LeakType              string      `json:"leak_type"`
	DateLeaked            string      `json:"date_leaked"`
	Notes                 string      `json:"notes"`
	ImageURL              string      `json:"image_url"`
	SessionTitles         string      `json:"session_titles"`
	SessionTracking       string      `json:"session_tracking"`
	InstrumentalNames     string      `json:"instrumental_names"`
	PublicID              interface{} `json:"public_id"`
}

type FileInfo struct {
	Name      string        `json:"name"`
	Type      string        `json:"type"`
	Size      int64         `json:"size"`
	SizeHuman string        `json:"size_human"`
	Path      string        `json:"path"`
	Extension string        `json:"extension"`
	MimeType  string        `json:"mime_type"`
	Created   *FlexibleTime `json:"created"`
	Modified  *FlexibleTime `json:"modified"`
	Encoding  *string       `json:"encoding"`
}

type DirectoryInfo struct {
	CurrentPath       string              `json:"current_path"`
	PathParts         []map[string]string `json:"path_parts"`
	Items             []FileInfo          `json:"items"`
	TotalFiles        int                 `json:"total_files"`
	TotalDirectories  int                 `json:"total_directories"`
	SearchQuery       *string             `json:"search_query"`
	IsRecursiveSearch bool                `json:"is_recursive_search"`
}

type SearchResult struct {
	Songs     []Song  `json:"songs"`
	Total     int     `json:"total"`
	Category  *string `json:"category"`
	QueryTime string  `json:"query_time"`
}

type PaginatedSongsResponse struct {
	Results  []Song  `json:"results"`
	Count    int     `json:"count"`
	Next     *string `json:"next"`
	Previous *string `json:"previous"`
}

type Stats struct {
	TotalSongs    int            `json:"total_songs"`
	CategoryStats map[string]int `json:"category_stats"`
	EraStats      map[string]int `json:"era_stats"`
}
