package route

import (
	"producer/controller"
	"producer/service"
	"strings"

	"github.com/Shopify/sarama"
	"github.com/gofiber/fiber/v2"
	"github.com/spf13/viper"
)

func init() {
	viper.SetConfigName("configs")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}
}
func Public_route(app *fiber.App) {
	producer, err := sarama.NewSyncProducer(viper.GetStringSlice("kafka.servers"), nil)
	if err != nil {
		panic(err)
	}

	eventProducer := service.NewEventProducer(producer)
	bankService := service.NewBankService(eventProducer)
	ctrl := controller.NewAccountController(bankService)
	//	bankController := controller
	route := app.Group("/api/bank")
	route.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})
	route.Post("/openAccount", ctrl.OpenAccount)
	route.Post("/depositFund", ctrl.DepositFund)
	route.Post("/withdrawFund", ctrl.WithdrawFund)
	route.Post("/closeAccount", ctrl.CloseAccount)
	route.Get("/get", ctrl.ClosePrint)
	//defer producer.Close()
}
