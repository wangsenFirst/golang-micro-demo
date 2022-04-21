package initialize

import (
	"fmt"
	goredislib "github.com/go-redis/redis/v8"
	"github.com/go-redsync/redsync/v4"
	"github.com/go-redsync/redsync/v4/redis/goredis/v8"
	"mxshop/inventory_srv/global"
)

func InitRedis() {

	client := goredislib.NewClient(&goredislib.Options{
		//Addr: "192.168.0.104:6379",
		Addr: fmt.Sprintf("%s:%d", global.ServerConfig.RedisInfo.Host, global.ServerConfig.RedisInfo.Port),
	})
	pool := goredis.NewPool(client) // or, pool := redigo.NewPool(...)
	global.RedisDB = redsync.New(pool)
}
