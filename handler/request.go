package handler

import (
	"errors"
	"io/ioutil"
	"net/http"
	"strings"
	"time"

	"github.com/afief/mockidi/entity"
)

func (h *handlers) HandleRequest(w http.ResponseWriter, r *http.Request) (*entity.HTTPResponse, error) {
	paths := strings.Split(r.URL.Path, "/")
	if len(paths) <= 1 {
		return nil, errors.New("Not Found")
	}

	hash := paths[1]
	pathInfo := h.store.Get(h.ctx, hash)
	if pathInfo == nil {
		return nil, errors.New("Not Found")
	}

	request := buildHTTPRequest(r)
	if err := h.store.PushRequest(h.ctx, hash, request); err != nil {
		return nil, err
	}

	return &entity.HTTPResponse{
		Status:  pathInfo.Status,
		Body:    pathInfo.Body,
		Headers: pathInfo.Headers,
	}, nil
}

func buildHTTPRequest(r *http.Request) *entity.HTTPRequest {
	requestHeaders := map[string]string{}
	for key, values := range r.Header {
		requestHeaders[key] = values[0]
	}

	requestBody, _ := ioutil.ReadAll(r.Body)

	return &entity.HTTPRequest{
		ContentLength: r.ContentLength,
		Method:        r.Method,
		Path:          r.URL.Path,
		Headers:       requestHeaders,
		QueryString:   r.URL.Query().Encode(),
		Body:          string(requestBody),
		RequestTime:   time.Now(),
	}
}
