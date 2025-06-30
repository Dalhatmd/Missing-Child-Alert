package alert

import (
	"time"
	"errors"
	"github.com/google/uuid"
)

type AlertStatus string

const (
	// ActiveStatus indicates that the alert is currently active.
	ActiveStatus AlertStatus = "active"
	// ResolvedStatus indicates that the alert has been resolved.
	ResolvedStatus AlertStatus = "resolved"
	// ClosedStatus indicates that the alert has been closed.
	ClosedStatus AlertStatus = "closed"
)


type Alert struct {
	ID			string 		`gorm:"PrimaryKey" json:"id"`
	ChildName	string 		`json:"child_name"`
	Age			int			`json:"Age"`
	Gender		string		`json:"gender"`
	Description	string		`json:"description"`
	LastSeenLocation	string	`json:"last_seen_location"`
	PhotoUrl 	string		`json:"photo_url"`
	ReporterContact	string   `json:"reporter_contact"`
	CreatedAt	time.Time	  `json:"created_at"`
	TimeLost	time.Time	  `json:"time_lost"`
	UserId		string		`json:"user_id"`
	Status AlertStatus `json:"status,omitempty"`
}

type UpdateAlertRequest struct {
    ChildName        *string    `json:"child_name,omitempty"`
    Age              *int       `json:"age,omitempty"`
    LastSeenLocation *string    `json:"last_seen_location,omitempty"`
    Description      *string    `json:"description,omitempty"`
    PhotoUrl         *string    `json:"photo_url,omitempty"`
    ReporterContact  *string    `json:"reporter_contact,omitempty"`
    Status           *AlertStatus `json:"status,omitempty"`
    TimeLost         *time.Time `json:"time_lost,omitempty"`
	Gender			 *string     `json:"gender,omitempty"`
}

func NewAlert(
			childName string,
			age int,
			gender string,
	        lastSeenlLocation string,
			description string,
			photoUrl string,
			reporterContact string,
			UserId string,
) (*Alert, error) {
	if childName == "" {
		return nil, errors.New("child name cannot be empty")
	}
	if age <= 0 {
		return nil, errors.New("Age must be a positive integer")
	}
	if lastSeenlLocation == "" {
		return nil, errors.New("Last seen location cannot be empty")
	}
	if reporterContact == "" {
		return nil, errors.New("Reporter contact cannot be empty")
	}
	if UserId == "" {
		return nil, errors.New("User ID cannot be empty")
	}

	newID, err := uuid.NewRandom()
	if err != nil {
		return nil, errors.New("failed to generate uid")
	}

	return &Alert{
		ID:				  newID.String(),
		ChildName:        childName,
		Age:              age,
		LastSeenLocation: lastSeenlLocation,
		Description:      description,
		PhotoUrl:         photoUrl,
		ReporterContact:  reporterContact,
		Status:           ActiveStatus,
		CreatedAt:        time.Now(),
		TimeLost:         time.Time{},
	}, nil
}

// Update applies the provided partial updates to the Alert.
// It returns an error if any of the provided values are invalid.
func (a *Alert) Update(info UpdateAlertRequest) error {
	if info.ChildName != nil {
		if *info.ChildName == "" {
			return errors.New("child name cannot be empty")
		}
		a.ChildName = *info.ChildName
	}
	if info.Age != nil {
		if *info.Age <= 0 || *info.Age > 18 {
			return errors.New("age must be a positive number between 1 and 18")
		}
		a.Age = *info.Age
	}
	if info.Gender != nil {
		if info.Gender == nil || *info.Gender == "" {
			return errors.New("Please provide gender")
		}
		a.Gender = *info.Gender
	}
	if info.LastSeenLocation != nil {
		if *info.LastSeenLocation == "" {
			return errors.New("last seen location cannot be empty")
		}
		a.LastSeenLocation = *info.LastSeenLocation
	}
	if info.Description != nil {
		a.Description = *info.Description
	}
	if info.PhotoUrl != nil {
		a.PhotoUrl = *info.PhotoUrl
	}
	if info.ReporterContact != nil {
		if *info.ReporterContact == "" {
			return errors.New("reporter contact cannot be empty")
		}
		a.ReporterContact = *info.ReporterContact
	}
	if info.Status != nil {
		switch *info.Status {
		case ActiveStatus, ResolvedStatus, ClosedStatus:
			a.Status = *info.Status
		default:
			return errors.New("invalid alert status provided")
		}
	}
	if info.TimeLost != nil {
		a.TimeLost = *info.TimeLost
	}
	return nil
}
