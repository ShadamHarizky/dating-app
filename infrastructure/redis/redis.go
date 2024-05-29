package redis

import "github.com/go-redis/redis/v8"

type RedisDb struct {
	Client *redis.Client
}

func NewRedisDB(host, port, password string) (*RedisDb, error) {
	redisClient := redis.NewClient(&redis.Options{
		Addr:     host + ":" + port,
		Password: password,
		DB:       0,
	})
	return &RedisDb{
		Client: redisClient,
	}, nil
}
