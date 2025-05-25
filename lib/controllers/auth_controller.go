package controllers

import (
	"fmt"
	"github.com/fatih/color"
	"tugas-besar/lib/model"
	"tugas-besar/lib/services"
)

// AuthController handles authentication-related operations by delegating
// to the AuthService layer.
type AuthController struct {
	authService services.AuthService
}

// NewAuthController creates a new AuthController with the provided AuthService.
// Parameters:
//   - service: An implementation of the AuthService interface
//
// Returns:
//   - *AuthController: A pointer to the newly created controller
func NewAuthController(service services.AuthService) *AuthController {
	return &AuthController{
		authService: service,
	}
}

// Login attempts to authenticate a user with the provided credentials.
// It displays an error message if authentication fails.
//
// Parameters:
//   - user: A pointer to a User model containing login credentials
func (c *AuthController) Login(user *model.User) {
	for {
		err := c.authService.Login(user)
		if err != nil {
			if err.Error() == "back" {
				break
			}

			if err.Error() == "continue" {
				continue
			}

			color.Red(err.Error())
			fmt.Scanln()
			continue
		} else {
			break
		}
	}
}

// Register handles the user registration process.
// It displays an error message if registration fails.
//
// Returns:
//   - None, but prompts for user input and handles errors internally
func (c *AuthController) Register() {
	//err := c.authService.Register()
	//if err != nil {
	//	color.Red(err.Error())
	//	fmt.Scanln()
	//	return
	//}
	//
	//color.Green("Registration successful! Please login to continue.")
	//fmt.Scanln()

	for {
		err := c.authService.Register()
		if err != nil {
			if err.Error() == "back" {
				break
			}

			if err.Error() == "continue" {
				continue
			}

			color.Red(err.Error())
			fmt.Scanln()
			continue
		} else {
			color.Green("Registration successful! Please login to continue.")
			fmt.Scanln()
			break
		}
	}
}
