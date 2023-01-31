package mapping

import (
	"errors"
	"game.sdk.center/internal/model/system"
	"github.com/go-redis/redis/v8"
)

const GroupCacheKey = "group_list_cache"

func Group() (map[int]string, error) {

	data, err := Get(GroupCacheKey)
	if err != nil && !errors.Is(err, redis.Nil) {
		return nil, err
	}
	if data != nil {
		return data, err
	}

	group := system.NewGroup()

	groups, err := group.GetAll()
	if err != nil {
		return nil, err
	}

	data = make(map[int]string)
	for _, v := range groups {
		data[v.Id] = v.Name
	}

	if err = Set(GroupCacheKey, data); err != nil {
		return nil, err
	}

	return data, nil

}
