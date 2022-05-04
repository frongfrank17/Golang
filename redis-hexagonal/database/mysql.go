package databese

import (
	"log"
	"os"
	"redishex/repository"
	"time"

	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func SQLInit(dial gorm.Dialector) (db *gorm.DB, err error) {
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

	mydb.AutoMigrate(&repository.Products{})
	return mydb, nil
}
