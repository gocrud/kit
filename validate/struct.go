package validate

import (
	"errors"

	"github.com/go-playground/validator/v10"
)

type Validator interface {
	ValidateInfo() map[string]string
}

func Struct(v Validator) error {
	err := validate.Struct(v)
	if err == nil {
		return nil
	}
	var errs validator.ValidationErrors
	if ok := errors.As(err, &errs); ok {
		infos := v.ValidateInfo()
		for _, e := range errs {
			key := e.Field() + "." + e.Tag()
			if tip, ok := infos[key]; ok {
				return newError(e.Field(), e.Tag(), tip)
			} else {
				return newError(e.Field(), e.Tag(), e.Field())
			}
		}
	}
	return err
}
