package model

import "time"

type TrustedRestaurant struct {
	ID        uint      `gorm:"primaryKey;autoIncrement" json:"id"`
	Name      string    `gorm:"not null"                 json:"name"`
	Emoji     string    `gorm:"not null;default:'🍽️'"    json:"emoji"`
	Rating    string    `gorm:"not null;default:'5.0'"   json:"rating"`
	IsActive  bool      `gorm:"default:true"             json:"is_active"`
	SortOrder int       `gorm:"default:0"                json:"sort_order"`
	CreatedAt time.Time `                                json:"created_at"`
	UpdatedAt time.Time `                                json:"updated_at"`
}

func (TrustedRestaurant) TableName() string { return "page_trusted_restaurants" }
