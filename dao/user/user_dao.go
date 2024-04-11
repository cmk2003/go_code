package user

import (
	"ginEssential/common"
	"ginEssential/model"
)

type DAO struct {
}

func (dao *DAO) GetAllUsers() ([]model.User, error) {
	db := common.GetDB()
	var users []model.User
	// 假设你有一个已经初始化的GORM DB连接对象 db
	result := db.Find(&users)
	return users, result.Error
}
