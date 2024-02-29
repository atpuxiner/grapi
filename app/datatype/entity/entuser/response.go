package entuser

import (
	"github.com/atpuxiner/grapi/app/datatype/entity"
	"github.com/atpuxiner/grapi/app/datatype/model"
)

type GetResp struct {
	model.BaseUser
	entity.DefaultEntityWithoutDlt
	Phone  string `json:"phone"`
	Name   string `json:"name"`
	Age    int    `json:"age"`
	Gender int    `json:"gender"`
}

type GetListResp struct {
	model.BaseUser
	entity.DefaultEntityWithoutDlt
	Phone  string `json:"phone"`
	Name   string `json:"name"`
	Age    int    `json:"age"`
	Gender int    `json:"gender"`
}
