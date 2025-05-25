package lib

import (
	"tugas-besar/lib/config"
	"tugas-besar/lib/model"
)

// Bootstrap initializes the application by loading environment configurations.
// It calls config.GetEnvConfig() to load environment variables from the .env file.
// After initializing configurations, it enters an infinite loop to keep the
// application running. This function is called from the main function to start
// the application processes.
//
// The function does not accept any parameters and does not return any values.
func Bootstrap() {
	var result string
	var user model.User

	// Configuration
	config.GetEnvConfig()

	// Dependency Injection
	container := config.DependencyConfig()

	for {
		container.MainController.MainMenu(&result)

		if result == "Exit" {
			break
		}

		switch result {
		case "Login":
			container.AuthController.Login(&user)
			if user.Username != "" {
				for {
					err := container.UserController.UserPage(&result)
					if err != nil {
						break
					}

					if result == "Exit" {
						user.Username = ""
						user.Password = ""
						break
					}

					switch result {
					case "Tambah Komentar":
						container.CommentController.CommentInputPage(user)
					case "Lihat Komentar":
						container.CommentController.CommentView()
					case "Edit Komentar":
						container.CommentController.EditComment(user)
					case "Delete Komentar":
						container.CommentController.DeleteComment(user)
					}
				}
			}
		case "Register":
			container.AuthController.Register()
		case "Admin":
			container.AdminController.AdminMenu()
		}
	}

}
