package models

type User struct {
	ID        int     // The user's ID
	FirstName string  // The user's first name
	Surname   string  // The user's last name
	Email     string  // The user's email address
	DOB       string  // The user's date of birth
	Password  string  // The user's password (hashed)
	FilePath  *string // The path to the user's file on disk (Allows for Null values)
}
