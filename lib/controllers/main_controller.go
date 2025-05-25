package controllers

import (
	"fmt"
	"github.com/fatih/color"
	"tugas-besar/lib/services"
)

// MainController handles application requests and delegates operations to the main service.
// It implements the controller logic for main functionality of the application.
type MainController struct {
	mainService services.MainService
}

// NewMainController creates a new MainController instance with the provided service dependency.
// It follows the dependency injection pattern, allowing for better testability and modular design.
//
// Parameters:
//   - service: An implementation of the MainService interface
//
// Returns:
//   - A pointer to the newly created MainController
func NewMainController(service services.MainService) *MainController {
	return &MainController{
		mainService: service,
	}
}

// MainMenu displays the main menu of the application and captures the user's choice.
// It delegates to the mainService to handle menu display and selection logic.
//
// Parameters:
//   - result: A pointer to a string that will store the user's menu selection
//
// The function displays errors in red if any occur during menu operations
// and waits for user acknowledgment by pressing Enter before returning.
func (c *MainController) MainMenu(result *string) {
	err := c.mainService.MainMenu(result)

	if err != nil {
		color.Red(err.Error())
		fmt.Scanln()
		return
	}
}
