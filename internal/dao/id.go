package dao

import (
	"apps/stardust/internal/snowflake"
	"context"
	"strconv"
	"time"

	"github.com/UnderTreeTech/waterdrop/pkg/log"
	"github.com/gomodule/redigo/redis"
)

//同步时钟
func (d *dao) SyncWorkerInfo(worker *snowflake.IdWorker) {
	ctx := context.Background()
	redisKey := "id_pumper_" + strconv.Itoa(int(worker.GetDataCenter())) + "_" + strconv.Itoa(int(worker.GetWorkerId()))

	workerTimestamp, err := redis.Int64(d.redis.Do(ctx, "get", redisKey))
	if err != nil && err != redis.ErrNil {
		panic(err)
	}

	if worker.GetCurrentTimestamp() < workerTimestamp {
		panic("fatal error:server time backwards")
	}

	go func(ctx context.Context, key string) {
		ticker := time.NewTicker(2 * time.Second)
		defer ticker.Stop()

		for {
			select {
			case <-ticker.C:
				_, err := d.redis.Do(ctx, "set", redisKey, worker.GetCurrentTimestamp())
				if err != nil {
					log.Error(ctx, "update id pumper fail", log.String("error", err.Error()))
				}
			}
		}
	}(ctx, redisKey)
}
