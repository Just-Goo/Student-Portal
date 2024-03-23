package model

import "time"

type Student struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	Age       int       `json:"age"`
	RegNumber string    `json:"reg_number"`
	Gender    string    `json:"gender"`
	CreatedAt time.Time `json:"created_at"`
}
