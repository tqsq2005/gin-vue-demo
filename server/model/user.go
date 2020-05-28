package model

import "github.com/jinzhu/gorm"

type User struct {
	gorm.Model
	Name     string `form:"name" gorm:"type:varchar(30); not null" valid:"alphanum,required,stringlength(6|30)"`
	Password string `form:"password" gorm:"type:varchar(191); not null" valid:"required,minstringlength(6)"`
	Tel      string `form:"tel" gorm:"type:varchar(20); not null; unique" valid:"numeric,required"`
}
