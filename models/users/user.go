package users

import "time"

// User struct is model using for database & response
type User struct {
	UserID        int       `json:"user_id"`
	InstitutionID int       `json:"institution_id"`
	UserFullname  string    `json:"user_fullname"`
	UserEmail     string    `json:"user_email"`
	UserName      string    `json:"user_name"`
	UserPassword  string    `json:"-"`
	UserGender    string    `json:"user_gender"`
	UserBirthdate string    `json:"user_birthdate"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
	TeamID        string    `json:"team_id"`
}
