package entuser

import (
	"github.com/atpuxiner/grapi/app/datatype/entity"
	"github.com/atpuxiner/grapi/app/datatype/model"
)

type User struct {
	model.BaseUser
	entity.DefaultEntity
	Phone    string `json:"phone"`
	Password string `json:"password"`
	Name     string `json:"name"`
	Age      int    `json:"age"`
	Gender   int    `json:"gender"`
}
