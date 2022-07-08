package models

import (
	"time"

	"gorm.io/gorm"
)

type Model struct {
	ID        uint `gorm:"primaryKey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

type Person struct {
	Name string `json:"name"`
	Pass string `json:"pass"`
}

type Dogs struct {
	gorm.Model
	Name  string `json:"name"`
	DogID int    `json:"dog_id"`
}

type Users struct {
	gorm.Model
	Employee_id int    `json:"employee_id"`
	Name        string `json:"name"`
	Lastname    string `json:"lastname"`
	Birthday    string `json:"birthday"`
	Age         int    `json:"age"`
	Email       string `json:"email"`
	Tel         string `json:"tel"`
}
