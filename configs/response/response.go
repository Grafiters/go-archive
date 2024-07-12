package response

type Errors struct {
	Code   int16    `json:"code"`
	Status bool     `json:"status"`
	Errors []string `json:"errors"`
}

type Success struct {
	Code   int16       `json:"code"`
	Status bool        `json:"status"`
	Data   interface{} `json:"data"`
}

func (e Errors) Size() int {
	return len(e.Errors)
}
