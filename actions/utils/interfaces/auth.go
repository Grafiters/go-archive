package interfaces

type LoginGoogle struct {
	RedirectUrl string `json:"redirect_url" binding:"required" swag:"form"`
	Code        string `json:"code" binding:"required" swag:"form"`
}

type GoogleData struct {
	Email     string `json:"email" binding:"optional" swag:"form"`
	Name      string `json:"name" binding:"optional" swag:"form"`
	GivenName string `json:"given_name" binding:"optional" swag:"form"`
	Picture   string `json:"picture" binding:"optional" swag:"form"`
}
