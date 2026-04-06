package model

import "time"

type Lead struct {
	ID             uint      `gorm:"primaryKey;autoIncrement" json:"id"`
	FullName       string    `gorm:"not null"                 json:"full_name"`
	RestaurantName string    `gorm:"not null"                 json:"restaurant_name"`
	Email          string    `gorm:"not null"                 json:"email"`
	Phone          string    `gorm:"not null;default:''"      json:"phone"`
	CreatedAt      time.Time `                                json:"created_at"`
}

func (Lead) TableName() string { return "page_leads" }
