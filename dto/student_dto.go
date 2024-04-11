package dto

import "ginEssential/model"

type StudentDto struct {
	Name      string `json:"name"`
	StudentId string `json:"studentId"`
	Teacher   string `json:"teacher"`
}

func ToStudentDto(student *model.Student) StudentDto {
	return StudentDto{
		Name:      student.Name,
		StudentId: student.StudentId,
		Teacher:   student.Teacher,
	}
}
