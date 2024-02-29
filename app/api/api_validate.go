package api

import (
	"errors"
	"github.com/atpuxiner/grapi/app/datatype/entity/entuser"
	"github.com/atpuxiner/grapi/app/utils"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	"reflect"
	"strings"
	"sync"
)

var (
	once     sync.Once
	validate *validator.Validate
)

type CustomValidateMap map[string]validator.Func

func (r BaseApi) ValidateWithBindQuery(c *gin.Context, obj any) error {
	err := c.ShouldBindWith(obj, binding.Query)
	if err != nil {
		return err
	}
	return r.ValidateStruct(obj)
}

func (r BaseApi) ValidateWithBindJson(c *gin.Context, obj any) error {
	err := c.ShouldBindWith(obj, binding.JSON)
	if err != nil {
		return err
	}
	return r.ValidateStruct(obj)
}

func (r BaseApi) ValidateStruct(obj any) error {
	if obj == nil {
		return nil
	}
	lazyValidator()
	err := validate.Struct(obj)
	if err == nil {
		return nil
	}
	return r.parseValidateErrs(err, obj)
}

func (r BaseApi) parseValidateErrs(err error, obj any) error {
	if validErrs, ok := err.(validator.ValidationErrors); ok {
		var (
			errMsgs error
			errMsg  string
		)
		structElem := reflect.TypeOf(obj).Elem()
		for _, fieldErr := range validErrs {
			structField, _ := structElem.FieldByName(fieldErr.Field())
			if errTagMsg := structField.Tag.Get(fieldErr.Tag() + "_err"); errTagMsg != "" {
				errMsg = errTagMsg
			} else if fieldErr.Tag() == "required" {
				errMsg = fieldErr.Field() + ": 不能为空"
			} else if msgTagMsg := structField.Tag.Get("msg"); msgTagMsg != "" {
				errMsg = msgTagMsg
			} else {
				errMsg = fieldErr.Field() + ": 格式不正确"
			}
			if !strings.HasPrefix(errMsg, fieldErr.Field()) {
				errMsg = fieldErr.Field() + ": " + errMsg
			}
			errMsgs = utils.AppendErr(errMsgs, errors.New(errMsg))
		}
		return errMsgs
	}
	return err
}

func lazyValidator() {
	once.Do(func() {
		validate = validator.New()
		// register
		customValidateMap := registerCustomValidate()
		for tag, fn := range customValidateMap {
			_ = validate.RegisterValidation(tag, fn)
		}
	})
}

func registerCustomValidate() CustomValidateMap {
	return CustomValidateMap{
		"phone":    entuser.ValidPhone,
		"password": entuser.ValidPassword,
	}
}
