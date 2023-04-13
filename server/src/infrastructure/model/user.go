package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	ID       uint64 `gorm:"primary_key;auto_increment"`
	UserID   []byte
	Name     string
	Email    string
	Password []byte
}
