package user

import (
	"github.com/dalhatmd/Missing-Child-Alert/alert"
	"golang.org/x/crypto/bcrypt"
	"errors"
	"fmt"
)

type User struct {
	ID       string `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	PasswordHash string `json:"password"`
	Location string `json:"location"`
	PhoneNumber string `json:"phone_number,omitempty"`
	Alerts []alert.Alert `gorm:"foreignKey:UserID" json:"alerts,omitempty"`
}

// newUser creates a new User instance with the provided details.

func NewUser(id string, username, email, password, location string) (*User, error) {
	if username == "" {
		return nil, errors.New("username cannot be empty")
	}
	if email == "" {
		return nil, errors.New("email cannot be empty")
	}
	if password == "" {
		return nil, errors.New("password cannot be empty")
	}
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, fmt.Errorf("failed to hash password: %v", err)
		}
	return &User{
		ID:       id,
		Username: username,
		Email:    email,
		PasswordHash: string(hashedPassword),
		Location: location,
	}, nil
}

// DTO for User represents the data transfer object for a user.
type UserDTO struct {
	Username *string `json:"username"`
	Email    *string `json:"email"`
	Password *string `json:"password,omitempty"`
	Location *string `json:"location"`
	PhoneNumber *string `json:"phone_number,omitempty"`
}

func (u *User) Update(info UserDTO) error {
	if info.Username != nil {
		if *info.Username == "" {
			return errors.New("username cannot be empty")
		}
		u.Username = *info.Username
	}
	if info.Email != nil {
		if *info.Email == ""{
			return errors.New("Email cannot be empty")
		}
		u.Email = *info.Email
	}
	if info.Password != nil {
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(*info.Password), bcrypt.DefaultCost)
		if err != nil {
			return fmt.Errorf("failed to hash password: %v", err)
		}
		u.PasswordHash = string(hashedPassword)
	}
	if info.Location != nil {
		u.Location = *info.Location
	}
	if info.PhoneNumber != nil {
		u.PhoneNumber = *info.PhoneNumber
	}
	return nil
}
