package models

type Permission struct {
	ID        string `json:"id"`
	Role      string `json:"role"`
	Action    string `json:"action"`
	Verb      string `json:"verb"`
	Path      string `json:"path"`
	CreatedAT string `json:"created_at"`
	UpdatedAT string `json:"updated_at"`
}

func (Permission) TableName() string {
	return "user_service.permissions"
}
