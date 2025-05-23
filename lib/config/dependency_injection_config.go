package config

import (
	"tugas-besar/lib/controllers"
	"tugas-besar/lib/services"
)

// AppContainer holds references to controllers that have been initialized with
// their required dependencies. It serves as a central access point for all
// properly configured controllers in the application.
type AppContainer struct {
	MainController *controllers.MainController
}

// DependencyConfig initializes and wires all application dependencies.
// It creates service instances and injects them into the appropriate controllers,
// following the dependency injection pattern.
// Returns an AppContainer with all initialized controllers ready for use.
func DependencyConfig() *AppContainer {
	mainService := services.NewMainService()
	mainController := controllers.NewMainController(mainService)

	return &AppContainer{
		MainController: mainController,
	}
}
