package models

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"os"
	"log"
// 	"gorm.io/gorm/schema"
)

var DB *gorm.DB

func ConnectDatabase() {
    var (
        dbHost     = os.Getenv("DB_HOST")
        dbUserName = os.Getenv("DB_USERNAME")
        dbDatabase = os.Getenv("DB_DATABASE")
        dbPassword = os.Getenv("DB_PASSWORD")
    )
    newLogger := logger.New(
        log.New(os.Stdout, "\r\n", log.LstdFlags), // выводим лог в консоль
        logger.Config{
            LogLevel: logger.Info, // уровень отладки
        },
    )

	db, err := gorm.Open(mysql.Open("" + dbUserName + ":" + dbPassword + "@tcp(" + dbHost + ":3306)/" + dbDatabase + ""),
	 &gorm.Config{
          Logger: newLogger,
//           NamingStrategy: schema.NamingStrategy{
//             NoLowerCase: true,
//             SingularTable: true,
//           },
    })
	if err != nil {
		panic(err)
	}

// 	db.AutoMigrate(&Studio{})
// 	db.Migrator().DropColumn(&Setting{}, "name_two")
	DB = db
}



