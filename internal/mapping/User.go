package mapping

import (
	"context"
	"encoding/json"
	"errors"
	"game.sdk.center/internal/model/system"
	"game.sdk.center/tool"
	"github.com/go-redis/redis/v8"
	"time"
)

func User() (userMap map[int]string, err error) {

	result, err := tool.RedisClient.Get(context.Background(), "userNameMap").Result()
	if err != nil && !errors.Is(err, redis.Nil) {
		return
	}
	userMap = make(map[int]string)
	if result != "" {
		if err = json.Unmarshal([]byte(result), &userMap); err != nil {
			return
		}

		return
	}

	user := system.NewUser()
	users, err := user.All()
	if err != nil {
		return
	}

	for _, v := range users {
		userMap[int(v.Id)] = v.Nickname
	}
	marshal, err := json.Marshal(&userMap)
	if err != nil {
		return
	}

	err = tool.RedisClient.Set(context.Background(), "userNameMap", marshal, time.Second*60*60*2).Err()

	return
}
