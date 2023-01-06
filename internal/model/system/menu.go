package system

import (
	"context"
	"encoding/json"
	"errors"
	"game.sdk.center/internal/model/common"
	"game.sdk.center/tool"
	"time"
)

type Menu struct {
	common.Model
	Title     string `gorm:"title" form:"title" json:"title"`
	Flag      string `gorm:"flag" form:"flag" json:"flag"`
	Parent    int    `gorm:"parent" form:"parent" json:"parent"`
	Path      string `gorm:"path" form:"path" json:"path"`
	Redirect  string `gorm:"path" form:"path" json:"redirect"`
	Component string `gorm:"component" form:"component" json:"component"`
	Icon      string `gorm:"icon" form:"icon" json:"icon"`
	Sort      int    `json:"sort" form:"sort" json:"sort"`
}

// MenuTree
//
//		{
//	   path: '/',
//	   component: Layout,
//	   redirect: '/dashboard',
//	   children: [
//	     {
//	       path: 'dashboard',
//	       component: () => import('@/views/dashboard/index'),
//	       name: 'Dashboard',
//	       meta: { title: 'Dashboard', icon: 'dashboard', affix: true }
//	     }
//	   ]
//	 }
type MenuTree struct {
	Id         int64  `json:"id"`
	Parent     int    `json:"parent"`
	Path       string `json:"path"`
	Redirect   string `json:"redirect"`
	Component  string `json:"component"`
	AlwaysShow bool   `json:"alwaysShow"`
	Name       string `json:"name"`
	Meta       `json:"meta"`
	Children   []*MenuTree `json:"children,omitempty"`
}

type Meta struct {
	Title string `gorm:"title" form:"title" json:"title"`
	Icon  string `gorm:"icon" form:"icon" json:"icon"`
}

func NewMenu() *Menu {
	return &Menu{}
}

func (m *Menu) Create() error {
	if m.Title == "" {
		return errors.New("标题不能空")
	}
	if m.Flag == "" {
		return errors.New("标识不能空")
	}
	if m.Path == "" {
		return errors.New("路径不能空")
	}
	if m.Component == "" {
		return errors.New("组件不能空")
	}

	if err := tool.MysqlDb.Model(m).Where("flag", m.Flag).First(&Menu{}).Error; err == nil {
		return errors.New("标识已存在")
	}

	if err := tool.MysqlDb.Model(m).Where("component", m.Component).First(&Menu{}).Error; err == nil {
		return errors.New("组件已存在")
	}

	if err := tool.MysqlDb.Model(m).Create(m).Error; err != nil {
		return err
	}

	return nil
}

func (m *Menu) Update() error {
	if m.Id <= 0 {
		return errors.New("id 无效")
	}

	if err := tool.MysqlDb.Model(m).Where("id", m.Id).Updates(m).Error; err != nil {
		return err
	}

	return nil
}

func (m *Menu) GetAll() (menus []*Menu, err error) {

	if err = tool.MysqlDb.Model(&m).Where("status", 0).Order("sort desc").Order("created_at desc").Find(&menus).Error; err != nil {
		return
	}

	return
}

func (m *Menu) List(params *common.Params) (menus []*Menu, total int64, err error) {
	tx := tool.MysqlDb.Model(&m)
	if m.Id > 0 {
		tx = tx.Where("id", m.Id)
	}
	if m.Title != "" {
		tx = tx.Where("title", m.Title)
	}
	if m.Flag != "" {
		tx = tx.Where("id", m.Flag)
	}
	if err = tx.Count(&total).Error; err != nil {
		if err != nil {
			return
		}
	}
	if err = tx.Offset(params.Offse).Limit(params.Limit).Find(&menus).Error; err != nil {
		if err != nil {
			return
		}
	}
	return
}

func (m *Menu) GetTree() ([]*MenuTree, error) {

	// result, err := tool.RedisClient.Get(context.Background(), "menus").Result()
	// if err != nil && !errors.Is(err, redis.Nil) {
	// 	return nil, err
	// }
	//
	var trees []*MenuTree
	// if result != "" {
	// 	if err = json.Unmarshal([]byte(result), &trees); err != nil {
	// 		return nil, err
	// 	}
	// 	return trees, err
	// }

	menus, err := m.GetAll()
	if err != nil {
		return nil, err
	}

	trees = tree(menus, 0)
	jsonTress, err := json.Marshal(&trees)
	if err != nil {
		return nil, err
	}

	if err = tool.RedisClient.Set(context.Background(), "menus", string(jsonTress), time.Second*2*360).Err(); err != nil {
		return nil, err
	}

	return trees, nil
}

func tree(menus []*Menu, pid int) []*MenuTree {
	var Trees []*MenuTree
	for _, v := range menus {
		if v.Parent == pid {
			menuTree := &MenuTree{
				Id:         v.Id,
				Parent:     v.Parent,
				Path:       v.Path,
				Redirect:   v.Redirect,
				Component:  v.Component,
				Name:       v.Title,
				AlwaysShow: true,
				Meta: Meta{
					Title: v.Title,
					Icon:  v.Icon,
				},
				Children: tree(menus, int(v.Id)),
			}

			Trees = append(Trees, menuTree)
		}
	}

	return Trees
}
