package helpers

import (
	"net/url"

	"github.com/Grafiters/archive/configs/response"
	"github.com/gookit/validate"
)

func Vaildate(payload interface{}, err_src *response.Errors) {
	v := validate.Struct(payload)
	if !v.Validate() {
		for _, errs := range v.Errors.All() {
			for _, err := range errs {
				err_src.Errors = append(err_src.Errors, err)
			}
		}
	}
}

func IsURL(urlCheck string) bool {
	_, err := url.ParseRequestURI(urlCheck)
	if err == nil {
		return true
	} else {
		return false
	}
}
