package database

import (
	"fmt"

	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	Db  *gorm.DB
	err error
)

//SetupDatabase sets the DB configuration
func SetupDatabase() error {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		viper.Get("DB.USER"),
		viper.Get("DB.PASSWORD"),
		viper.Get("DB.HOST"),
		viper.Get("DB.PORT"),
		viper.Get("DB.NAME"))
	Db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return err
	}
	return nil
}
