package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Attendance struct {
	ID        uuid.UUID       `gorm:"type:char(36);primaryKey" json:"id"`
	UserID    uint            `json:"user_id"`
	CiDate    time.Time       `json:"ci_date"`
	CiType    AttendanceType  `gorm:"type:enum('wfa','wfo');defaul:'wfa'" json:"ci_type"`
	CiFoto    string          `json:"ci_foto"`
	CiLat     float64         `json:"ci_lat"`
	CiLon     float64         `json:"ci_lon"`
	CoDate    *time.Time      `json:"co_date"`
	CoType    *AttendanceType `gorm:"type:enum('wfa','wfo');defaul:'wfa'" json:"co_type"`
	CoFoto    *string         `json:"co_foto"`
	CoLat     float64         `json:"co_lat"`
	CoLon     float64         `json:"co_lon"`
	CreatedAt time.Time       `json:"created_at"`
	UpdatedAt time.Time       `json:"updated_at"`
}

type AttendanceError struct {
	Status  uint   `json:"status"`
	Message string `json:"message"`
}

type AttendanceResponse struct {
	Status       uint   `json:"status"`
	Message      string `json:"message"`
	AttendanceID string `json:"attendance_id"`
}

type AttendanceType string

const (
	WFA AttendanceType = "wfa"
	WFO AttendanceType = "wfo"
)

func (u *Attendance) BeforeCreate(tx *gorm.DB) (err error) {
	u.ID = uuid.New()
	return
}
