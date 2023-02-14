package tools

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/go-playground/validator/v10"
)

func ValidateStruct(s interface{}) error {
	validate := validator.New()
	validate.RegisterTagNameFunc(func(fld reflect.StructField) string {
		name := strings.SplitN(fld.Tag.Get("json"), ",", 2)[0]
		if name == "-" {
			return ""
		}
		return name
	})
	if err := validate.Struct(s); err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			return fmt.Errorf("Field validation for '%s' failed on the '%s' tag", err.Field(), err.Tag())
		}
	}
	return nil
}
