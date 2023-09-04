package utils

import (
	"reflect"

	CN_ZH "github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	zhTranslations "github.com/go-playground/validator/v10/translations/zh"
	"github.com/sirupsen/logrus"
)

var Validate *validator.Validate

// Validate/v10 全局验证器
var trans ut.Translator

// 初始化Validate/v10国际化
func init() {
	zh := CN_ZH.New()
	uni := ut.New(zh, zh)
	trans, _ = uni.GetTranslator("zh")

	Validate = validator.New()

	//通过label标签返回自定义错误内容
	// Validate.RegisterTagNameFunc(func(field reflect.StructField) string {
	// 	label := field.Tag.Get("label")
	// 	if label == "" {
	// 		return field.Name
	// 	}
	// 	return label
	// })
	zhTranslations.RegisterDefaultTranslations(Validate, trans)

	//自定义required_if错误内容
	// Validate.RegisterTranslation("required_if", trans, func(ut ut.Translator) error {
	// 	return ut.Add("required_if", "{0}为必填字段!", false) // see universal-translator for details
	// }, func(ut ut.Translator, fe validator.FieldError) string {
	// 	t, _ := ut.T("required_if", fe.Field())
	// 	return t
	// })

}

func processErr(u interface{}, err error) string {
	if err == nil { //如果为nil 说明校验通过
		return ""
	}

	invalid, ok := err.(*validator.InvalidValidationError) //如果是输入参数无效，则直接返回输入参数错误
	if ok {
		return "输入参数错误：" + invalid.Error()
	}
	validationErrs := err.(validator.ValidationErrors) //断言是ValidationErrors
	for _, validationErr := range validationErrs {
		fieldName := validationErr.Field() //获取是哪个字段不符合格式
		logrus.Info("filedName", fieldName)
		field, ok := reflect.TypeOf(u).FieldByName(fieldName) //通过反射获取filed
		if ok {
			errorInfo := field.Tag.Get("field_error_info") //获取field对应的reg_error_info tag值
			if errorInfo == "" {
				return validationErr.Translate(trans) //如果reg_error_info tag为空，则直接返回错误信息
			}
			return errorInfo //返回错误
		} else {
			return "缺失field_error_info"
		}
	}
	return ""
}

// 检验并返回检验错误信息
func GetErrorMessage(u interface{}) string {
	err := Validate.Struct(u)
	if err != nil {
		return processErr(u, err)
	}
	return ""
}
