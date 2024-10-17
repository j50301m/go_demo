package redis_cache

import (
	"context"
	"encoding/json"
	"hype-casino-platform/pkg/db"
	"hype-casino-platform/pkg/kgserr"
	"hype-casino-platform/pkg/kgsotel"
	"time"

	"github.com/redis/go-redis/v9"
)

type RedisCache struct {
	client *redis.Client
}

var _ db.Cache = (*RedisCache)(nil)

func NewRedisCache(client *redis.Client) *RedisCache {
	return &RedisCache{
		client: client,
	}
}

func (r *RedisCache) Get(ctx context.Context, key string) (string, *kgserr.KgsError) {
	// Start trace
	ctx, span := kgsotel.StartTrace(ctx)
	defer span.End()

	// Get value from Redis
	val, err := r.client.Get(ctx, key).Result()
	if err != nil {
		// When key not found return ResourceNotFound error.
		if err == redis.Nil {
			kgsErr := kgserr.New(kgserr.ResourceNotFound, "Failed to get key", err)
			kgsotel.Error(ctx, kgsErr.Error())
			return "", kgsErr
		}
		// Otherwise, return InternalServerError error.
		kgsErr := kgserr.New(kgserr.InternalServerError, "Failed to get key", err)
		kgsotel.Error(ctx, kgsErr.Error())
		return "", kgsErr
	}

	return val, nil
}

func (r *RedisCache) Set(ctx context.Context, key string, value string, expiration time.Duration) *kgserr.KgsError {
	// Start trace
	ctx, span := kgsotel.StartTrace(ctx)
	defer span.End()

	// Set value in Redis
	err := r.client.Set(ctx, key, value, expiration).Err()
	if err != nil {
		// Return InternalServerError error.
		kgsErr := kgserr.New(kgserr.InternalServerError, "Failed to set key", err)
		kgsotel.Error(ctx, kgsErr.Error())
		return kgsErr
	}

	return nil
}

func (r *RedisCache) GetObject(ctx context.Context, key string, dest any) *kgserr.KgsError {
	// Start trace
	ctx, span := kgsotel.StartTrace(ctx)
	defer span.End()

	// Get value from Redis
	val, err := r.client.Get(ctx, key).Result()
	if err != nil {
		// When key not found return ResourceNotFound error.
		if err == redis.Nil {
			kgsErr := kgserr.New(kgserr.ResourceNotFound, "Failed to get key", err)
			kgsotel.Error(ctx, kgsErr.Error())
			return kgsErr
		}
		// Otherwise, return InternalServerError error.
		kgsErr := kgserr.New(kgserr.InternalServerError, "Failed to get key", err)
		kgsotel.Error(ctx, kgsErr.Error())
		return kgsErr
	}

	// Unmarshal value to dest
	err = json.Unmarshal([]byte(val), dest)
	if err != nil {
		// Return InternalServerError error.
		kgsErr := kgserr.New(kgserr.InternalServerError, "Failed to unmarshal value", err)
		kgsotel.Error(ctx, kgsErr.Error())
		return kgsErr
	}

	return nil
}

func (r *RedisCache) SetObject(ctx context.Context, key string, value any, expiration time.Duration) *kgserr.KgsError {
	// Start trace
	ctx, span := kgsotel.StartTrace(ctx)
	defer span.End()

	// Marshal value
	val, err := json.Marshal(value)
	if err != nil {
		// Return InternalServerError error.
		kgsErr := kgserr.New(kgserr.InternalServerError, "Failed to marshal value", err)
		kgsotel.Error(ctx, kgsErr.Error())
		return kgsErr
	}

	// Set value in Redis
	err = r.client.Set(ctx, key, val, expiration).Err()
	if err != nil {
		// Return InternalServerError error.
		kgsErr := kgserr.New(kgserr.InternalServerError, "Failed to set key", err)
		kgsotel.Error(ctx, kgsErr.Error())
		return kgsErr
	}

	return nil
}

func (r *RedisCache) Delete(ctx context.Context, keys ...string) *kgserr.KgsError {
	// Start trace
	ctx, span := kgsotel.StartTrace(ctx)
	defer span.End()

	// Delete key from Redis
	result, err := r.client.Del(ctx, keys...).Result()
	if err != nil {
		// Return InternalServerError error.
		kgsErr := kgserr.New(kgserr.InternalServerError, "Failed to delete key", err)
		kgsotel.Error(ctx, kgsErr.Error())
		return kgsErr
	}

	// If no key was deleted, return ResourceNotFound error
	if result == 0 {
		kgsErr := kgserr.New(kgserr.ResourceNotFound, "Key not found", nil)
		kgsotel.Error(ctx, kgsErr.Error())
		return kgsErr
	}

	return nil
}
