package models

import "time"

type BaseModel struct {
	ID        uint      `json:"id" gorm:"column:id;primaryKey;unique;autoIncrement;"`
	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt time.Time `json:"updated_at" gorm:"autoUpdateTime;default:null"`
}
