package models

type User struct {
	// Databse fields
	ID         int
	TelegramID int

	// Generic fields
	Username  string
	Firstname string
	Lastname  string
}
