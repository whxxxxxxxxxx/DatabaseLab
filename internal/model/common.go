package model

import (
	"github.com/oklog/ulid/v2"
	"gorm.io/gorm"
	"time"
)

type Base struct {
	ID        string         `json:"id" gorm:"primary_key;type:char(26)"`
	CreatedAt time.Time      `json:"-"`
	UpdatedAt time.Time      `json:"-"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}

func (u *Base) BeforeCreate(tx *gorm.DB) (err error) {
	u.ID = ulid.Make().String()
	return
}
