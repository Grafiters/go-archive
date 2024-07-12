package models

import (
	"time"

	"github.com/Grafiters/archive/configs/types"
)

type User struct {
	ID        int              `json:"id" gorm:"column:id;type:bigserial;primaryKey;autoIncrement"`
	UID       string           `json:"uid"`
	Email     string           `json:"email" gorm:"unique" validate:"required"`
	GoogleID  string           `json:"google_id"`
	Role      types.MemberRole `json:"role" gorm:"default:'member'"`
	Password  string           `json:"password"`
	CreatedAT time.Time        `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAT time.Time        `json:"updated_at" gorm:"autoUpdateTime"`
}

func (User) TableName() string {
	return "user_service.users"
}
