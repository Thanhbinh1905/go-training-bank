package api

import (
	"github.com/Thanhbinh1905/go-training-bank/pkg/util"
	"github.com/go-playground/validator/v10"
)

var validCurrency validator.Func = func(fl validator.FieldLevel) bool {
	if currency, ok := fl.Field().Interface().(string); ok {
		// check is currency supported
		return util.IsSupportedCurrency(currency)
	}
	return false
}
