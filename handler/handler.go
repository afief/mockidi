package handler

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/afief/mockidi/entity"
)

type handlers struct {
	ctx   context.Context
	store entity.Store
}

// NewHandlers returns Handlers interface
func NewHandlers(ctx context.Context, store entity.Store) entity.Handlers {
	return &handlers{
		ctx:   ctx,
		store: store,
	}
}

// Initial ...
func (h *handlers) Init(w http.ResponseWriter, r *http.Request) {
	var resp *entity.HTTPResponse
	var err error

	switch path := r.URL.Path; path {
	case "/create":
		resp, err = h.HandleCreate(w, r)
	default:
		resp, err = h.HandleRequest(w, r)
	}

	if err != nil {
		resp = &entity.HTTPResponse{
			Status: 400,
			Body: map[string]string{
				"errorMessage": err.Error(),
			},
		}
	}

	w.Header().Set("Content-Type", "application/json")
	for k, v := range resp.Headers {
		w.Header().Set(k, v)
	}
	w.WriteHeader(resp.Status)

	if strBody, ok := resp.Body.(string); ok {
		fmt.Fprint(w, strBody)
		return
	}

	json.NewEncoder(w).Encode(resp.Body)
}
