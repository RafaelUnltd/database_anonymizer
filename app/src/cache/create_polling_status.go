package cache

import (
	"context"
	"database_anonymizer/app/src/common"
	"database_anonymizer/app/src/structs"
	"encoding/json"
)

func (c CacheManager) CreatePollingStatus(ctx context.Context, pollingKey string) error {
	status := structs.PollingStatus{
		Key:      pollingKey,
		Finished: false,
		Status:   structs.StatusPending,
		Progress: make(map[string]structs.TableStatus),
	}

	jsonStatus, err := json.Marshal(status)
	if err != nil {
		return err
	}

	err = c.redisClient.Set(ctx, pollingKey, jsonStatus, common.CacheDuration).Err()

	return err
}

func (c CacheManager) ReadPollingStatus(ctx context.Context, pollingKey string) (structs.PollingStatus, error) {
	status, err := c.redisClient.Get(ctx, pollingKey).Result()
	if err != nil {
		return structs.PollingStatus{}, err
	}

	unmarshalledStatus := structs.PollingStatus{}
	err = json.Unmarshal([]byte(status), &unmarshalledStatus)
	if err != nil {
		return structs.PollingStatus{}, err
	}

	return unmarshalledStatus, err
}

func (c CacheManager) UpdatePollingStatus(ctx context.Context, pollingKey string, status structs.PollingStatus) error {
	jsonStatus, err := json.Marshal(status)
	if err != nil {
		return err
	}

	err = c.redisClient.Set(ctx, pollingKey, jsonStatus, common.CacheDuration).Err()
	if err != nil {
		return err
	}

	return nil
}
