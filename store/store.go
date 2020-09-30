package store

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/afief/mockidi/entity"
	"github.com/go-redis/redis/v8"
)

type store struct {
	client *redis.Client
}

// NewStore initiate new store
func NewStore() entity.Store {
	return &store{
		client: redis.NewClient(&redis.Options{
			Addr:     "localhost:6379",
			Password: "",
			DB:       0,
		}),
	}
}

func (s *store) Save(ctx context.Context, hash string, data *entity.PathInfo) error {
	encoded, err := json.Marshal(data)
	if err != nil {
		return err
	}

	if err := s.client.HSet(ctx, fmt.Sprintf("mockidi:%s", hash), "default", encoded).Err(); err != nil {
		return nil
	}

	return nil
}

func (s *store) Get(ctx context.Context, hash string) *entity.PathInfo {
	encoded, err := s.client.HGet(ctx, fmt.Sprintf("mockidi:%s", hash), "default").Bytes()
	if err != nil {
		return nil
	}
	var decoded entity.PathInfo
	if err := json.Unmarshal(encoded, &decoded); err != nil {
		return nil
	}
	return &decoded
}
