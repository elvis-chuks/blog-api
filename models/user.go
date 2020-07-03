package models

type User struct {
	ID        string `json:"id,omitempty"`
	Email     string `json:"email"`
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
	Password  string `json:"password"`
	Pin       string `json:"pin"`
}
