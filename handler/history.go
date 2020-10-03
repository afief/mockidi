package handler

import (
	"errors"
	"net/http"
	"strconv"
	"strings"

	"github.com/afief/mockidi/entity"
)

func (h *handlers) HandleHistory(w http.ResponseWriter, r *http.Request) (*entity.HTTPResponse, error) {
	paths := strings.Split(r.URL.Path, "/")
	if len(paths) <= 2 {
		return nil, errors.New("Not Found")
	}

	hash := paths[2]
	start, _ := strconv.ParseInt(r.URL.Query().Get("start"), 10, 64)
	stop, _ := strconv.ParseInt(r.URL.Query().Get("stop"), 10, 64)
	if stop == 0 {
		stop = 10
	}

	httpRequests, err := h.store.GetRequests(h.ctx, hash, start, stop)
	if err != nil {
		return nil, err
	}

	return &entity.HTTPResponse{
		Status: 200,
		Body:   httpRequests,
	}, nil
}
