package controllers

import (
	"fmt"
	"github.com/fatih/color"
	"tugas-besar/lib/services"
)

// AdminController manages administrative operations through the admin service.
// It provides methods for user management, authentication, and other admin tasks.
type AdminController struct {
	// adminService handles the business logic for admin operations
	adminService services.AdminService
}

// NewAdminController creates and returns a new AdminController instance.
// It takes a services.AdminService implementation as a dependency for performing
// admin-related operations.
func NewAdminController(service services.AdminService) *AdminController {
	return &AdminController{
		adminService: service,
	}
}

// AdminMenu displays and handles the administrative menu interface.
// It manages authentication flow and provides access to administrative functions.
//
// The function runs in a continuous loop until the user selects "Exit" from the menu.
// It first checks if the user is authenticated, and if not, prompts for admin credentials.
// After successful authentication, it displays the admin menu and processes user selections.
//
// The menu supports the following operations:
// - "Lihat User": View and manage user accounts
// - "Exit": Return to the previous menu
//
// Authentication errors with message "back" will cause immediate return from the function.
// Other errors are displayed to the user in red text.
func (c *AdminController) AdminMenu() {
	var result string
	var isAuthenticated bool

	for {
		if !isAuthenticated {
			err := c.adminService.AdminPassword()
			if err != nil {
				if err.Error() == "back" {
					return
				}

				color.Red(err.Error())
				fmt.Scanln()
				continue
			}
		}

		isAuthenticated = true

		err := c.adminService.AdminMenu(&result)
		if err != nil {
			color.Red(err.Error())
			fmt.Scanln()
		}

		if result == "Exit" {
			break
		}

		switch result {
		case "Lihat User":
			c.adminLihatUser()
		}
	}
}

// adminLihatUser handles the user management menu in the admin interface.
//
// It displays a menu for managing user accounts through the admin service and processes
// the user's selection in a continuous loop until "Exit" is chosen.
//
// The method supports the following operations:
// - "Search": Search for users
// - "Add": Create a new user
// - "Edit": Modify an existing user
// - "Delete": Remove a user
// - "Exit": Return to the previous menu
//
// Any errors encountered while displaying the menu are shown to the user in red text.
// The function handles navigation between different user management functions based on
// the selected option.
func (c *AdminController) adminLihatUser() {
	var result string

	for {
		err := c.adminService.LihatUser(&result)
		if err != nil {
			color.Red(err.Error())
			fmt.Scanln()
		}

		if result == "Exit" {
			break
		}

		switch result {
		case "Search":
			c.userSearch()
		case "Add":
			c.CreateUser()
		case "Edit":
			c.EditUser()
		case "Delete":
			c.DeleteUser()
		}
	}
}

// userSearch handles the user search functionality in the admin interface.
//
// It runs in a continuous loop, calling the SearchUsers method from the admin service
// until a terminating condition is met. The function processes different error types:
//
// Error handling:
//   - "back": Returns to the previous menu
//   - "continue": Restarts the search process
//   - Other errors: Displays the error message in red text, waits for user input,
//     and returns to the previous menu
//
// The function terminates when either a "back" error is received, a non-"continue"
// error occurs, or when the SearchUsers method completes successfully.
func (c *AdminController) userSearch() {
	for {
		err := c.adminService.SearchUsers()
		if err != nil {
			if err.Error() == "back" {
				break
			}

			if err.Error() == "continue" {
				continue
			}

			color.Red(err.Error())
			fmt.Scanln()
			break
		}
	}
}

// CreateUser handles the user creation functionality in the admin interface.
//
// It runs in a continuous loop, calling the CreateUser method from the admin service
// until a terminating condition is met. The function processes different error types:
//
// Error handling:
//   - "back": Returns to the previous menu
//   - "continue": Restarts the user creation process
//   - Other errors: Displays the error message in red text, waits for user input,
//     and returns to the previous menu
//
// On successful user creation, the function displays a success message in green,
// waits for user input, and returns to the previous menu.
func (c *AdminController) CreateUser() {
	for {
		err := c.adminService.CreateUser()
		if err != nil {
			if err.Error() == "back" {
				break
			}

			if err.Error() == "continue" {
				continue
			}

			color.Red(err.Error())
			fmt.Scanln()
			break
		}

		color.Green("User created successfully!")
		fmt.Scanln()
		break
	}
}

// EditUser handles the user editing functionality in the admin interface.
//
// It runs in a continuous loop, calling the EditUser method from the admin service
// until a terminating condition is met. The function processes different error types:
//
// Error handling:
//   - "back": Returns to the previous menu
//   - "continue": Restarts the user editing process
//   - Other errors: Displays the error message in red text, waits for user input,
//     and returns to the previous menu
//
// On successful user editing, the function displays a success message in green,
// waits for user input, and returns to the previous menu.
func (c *AdminController) EditUser() {
	for {
		err := c.adminService.EditUser()
		if err != nil {
			if err.Error() == "back" {
				break
			}

			if err.Error() == "continue" {
				continue
			}

			color.Red(err.Error())
			fmt.Scanln()
			break
		}

		color.Green("User edited successfully!")
		fmt.Scanln()
		break
	}
}

// DeleteUser handles the user deletion functionality in the admin interface.
//
// It runs in a continuous loop, calling the DeleteUser method from the admin service
// until a terminating condition is met. The function processes different error types:
//
// Error handling:
//   - "back": Returns to the previous menu
//   - "continue": Restarts the user deletion process
//   - Other errors: Displays the error message in red text, waits for user input,
//     and returns to the previous menu
//
// On successful user deletion, the function displays a success message in green,
// waits for user input, and returns to the previous menu.
func (c *AdminController) DeleteUser() {
	for {
		err := c.adminService.DeleteUser()
		if err != nil {
			if err.Error() == "back" {
				break
			}

			if err.Error() == "continue" {
				continue
			}

			color.Red(err.Error())
			fmt.Scanln()
			break
		}

		color.Green("User deleted successfully!")
		fmt.Scanln()
		break
	}
}
