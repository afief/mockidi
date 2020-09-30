package handler

import (
	"errors"
	"net/http"
	"strings"

	"github.com/afief/mockidi/entity"
)

func (h *handlers) HandleRequest(w http.ResponseWriter, r *http.Request) (*entity.HTTPResponse, error) {
	paths := strings.Split(r.URL.Path, "/")
	if len(paths) <= 1 {
		return nil, errors.New("Not Found")
	}

	pathInfo := h.store.Get(h.ctx, paths[1])
	if pathInfo == nil {
		return nil, errors.New("Not Found")
	}
	return &entity.HTTPResponse{
		Status:  pathInfo.Status,
		Body:    pathInfo.Body,
		Headers: pathInfo.Headers,
	}, nil
}
