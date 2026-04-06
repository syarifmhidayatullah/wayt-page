package model

import "time"

type Testimonial struct {
	ID         uint      `gorm:"primaryKey;autoIncrement" json:"id"`
	Name       string    `gorm:"not null"                 json:"name"`
	Restaurant string    `gorm:"not null"                 json:"restaurant"`
	Quote      string    `gorm:"type:text;not null"       json:"quote"`
	Rating     int       `gorm:"not null;default:5"       json:"rating"`
	Phone      string    `gorm:"default:''"               json:"phone"`
	IsActive   bool      `gorm:"default:true"             json:"is_active"`
	SortOrder  int       `gorm:"default:0"                json:"sort_order"`
	CreatedAt  time.Time `                                json:"created_at"`
	UpdatedAt  time.Time `                                json:"updated_at"`
}

func (Testimonial) TableName() string { return "page_testimonials" }
