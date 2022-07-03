package model

import "time"

type User struct {
	ID          string    `json:"id" gorm:"primaryKey"`
	Username    string    `json:"username" gorm:"not null"`
	FirstName   string    `json:"first_name" gorm:"not null" db:"first_name"`
	LastName    string    `json:"last_name" gorm:"not null" db:"last_name"`
	Email       string    `json:"email" gorm:"not null"`
	DateOfBirth string    `json:"date_of_birth" gorm:"column:dob" db:"dob"`
	CreatedAt   time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt   time.Time `json:"updated_at" gorm:"autoUpdateTime"`
	DeletedAt   time.Time `json:"deleted_at"`
	Address     []Address `json:"addresses" db:"-"`
}
