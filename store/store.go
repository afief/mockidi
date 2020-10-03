package store

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"
	"time"

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

	ttl := time.Hour * 144 // 6 months
	if err := s.client.Set(ctx, hashKey(hash), encoded, ttl).Err(); err != nil {
		return err
	}

	return nil
}

func (s *store) Get(ctx context.Context, hash string) *entity.PathInfo {
	encoded, err := s.client.Get(ctx, hashKey(hash)).Bytes()
	if err != nil {
		return nil
	}
	var decoded entity.PathInfo
	if err := json.Unmarshal(encoded, &decoded); err != nil {
		return nil
	}
	return &decoded
}

func (s *store) PushRequest(ctx context.Context, hash string, data *entity.HTTPRequest) error {
	encoded, err := json.Marshal(data)
	if err != nil {
		return err
	}

	if err := s.client.RPush(ctx, hashKey(hash, "reqs"), encoded).Err(); err != nil {
		return err
	}

	return nil
}

func (s *store) GetRequests(ctx context.Context, hash string, start int64, stop int64) ([]*entity.HTTPRequest, error) {
	_start := start
	start = 0 - start - stop
	stop = 0 - _start - 1

	result := s.client.LRange(ctx, hashKey(hash, "reqs"), start, stop)
	if result.Err() != nil {
		return nil, result.Err()
	}

	var decodeds []*entity.HTTPRequest
	for _, encoded := range result.Val() {
		var decoded entity.HTTPRequest
		if err := json.Unmarshal([]byte(encoded), &decoded); err == nil {
			decodeds = append(decodeds, &decoded)
		}
	}
	return decodeds, nil
}

func hashKey(hash ...string) string {
	return fmt.Sprintf("mockidi:%s", strings.Join(hash, ":"))
}
