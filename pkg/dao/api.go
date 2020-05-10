package dao

import (
	"apimake/cmd/storage"
	"apimake/pkg/entity"
)

type ApiDao struct {
}

func Api() *ApiDao {
	return &ApiDao{}
}

func (dao ApiDao) GetItems() []entity.Api {
	var items []entity.Api

	if err := storage.Read(entity.TableApi, &items); err != nil {
		return nil
	}

	return items
}

func (dao ApiDao) Get(id int) *entity.Api {
	if id < 1 {
		return nil
	}
	items := dao.GetItems()
	if len(items) < id {
		return nil
	}

	item := items[id-1]
	return &item
}

func (dao ApiDao) Create(api entity.Api) error {
	items := dao.GetItems()
	if items == nil {
		items = make([]entity.Api, 0)
	}

	items = append(items, api)
	return storage.Write(api.TableName(), items)
}

func (dao ApiDao) Update(id int, item *entity.Api) error {
	items := dao.GetItems()
	items[id-1] = *item
	return storage.Write(item.TableName(), items)
}
