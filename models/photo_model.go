package models

import "time"

type Photo struct {
	ID        uint      `json:"id,omitempty" gorm:"primaryKey"`
	Title     string    `json:"title" gorm:"not null"`
	Caption   string    `json:"caption" gorm:"null"`
	PhotoURL  string    `json:"photo_url" gorm:"not null"`
	UserID    uint      `json:"user_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	User      *User     `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE" json:",omitempty"`
}
