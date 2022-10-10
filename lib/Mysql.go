package lib

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var MysqlDb *gorm.DB

type MysqlConfig struct {
	Host     string
	Port     int
	User     string
	Password string
}

func init() {
	var err error
	MysqlDb, err = gorm.Open(mysql.Open(getDsn()), &gorm.Config{})
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
		"%s:%s@tcp(%s:%d)/dbname?charset=utf8mb4&parseTime=True&loc=Local",
		config.User, config.Password, config.Host, config.Port,
	)
}

// Mysql TODO 根据配置不同选择不同实例
func Mysql(config string) *gorm.DB {
	if config == "" {
		config = "default"
	}
	db, err := gorm.Open(mysql.Open(getDsn()), &gorm.Config{})
	if err != nil {
		fmt.Println(err)
	}
	return db
}
