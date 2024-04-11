package model

import "github.com/jinzhu/gorm"

type Student struct {
	gorm.Model
	Name      string `gorm:"type:varchar(20);not null"`
	StudentId string `gorm:"type:varchar(20);not null;unique"` //学生卡号
	Teacher   string `gorm:"type:varchar(20);not null"`        //user的name
	Grade     string `gorm:"type:varchar(20);not null"`        //年级 1 2 3 4 5
}
