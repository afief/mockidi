package handler

import (
	"context"
	"fmt"
	"net/http"
	"strings"

	"github.com/afief/mockidi/entity"
)

type ctxVal string

type handlers struct {
	ctx   context.Context
	store entity.Store
}

// NewHandlers returns Handlers interface
func NewHandlers(store entity.Store) func(w http.ResponseWriter, r *http.Request) {
	h := &handlers{
		store: store,
	}

	return func(w http.ResponseWriter, r *http.Request) {

	}
}

// Initial ...
func (h *handlers) Init(w http.ResponseWriter, r *http.Request) {
	fmt.Println(h.ctx)
	h.ctx = context.Background()
	var resp *entity.HTTPResponse
	var err error

	paths := strings.Split(r.URL.Path, "/")
	switch path := paths[1]; path {
	case "create":
		resp, err = h.HandleCreate(w, r)
	case "history":
		resp, err = h.HandleHistory(w, r)
	default:
		resp, err = h.HandleRequest(w, r)
	}
}
