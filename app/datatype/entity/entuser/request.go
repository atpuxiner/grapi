package entuser

import (
	"github.com/atpuxiner/grapi/app/datatype/model"
	"github.com/atpuxiner/grapi/app/utils"
	"gorm.io/gorm"
	"time"
)

type GetListParams struct {
	model.BaseUser
	Phone  string `form:"phone" validate:"max=11"`
	Name   string `form:"name" validate:"max=50"`
	Age    int    `form:"age" validate:"max=200"`
	Gender int    `form:"gender" validate:"oneof=0 1 2" oneof_err:"只支持: 0,1,2"`
}

type CreateBody struct {
	model.BaseUser
	ID        uint
	Phone     string `json:"phone" validate:"required,phone"`
	Password  string `json:"password" validate:"required,password"`
	Name      string `json:"name" validate:"required,max=50"`
	Age       int    `json:"age" validate:"required,max=200"`
	Gender    int    `json:"gender" validate:"required,oneof=0 1 2" oneof_err:"只支持: 0,1,2"`
	CreatedAt time.Time
}

func (r *CreateBody) BeforeCreate(_ *gorm.DB) error {
	hashedPassword, err := utils.HashPassword(r.Password)
	r.Password = hashedPassword
	return err
}

type UpdateBody struct {
	model.BaseUser
	ID        uint
	Name      string `json:"name" validate:"max=50"`
	Age       int    `json:"age" validate:"max=200"`
	Gender    int    `json:"gender" validate:"oneof=0 1 2" oneof_err:"只支持: 0,1,2"`
	UpdatedAt time.Time
}
