package storage

import "net/http"

type LinkStorage interface {
	Add(string, string, string) error
	AddBatch([]Batch, string) error
	Get(string) (Link, bool)
	Close() error
	PingHandler(http.ResponseWriter, *http.Request)
	GetUserURLs(string) []Link
	DeleteUserURLs(string, []string) error
}

type Link struct {
	UserID    string `json:"uuid,omitempty" db:"user_id"`
	ShortLink string `json:"short_url" db:"short_url"`
	FullLink  string `json:"original_url" db:"original_url"`
	Deleted   bool   `json:"deleted" db:"deleted"`
}

type Batch struct {
	ID       string `json:"correlation_id"`
	URL      string `json:"-"`
	ShortURL string `json:"-"`
	Result   string `json:"short_url"`
}
