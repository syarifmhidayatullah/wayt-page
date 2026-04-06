package model

import "time"

type Setting struct {
	ID        uint      `gorm:"primaryKey;autoIncrement" json:"id"`
	Key       string    `gorm:"uniqueIndex;not null"     json:"key"`
	Value     string    `gorm:"type:text;not null"       json:"value"`
	Label     string    `gorm:"default:''"               json:"label"`
	UpdatedAt time.Time `                                json:"updated_at"`
}

func (Setting) TableName() string { return "page_settings" }
