package configs

import (
	"consumer/repository"
	"log"
	"os"
	"time"

	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

const DNS = "root:root@tcp(localhost:3333)/DBProduct?charset=utf8mb4&parseTime=True&loc=Local"

func DatabaseInit(dial gorm.Dialector) (db *gorm.DB, err error) {
	loc, err := time.LoadLocation("Asia/Bangkok")
	if err != nil {
		return nil, err
	}
	time.Local = loc
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			SlowThreshold: time.Second,
			LogLevel:      logger.Error,
			Colorful:      true,
		},
	)

	mydb, err := gorm.Open(dial, &gorm.Config{
		Logger: newLogger,
	})
	if err != nil {
		return nil, err
	}
	sqlDB, err := mydb.DB()
	if err != nil {
		return nil, err
	}
	sqlDB.SetMaxIdleConns(50)
	sqlDB.SetConnMaxLifetime(time.Hour)
	sqlDB.SetMaxOpenConns(100)

	mydb.AutoMigrate(&repository.BankAccount{})
	return mydb, nil
}
