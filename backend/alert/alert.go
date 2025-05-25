package alert

import (
	"time"
)

type Alert struct {
	ID			string 		`gorm:"PrimaryKey" json:"id"`
	ChildName	string 		`json:"child_name"`
	Age			int			`json:"Age"`
	Gender		string		`json:"gender"`
	Description	string		`json:"description"`
	LastSeenLocation	string	`json:"last_seen_location"`
	PhotoUrl 	string		`json:"PhotoUrl"`
	ReporterContact	string   `json:"reporter_contact"`
	Status		string		 	`json:"status"`
	CreatedAt	time.Time	  `json:"created_at"`
	TimeLost	time.Time	  `json:"time_lost"`
}

func NewAlert(childName string, age int, last_seen_location string, description string, photoUrl string, reporterContact string) *Alert {
	return &Alert{
		ChildName:        childName,
		Age:              age,
		LastSeenLocation: last_seen_location,
		Description:      description,
		PhotoUrl:         photoUrl,
		ReporterContact:  reporterContact,
		Status:           "active",
		CreatedAt:        time.Now(),
	}
}

func (a *Alert) Update(info Alert) {
	if info.ChildName != "" {
		a.ChildName = info.ChildName
	}
	if info.Age != 0 {
		a.Age = info.Age
	}
	if info.LastSeenLocation != "" {
		a.LastSeenLocation = info.LastSeenLocation
	}
	if info.Description != "" {
		a.Description = info.Description
	}
	if info.PhotoUrl != "" {
		a.PhotoUrl = info.PhotoUrl
	}
	if info.ReporterContact != "" {
		a.ReporterContact = info.ReporterContact
	}
	if info.Status != "" {
		a.Status = info.Status
	}
}

