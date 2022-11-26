package models

import "time"

type User struct {
	UserId    int       `gorm:"not null;uniqueIndex;primaryKey;" json:"user_id"`
	Username  string    `gorm:"not null;uniqueIndex;size:256;" json:"username"`
	Password  string    `gorm:"not null;" json:"password"`
	Name      string    `gorm:"not null;size:256" json:"name"`
	Role      string    `gorm:"not null;size:5" json:"role"`
	CreatedAt time.Time `gorm:"not null;" json:"created_at"`
	UpdatedAt time.Time `gorm:"not null;" json:"updated_at"`
}
