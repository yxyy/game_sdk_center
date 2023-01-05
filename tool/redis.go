package tool

import (
	"context"
	"errors"
	"github.com/go-redis/redis/v8"
	"github.com/spf13/viper"
	"strconv"
)

var RedisClient *redis.Client

var RedisClientPoll = make(map[int]*redis.Client)

func InitRedis() error {
	RedisClient = redis.NewClient(&redis.Options{
		Addr:     viper.GetString("redis.host") + ":" + viper.GetString("redis.port"),
		Password: viper.GetString("redis.password"),
		DB:       viper.GetInt("redis.db"),
	})
	_, err := RedisClient.Ping(context.Background()).Result()
	if err != nil {
		return err
	}

	return nil
}

func InitRedisPoll() error {
	redisPoll := viper.GetStringMap("redisPoll")
	for k, v := range redisPoll {
		db, err := strconv.Atoi(v.(map[string]string)["db"])
		if err != nil {
			return errors.New(k + ":" + err.Error())
		}
		RedisClientPoll[db] = redis.NewClient(&redis.Options{
			Addr:     v.(map[string]string)["host"] + ":" + v.(map[string]string)["port"],
			Password: v.(map[string]string)["password"],
			DB:       db,
		})

		_, err = RedisClientPoll[db].Ping(context.Background()).Result()
		if err != nil {
			return errors.New(k + ":" + err.Error())
		}
	}

	return nil
}
