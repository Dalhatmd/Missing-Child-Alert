package user

import (
	"errors"
	"fmt"

	"github.com/dalhatmd/Missing-Child-Alert/alert"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID           string        `json:"id"`
	FirstName    string        `json:"first_name"`
	LastName     string        `json:"last_name"`
	Email        string        `json:"email"`
	PasswordHash string        `json:"password"`
	Location     string        `json:"location"`
	PhoneNumber  string        `json:"phone_number,omitempty"`
	Alerts       []alert.Alert `gorm:"foreignKey:UserId;references:ID" json:"alerts,omitempty"`
}

// newUser creates a new User instance with the provided details.

func NewUser(id string, first_name, last_name, email, password, location string) (*User, error) {
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
		ID:           id,
		FirstName:    first_name,
		LastName:     last_name,
		Email:        email,
		PasswordHash: string(hashedPassword),
		Location:     location,
	}, nil
}

// DTO for User represents the data transfer object for a user.
type UserDTO struct {
	FirstName   *string `json:"firstName"`
	LastName    *string `json:"lastName"`
	Email       *string `json:"email"`
	Password    *string `json:"password,omitempty"`
	Location    *string `json:"location"`
	PhoneNumber *string `json:"phone_number,omitempty"`
}

func (u *User) Update(info UserDTO) error {
	if info.FirstName != nil {
		u.FirstName = *info.FirstName
	}
	if info.LastName != nil {
		u.LastName = *info.LastName
	}
	if info.Email != nil {
		if *info.Email == "" {
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

// CheckPassword compares a hashed password with a plain password.
func CheckPassword(hashedPassword, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}

// UserResponse is a safe struct for returning user data in API responses (no password hash).
type UserResponse struct {
	ID          string        `json:"id"`
	FirstName   string        `json:"first_name"`
	LastName    string        `json:"last_name"`
	Email       string        `json:"email"`
	Location    string        `json:"location"`
	PhoneNumber string        `json:"phone_number,omitempty"`
	Alerts      []alert.Alert `json:"alerts,omitempty"`
}
