package confs

import "github.com/spf13/viper"

func InitConf() error {
	//设置配置文件名称
	viper.SetConfigName("confs")
	//	设置配置文件类型
	viper.SetConfigType("yaml")
	// 设置配置文件路径
	viper.AddConfigPath(".")
	viper.AddConfigPath("./confs")

	// 读取配置文件
	if err := viper.ReadInConfig(); err != nil {
		return err
	}

	return nil
}
