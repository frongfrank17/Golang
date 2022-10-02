package main

import (
	"fmt"
	"producer/route"

	"github.com/gofiber/fiber/v2"
)

func main() {
	fmt.Println("Start Service Producer")
	app := fiber.New()

	// producer, err := sarama.NewSyncProducer(viper.GetStringSlice("kafka.servers"), nil)
	// if err != nil {
	// 	panic(err)
	// }

	// defer producer.Close()
	// eventProducer := service.NewEventProducer(producer)
	// bankService := service.NewBankService(eventProducer)
	// ctrl := controller.NewAccountController(bankService)
	route.Public_route(app)
	app.Listen(":3002")
}
