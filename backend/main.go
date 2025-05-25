package main

import(
	"fmt"
	"github.com/dalhatmd/Missing-Child-Alert/alert"
)

func main() {
	newAlert := alert.NewAlert("John Doe", 12, "Abuja", "Taken in a car", "http://example.com/photo.jpg", "123-456-7890")
	fmt.Printf("Child Name: %s\n", newAlert.ChildName)
	fmt.Printf("Age: %d\n", newAlert.Age)
	fmt.Printf("Last Seen Location: %s\n", newAlert.LastSeenLocation)
	fmt.Printf("Description: %s\n", newAlert.Description)
	fmt.Printf("Photo URL: %s\n", newAlert.PhotoUrl)
	fmt.Printf("Reporter Contact: %s\n", newAlert.ReporterContact)
	fmt.Printf("Status: %s\n", newAlert.Status)

	newAlert.Update(alert.Alert{ChildName: "Jane Doe"})
	fmt.Printf("Updated Child Name: %s\n", newAlert.ChildName)
}
