package mapping

import (
	"game.sdk.center/internal/model/basics"
)

func Apps() (map[int]string, error) {
	app := basics.App{}
	apps, err := app.GetAll()
	if err != nil {
		return nil, err
	}

	appMap := make(map[int]string)
	for _, v := range apps {
		appMap[v.Id] = v.Name
	}

	return appMap, nil
}

func AppType() (map[int]string, error) {
	appType := basics.AppType{}
	companies, err := appType.GetAll()
	if err != nil {
		return nil, err
	}

	appTypes := make(map[int]string)
	for _, v := range companies {
		appTypes[v.Id] = v.Name
	}

	return appTypes, nil
}
