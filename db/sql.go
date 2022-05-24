package db

import (
	"chemex/config"
	"chemex/model"
	"fmt"
	"github.com/jinzhu/gorm"
	"log"

	_ "github.com/go-sql-driver/mysql"
)
var MySQL *gorm.DB

func Connect() error {
	log.Println("start connecting to Database...")
	connectionString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true",
		config.SQLUserName, config.SQLPassword, config.SQLHost, config.SQLPort, config.SQLDBName)
	var err error
	if MySQL == nil {
		MySQL, err = gorm.Open("mysql", connectionString)
		fmt.Println("mysql connected")
		if err != nil {
			// implement retry mechanism
			log.Println("Error on connecting to mysql")
			log.Fatal(err)
			return err
		}
	}
	return nil
}

func AutoMigrate() error{
	err := Connect()
	if err != nil {
		// retry mechanism
		log.Fatal(err)
		return err
	}
	MySQL.AutoMigrate(model.Garden{})
	MySQL.AutoMigrate(model.Tree{})
	MySQL.AutoMigrate(model.Comment{})
    MySQL.AutoMigrate(model.Tags{})
	MySQL.AutoMigrate(model.User{})
	MySQL.AutoMigrate(model.GardenLocation{})
 return nil
}

