package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID        int            `json:"id" gorm:"primaryKey"`
	Firstname string         `json:"firstname" gorm:"not null"`
	Lastname  string         `json:"lastname" gorm:"not null"`
	Username  string         `json:"username" gorm:"not null;unique"`
	Email     string         `json:"email" gorm:"not null;unique"`
	Password  string         `json:"-" gorm:"not null"`
	Picture   string         `json:"path,omitempty" gorm:"not null"`
	IsActive  bool           `json:"is_active" gorm:"not null"`
	CreatedAt time.Time      `json:"-"`
	UpdatedAt time.Time      `json:"-"`
	DeletedAt gorm.DeletedAt `json:"-" gorm:"index"`
}

type Account struct {
	ID        int            `json:"id" gorm:"primaryKey"`
	Name      string         `json:"name" gorm:"not null"`
	UserID    int            `json:"user_id" gorm:"references:users;not null"`
	Balance   float64        `json:"balance,omitempty" gorm:"not null"`
	Picture   string         `json:"path,omitempty" gorm:"not null"`
	IsActive  bool           `json:"is_active" gorm:"not null"`
	CreatedAt time.Time      `json:"-"`
	UpdatedAt time.Time      `json:"-"`
	DeletedAt gorm.DeletedAt `json:"-" gorm:"index"`
}

type Category struct {
	ID          int            `json:"id" gorm:"primaryKey"`
	Name        string         `json:"name" gorm:"not null"`
	Description string         `json:"description" gorm:"not null"`
	Price       float64        `json:"price,omitempty" gorm:"not null"`
	Picture     string         `json:"path,omitempty" gorm:"not null"`
	IsActive    bool           `json:"is_active" gorm:"not null"`
	CreatedAt   time.Time      `json:"-"`
	UpdatedAt   time.Time      `json:"-"`
	DeletedAt   gorm.DeletedAt `json:"-" gorm:"index"`
}

type Transaction struct {
	ID        int            `json:"id" gorm:"primaryKey"`
	From      int            `json:"from_id"`
	To        int            `json:"to_id"`
	ToType    string         `json:"to_type"`
	Comment   string         `json:"comment,omitempty" gorm:"not null"`
	Amount    float64        `json:"amount"`
	Type      string         `json:"type"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"-" gorm:"index"`
}

type Report struct {
	FromID int       `json:"from_id,omitempty"`
	ToID   int       `json:"to_id,omitempty"`
	ToType string    `json:"to_type,omitempty"`
	Limit  int       `json:"limit,omitempty"`
	Page   int       `json:"page,omitempty"`
	Type   string    `json:"type,omitempty"`
	From   time.Time `json:"-"`
	To     time.Time `json:"-"`
}
