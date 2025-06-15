package main

import (
	"fmt"
	"errors" // For errors.Is and errors.New comparisons
	"log"    // For fatal errors

	"github.com/dalhatmd/Missing-Child-Alert/user" // Adjust import path if necessary
	"golang.org/x/crypto/bcrypt" // To verify hashed passwords
	"github.com/google/uuid"    // To generate unique IDs for users
)

func main() {
	fmt.Println("--- Testing User Package ---")

	// --- Test NewUser ---
	fmt.Println("\n--- Testing NewUser ---")

	// Test Case 1: Valid User Creation
	fmt.Println("\nTesting NewUser with valid data:")
	userID1 := uuid.New().String()
	validUser, err := user.NewUser(
		userID1,
		"john_doe",
		"john.doe@example.com",
		"StrongP@ssw0rd!",
		"Abuja, FCT",
	)
	if err != nil {
		log.Fatalf("Error creating valid user (unexpected): %v", err)
	}
	fmt.Println("Valid user created successfully:")
	fmt.Printf("ID: %s\n", validUser.ID)
	fmt.Printf("Username: %s\n", validUser.Username)
	fmt.Printf("Email: %s\n", validUser.Email)
	// You cannot print PasswordHash directly as it's hashed
	fmt.Printf("PasswordHash (first 10 chars): %s...\n", validUser.PasswordHash[:10])
	fmt.Printf("Location: %s\n", validUser.Location)
	fmt.Printf("Phone Number: %s (should be empty)\n", validUser.PhoneNumber) // Not set in constructor

	// Verify password hash (example of how you'd check a password later)
	err = bcrypt.CompareHashAndPassword([]byte(validUser.PasswordHash), []byte("StrongP@ssw0rd!"))
	if err != nil {
		fmt.Printf("Password verification failed (unexpected): %v\n", err)
	} else {
		fmt.Println("Password verified successfully.")
	}

	// Test Case 2: NewUser with empty username
	fmt.Println("\nTesting NewUser with empty username:")
	userID2 := uuid.New().String()
	_, err = user.NewUser(userID2, "", "test@example.com", "password123", "Lagos")
	if err != nil && errors.Is(err, errors.New("username cannot be empty")) {
		fmt.Printf("Expected error for empty username: %v\n", err)
	} else {
		fmt.Printf("Unexpected result for empty username. Got: %v\n", err)
	}

	// Test Case 3: NewUser with empty email
	fmt.Println("\nTesting NewUser with empty email:")
	userID3 := uuid.New().String()
	_, err = user.NewUser(userID3, "jane_doe", "", "password123", "Kano")
	if err != nil && errors.Is(err, errors.New("email cannot be empty")) {
		fmt.Printf("Expected error for empty email: %v\n", err)
	} else {
		fmt.Printf("Unexpected result for empty email. Got: %v\n", err)
	}

	// Test Case 4: NewUser with empty password
	fmt.Println("\nTesting NewUser with empty password:")
	userID4 := uuid.New().String()
	_, err = user.NewUser(userID4, "bob_smith", "bob@example.com", "", "Port Harcourt")
	if err != nil && errors.Is(err, errors.New("password cannot be empty")) {
		fmt.Printf("Expected error for empty password: %v\n", err)
	} else {
		fmt.Printf("Unexpected result for empty password. Got: %v\n", err)
	}

	// --- Test Update Method ---
	fmt.Println("\n--- Testing Update Method ---")

	if validUser == nil {
		fmt.Println("Cannot test update as initial user creation failed.")
		return
	}

	// Test Case 5: Update Username only
	fmt.Println("\nTesting update: Change username to 'john_updated'")
	newUsername := "john_updated"
	updateReq1 := user.UserDTO{Username: &newUsername}
	err = validUser.Update(updateReq1)
	if err != nil {
		fmt.Printf("Error updating username (unexpected): %v\n", err)
	} else {
		fmt.Printf("Updated Username: %s (Expected: john_updated)\n", validUser.Username)
	}

	// Test Case 6: Update Email and Location
	fmt.Println("\nTesting update: Change email and location")
	newEmail := "john.doe.new@example.com"
	newLocation := "Lagos, Ikeja"
	updateReq2 := user.UserDTO{Email: &newEmail, Location: &newLocation}
	err = validUser.Update(updateReq2)
	if err != nil {
		fmt.Printf("Error updating email/location (unexpected): %v\n", err)
	} else {
		fmt.Printf("Updated Email: %s (Expected: %s)\n", validUser.Email, newEmail)
		fmt.Printf("Updated Location: %s (Expected: %s)\n", validUser.Location, newLocation)
	}

	// Test Case 7: Update Password
	fmt.Println("\nTesting update: Change password to 'NewS3cureP@ss'")
	newPassword := "NewS3cureP@ss"
	updateReq3 := user.UserDTO{Password: &newPassword}
	err = validUser.Update(updateReq3)
	if err != nil {
		fmt.Printf("Error updating password (unexpected): %v\n", err)
	} else {
		fmt.Println("Password update attempted.")
		// Verify the new password hash
		err = bcrypt.CompareHashAndPassword([]byte(validUser.PasswordHash), []byte(newPassword))
		if err != nil {
			fmt.Printf("New password verification failed (unexpected): %v\n", err)
		} else {
			fmt.Println("New password verified successfully.")
		}
	}

	// Test Case 8: Update Phone Number
	fmt.Println("\nTesting update: Set Phone Number")
	newPhoneNumber := "08011223344"
	updateReq4 := user.UserDTO{PhoneNumber: &newPhoneNumber}
	err = validUser.Update(updateReq4)
	if err != nil {
		fmt.Printf("Error setting phone number (unexpected): %v\n", err)
	} else {
		fmt.Printf("Updated Phone Number: %s (Expected: %s)\n", validUser.PhoneNumber, newPhoneNumber)
	}

	// Test Case 9: Attempt to update username to empty string (should fail due to validation)
	fmt.Println("\nTesting update with empty username (should error):")
	emptyString := ""
	updateReq5 := user.UserDTO{Username: &emptyString}
	err = validUser.Update(updateReq5)
	if err != nil && errors.Is(err, errors.New("username cannot be empty")) {
		fmt.Printf("Expected error for empty username update: %v\n", err)
	} else {
		fmt.Printf("Unexpected result for empty username update. Got: %v\n", err)
	}
	fmt.Printf("Current Username (should not have changed): %s\n", validUser.Username)

	// Test Case 10: Provide nil for a field (should not change)
	fmt.Println("\nTesting update: Provide nil for Email (should not change)")
	updateReq6 := user.UserDTO{Email: nil} // Email is nil
	err = validUser.Update(updateReq6)
	if err != nil {
		fmt.Printf("Error updating with nil (unexpected): %v\n", err)
	} else {
		fmt.Printf("Email after nil update (should be %s): %s\n", newEmail, validUser.Email)
	}

	fmt.Println("\n--- End of Testing ---")
}

// You'll need an isValidEmail function in your user/user.go for proper email validation
// For testing purposes, you can add a dummy or a simple regex-based one if you haven't already.
// Example for user/user.go:
// var emailRegex = regexp.MustCompile(`^[a-zA-Z0-9._%+\-]+@[a-zA-Z0-9.\-]+\.[a-zA-Z]{2,4}$`)
// func isValidEmail(email string) bool {
//     return emailRegex.MatchString(email)
// }
