package models

import "time"

type Person struct {
	Id        uint      `gorm:"primarykey" json:"id"`
	Name      string    `gorm:"type:varchar(100)" json:"name"`
	Email     string    `gorm:"type:varchar(100)" json:"email"`
	Umur      uint      `gorm:"type:integer" json:"umur"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"update_at"`
}
