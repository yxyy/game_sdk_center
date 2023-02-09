package mapping

import (
	"game.sdk.center/internal/model/basics"
)

func Company() (map[int]string, error) {
	company := basics.Company{}
	companies, err := company.GetAll()
	if err != nil {
		return nil, err
	}

	companys := make(map[int]string)
	for _, v := range companies {
		companys[v.Id] = v.Name
	}

	return companys, nil
}
