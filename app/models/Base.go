package models

import "time"

type ID struct {
	ID uint `gorm:"primary_key" json:"id,omitempty"`
}

type TimeStamps struct {
	CreatedAt time.Time `gorm:"not null" json:"created_at" sql:"DEFAULT:CURRENT_TIMESTAMP"`
	UpdatedAt time.Time `gorm:"not null" json:"updated_at" sql:"DEFAULT:CURRENT_TIMESTAMP"`
}
type SoftDeletes struct {
	DeletedAt *time.Time `sql:"index" json:"deleted_at,omitempty"`
}
