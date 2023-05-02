package database

import (
	"fmt"
	"log"
	"spotless/models"

	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type DBConfig struct {
	Host     string
	Port     string
	User     string
	Pass     string
	Database string
}

var dbConfig *DBConfig
var DB *gorm.DB

func init() {
	viper.SetDefault("database.host", "localhost")
	viper.SetDefault("database.port", "3306")
}

func Load() error {
	viper.SetConfigName("ConfigDB")
	viper.SetConfigType("toml")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()

	if err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			return err
		}
	}

	dbConfig = new(DBConfig)

	dbConfig = &DBConfig{
		Host:     viper.GetString("database.host"),
		Port:     viper.GetString("database.port"),
		User:     viper.GetString("database.user"),
		Pass:     viper.GetString("database.pass"),
		Database: viper.GetString("database.database"),
	}

	return nil
}

func ConnectDatabase() error {
	err := Load()
	if err != nil {
		log.Fatal(err)
	}

	stringConnection := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbConfig.User, dbConfig.Pass, dbConfig.Host, dbConfig.Port, dbConfig.Database)

	db, err := gorm.Open(mysql.Open(stringConnection))
	if err != nil {
		panic(err)
	}
	db.Migrator().DropTable(&models.User{})
	db.AutoMigrate(&models.User{})
	DB = db
	return nil
}
