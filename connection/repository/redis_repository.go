package repository

import (
	"github.com/go-redis/redis"
	"time"
)

type redisConnectionRepository struct {
	client *redis.Client
}

func NewRedisConnectionRepository(client *redis.Client) ConnectionRepository {
	return &redisConnectionRepository{client}
}

func (repo *redisConnectionRepository) StoreSecret(userId string, secret string) (*string, error) {
	var key = "connection:" + userId
	err := repo.client.Set(key, secret, time.Second*60).Err()
	if err != nil {
		return nil, err
	}
	return &secret, nil
}

func (repo *redisConnectionRepository) GetSecret(key string) (*string, error) {
	result, err := repo.client.Get(key).Result()
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (repo *redisConnectionRepository) DeleteKey(key string) (bool, error) {
	result, err := repo.client.Del(key).Result()
	if err != nil {
		return false, err
	}
	return result > 0, err
}

func (repo *redisConnectionRepository) GetAllConnection() ([]string, error) {
	scan := repo.client.Scan(0, "connection:*", 0)
	keys, _, err := scan.Result()
	if err != nil {
		return nil, err
	}
	return keys, err
}

func (repo *redisConnectionRepository) Subscribe(channel string) *redis.PubSub {
	return repo.client.Subscribe(channel)
}

func (repo *redisConnectionRepository) Publish(channel string, message string) (int, error) {
	result := repo.client.Publish(channel, message)
	err := result.Err()
	if err != nil {
		return 0, err
	}
	return int(result.Val()), nil
}
