package lib

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var Mysql *gorm.DB

type MysqlConfig struct {
	Host     string
	Port     int
	User     string
	Password string
}

func init() {
	var err error
	Mysql, err = gorm.Open(mysql.Open(getDsn()), &gorm.Config{})
	if err != nil {
		fmt.Println(err)
	}
}

func getDsn() string {
	var config = MysqlConfig{
		Host:     "127.0.0.1",
		Port:     3306,
		User:     "root",
		Password: "root",
	}

	return fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/dbname?charset=utf8mb4&parseTime=True&loc=Local",
		config.User, config.Password, config.Host, &config,
	)
}
