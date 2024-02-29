package model

const tblNameUser = "user"

type BaseUser struct{}

func (BaseUser) TableName() string {
	return tblNameUser
}

type User struct {
	BaseUser
	DefaultModel
	Phone    string `gorm:"column:phone;type:varchar(11);unique_index;not null;comment:手机号"`
	Password string `gorm:"column:password;type:varchar(255);not null;comment:密码"`
	Name     string `gorm:"column:name;type:varchar(50);not null;comment:名字"`
	Age      int    `gorm:"column:age;type:int;not null;comment:年龄"`
	Gender   int    `gorm:"column:gender;type:int;not null;comment:性别(1-男，2-女)"`
}

func init() {
	registerModel(&User{})
}
