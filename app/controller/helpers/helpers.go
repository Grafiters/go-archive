package helpers

import (
	"net/url"

	"github.com/gookit/validate"
)

type Errors struct {
	Code   int16    `json:"code"`
	Status bool     `json:"status"`
	Errors []string `json:"errors"`
}

type Response struct {
	Code   int16       `json:"code"`
	Status bool        `json:"status"`
	Data   interface{} `json:"data"`
}

func (e Errors) Size() int {
	return len(e.Errors)
}

func Vaildate(payload interface{}, err_src *Errors) {
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
	if err != nil {
		return false
	}

	return true
}
