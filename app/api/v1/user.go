package v1

import (
	"github.com/atpuxiner/grapi/app/api"
	"github.com/atpuxiner/grapi/app/business/user"
	"github.com/atpuxiner/grapi/app/datatype/entity/entuser"
	"github.com/atpuxiner/grapi/app/utils/status"
	"github.com/gin-gonic/gin"
)

type UserApi struct {
	api.BaseApi
}

func NewUserApi() UserApi {
	return UserApi{}
}

var userBusiness = user.NewBusiness()

// @Tags Api.User
// @Summary 查询user详情
// @Description 查询user详情
// @Param id path uint true "业务id"
// @Success 0 {object} api.ResponseJson "请求成功"
// @Failure 1 {object} api.ResponseJson "请求失败"
// @Router /api/v1/user/{id} [get]
func (r UserApi) Get(c *gin.Context) {
	// 参数处理
	id, err := r.GetParamUint(c, "id")
	if err != nil {
		r.Failure(c, status.CodeBadParam, err)
		return
	}
	// 数据处理
	data, err := userBusiness.Get(id)
	r.Response(c, data, err)
}

// @Tags Api.User
// @Summary 查询user列表
// @Description 查询user列表
// @Param phone formData string false "手机号"
// @Param name formData string false "名字"
// @Param age formData int false "年龄"
// @Param gender formData int false "性别"
// @Success 0 {object} api.ResponseJson "请求成功"
// @Failure 1 {object} api.ResponseJson "请求失败"
// @Router /api/v1/user [get]
func (r UserApi) GetList(c *gin.Context) {
	// 参数处理
	var req entuser.GetListParams
	err := r.ValidateWithBindQuery(c, &req)
	if err != nil {
		r.Failure(c, status.CodeBadParam, err)
		return
	}
	pinfo, err := r.GetQueryPageInfo(c)
	if err != nil {
		r.Failure(c, status.CodeBadParam, err)
		return
	}
	// 数据处理
	data, err := userBusiness.GetList(req, pinfo)
	r.Response(c, data, err)
}

// @Tags Api.User
// @Summary 新增user
// @Description 新增user
// @Param body body entuser.CreateBody true "body"
// @Success 0 {object} api.ResponseJson "请求成功"
// @Failure 1 {object} api.ResponseJson "请求失败"
// @Router /api/v1/user [post]
func (r UserApi) Create(c *gin.Context) {
	// 参数处理
	var req entuser.CreateBody
	err := r.ValidateWithBindJson(c, &req)
	if err != nil {
		r.Failure(c, status.CodeBadParam, err)
		return
	}
	// 数据处理
	data, err := userBusiness.Create(&req)
	r.Response(c, data, err)
}

// @Tags Api.User
// @Summary 更改user
// @Description 更改user
// @Param id path uint true "业务id"
// @Param body body entuser.UpdateBody true "body"
// @Success 0 {object} api.ResponseJson "请求成功"
// @Failure 1 {object} api.ResponseJson "请求失败"
// @Router /api/v1/user/{id} [put]
func (r UserApi) Update(c *gin.Context) {
	// 参数处理
	var req entuser.UpdateBody
	id, err := r.GetParamUint(c, "id")
	if err != nil {
		r.Failure(c, status.CodeBadParam, err)
		return
	}
	err = r.ValidateWithBindJson(c, &req)
	if err != nil {
		r.Failure(c, status.CodeBadParam, err)
		return
	}
	// 数据处理
	data, err := userBusiness.Update(id, &req)
	r.Response(c, data, err)
}

// @Tags Api.User
// @Summary 删除user
// @Description 删除user
// @Param id path uint true "业务id"
// @Success 0 {object} api.ResponseJson "请求成功"
// @Failure 1 {object} api.ResponseJson "请求失败"
// @Router /api/v1/user/{id} [delete]
func (r UserApi) Delete(c *gin.Context) {
	// 参数处理
	id, err := r.GetParamUint(c, "id")
	if err != nil {
		r.Failure(c, status.CodeBadParam, err)
		return
	}
	// 数据处理
	data, err := userBusiness.Delete(id)
	r.Response(c, data, err)
}
