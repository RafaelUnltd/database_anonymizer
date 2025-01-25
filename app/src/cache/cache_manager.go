package cache

import "github.com/redis/go-redis/v9"

type CacheManager struct {
	redisClient *redis.Client
}

// NewCacheManager creates a new CacheManager
func NewCacheManager() CacheManager {
	return CacheManager{
		redisClient: nil,
	}
}

func (c *CacheManager) StartCacheManger(address, password string, db int) {
	c.redisClient = redis.NewClient(&redis.Options{
		Addr:     address,
		Password: password,
		DB:       db,
	})
}
