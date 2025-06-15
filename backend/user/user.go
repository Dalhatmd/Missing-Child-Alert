package user

import (
	"github.com/dalhatmd/Missing-Child-Alert/alert"
)

type User struct {
	ID       int `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Location string `json:"location"`
	PhoneNumber string `json:"phone_number,omitempty"`
	Alerts []alert.Alert `gorm:"foreignKey:UserID" json:"alerts,omitempty"`
}

// newUser creates a new User instance with the provided details.

func newUser(id int, username, email, password, location string) *User {
	return &User{
		ID:       id,
		Username: username,
		Email:    email,
		Password: password,
		Location: location,
	}
}

func UpdateUser(user *User, username, email, password, location string) {
	if username != "" {
		user.Username = username
	}
	if email != "" {
		user.Email = email
	}
	if password != "" {
		user.Password = password
	}
	if location != "" {
		user.Location = location
	}
}
