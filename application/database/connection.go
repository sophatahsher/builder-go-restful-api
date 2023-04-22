package database

import (
	util "builder/restful-api-gogin/utils"
	"os"

	"github.com/sirupsen/logrus"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Connection() *gorm.DB {
	databaseURI := make(chan string, 1)

	if os.Getenv("GO_ENV") != "production" {
		databaseURI <- util.GodotEnv("DATABASE_URI_DEV")
	} else {
		databaseURI <- os.Getenv("DATABASE_URI_PROD")
	}

	//dsn := "host=localhost user=postgres password=postgres dbname=acl_db port=5432 sslmode=disable TimeZone=Asia/Phnom_Penh"
	//db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	db, err := gorm.Open(postgres.Open(<-databaseURI), &gorm.Config{})

	if err != nil {
		defer logrus.Info("Connection to Database Failed")
		logrus.Fatal(err.Error())
	}

	if os.Getenv("GO_ENV") != "production" {
		logrus.Info("Connection to Database Successfully")
	}

	/*
		err = db.AutoMigrate(
			&model.EntityUsers{},
			&model.EntityStudent{},
		)
	*/
	if err != nil {
		logrus.Fatal(err.Error())
	}

	return db
}
