package controllers

import (
	"tugas-besar/lib/services"
)

// UserController handles application requests and delegates operations to the user service.
// It implements the controller logic for user functionality of the application.
type UserController struct {
	userService services.UserService
}

// NewUserController creates and returns a new UserController instance.
// It initializes the controller with the provided user service which will handle
// business logic operations for user-related functionality.
//
// Parameters:
//   - service: A services.UserService implementation that provides user operations
//
// Returns:
//   - *UserController: A pointer to the newly created UserController instance
func NewUserController(service services.UserService) *UserController {
	return &UserController{
		userService: service,
	}
}

// UserPage displays the user menu interface and captures the user's selection.
// This method delegates to the userService to display the menu and handle the user's choice.
//
// Parameters:
//   - chose: A pointer to a string that will store the user's menu selection
//
// Returns:
//   - error: An error if displaying the menu or capturing the selection fails, nil on success
func (c *UserController) UserPage(chose *string) error {
	err := c.userService.UserPage(chose)
	if err != nil {
		return err
	}
	return nil
}
