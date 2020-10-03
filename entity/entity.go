package entity

import (
	"context"
	"net/http"
	"time"
)

// PathInfo represent information stored in redis for related path
type PathInfo struct {
	Body    string            `json:"body"`
	Status  int               `json:"status"`
	Headers map[string]string `json:"headers"`
}

// Store represent store repository
type Store interface {
	Save(ctx context.Context, path string, data *PathInfo) error
	Get(ctx context.Context, hash string) *PathInfo
	PushRequest(ctx context.Context, hash string, data *HTTPRequest) error
	GetRequests(ctx context.Context, hash string, start int64, stop int64) ([]*HTTPRequest, error)
}

// Handlers contains all handlers
type Handlers interface {
	Init(w http.ResponseWriter, r *http.Request)
	HandleCreate(w http.ResponseWriter, r *http.Request) (*HTTPResponse, error)
	HandleRequest(w http.ResponseWriter, r *http.Request) (*HTTPResponse, error)
	HandleHistory(w http.ResponseWriter, r *http.Request) (*HTTPResponse, error)
}

// HTTPRequest ...
type HTTPRequest struct {
	Path          string            `json:"path"`
	Method        string            `json:"method"`
	QueryString   string            `json:"query_string"`
	Body          interface{}       `json:"body"`
	Headers       map[string]string `json:"headers"`
	ContentLength int64             `json:"content_length"`
	RequestTime   time.Time         `json:"request_time"`
}

// HTTPResponse ...
type HTTPResponse struct {
	Status  int
	Body    interface{}
	Headers map[string]string
}
