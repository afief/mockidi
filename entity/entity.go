package entity

import (
	"context"
	"net/http"
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
}

// Handlers contains all handlers
type Handlers interface {
	Init(w http.ResponseWriter, r *http.Request)
	HandleCreate(w http.ResponseWriter, r *http.Request) (*HTTPResponse, error)
	HandleRequest(w http.ResponseWriter, r *http.Request) (*HTTPResponse, error)
}

// HTTPResponse ...
type HTTPResponse struct {
	Status  int
	Body    interface{}
	Headers map[string]string
}
