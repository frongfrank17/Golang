package main

import (
	"consumer/configs"
	"consumer/repository"
	"consumer/service"
	"context"
	"events"
	"fmt"
	"strings"

	"github.com/Shopify/sarama"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
)

const DNS = "root:root@tcp(localhost:3333)/DBProduct?charset=utf8mb4&parseTime=True&loc=Local"

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
func main() {

	consumer, err := sarama.NewConsumerGroup(viper.GetStringSlice("kafka.servers"), viper.GetString("kafka.group"), nil)
	fmt.Println("Start Sevice Consumer")
	if err != nil {
		panic(err)
	}
	defer consumer.Close()
	dial := mysql.Open(DNS)
	db, err := configs.DatabaseInit(dial)
	if err != nil {
		panic("Database Error")
	}
	repo := repository.NewAccountRepository(db)
	srv := service.NewAccountService(repo)
	serviceConsumer := service.NewConsumerHandler(srv)
	fmt.Println("Consumer")
	for {
		consumer.Consume(context.Background(), events.Topics, serviceConsumer)
	}
	fmt.Println("Consumer")

}
