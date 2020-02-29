package entity

import "github.com/jinzhu/gorm"

// User is user models property
type User struct {
	gorm.Model
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
}
