package main

import (
	"fmt"
	db "redishex/database"
	repository "redishex/repository"
	service "redishex/service"

	fiber "github.com/gofiber/fiber/v2"
	"gorm.io/driver/mysql"
)

func main() {
	fmt.Println("Redis Golang")
	server := fiber.New()
	dns := "root:root@tcp(localhost:3333)/DBProduct?charset=utf8mb4&parseTime=True&loc=Local"
	uri_redis := "127.0.0.1:6379"
	dial := mysql.Open(dns)
	sql, _ := db.SQLInit(dial)
	mem := db.Redisinit(uri_redis)
	domian := repository.NewRedisRepositoryRedis(sql, mem)
	srv := service.NewRedisService(domian)
	data, err := srv.CreateData("000321", "IPAD PRO", 32000)
	if err != nil {
		panic(err.Error())
	}
	fmt.Println(data)
	server.Listen("8000")

}
