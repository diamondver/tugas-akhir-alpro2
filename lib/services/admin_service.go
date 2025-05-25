package services

import (
	"fmt"
	"os"
	"strconv"

	"github.com/fatih/color"
	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/manifoldco/promptui"

	"tugas-besar/lib/global"
	"tugas-besar/lib/helper"
	"tugas-besar/lib/model"
)

// AdminService defines the interface for administrative operations in the application.
//
// This service provides functionality for admin authentication, user management,
// and navigation through the admin interface. It handles displaying menus,
// processing user selections, and performing CRUD operations on user accounts.
type AdminService interface {
	// AdminMenu displays the main admin menu and captures the user's selection.
	AdminMenu(result *string) error

	// AdminPassword validates the admin password for authentication.
	AdminPassword() error

	// LihatUser displays the user management menu and captures the user's selection.
	LihatUser(result *string) error

	// SearchUsers handles the user search functionality.
	SearchUsers() error

	// CreateUser handles the user creation process.
	CreateUser() error

	// EditUser handles the user editing process.
	EditUser() error

	// DeleteUser handles the user deletion process.
	DeleteUser() error
}

// adminService implements the AdminService interface and provides
// functionality for administrative operations in the application.
// It manages user-related administration tasks through the embedded UserService.
type adminService struct {
	userService UserService
}

// NewAdminService creates and returns a new AdminService implementation.
//
// Parameters:
//   - userService: The UserService implementation used to perform user-related operations
//
// Returns:
//   - AdminService: A new AdminService implementation backed by the provided UserService
func NewAdminService(userService UserService) AdminService {
	return &adminService{
		userService: userService,
	}
}

// AdminPassword validates the admin password for authentication.
//
// It retrieves the admin password from environment variables and prompts the user
// to enter the password for validation. If no password is set in the environment,
// authentication is skipped. The function handles different scenarios:
//
// - When password matches: Displays success message and returns nil
// - When password doesn't match: Offers the user to try again
//   - If user chooses to try again: Returns "continue" error
//   - If user chooses not to try again: Returns "back" error
//
// Returns:
//   - nil: When authentication succeeds or no password is required
//   - error: Authentication errors or user navigation commands ("back", "continue")
func (service *adminService) AdminPassword() error {
	var password = helper.GetEnv("ADMIN_PASS", "")

	helper.ClearScreen()
	color.Yellow("Main Menu > Admin Menu")
	color.Yellow("========================================")
	color.Yellow("=              ADMIN MENU              =")
	color.Yellow("========================================")

	if password == "" {
		return nil
	}

	prompt := promptui.Prompt{
		Label: "Masukkan Password Admin",
		Mask:  '*',
	}

	result, err := prompt.Run()
	if err != nil {
		return err
	}

	if result == password {
		color.Green("Password matched successfully!")
		fmt.Scanln()
		return nil
	}

	color.Red("Passwords do not match")

	askPrompt := promptui.Prompt{
		Label:     "Apakah Anda ingin mencoba lagi?",
		IsConfirm: true,
	}

	_, err = askPrompt.Run()
	if err != nil {
		return fmt.Errorf("back")
	}

	return fmt.Errorf("continue")
}

// AdminMenu displays the main admin menu and captures the user's selection.
//
// It clears the screen, displays a formatted menu header, and presents
// a selection interface with various admin options (Lihat Komentar, Lihat User,
// Lihat Grafik, Exit). The function uses promptui to create an interactive
// selection interface with custom styling for menu items.
//
// Parameters:
//   - result: Pointer to store the selected menu option as a string
//
// Returns:
//   - error: Any error encountered during menu display or selection process
func (service *adminService) AdminMenu(result *string) error {
	helper.ClearScreen()
	color.Yellow("Main Menu > Admin Menu")
	color.Yellow("========================================")
	color.Yellow("=              ADMIN MENU              =")
	color.Yellow("========================================")

	prompt := promptui.Select{
		Label: "Pilih Menu",
		Items: []string{"Lihat Komentar", "Lihat User", "Lihat Grafik", "Exit"},
		Templates: &promptui.SelectTemplates{
			Label:    "{{ . | blue }}:",
			Active:   "\u27A1 {{ . | cyan }}",
			Inactive: "  {{ . | cyan }}",
			Selected: "\u2705 {{ . | blue | cyan }}",
		},
	}

	_, resultInput, err := prompt.Run()
	if err != nil {
		return err
	}

	*result = resultInput

	return nil
}

// LihatUser displays the user management menu and captures the user's selection.
//
// It clears the screen, displays a formatted header for the user data view,
// shows the current user table by calling ShowUserTable(), and presents an
// interactive menu with user management options (Search, Add, Edit, Delete, Exit).
// The function uses promptui to create an interactive selection interface with
// custom styling for menu items.
//
// Parameters:
//   - result: Pointer to store the selected menu option as a string
//
// Returns:
//   - error: Any error encountered during displaying the user table or menu selection
func (service adminService) LihatUser(result *string) error {
	helper.ClearScreen()
	color.Yellow("Main Menu > Admin Menu > Lihat User")
	color.Yellow("========================================")
	color.Yellow("=              DATA USER               =")
	color.Yellow("========================================")

	err := service.ShowUserTable()
	if err != nil {
		return err
	}

	prompt := promptui.Select{
		Label: "Pilih Menu",
		Items: []string{"Search", "Add", "Edit", "Delete", "Exit"},
		Templates: &promptui.SelectTemplates{
			Label:    "{{ . | blue }}:",
			Active:   "\u27A1 {{ . | cyan }}",
			Inactive: "  {{ . | cyan }}",
			Selected: "\u2705 {{ . | blue | cyan }}",
		},
	}

	_, resultPrompt, err := prompt.Run()
	if err != nil {
		return err
	}

	*result = resultPrompt

	return nil
}

// SearchUsers handles the user search functionality.
//
// It displays a search interface that prompts the user to enter a username
// to search for, performs the search using the underlying userService, and
// displays the filtered results in a table. After showing the results, it
// asks if the user wants to search again, handling navigation accordingly.
//
// The function follows this workflow:
// 1. Clear screen and display the search interface header
// 2. Prompt user to enter a username to search for
// 3. Execute the search via userService.SearchUsers
// 4. Display results in a table via ShowUserTable
// 5. Ask if user wants to search again
//   - If yes: Return "continue" error to loop back to search
//   - If no: Return "back" error to go back to previous menu
//
// Returns:
//   - error: Search errors or user navigation commands ("back", "continue")
func (service *adminService) SearchUsers() error {
	helper.ClearScreen()
	color.Yellow("Main Menu > Admin Menu > Lihat User > Search")
	color.Yellow("========================================")
	color.Yellow("=              DATA USER               =")
	color.Yellow("========================================")

	prompt := promptui.Prompt{
		Label: "Masukkan Username yang ingin dicari",
	}

	askPrompt := promptui.Prompt{
		Label:     "Search Again?",
		IsConfirm: true,
	}

	search, err := prompt.Run()
	if err != nil {
		return err
	}

	var users [255]model.User
	err = service.userService.SearchUsers(search, &users)
	if err != nil {
		return err
	}

	helper.ClearScreen()
	color.Yellow("Main Menu > Admin Menu > Lihat User > Search")
	color.Yellow("========================================")
	color.Yellow("=              DATA USER               =")
	color.Yellow("========================================")

	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	t.AppendHeader(table.Row{"#", "Username"})
	var j int
	for i := 0; i < global.UserCount; i++ {
		if users[i].Username != "" {
			j++
			t.AppendRow(table.Row{j, users[i].Username})
		}
	}
	t.SetStyle(table.StyleColoredBright)
	t.Render()

	_, err = askPrompt.Run()
	if err != nil {
		return fmt.Errorf("back")
	}

	return fmt.Errorf("continue")
}

// CreateUser handles the user creation process.
//
// It displays a user creation interface where admins can add new users to the system.
// The function follows this workflow:
// 1. Clear screen and display the user creation interface header
// 2. Prompt admin to enter username, password, and confirm password via createUserForm
// 3. Validate the inputs:
//   - Check if username already exists using userService.IsUserExists
//   - Verify that password and confirmPassword match
//
// 4. If validation fails:
//   - Display appropriate error message
//   - Prompt admin to try again
//   - Return "continue" to retry or "back" to return to previous menu
//
// 5. If validation passes, create the user via userService.CreateUser
//
// Returns:
//   - nil: When user creation succeeds
//   - error: Creation errors or user navigation commands ("back", "continue")
func (service *adminService) CreateUser() error {
	helper.ClearScreen()
	color.Yellow("Main Menu > Admin Menu > Lihat User > Add")
	color.Yellow("========================================")
	color.Yellow("=              DATA USER               =")
	color.Yellow("========================================")

	var username, password, confirmPassword string

	err := createUserForm(&username, &password, &confirmPassword)
	if err != nil {
		return err
	}

	askPrompt := promptui.Prompt{
		Label:     "Try Again?",
		IsConfirm: true,
	}

	if service.userService.IsUserExists(username, -1) {
		color.Red("User %s already exists", username)
		_, err = askPrompt.Run()
		if err != nil {
			return fmt.Errorf("back")
		}

		return fmt.Errorf("continue")
	}

	if password != confirmPassword {
		color.Red("Password does not match")
		_, err = askPrompt.Run()
		if err != nil {
			return fmt.Errorf("back")
		}

		return fmt.Errorf("continue")
	}

	err = service.userService.CreateUser(&model.User{
		Username: username,
		Password: password,
	})
	if err != nil {
		return err
	}

	return nil
}

// createUserForm collects user credentials through an interactive command-line interface.
//
// This helper function creates a series of prompts for username, password, and password
// confirmation. It uses the promptui library to display labeled prompts with appropriate
// masking for password fields. The collected inputs are assigned to the provided pointers.
//
// Parameters:
//   - username: Pointer to store the collected username
//   - password: Pointer to store the collected password
//   - confirmPassword: Pointer to store the password confirmation input
//
// Returns:
//   - error: Any error encountered during the prompt process
func createUserForm(username, password, confirmPassword *string) error {
	usernamePrompt := promptui.Prompt{Label: "Username"}
	passwordPrompt := promptui.Prompt{Label: "Password", Mask: '*'}
	confirmPasswordPrompt := promptui.Prompt{Label: "Confirm Password", Mask: '*'}

	usernameInput, err := usernamePrompt.Run()
	if err != nil {
		return err
	}

	passwordInput, err := passwordPrompt.Run()
	if err != nil {
		return err
	}

	confirmPasswordInput, err := confirmPasswordPrompt.Run()
	if err != nil {
		return err
	}

	*username = usernameInput
	*password = passwordInput
	*confirmPassword = confirmPasswordInput

	return nil
}

// EditUser handles the user editing process.
//
// It displays the user editing interface where admins can modify existing user information.
// The function follows this workflow:
// 1. Clear screen and display the edit interface header
// 2. Show the current user table via ShowUserTable
// 3. Prompt admin to select a user by number with input validation:
//   - Ensure input is not empty
//   - Verify input is a valid number within the range of existing users
//
// 4. Collect updated information (username, password) via editUserForm
// 5. Validate the inputs:
//   - If username is provided, check it doesn't conflict with existing users
//   - If password is provided, verify that it matches the confirmation
//
// 6. If validation fails:
//   - Display appropriate error message
//   - Prompt admin to try again
//   - Return "continue" to retry or "back" to return to previous menu
//
// 7. If validation passes, update the user via userService.EditUser
//
// Returns:
//   - nil: When user editing succeeds
//   - error: Editing errors or user navigation commands ("back", "continue")
func (service *adminService) EditUser() error {
	helper.ClearScreen()
	color.Yellow("Main Menu > Admin Menu > Lihat User > Edit")
	color.Yellow("========================================")
	color.Yellow("=              DATA USER               =")
	color.Yellow("========================================")

	err := service.ShowUserTable()
	if err != nil {
		return err
	}

	prompt := promptui.Prompt{
		Label: "Masukkan Nomor User yang ingin diubah",
		Validate: func(input string) error {
			if input == "" {
				return fmt.Errorf("input cannot be empty")
			}

			index, err := strconv.Atoi(input)
			if err != nil || index < 1 || index > global.UserCount {
				return fmt.Errorf("invalid user number")
			}

			return nil
		},
	}

	askPrompt := promptui.Prompt{
		Label:     "Try Again?",
		IsConfirm: true,
	}

	indexInput, err := prompt.Run()
	if err != nil {
		color.Red(err.Error())

		_, err = askPrompt.Run()
		if err != nil {
			return fmt.Errorf("back")
		}

		return fmt.Errorf("continue")
	}

	index, err := strconv.Atoi(indexInput)
	if err != nil {
		color.Red(err.Error())

		_, err = askPrompt.Run()
		if err != nil {
			return fmt.Errorf("back")
		}

		return fmt.Errorf("continue")
	}

	index--

	var username, password, confirmPassword string
	err = editUserForm(&username, &password, &confirmPassword)
	if err != nil {
		return err
	}

	if username != "" && service.userService.IsUserExists(username, index) {
		color.Red("User %s already exists", username)

		_, err = askPrompt.Run()
		if err != nil {
			return fmt.Errorf("back")
		}

		return fmt.Errorf("continue")
	}

	if password != "" && password != confirmPassword {
		color.Red("Password does not match")

		_, err = askPrompt.Run()
		if err != nil {
			return fmt.Errorf("back")
		}

		return fmt.Errorf("continue")
	}

	err = service.userService.EditUser(index, model.User{
		Username: username,
		Password: password,
	})
	if err != nil {
		return err
	}

	return nil
}

// editUserForm collects updated user credentials through an interactive command-line interface.
//
// This helper function creates a series of prompts for username, password, and password
// confirmation during the user editing process. It uses the promptui library to display
// labeled prompts with appropriate masking for password fields. Unlike createUserForm,
// this function is specifically designed for the edit workflow where fields can be
// left empty to keep existing values.
//
// Parameters:
//   - username: Pointer to store the collected username (can be empty to keep existing)
//   - password: Pointer to store the collected password (can be empty to keep existing)
//   - confirmPassword: Pointer to store the password confirmation input
//
// Returns:
//   - error: Any error encountered during the prompt process
func editUserForm(username, password, confirmPassword *string) error {
	usernamePrompt := promptui.Prompt{Label: "Username"}
	passwordPrompt := promptui.Prompt{Label: "Password", Mask: '*'}
	confirmPasswordPrompt := promptui.Prompt{Label: "Confirm Password", Mask: '*'}

	usernameInput, err := usernamePrompt.Run()
	if err != nil {
		return err
	}

	passwordInput, err := passwordPrompt.Run()
	if err != nil {
		return err
	}

	confirmPasswordInput, err := confirmPasswordPrompt.Run()
	if err != nil {
		return err
	}

	*username = usernameInput
	*password = passwordInput
	*confirmPassword = confirmPasswordInput

	return nil
}

// DeleteUser handles the user deletion process.
//
// It displays the user deletion interface where admins can remove existing users from the system.
// The function follows this workflow:
// 1. Clear screen and display the delete interface header
// 2. Show the current user table via ShowUserTable
// 3. Prompt admin to select a user by number with input validation:
//   - Ensure input is not empty
//   - Verify input is a valid number within the range of existing users
//
// 4. If validation fails:
//   - Display appropriate error message
//   - Prompt admin to try again
//   - Return "continue" to retry or "back" to return to previous menu
//
// 5. If validation passes, delete the user via userService.DeleteUser
// 6. Display success message
//
// Returns:
//   - nil: When user deletion succeeds
//   - error: Deletion errors or user navigation commands ("back", "continue")
func (service *adminService) DeleteUser() error {
	helper.ClearScreen()
	color.Yellow("Main Menu > Admin Menu > Lihat User > Delete")
	color.Yellow("========================================")
	color.Yellow("=              DATA USER               =")
	color.Yellow("========================================")

	err := service.ShowUserTable()
	if err != nil {
		return err
	}

	prompt := promptui.Prompt{
		Label: "Masukkan Nomor User yang ingin dihapus",
		Validate: func(input string) error {
			if input == "" {
				return fmt.Errorf("input cannot be empty")
			}

			index, err := strconv.Atoi(input)
			if err != nil || index < 1 || index > global.UserCount {
				return fmt.Errorf("invalid user number")
			}

			return nil
		},
	}

	askPrompt := promptui.Prompt{
		Label:     "Try Again?",
		IsConfirm: true,
	}

	indexInput, err := prompt.Run()
	if err != nil {
		color.Red(err.Error())

		_, err = askPrompt.Run()
		if err != nil {
			return fmt.Errorf("back")
		}

		return fmt.Errorf("continue")
	}

	index, err := strconv.Atoi(indexInput)
	if err != nil {
		color.Red(err.Error())

		_, err = askPrompt.Run()
		if err != nil {
			return fmt.Errorf("back")
		}

		return fmt.Errorf("continue")
	}

	index--

	err = service.userService.DeleteUser(index)
	if err != nil {
		return err
	}

	color.Green("User deleted successfully")
	return nil
}

// ShowUserTable displays a formatted table of all users in the system.
//
// It retrieves all users from the userService and renders them as a table
// to standard output using the go-pretty/table package. The table includes
// row numbers and usernames with colored formatting for better readability.
//
// Returns:
//   - error: Any error encountered during user data retrieval
func (service *adminService) ShowUserTable() error {
	var users [255]model.User

	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	t.AppendHeader(table.Row{"#", "Username"})

	err := service.userService.GetAllUsers(&users)
	if err != nil {
		return err
	}

	for i := 0; i < global.UserCount; i++ {
		t.AppendRow(table.Row{i + 1, users[i].Username})
	}

	t.SetStyle(table.StyleColoredBright)
	t.Render()

	return nil
}
