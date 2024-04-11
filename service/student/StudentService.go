package student

import (
	"errors"
	"ginEssential/common"
	"ginEssential/dao/student"
	"ginEssential/dto"
	"ginEssential/model"
	"github.com/jinzhu/gorm"
)

type Service struct {
	studentDAO student.DAO
}

func (s *Service) GetStudentListByID(id uint) ([]dto.StudentDto, error) {
	studentList, err := s.studentDAO.GetStudentListByID(id)
	if err != nil {
		return nil, err
	}
	// 转换为DTO列表
	studentDtos := make([]dto.StudentDto, len(studentList))
	for i, student_ := range studentList {
		studentDtos[i] = dto.ToStudentDto(&student_)
	}
	return studentDtos, nil
}
func isStudentIdExist(db *gorm.DB, studentId string) bool {
	var s model.Student
	db.Where("student_id=?", studentId).First(&s)
	return s.ID != 0
}

func (s *Service) Add(student model.Student) error {
	db := common.GetDB()
	//判断studentId是否唯一
	studentId := student.StudentId
	if isStudentIdExist(db, studentId) {
		//创建error，抛出去
		return errors.New("学生id已经存在，请勿重复添加")
	}
	err := s.studentDAO.Add(student)
	return err
}
