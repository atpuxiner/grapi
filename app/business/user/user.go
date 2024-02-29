package user

import (
	"github.com/atpuxiner/grapi/app/datatype/entity"
	"github.com/atpuxiner/grapi/app/datatype/entity/entuser"
	"github.com/atpuxiner/grapi/app/initializer/db"
	"github.com/atpuxiner/grapi/app/utils/errs"
	"gorm.io/gorm"
)

type Business struct{}

func NewBusiness() Business {
	return Business{}
}

func (r Business) Get(id uint) (any, error) {
	var user entuser.GetResp
	if err := db.DB.First(&user, id).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func (r Business) GetList(req entuser.GetListParams, pinfo entity.PageInfo) (any, error) {
	var (
		userList []entuser.GetListResp
		total    int64
	)
	tx := db.DB.Model(entuser.User{})
	if req.Phone != "" {
		tx.Where("phone LIKE ?", "%"+req.Phone+"%")
	}
	if req.Name != "" {
		tx.Where("LOWER(name) LIKE ?", "%"+req.Name+"%")
	}
	if req.Age != 0 {
		tx.Where("age >= ?", req.Age)
	}
	if req.Gender != 0 {
		tx.Where("gender = ?", req.Gender)
	}
	if err := tx.Count(&total).Limit(pinfo.Limit()).Offset(pinfo.Offset()).Find(&userList).Error; err != nil {
		return nil, err
	}
	if tx.RowsAffected == 0 {
		return nil, gorm.ErrRecordNotFound
	}
	res := map[string]any{
		"data":     userList,
		"total":    total,
		"page":     pinfo.Page,
		"pageSize": pinfo.PageSize,
	}
	return res, nil
}

func (r Business) Create(req *entuser.CreateBody) (any, error) {
	var user entuser.User
	if err := db.DB.Model(&user).Where("phone = ?", req.Phone).First(&user).Error; err == nil {
		return nil, errs.ErrRecordExisted
	}
	if err := db.DB.Create(req).Error; err != nil {
		return nil, err
	}
	return req, nil
}

func (r Business) Update(id uint, req *entuser.UpdateBody) (any, error) {
	var user entuser.User
	if err := db.DB.First(&user, id).Error; err != nil {
		return nil, err
	}
	req.ID = id
	if err := db.DB.Updates(req).Error; err != nil {
		return nil, err
	}
	return req, nil
}

func (r Business) Delete(id uint) (any, error) {
	var user entuser.User
	if err := db.DB.First(&user, id).Error; err != nil {
		return nil, err
	}
	if err := db.DB.Delete(&user).Error; err != nil {
		return nil, err
	}
	return user, nil
}
