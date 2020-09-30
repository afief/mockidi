package handler

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"

	"github.com/afief/mockidi/entity"
	"github.com/google/uuid"
)

// HandleCreate save path to redis
func (h *handlers) HandleCreate(w http.ResponseWriter, r *http.Request) (*entity.HTTPResponse, error) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return nil, err
	}

	var pathInfo entity.PathInfo
	if err := json.Unmarshal(body, &pathInfo); err != nil {
		return nil, err
	}

	if pathInfo.Status == 0 {
		return nil, errors.New("Invalid Parameters")
	}

	hash := uuid.Must(uuid.NewRandom()).String()
	if err := h.store.Save(h.ctx, hash, &pathInfo); err != nil {
		return nil, err
	}

	return &entity.HTTPResponse{
		Status: 200,
		Body: map[string]string{
			"hash": hash,
		},
	}, nil
}
