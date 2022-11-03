package config

import (
	"fmt"
	"log"

	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/rest-api/models"

	"github.com/jinzhu/gorm"
)

var (
	db  *gorm.DB
	err error
)

// Connection create database connection
func GetDb(cfg Config) (*gorm.DB, error) {

	// mainDSN := fmt.Sprintf(
	// 	"host=%s user=%s password=%s dbname=%s port=%s sslmode=%s",
	// 	cfg.Database.Host, cfg.Database.Username, cfg.Database.Password, cfg.Database.Name, cfg.Database.Port, cfg.Database.SSLMode,
	// )
	// replicaDSN := fmt.Sprintf(
	// 	"host=%s user=%s password=%s dbname=%s port=%s sslmode=%s",
	// 	cfg.Database.Host, cfg.Database.Username, cfg.Database.Password, cfg.Database.Name, cfg.Database.Port, cfg.Database.SSLMode,
	// )

	// logMode := cfg.Database.LogMode
	// debug := cfg.Server.Debug

	// loglevel := logger.Silent
	// if logMode {
	// 	loglevel = logger.Info
	// }

	// db, err = gorm.Open(postgres.Open(mainDSN), &gorm.Config{
	// 	Logger: logger.Default.LogMode(loglevel),
	// })
	// if !debug {
	// 	db.Use(dbresolver.Register(dbresolver.Config{
	// 		Replicas: []gorm.Dialector{
	// 			postgres.Open(replicaDSN),
	// 		},
	// 		Policy: dbresolver.RandomPolicy{},
	// 	}))
	// }

	// App.Config.Database.DatabaseConfiguration.Host

	fmt.Println("Heeeeee")
	fmt.Println(cfg.Database.UserName)

	DBURL := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", cfg.Database.UserName, cfg.Database.Password, cfg.Database.Host, cfg.Database.Port, cfg.Database.Name)

	driver := cfg.Database.DbDriver

	db, err = gorm.Open(driver, DBURL)

	if err != nil {
		fmt.Println("Cannot connect to database ", driver)
		log.Fatal("connection error:", err)
	} else {
		fmt.Println("We are connected to the database ", driver)
	}

	db.AutoMigrate(&models.User{})

	// Migrate tables and data
	//	err = db.AutoMigrate(&models.User{})
	if err != nil {
		log.Fatalf("database migration error: %s", err)
	}

	return db, nil
}
