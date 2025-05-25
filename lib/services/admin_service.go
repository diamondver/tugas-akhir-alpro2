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
	"tugas-besar/lib/repository"
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

	// LihatComment displays the comment management menu and captures the user's selection.
	// It clears the screen, displays a formatted header for the comment data view,
	// shows the current comment table, and presents an interactive menu with comment
	// management options (Search, Sorting, Add, Edit, Delete, Exit).
	LihatComment(result *string) error

	// SearchAdminComment handles the comment search functionality in the admin interface.
	// It displays a search interface that prompts the user to enter a keyword to search for,
	// performs the search using the comment repository, and displays the filtered results
	// in a table. After showing the results, it asks if the user wants to search again.
	SearchAdminComment() error

	// AddComment handles the comment creation process in the admin interface.
	// It displays a comment creation interface where admins can add new comments to the system.
	// The function collects comment text and category through a form, validates the inputs,
	// and creates a new comment record using the comment repository.
	AddComment() error

	// EditComment handles the comment editing process in the admin interface.
	// It displays the comment editing interface where admins can modify existing comments.
	// The function shows the current comment table, prompts the admin to select a comment
	// by ID, collects updated information, and saves the changes using the comment service.
	EditComment() error

	// DeleteComment handles the comment deletion process in the admin interface.
	// It displays the comment deletion interface where admins can remove existing comments.
	// The function shows the current comment table, prompts the admin to select a comment
	// by ID, and deletes the selected comment using the comment repository.
	DeleteComment() error

	// Grafik displays statistics and data visualization about comments and users.
	// It shows a summary screen with counts of total users, total comments, and comments
	// categorized by sentiment (positive, neutral, negative). The data is retrieved
	// from the comment repository and presented in a formatted display.
	Grafik() error

	// SortingKomentar handles the comment sorting functionality in the admin interface.
	// It presents an interface for selecting sorting criteria (by comment text or category)
	// and sorting mode (ascending or descending). After user selection, it retrieves
	// sorted comments from the repository and displays them in a table format.
	SortingKomentar() error
}

// adminService implements the AdminService interface and provides
// functionality for administrative operations in the application.
// It manages user-related administration tasks through the embedded UserService.
type adminService struct {
	userService    UserService
	commentService CommentService
	commentRepo    repository.CommentRepository
}

// NewAdminService creates and returns a new AdminService implementation.
//
// Parameters:
//   - userService: The UserService implementation used to perform user-related operations
//
// Returns:
//   - AdminService: A new AdminService implementation backed by the provided UserService
func NewAdminService(userService UserService, commentService CommentService, commentRepo repository.CommentRepository) AdminService {
	return &adminService{
		userService:    userService,
		commentService: commentService,
		commentRepo:    commentRepo,
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
func (a *adminService) AdminPassword() error {
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
func (a *adminService) AdminMenu(result *string) error {
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
func (a adminService) LihatUser(result *string) error {
	helper.ClearScreen()
	color.Yellow("Main Menu > Admin Menu > Lihat User")
	color.Yellow("========================================")
	color.Yellow("=              DATA USER               =")
	color.Yellow("========================================")

	err := a.ShowUserTable()
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
func (a *adminService) SearchUsers() error {
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
	err = a.userService.SearchUsers(search, &users)
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
func (a *adminService) CreateUser() error {
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

	if a.userService.IsUserExists(username, -1) {
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

	err = a.userService.CreateUser(&model.User{
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
func (a *adminService) EditUser() error {
	helper.ClearScreen()
	color.Yellow("Main Menu > Admin Menu > Lihat User > Edit")
	color.Yellow("========================================")
	color.Yellow("=              DATA USER               =")
	color.Yellow("========================================")

	err := a.ShowUserTable()
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

	if username != "" && a.userService.IsUserExists(username, index) {
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

	err = a.userService.EditUser(index, model.User{
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
func (a *adminService) DeleteUser() error {
	helper.ClearScreen()
	color.Yellow("Main Menu > Admin Menu > Lihat User > Delete")
	color.Yellow("========================================")
	color.Yellow("=              DATA USER               =")
	color.Yellow("========================================")

	err := a.ShowUserTable()
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

	err = a.userService.DeleteUser(index)
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
func (a *adminService) ShowUserTable() error {
	var users [255]model.User

	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	t.AppendHeader(table.Row{"#", "Username"})

	err := a.userService.GetAllUsers(&users)
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

// LihatComment displays the comment management menu and captures the user's selection.
//
// It clears the screen, displays a formatted header for the comment data view,
// shows the current comment table, and presents an interactive menu with comment
// management options (Search, Sorting, Add, Edit, Delete, Exit).
//
// Parameters:
//   - result: Pointer to store the selected menu option as a string
//
// Returns:
//   - error: Any error encountered during displaying the comment table or menu selection
func (a *adminService) LihatComment(result *string) error {
	helper.ClearScreen()
	color.Yellow("* MAIN MENU > ADMIN > LIHAT KOMENTAR")
	color.Yellow("========================================")
	color.Yellow("=            DATA KOMENTAR             =")
	color.Yellow("========================================")

	err := a.commentService.ShowTable()
	if err != nil {
		return err
	}

	prompt := promptui.Select{
		Label: "Pilih Menu",
		Items: []string{"Search", "Sorting", "Add", "Edit", "Delete", "Exit"},
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

// SearchAdminComment handles the comment search functionality in the admin interface.
//
// It displays a search interface that prompts the user to enter a keyword to search for,
// performs the search using the comment repository, and displays the filtered results
// in a table format. The function follows this workflow:
//
// 1. Clears the screen and displays the search interface header
// 2. Prompts user to enter a search keyword
// 3. Searches comments via commentRepo.SearchComments
// 4. Displays matching results in a formatted table
// 5. Asks if user wants to search again
//   - If yes: Returns "continue" error to loop back to search
//   - If no: Returns "back" error to go back to previous menu
//
// Returns:
//   - error: Search errors or user navigation commands ("back", "continue")
func (a *adminService) SearchAdminComment() error {
	helper.ClearScreen()
	color.Yellow("* MENU > ADMIN > LIHAT KOMENTAR > CARI KOMENTAR")
	color.Yellow("========================================")
	color.Yellow("=           CARI KOMENTAR              =")
	color.Yellow("========================================")

	searchPrompt := promptui.Prompt{
		Label: "Masukkan kata kunci untuk mencari komentar",
	}

	searchInput, err := searchPrompt.Run()
	if err != nil {
		return err
	}

	var comments [255]model.Comment
	err = a.commentRepo.SearchComments(searchInput, &comments)
	if err != nil {
		return err
	}

	helper.ClearScreen()
	color.Yellow("* MENU > ADMIN > LIHAT KOMENTAR > CARI KOMENTAR")
	color.Yellow("========================================")
	color.Yellow("=           CARI KOMENTAR              =")
	color.Yellow("========================================")
	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	t.AppendHeader(table.Row{"#", "Komentar", "Kategori"})
	var j int
	for i := 0; i < global.CommentCount; i++ {
		if comments[i].Komentar != "" {
			j++
			t.AppendRow(table.Row{
				j,
				comments[i].Komentar,
				comments[i].Kategori,
			})
		}
	}
	t.SetStyle(table.StyleColoredBright)
	t.Render()

	askPrompt := promptui.Prompt{
		Label:     "Search Again?",
		IsConfirm: true,
	}

	_, err = askPrompt.Run()
	if err != nil {
		return fmt.Errorf("back")
	}

	return fmt.Errorf("continue")
}

// AddComment handles the comment creation process in the admin interface.
//
// It displays a comment creation interface where admins can add new comments to the system.
// The function follows this workflow:
// 1. Clears the screen and displays the comment creation interface header
// 2. Collects comment text and category through CreateCommentForm
// 3. Creates a new comment record using the comment repository
//
// Error handling:
//   - Form errors: Displays the error message in red text and offers to try again
//   - If user chooses to try again: Returns "continue" error to restart the process
//   - If user chooses not to try again: Returns "back" error to go to previous menu
//   - Creation errors: Follows the same error handling pattern as form errors
//
// Returns:
//   - nil: When comment creation succeeds
//   - error: Creation errors or user navigation commands ("back", "continue")
func (a *adminService) AddComment() error {
	helper.ClearScreen()
	color.Yellow("* MENU > ADMIN > LIHAT KOMENTAR > TAMBAH KOMENTAR")
	color.Yellow("========================================")
	color.Yellow("=           TAMBAH KOMENTAR            =")
	color.Yellow("========================================")

	var komentar, kategori string

	askPrompt := promptui.Prompt{
		Label:     "Try Again?",
		IsConfirm: true,
	}

	err := a.commentService.CreateCommentForm(&komentar, &kategori)
	if err != nil {
		color.Red(err.Error())

		_, err = askPrompt.Run()
		if err != nil {
			return fmt.Errorf("back")
		}

		return fmt.Errorf("continue")
	}

	err = a.commentRepo.Create(&model.Comment{
		Komentar: komentar,
		Kategori: kategori,
	}, 0)
	if err != nil {
		color.Red(err.Error())

		_, err = askPrompt.Run()
		if err != nil {
			return fmt.Errorf("back")
		}

		return fmt.Errorf("continue")
	}

	return nil
}

// EditComment handles the comment editing process in the admin interface.
//
// It displays the comment editing interface where admins can modify existing comments.
// The function follows this workflow:
// 1. Clears the screen and displays the edit interface header
// 2. Shows the current comment table via commentService.ShowTable
// 3. Prompts admin to select a comment by ID with input validation:
//   - Ensures input is not empty
//   - Verifies input is a valid number within the range of existing comments
//
// 4. Collects updated information (comment text and category) via EditForm
// 5. Updates the comment via commentService.EditComment
// 6. Asks if admin wants to try editing again
//   - If yes: Returns "continue" error to restart the process
//   - If no: Returns "back" error to go back to previous menu
//
// Returns:
//   - error: Editing errors or user navigation commands ("back", "continue")
func (a *adminService) EditComment() error {
	helper.ClearScreen()
	color.Yellow("* MENU > ADMIN > LIHAT KOMENTAR > EDIT KOMENTAR")
	color.Yellow("========================================")
	color.Yellow("=            EDIT KOMENTAR             =")
	color.Yellow("========================================")

	err := a.commentService.ShowTable()
	if err != nil {
		return err
	}

	prompt := promptui.Prompt{
		Label: "Masukkan Id Komentar yang ingin diubah",
		Validate: func(input string) error {
			if input == "" {
				return fmt.Errorf("input tidak boleh kosong")
			}

			id, err := strconv.Atoi(input)
			if err != nil || id < 1 || id > global.IdCommentIncrement {
				return fmt.Errorf("id komentar tidak valid")
			}

			return nil
		},
	}

	idInput, err := prompt.Run()
	if err != nil {
		return err
	}

	id, err := strconv.Atoi(idInput)
	if err != nil {
		return err
	}

	var komentar, kategori string

	err = a.commentService.EditForm(&komentar, &kategori)
	if err != nil {
		return err
	}

	err = a.commentService.EditComment(id, model.Comment{
		Komentar: komentar,
		Kategori: kategori,
	})
	if err != nil {
		return err
	}

	askPrompt := promptui.Prompt{
		Label:     "Try Again?",
		IsConfirm: true,
	}

	_, err = askPrompt.Run()
	if err != nil {
		return fmt.Errorf("back")
	}

	return fmt.Errorf("continue")
}

// DeleteComment handles the comment deletion process in the admin interface.
//
// It displays the comment deletion interface where admins can remove existing comments.
// The function follows this workflow:
// 1. Clears the screen and displays the deletion interface header
// 2. Shows the current comment table via commentService.ShowTable
// 3. Prompts admin to select a comment by ID with input validation:
//   - Ensures input is not empty
//   - Verifies input is a valid number within the range of existing comments
//
// 4. Deletes the selected comment using the comment repository
// 5. If deletion fails:
//   - Displays the error message in red text
//   - Asks if admin wants to try again
//   - Returns "continue" to retry or "back" to return to previous menu
//
// Returns:
//   - nil: When comment deletion succeeds
//   - error: Deletion errors or user navigation commands ("back", "continue")
func (a *adminService) DeleteComment() error {
	helper.ClearScreen()
	color.Yellow("* MENU > ADMIN > LIHAT KOMENTAR > DELETE KOMENTAR")
	color.Yellow("========================================")
	color.Yellow("=           DELETE KOMENTAR            =")
	color.Yellow("========================================")

	err := a.commentService.ShowTable()
	if err != nil {
		return err
	}

	prompt := promptui.Prompt{
		Label: "Masukkan Id Komentar yang ingin dihapus",
		Validate: func(input string) error {
			if input == "" {
				return fmt.Errorf("input tidak boleh kosong")
			}

			id, err := strconv.Atoi(input)
			if err != nil || id < 1 || id > global.IdCommentIncrement {
				return fmt.Errorf("id komentar tidak valid")
			}

			return nil
		},
	}

	idInput, err := prompt.Run()
	if err != nil {
		return err
	}

	id, err := strconv.Atoi(idInput)
	if err != nil {
		return err
	}

	askPrompt := promptui.Prompt{
		Label:     "Try Again?",
		IsConfirm: true,
	}

	err = a.commentRepo.DeleteComment(id)
	if err != nil {
		color.Red(err.Error())

		_, err = askPrompt.Run()
		if err != nil {
			return fmt.Errorf("back")
		}

		return fmt.Errorf("continue")
	}

	return nil
}

// SortingKomentar handles the comment sorting functionality in the admin interface.
//
// It displays a sorting interface where admins can select sorting criteria and order.
// The function follows this workflow:
// 1. Clears the screen and displays the sorting interface header
// 2. Presents two selection menus to the admin:
//   - First menu: Select sorting criteria (by comment text "Komentar" or by category "Kategori")
//   - Second menu: Select sorting order (Ascending or Descending)
//
// 3. Based on the selections, calls the appropriate sorting method:
//   - sortCommentByKomentar: Sorts comments by their text content
//   - sortCommentByKategori: Sorts comments by their category
//
// The sorting mode is converted to an integer (0 for Ascending, 1 for Descending)
// before being passed to the sorting functions.
//
// Returns:
//   - error: Any error encountered during the sorting process or menu navigation
func (a *adminService) SortingKomentar() error {
	helper.ClearScreen()
	color.Yellow("* MENU > ADMIN > LIHAT KOMENTAR > SORTING")
	color.Yellow("========================================")
	color.Yellow("=               SORTING                =")
	color.Yellow("========================================")

	prompt := promptui.Select{
		Label: "Pilih Berdasarkan",
		Items: []string{"Komentar", "Kategori"},
		Templates: &promptui.SelectTemplates{
			Label:    "{{ . | blue }}:",
			Active:   "\u27A1 {{ . | cyan }}",
			Inactive: "  {{ . | cyan }}",
			Selected: "\u2705 {{ . | blue | cyan }}",
		},
	}

	promptMode := promptui.Select{
		Label: "Pilih Mode",
		Items: []string{"Ascending", "Descending"},
		Templates: &promptui.SelectTemplates{
			Label:    "{{ . | blue }}:",
			Active:   "\u27A1 {{ . | cyan }}",
			Inactive: "  {{ . | cyan }}",
			Selected: "\u2705 {{ . | blue | cyan }}",
		},
	}

	_, sortBy, err := prompt.Run()
	if err != nil {
		return err
	}

	_, sortMode, err := promptMode.Run()
	if err != nil {
		return err
	}

	modeInt := 0
	if sortMode == "Descending" {
		modeInt = 1
	}

	switch sortBy {
	case "Komentar":
		err = a.sortCommentByKomentar(modeInt)
	case "Kategori":
		err = a.sortCommentByKategori(modeInt)
	}
	if err != nil {
		return err
	}

	return nil
}

// sortCommentByKomentar sorts and displays comments based on their text content.
//
// This method sorts the comments using the comment repository's SortCommentsByComment
// function, then displays the results in a formatted table. The sorting direction
// is determined by the mode parameter.
//
// Parameters:
//   - mode: Integer determining sort order (0 for ascending, 1 for descending)
//
// The function workflow:
// 1. Retrieves sorted comments from the repository
// 2. Clears the screen and displays sorting interface header
// 3. Creates and populates a table with the sorted comments
// 4. Renders the table to standard output
// 5. Waits for user input (via Scanln) before returning
//
// Returns:
//   - error: Any error encountered during the sorting process or display
func (a *adminService) sortCommentByKomentar(mode int) error {
	var comments [255]model.Comment

	err := a.commentRepo.SortCommentsByComment(&comments, mode)
	if err != nil {
		return err
	}

	helper.ClearScreen()
	color.Yellow("* MENU > ADMIN > LIHAT KOMENTAR > SORTING")
	color.Yellow("========================================")
	color.Yellow("=               SORTING                =")
	color.Yellow("========================================")

	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	t.AppendHeader(table.Row{"#", "Komentar", "Kategori"})
	j := 0
	for i := 0; i < global.CommentCount; i++ {
		j++
		t.AppendRow(table.Row{
			j,
			comments[i].Komentar,
			comments[i].Kategori,
		})
	}
	t.SetStyle(table.StyleColoredBright)
	t.Render()

	fmt.Scanln()

	return nil
}

// sortCommentByKategori sorts and displays comments based on their category.
//
// This method sorts the comments using the comment repository's SortCommentsByKategori
// function, then displays the results in a formatted table. The sorting direction
// is determined by the mode parameter.
//
// Parameters:
//   - mode: Integer determining sort order (0 for ascending, 1 for descending)
//
// The function workflow:
// 1. Retrieves sorted comments from the repository
// 2. Clears the screen and displays sorting interface header
// 3. Creates and populates a table with the sorted comments
// 4. Renders the table to standard output
// 5. Waits for user input (via Scanln) before returning
//
// Returns:
//   - error: Any error encountered during the sorting process or display
func (a *adminService) sortCommentByKategori(mode int) error {
	var comments [255]model.Comment

	err := a.commentRepo.SortCommentsByKategori(&comments, mode)
	if err != nil {
		return err
	}

	helper.ClearScreen()
	color.Yellow("* MENU > ADMIN > LIHAT KOMENTAR > SORTING")
	color.Yellow("========================================")
	color.Yellow("=               SORTING                =")
	color.Yellow("========================================")

	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	t.AppendHeader(table.Row{"#", "Komentar", "Kategori"})
	j := 0
	for i := 0; i < global.CommentCount; i++ {
		j++
		t.AppendRow(table.Row{
			j,
			comments[i].Komentar,
			comments[i].Kategori,
		})
	}
	t.SetStyle(table.StyleColoredBright)
	t.Render()

	fmt.Scanln()

	return nil
}

// Grafik displays statistics and data visualization about comments and users.
//
// This method displays a statistical summary of the application data, including:
// - Total number of users in the system
// - Total number of comments across all categories
// - Comment distribution by sentiment categories (positive, neutral, negative)
//
// The function workflow:
// 1. Clears the screen and displays the statistics interface header
// 2. Shows the total user and comment counts from global variables
// 3. Retrieves and displays comment counts for each sentiment category:
//   - Positive comments via commentRepo.GetCommentByKategori("positif")
//   - Neutral comments via commentRepo.GetCommentByKategori("netral")
//   - Negative comments via commentRepo.GetCommentByKategori("negatif")
//
// 4. Waits for user input (via Scanln) before returning
//
// Each count is displayed in cyan text for visual clarity. If any error occurs
// during data retrieval, the function immediately returns the error.
//
// Returns:
//   - error: Any error encountered during data retrieval or display
func (a *adminService) Grafik() error {
	var comments [255]model.Comment

	helper.ClearScreen()
	color.Yellow("* MENU > ADMIN > GRAFIK")
	color.Yellow("========================================")
	color.Yellow("=                GRAFIK                =")
	color.Yellow("========================================")
	color.Cyan("Jumlah User: %d", global.UserCount)
	color.Cyan("Jumlah Komentar: %d", global.CommentCount)

	positif, err := a.commentRepo.GetCommentByKategori("Positif", &comments)
	if err != nil {
		return err
	}
	color.Cyan("Jumlah Komentar Positif: %d", positif)

	netral, err := a.commentRepo.GetCommentByKategori("Netral", &comments)
	if err != nil {
		return err
	}
	color.Cyan("Jumlah Komentar Netral: %d", netral)

	negatif, err := a.commentRepo.GetCommentByKategori("Negatif", &comments)
	if err != nil {
		return err
	}
	color.Cyan("Jumlah Komentar Negatif: %d", negatif)

	fmt.Scanln()

	return nil
}
