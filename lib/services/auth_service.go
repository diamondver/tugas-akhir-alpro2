package services

import (
	"fmt"
	"github.com/fatih/color"
	"github.com/manifoldco/promptui"
	"tugas-besar/lib/helper"
	"tugas-besar/lib/model"
)

// AuthService defines the interface for authentication operations
// in the application, providing methods for user login and registration.
type AuthService interface {
	// Login authenticates a user with the provided credentials.
	// It takes a user model pointer that will be populated with user data on success.
	// Returns an error if authentication fails, nil otherwise.
	Login(user *model.User) error

	// Register handles the user registration process.
	// It collects and validates user information before creating a new account.
	// Returns an error if registration fails, nil otherwise.
	Register() error
}

// authService implements the AuthService interface and handles
// authentication logic by delegating user operations to UserService.
type authService struct {
	userService UserService
}

// NewAuthService creates and returns a new AuthService implementation.
// Parameters:
//   - userService: The UserService implementation to use for user operations
//
// Returns:
//   - AuthService: A new AuthService implementation
func NewAuthService(userService UserService) AuthService {
	return &authService{
		userService: userService,
	}
}

// Login handles the user authentication process.
// It displays a login form, clears the screen, and presents a formatted login interface.
// The method collects user credentials, validates them against stored user data,
// and checks password correctness.
//
// Parameters:
//   - user: A pointer to the User model that will be populated with user data on successful login
//
// Returns:
//   - error: An error if login fails (form interaction, user not found, or incorrect password), nil otherwise
func (service *authService) Login(user *model.User) error {
	var username, password string

	helper.ClearScreen()
	color.Yellow("Main Menu > Login")
	color.Yellow("=========================================")
	color.Yellow("=                LOGIN                  =")
	color.Yellow("=========================================")

	err := loginForm(&username, &password)
	if err != nil {
		return err
	}

	askPrompt := promptui.Prompt{
		Label:     "Do you want to try again?",
		IsConfirm: true,
	}

	err = service.userService.FindUserByUsername(username, user)
	if err != nil {
		color.Red("User not found: %s", username)
		_, err = askPrompt.Run()
		if err != nil {
			return fmt.Errorf("back")
		}

		return fmt.Errorf("continue")
	}

	if user.Password != password {
		color.Red("Password does not match")
		_, err = askPrompt.Run()
		if err != nil {
			return fmt.Errorf("back")
		}

		return fmt.Errorf("continue")
	}

	color.Green("Login successful! Welcome, %s!", user.Username)
	fmt.Scanln()

	return nil
}

// loginForm displays interactive prompts to collect username and password.
// It uses promptui to create formatted input fields with appropriate masking for the password.
//
// Parameters:
//   - username: A pointer to a string that will be populated with the entered username
//   - password: A pointer to a string that will be populated with the entered password
//
// Returns:
//   - error: An error if the prompt interaction fails, nil otherwise
func loginForm(username, password *string) error {
	usernamePrompt := promptui.Prompt{Label: "Username"}
	passwordPrompt := promptui.Prompt{Label: "Password", Mask: '*'}

	usernameInput, err := usernamePrompt.Run()
	if err != nil {
		return err
	}

	passwordInput, err := passwordPrompt.Run()
	if err != nil {
		return err
	}

	*username = usernameInput
	*password = passwordInput

	return nil
}

// Register handles the user registration process.
// It displays a registration form, clears the screen, and presents a formatted registration interface.
// The method collects user credentials, validates password confirmation,
// and creates a new user account.
//
// Returns:
//   - error: An error if registration fails (form interaction, password mismatch,
//     or user creation error), nil otherwise
func (service *authService) Register() error {
	var username, password, confirmPassword string

	helper.ClearScreen()
	color.Yellow("Main Menu > Register")
	color.Yellow("=========================================")
	color.Yellow("=                REGISTER               =")
	color.Yellow("=========================================")

	err := registerForm(&username, &password, &confirmPassword)
	if err != nil {
		return err
	}

	askPrompt := promptui.Prompt{
		Label:     "Do you want to try again?",
		IsConfirm: true,
	}

	if service.userService.IsUserExists(username) {
		color.Red("User with username %s already exists", username)
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

// registerForm displays interactive prompts to collect registration information.
// It uses promptui to create formatted input fields with appropriate masking for passwords.
//
// Parameters:
//   - username: A pointer to a string that will be populated with the entered username
//   - password: A pointer to a string that will be populated with the entered password
//   - confirmPassword: A pointer to a string that will be populated with the password confirmation
//
// Returns:
//   - error: An error if the prompt interaction fails, nil otherwise
func registerForm(username, password, confirmPassword *string) error {
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
