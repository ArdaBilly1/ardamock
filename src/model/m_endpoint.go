package model

import "time"

type Endpoint struct {
	ID        int       `json:"id" gorm:"primaryKey"`
	Method    string    `json:"method" gorm:"size:100"`
	Name      string    `json:"name" gorm:"size:100"`
	ExpiredAt time.Time `json:"expired_at"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at" gorm:"autoUpdateTime"`
}

func (Endpoint) TableName() string {
	return "endpoint"
}
