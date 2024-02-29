// 注册到 api/api_validate.go registerCustomValidate

package entuser

import (
	"github.com/go-playground/validator/v10"
	"regexp"
	"strings"
)

func ValidPhone(fl validator.FieldLevel) bool {
	re := regexp.MustCompile(`^(\+86|0086)?1[3-9]\d{9}$`)
	return re.MatchString(fl.Field().String())
}

func ValidPassword(fl validator.FieldLevel) bool {
	// 密码：8-20位，只包含字母、数字、指定的特殊字符，且至少有一个字母和一个数字
	password := fl.Field().String()
	re := regexp.MustCompile(`^[a-zA-Z0-9!@#$%^&*()]{8,20}$`)
	if !re.MatchString(password) {
		return false
	}
	hasLetter := strings.ContainsAny(password, "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	hasNumber := strings.ContainsAny(password, "0123456789")
	return hasLetter && hasNumber
}
