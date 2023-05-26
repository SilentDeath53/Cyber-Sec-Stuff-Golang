package main

import (
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"strings"
)

func checkPasswordStrength(password string) error {
	// Check if password length is at least 8 characters
	if len(password) < 8 {
		return fmt.Errorf("password should be at least 8 characters long")
	}

	if strings.ToUpper(password) == password || strings.ToLower(password) == password {
		return fmt.Errorf("password should contain both uppercase and lowercase letters")
	}

	hasDigit := false
	for _, char := range password {
		if char >= '0' && char <= '9' {
			hasDigit = true
			break
		}
	}
	if !hasDigit {
		return fmt.Errorf("password should contain at least one digit")
	}

	hasSpecialChar := false
	specialChars := "!@#$%^&*()-=_+[]{}|;:,.<>/?"
	for _, char := range password {
		if strings.ContainsRune(specialChars, char) {
			hasSpecialChar = true
			break
		}
	}
	if !hasSpecialChar {
		return fmt.Errorf("password should contain at least one special character")
	}

	return nil
}

func main() {
	password := "SilentDeath53!"

	err := checkPasswordStrength(password)
	if err != nil {
		fmt.Println("Password is not strong:", err)
	} else {
		fmt.Println("Password is strong")
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		fmt.Println("Error hashing password:", err)
		return
	}

	fmt.Println("Hashed password:", string(hashedPassword))

	err = bcrypt.CompareHashAndPassword(hashedPassword, []byte(password))
	if err != nil {
		fmt.Println("Password does not match the hashed password")
	} else {
		fmt.Println("Password matches the hashed password")
	}
}
