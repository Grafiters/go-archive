package models

import "time"

type User struct {
	ID        int       `json:"id" gorm:"column:id;type:bigserial;primaryKey;autoIncrement"`
	UID       string    `json:"uid"`
	Email     string    `json:"email"`
	GoogleID  string    `json:"google_id"`
	Password  string    `json:"password"`
	CreatedAT time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAT time.Time `json:"updated_at" gorm:"autoUpdateTime"`
}
