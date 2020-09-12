package users

import (
	"go-blueprint/models"
	"go-blueprint/repository"
	"net/http"

	"github.com/labstack/echo"
	"github.com/thedevsaddam/govalidator"
)

// GetAll function is for handling all user data and handle response
func GetAll(context echo.Context) error {
	result, _, err := repository.FetchAllUser()

	if err != nil {
		return context.JSON(http.StatusInternalServerError, models.Response{
			Status:  false,
			Message: err.Error(),
			Data:    nil,
		})
	}

	return context.JSON(http.StatusOK, result)
}

// Login function is for handling request data for login and handle response
func Login(context echo.Context) error {
	rules := govalidator.MapData{
		"email":    []string{"required", "email"},
		"password": []string{"required", "min:8"},
	}

	v := govalidator.New(govalidator.Options{
		Request:         context.Request(),
		Rules:           rules,
		RequiredDefault: true,
	})

	e := v.Validate()

	if e.Encode() != "" {
		return context.JSON(http.StatusBadRequest, models.Response{
			Status:  false,
			Message: "Input error",
			Data:    nil,
		})
	}

	result, code, err := repository.GetUserDetail(context.FormValue("email"), context.FormValue("password"))

	if err != nil {
		return context.JSON(http.StatusInternalServerError, models.Response{
			Status:  false,
			Message: err.Error(),
			Data:    nil,
		})
	} else if code != 200 {
		return context.JSON(code, result)
	}

	return context.JSON(http.StatusOK, result)
}
