package api

import (
	"errors"
	"github.com/atpuxiner/grapi/app/datatype/entity"
	"github.com/gin-gonic/gin"
	"strconv"
)

func (r BaseApi) GetParamInt(c *gin.Context, key string) (int, error) {
	res, err := strconv.Atoi(c.Param(key))
	if err != nil {
		return 0, err
	}
	return res, nil
}

func (r BaseApi) GetParamUint(c *gin.Context, key string) (uint, error) {
	res, err := strconv.ParseUint(c.Param(key), 10, 64)
	if err != nil {
		return 0, err
	}
	return uint(res), nil
}

func (r BaseApi) GetQueryPageInfo(c *gin.Context) (entity.PageInfo, error) {
	var pinfo entity.PageInfo
	page, err := strconv.Atoi(c.DefaultQuery("page", "1"))
	if err != nil {
		return pinfo, err
	}
	pageSize, err := strconv.Atoi(c.DefaultQuery("pageSize", "10"))
	if err != nil {
		return pinfo, err
	}
	if page < 0 || pageSize < 0 {
		err = errors.New("page or pageSize should be greater than 0")
		return pinfo, err
	}
	pinfo.Page = page
	pinfo.PageSize = pageSize
	return pinfo, nil
}
