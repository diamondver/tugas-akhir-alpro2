package controllers

import (
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

func (c *MainController) MainMenu() {
	c.mainService.MainMenu()
}
