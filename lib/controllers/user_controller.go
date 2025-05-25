package controllers

import (
	"fmt"
	"tugas-besar/lib/services"

	"github.com/fatih/color"
)

// UserController handles application requests and delegates operations to the user service.
// It implements the controller logic for user functionality of the application.
type UserController struct {
	userService services.UserService
}

// NewMainController creates a new MainController instance with the provided service dependency.
// It follows the dependency injection pattern, allowing for better testability and modular design.
//
// Parameters:
//   - service: An implementation of the MainService interface
//
// Returns:
//   - A pointer to the newly created MainController
func NewUserController(service services.UserService) *UserController {
	return &UserController{
		userService: service,
	}
}

// MainMenu displays the main menu of the application and captures the user's choice.
// It delegates to the userService to handle menu display and selection logic.
//
// Parameters:
//   - result: A pointer to a string that will store the user's menu selection
//
// The function displays errors in red if any occur during menu operations
// and waits for user acknowledgment by pressing Enter before returning.
func (c *UserController) UserPage(result *string) {
	err := c.userService.UserPage(result)

	if err != nil {
		color.Red(err.Error())
		fmt.Scanln()
		return
	}
}
