package initialize

import (
	"context"
	"fmt"
	"user-service/global"

	"github.com/redis/go-redis/v9"
	"go.uber.org/zap"
)

var ctx = context.Background()

func InitRedis() {
	r := global.Config.Redis

	redisDb := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%v", r.Host, r.Port),
		Password: r.Password,
		DB:       r.Database,
		PoolSize: 10,
	})

	_, err := redisDb.Ping(ctx).Result()
	if err != nil {
		global.Logger.Error("Redis initialization Error:", zap.Error(err))
	}

	fmt.Println("Initializing Redis Successfully")
	global.Rdb = redisDb
}
