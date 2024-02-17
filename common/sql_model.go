package common

import "time"

type SQLModel struct {
	Id int `json:"id" gorm:column"id;"`
	// FakeId    *UID       `json:"id" gorm:column"-;"`
	Status    int        `json:"status" gorm:column"status;default:1"`
	CreatedAt *time.Time `json:"created_at,omitempty" gorm:column"created_at;"`
	UpdatedAt *time.Time `json:"updated_at,omitempty" gorm:column"updated_at;"`
}