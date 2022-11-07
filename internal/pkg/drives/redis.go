package drives

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v9"
	"github.com/pro911/authApis/internal/pkg/conf"
)

var (
	rdb *redis.Client
)

func MustInitRedis() {
	rdb = redis.NewClient(&redis.Options{
		Addr: fmt.Sprintf("%s:%d", conf.Conf.Redis.Host, conf.Conf.Redis.Port),
		//Username:              conf.Conf.Redis.UserName,
		Password: conf.Conf.Redis.Password,
		DB:       conf.Conf.Redis.DataBase,
	})

	//defer rdb.Close()

	rdb.Ping(context.Background())

	fmt.Println("redis initialized")
}

func GetDB() *redis.Client {
	return rdb
}
