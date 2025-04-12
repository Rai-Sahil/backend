package domain

import "time"

type User struct {
	ID			uint		`json:"id" gorm:"primaryKey"`
	Name		string		`json:"name" gorm:"not null"`
	Email		string		`json:"email" gorm:"unique;not null"`
	Password	string		`json:"password" gorm:"type:varchar(255);not null"`
	CreatedAt	time.Time	`json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt	time.Time	`json:"updated_at" gorm:"autoUpdateTime"`
}
