package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Book struct {
	ID     uuid.UUID `gorm:"type:char(36);primaryKey" json:"id"`
	Title  string    `json:"title"`
	Author string    `json:"author"`
}

type Response struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
	Data    any    `json:"data"`
}

func (u *Book) BeforeCreate(tx *gorm.DB) (err error) {
	u.ID = uuid.New()
	return
}
