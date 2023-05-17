package models

import (
	"time"
)

type User struct {
	ID        string    // The user's ID
	FirstName string    // The user's first name
	Surname   string    // The user's last name
	Email     string    // The user's email address
	DOB       time.Time // The user's date of birth
	Password  string    // The user's password (this should be hashed and salted in a real application)
	FilePath  string    // The path to the user's file on disk
}
