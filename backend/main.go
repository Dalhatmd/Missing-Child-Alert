package main

import (
	"fmt"
	"time"
	"github.com/dalhatmd/Missing-Child-Alert/alert" // Adjust import path if necessary
)

func main() {
	fmt.Println("--- Testing NewAlert ---")

	// Test Case 1: Valid Alert Creation
	fmt.Println("\nTesting NewAlert with valid data:")
	validAlert, err := alert.NewAlert(
		"John Doe",
		12,
		"Male", // Added gender
		"Abuja, Wuse II",
		"Last seen wearing a blue t-shirt and red shorts.",
		"http://example.com/john_doe.jpg",
		"08012345678",
		"user123", // Added user ID
	)
	if err != nil {
		fmt.Printf("Error creating valid alert (unexpected): %v\n", err)
	} else {
		fmt.Println("Valid alert created successfully:")
		fmt.Printf("ID: %s\n", validAlert.ID)
		fmt.Printf("Child Name: %s\n", validAlert.ChildName)
		fmt.Printf("Age: %d\n", validAlert.Age)
		fmt.Printf("Gender: %s\n", validAlert.Gender)
		fmt.Printf("Last Seen Location: %s\n", validAlert.LastSeenLocation)
		fmt.Printf("Description: %s\n", validAlert.Description)
		fmt.Printf("Photo URL: %s\n", validAlert.PhotoUrl)
		fmt.Printf("Reporter Contact: %s\n", validAlert.ReporterContact)
		fmt.Printf("Status: %s\n", validAlert.Status)
		fmt.Printf("Created At: %s\n", validAlert.CreatedAt.Format(time.RFC3339))
		fmt.Printf("Time Lost: %s (zero value if not set)\n", validAlert.TimeLost.Format(time.RFC3339))
		fmt.Printf("User ID: %s\n", validAlert.UserId)
	}

	// Test Case 2: NewAlert with Invalid Child Name
	fmt.Println("\nTesting NewAlert with empty child name:")
	_, err = alert.NewAlert(
		"", // Empty child name
		5,
		"Female",
		"Lagos, Ikeja",
		"Wore a yellow dress.",
		"http://example.com/jane_doe.jpg",
		"09098765432",
		"user456",
	)
	if err != nil {
		fmt.Printf("Expected error for empty child name: %v\n", err)
	} else {
		fmt.Println("Unexpected: Alert created with empty child name.")
	}

	// Test Case 3: NewAlert with Invalid Age
	fmt.Println("\nTesting NewAlert with invalid age (-1):")
	_, err = alert.NewAlert(
		"Alice Smith",
		-1, // Invalid age
		"Female",
		"Kano, Gwagwarwa",
		"Has a distinctive birthmark.",
		"http://example.com/alice_smith.jpg",
		"07011223344",
		"user789",
	)
	if err != nil {
		fmt.Printf("Expected error for invalid age: %v\n", err)
	} else {
		fmt.Println("Unexpected: Alert created with invalid age.")
	}

	// Test Case 4: NewAlert with Invalid UserID
	fmt.Println("\nTesting NewAlert with empty UserID:")
	_, err = alert.NewAlert(
		"Bob Johnson",
		8,
		"Male",
		"Port Harcourt, Rumuola",
		"Lost near the market.",
		"http://example.com/bob_johnson.jpg",
		"08155667788",
		"", // Empty UserID
	)
	if err != nil {
		fmt.Printf("Expected error for empty UserID: %v\n", err)
	} else {
		fmt.Println("Unexpected: Alert created with empty UserID.")
	}

	fmt.Println("\n--- Testing Alert Update Method ---")

	if validAlert == nil {
		fmt.Println("Cannot test update as initial alert creation failed.")
		return
	}

	// Test Case 5: Update Child Name
	fmt.Println("\nTesting update: Change child name to 'Jane Doe'")
	newName := "Jane Doe"
	updateReq1 := alert.UpdateAlertRequest{ChildName: &newName}
	err = validAlert.Update(updateReq1)
	if err != nil {
		fmt.Printf("Error updating child name (unexpected): %v\n", err)
	} else {
		fmt.Printf("Updated Child Name: %s (Expected: Jane Doe)\n", validAlert.ChildName)
	}

	// Test Case 6: Update Age and Status
	fmt.Println("\nTesting update: Change age to 13 and status to 'resolved'")
	newAge := 13
	newStatus := alert.ResolvedStatus
	updateReq2 := alert.UpdateAlertRequest{Age: &newAge, Status: &newStatus}
	err = validAlert.Update(updateReq2)
	if err != nil {
		fmt.Printf("Error updating age and status (unexpected): %v\n", err)
	} else {
		fmt.Printf("Updated Age: %d (Expected: 13)\n", validAlert.Age)
		fmt.Printf("Updated Status: %s (Expected: resolved)\n", validAlert.Status)
	}

	// Test Case 7: Update TimeLost
	fmt.Println("\nTesting update: Set TimeLost to now")
	currentTime := time.Now()
	updateReq3 := alert.UpdateAlertRequest{TimeLost: &currentTime}
	err = validAlert.Update(updateReq3)
	if err != nil {
		fmt.Printf("Error setting TimeLost (unexpected): %v\n", err)
	} else {
		fmt.Printf("Updated Time Lost: %s\n", validAlert.TimeLost.Format(time.RFC3339))
	}

	// Test Case 8: Attempt to update with invalid Age
	fmt.Println("\nTesting update with invalid age (-5):")
	invalidAge := -5
	updateReq4 := alert.UpdateAlertRequest{Age: &invalidAge}
	err = validAlert.Update(updateReq4)
	if err != nil {
		fmt.Printf("Expected error for invalid age update: %v\n", err)
	} else {
		fmt.Println("Unexpected: Age updated with invalid value.")
	}
	fmt.Printf("Current Age (should not have changed): %d\n", validAlert.Age)

	// Test Case 9: Attempt to update with invalid Status
	fmt.Println("\nTesting update with invalid status ('unknown'):")
	invalidStatus := alert.AlertStatus("unknown")
	updateReq5 := alert.UpdateAlertRequest{Status: &invalidStatus}
	err = validAlert.Update(updateReq5)
	if err != nil {
		fmt.Printf("Expected error for invalid status update: %v\n", err)
	} else {
		fmt.Println("Unexpected: Status updated with invalid value.")
	}
	fmt.Printf("Current Status (should not have changed): %s\n", validAlert.Status)

	// Test Case 10: Update multiple fields, including optional ones
	fmt.Println("\nTesting update: Multiple fields including description and photo URL")
	newDescription := "Last seen near the school bus stop, now wearing a green jacket."
	newPhotoUrl := "http://example.com/jane_doe_updated.jpg"
	updateReq6 := alert.UpdateAlertRequest{
		Description: &newDescription,
		PhotoUrl:    &newPhotoUrl,
	}
	err = validAlert.Update(updateReq6)
	if err != nil {
		fmt.Printf("Error updating multiple fields (unexpected): %v\n", err)
	} else {
		fmt.Printf("Updated Description: %s\n", validAlert.Description)
		fmt.Printf("Updated Photo URL: %s\n", validAlert.PhotoUrl)
	}

	// Test Case 11: Attempt to set a string field to empty (if your validation allows this based on field)
	// For example, if LastSeenLocation could be updated to empty, but your current validation prevents it.
	// This tests if providing a nil pointer for a field means "don't update" vs. an empty string means "set to empty".
	fmt.Println("\nTesting update: Provide nil for ChildName (should not change)")
	updateReq7 := alert.UpdateAlertRequest{ChildName: nil} // ChildName is nil
	err = validAlert.Update(updateReq7)
	if err != nil {
		fmt.Printf("Error updating with nil (unexpected): %v\n", err)
	} else {
		fmt.Printf("Child Name after nil update (should be Jane Doe): %s\n", validAlert.ChildName)
	}

	fmt.Println("\n--- End of Testing ---")
}
