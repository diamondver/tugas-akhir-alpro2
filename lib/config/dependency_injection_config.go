package config

import (
	"tugas-besar/lib/controllers"
	"tugas-besar/lib/repository"
	"tugas-besar/lib/services"
)

// AppContainer holds references to controllers that have been initialized with
// their required dependencies. It serves as a central access point for all
// properly configured controllers in the application.
type AppContainer struct {
	MainController  *controllers.MainController
	AuthController  *controllers.AuthController
	AdminController *controllers.AdminController
}

// DependencyConfig initializes and wires all application dependencies.
// It creates service instances and injects them into the appropriate controllers,
// following the dependency injection pattern.
// Returns an AppContainer with all initialized controllers ready for use.
func DependencyConfig() *AppContainer {
	mainService := services.NewMainService()
	mainController := controllers.NewMainController(mainService)

	userService := services.NewUserService(repository.NewUserRepository())

	authService := services.NewAuthService(userService)
	authController := controllers.NewAuthController(authService)

	adminService := services.NewAdminService(userService)
	adminController := controllers.NewAdminController(adminService)

	return &AppContainer{
		MainController:  mainController,
		AuthController:  authController,
		AdminController: adminController,
	}
}
