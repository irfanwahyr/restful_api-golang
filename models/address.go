package models

import "time"

type Address struct {
	Id        uint      `gorm:"primarykey" json:"id"`
	Street    string    `json:"street"`
	Village   string    `json:"village"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
