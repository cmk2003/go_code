package student

import (
	"ginEssential/common"
	"ginEssential/model"
)

type DAO struct {
}

func (d DAO) GetStudentListByID(id uint) ([]model.Student, error) {
	db := common.GetDB()
	var students []model.Student
	result := db.Where("teacher=?", id).Find(&students)
	return students, result.Error

}

func (d DAO) Add(student model.Student) error {
	db := common.GetDB()
	//var students []model.Student
	if err := db.Create(&student).Error; err != nil {
		return err
	}
	return nil

}
