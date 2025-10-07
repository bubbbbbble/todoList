package redis

import (
	"github.com/go-redis/redis"
	"github.com/spf13/viper"
	"fmt"
)

var rdb *redis.Client

func Init() {
	rdb = redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", viper.GetString("redis.addr"), viper.GetInt("redis.port")),
		Password: viper.GetString("redis.password"),
		DB:       viper.GetInt("redis.db"),
	})
	rdb.Set("name", "wxw", 0)
	val := rdb.Get("name").Val()
	fmt.Println("redis connect success", val)

}