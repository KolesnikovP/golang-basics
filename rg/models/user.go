package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

// Saves a user to the database
func (user *User) Create() (*User, error) {
	err := Database.Model(&user).Create(&user).Error
	if err != nil {
		return &User{}, err
	}
	return user, nil
}

// Fetches a user from the database
func FetchUser(id string) (*User, error) {
	var user User
	err := Database.Where("id = ?", id).First(&user).Error
	if err != nil {
		return &User{}, err
	}
	return &user, nil
}

// Updates a user in the database
func (user *User) UpdateUser(id string) (*User, error) {
	err := Database.Model(&User{}).Where("id = ?", id).Updates(user).Error
	if err != nil {
		return &User{}, err
	}
	return user, nil
}
