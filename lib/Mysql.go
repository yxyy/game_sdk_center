package lib

import (
	"fmt"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var MysqlDb *gorm.DB

var Db = make(map[string]*gorm.DB)

type MysqlConfig struct {
	Host     string
	Port     int
	User     string
	Password string
	Database string
}

func InitMysql() (err error) {
	MysqlDb, err = gorm.Open(mysql.Open(getDsn()), &gorm.Config{})
	if err != nil {
		return err
	}

	return nil
}

func getDsn() string {
	var config = MysqlConfig{
		Host:     viper.GetString("mysql.master.host"),
		Port:     viper.GetInt("mysql.master.port"),
		User:     viper.GetString("mysql.master.user"),
		Password: viper.GetString("mysql.master.password"),
		Database: viper.GetString("mysql.master.database"),
	}

	return fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		config.User, config.Password, config.Host, config.Port, config.Database,
	)
}
func getDsnConfig(conf string) string {
	mysqlMap := viper.GetStringMap("mysql." + conf)
	return fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		mysqlMap["user"], mysqlMap["password"], mysqlMap["host"], mysqlMap["port"], mysqlMap["database"],
	)
}

// Mysql TODO 根据配置不同选择不同实例
func Mysql(config string) *gorm.DB {
	if config == "" {
		config = "master"
	}
	if Db[config] != nil {
		return Db[config]
	}
	db, err := gorm.Open(mysql.Open(getDsnConfig(config)), &gorm.Config{})
	if err != nil {
		fmt.Println(err)
	}
	Db[config] = db

	return Db[config]
}
