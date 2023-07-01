package models

import "time"

// Photo represents the Photo model
type Photo struct {
	ID       uint      `gorm:"primary_key;auto_increment" json:"id"`
	Title    string    `gorm:"not null" json:"title"`
	Caption  string    `json:"caption"`
	PhotoURL string    `gorm:"not null" json:"photo_url"`
	UserID   uint      `gorm:"not null" json:"user_id"`
	User     User      `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"user"`
	CreatedAt time.Time `gorm:"type:timestamp;default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt time.Time `gorm:"type:timestamp;default:CURRENT_TIMESTAMP" json:"updated_at"`
}