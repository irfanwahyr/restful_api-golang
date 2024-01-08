package models

import "time"

type Address struct {
	Id        uint      `gorm:"primarykey" json:"id"`
	Person_id uint      `json:"person_id"`
	Person    Person    `gorm:"foreignkey:Person_id" json:"person"`
	Street    string    `json:"street"`
	Village   string    `json:"village"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type AddressResponse struct {
	Id        uint          `json:"id"`
	Person_id uint          `json:"-"`
	Person    PersonAddress `gorm:"foreignkey:Person_id" json:"person"`
	Street    string        `json:"street"`
	Village   string        `json:"village"`
	CreatedAt time.Time     `json:"created_at"`
	UpdatedAt time.Time     `json:"updated_at"`
}
