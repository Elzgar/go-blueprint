package repository

import (
	"go-blueprint/database"
	"go-blueprint/models"
	"go-blueprint/models/users"

	"golang.org/x/crypto/bcrypt"
)

// FetchAllUser is function for get users data direct from database for
func FetchAllUser() (result models.Response, code int, error error) {
	db := database.CreateConn()

	users := []users.User{}
	db.Find(&users)

	return models.Response{
		Status:  true,
		Message: "Success fetch users",
		Data:    users,
	}, 200, nil
}

// GetUserDetail is function for doing login with parameter email & password
func GetUserDetail(email, password string) (result models.Response, code int, error error) {
	db := database.CreateConn()

	u := users.User{}
	db.Where(&users.User{UserEmail: email}).Find(&u)

	hp, _ := Hash(password)

	if !IsSame(u.UserPassword, hp) {
		return models.Response{
			Status:  false,
			Message: "Password is wrong!",
			Data:    hp,
		}, 400, nil
	}

	return models.Response{
		Status:  true,
		Message: "Success get user",
		Data:    u,
	}, 200, nil
}

// Hash is function for encrypt string with return type string and error
func Hash(str string) (string, error) {
	hashed, err := bcrypt.GenerateFromPassword([]byte(str), bcrypt.DefaultCost)
	return string(hashed), err
}

// IsSame is function for comparing hashed string with return type boolean
func IsSame(str string, hashed string) bool {
	return bcrypt.CompareHashAndPassword([]byte(hashed), []byte(str)) == nil
}
