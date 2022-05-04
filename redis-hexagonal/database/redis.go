package databese

import (
	"fmt"

	redis "github.com/go-redis/redis/v8"
)

func Redisinit(uri string) *redis.Client {
	con := redis.NewClient(&redis.Options{
		Addr: uri,
	})
	fmt.Println(con)
	return con

}
