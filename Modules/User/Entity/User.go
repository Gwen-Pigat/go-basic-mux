package userEntity

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name  string `json:"name"`
	Alias string `json:"alias"`
}
