package models

import "time"

// User represents the User model
type User struct {
	ID        uint       `gorm:"primary_key;auto_increment" json:"id"`
	Username  string     `gorm:"not null" json:"username"`
	Email     string     `gorm:"unique;not null" json:"email"`
	Password  string     `gorm:"not null" json:"password"`
	CreatedAt time.Time  `gorm:"type:timestamp;default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt time.Time  `gorm:"type:timestamp;default:CURRENT_TIMESTAMP" json:"updated_at"`
	Photos    []Photo    `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"photos"`
}