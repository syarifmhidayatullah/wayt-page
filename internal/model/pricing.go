package model

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"time"
)

type StringSlice []string

func (s StringSlice) Value() (driver.Value, error) {
	b, err := json.Marshal(s)
	return string(b), err
}

func (s *StringSlice) Scan(value any) error {
	bytes, ok := value.([]byte)
	if !ok {
		return errors.New("type assertion to []byte failed")
	}
	return json.Unmarshal(bytes, s)
}

type PricingPlan struct {
	ID            uint        `gorm:"primaryKey;autoIncrement" json:"id"`
	Name          string      `gorm:"not null"                 json:"name"`
	PriceMonthly  int         `gorm:"not null;default:0"       json:"price_monthly"`
	PriceYearly   int         `gorm:"not null;default:0"       json:"price_yearly"`
	Features      StringSlice `gorm:"type:jsonb;not null"      json:"features"`
	IsPopular     bool        `gorm:"default:false"            json:"is_popular"`
	IsActive      bool        `gorm:"default:true"             json:"is_active"`
	SortOrder     int         `gorm:"default:0"                json:"sort_order"`
	CreatedAt     time.Time   `                                json:"created_at"`
	UpdatedAt     time.Time   `                                json:"updated_at"`
}

func (PricingPlan) TableName() string { return "page_pricing_plans" }
