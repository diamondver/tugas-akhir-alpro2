package services

import (
	"tugas-besar/lib/model"
	"tugas-besar/lib/repository"
)

// UserService defines the interface for user management operations.
// It provides methods to create, find, and check the existence of users.
type UserService interface {
	// CreateUser adds a new user to the system.
	// Returns an error if the creation fails, nil otherwise.
	CreateUser(user *model.User) error

	// FindUserByUsername retrieves a user by their username.
	// It populates the provided user model with data if found.
	// Returns an error if the user is not found, nil otherwise.
	FindUserByUsername(username string, user *model.User) error

	// IsUserExists checks if a user with the specified username exists.
	// Returns true if a user with the given username exists, false otherwise.
	IsUserExists(username string, exceptId int) bool

	// GetAllUsers retrieves all users stored in the system.
	GetAllUsers(*[255]model.User) error

	// SearchUsers finds users whose usernames contain the search string.
	SearchUsers(search string, users *[255]model.User) error

	// EditUser updates a user's information at the specified index.
	// Only non-empty fields in data will overwrite existing values.
	EditUser(index int, data model.User) error

	// DeleteUser removes a user from the system.
	DeleteUser(id int) error
}

// userService implements the UserService interface.
// It acts as a service layer between the application and the repository.
type userService struct {
	userRepo repository.UserRepository
}

// NewUserService creates and returns a new UserService implementation.
//
// Parameters:
//   - userRepo: The user repository implementation to use for data operations
//
// Returns:
//   - UserService: A new instance of the userService implementation
func NewUserService(userRepo repository.UserRepository) UserService {
	return &userService{
		userRepo: userRepo,
	}
}

// CreateUser adds a new user to the system.
// It delegates the creation operation to the underlying repository.
//
// Parameters:
//   - user: A pointer to the User model to be created
//
// Returns:
//   - error: An error if the creation fails, nil otherwise
func (userService *userService) CreateUser(user *model.User) error {
	return userService.userRepo.Create(user)
}

// FindUserByUsername retrieves a user by their username.
// It delegates the search operation to the underlying repository.
//
// Parameters:
//   - username: The username to search for
//   - user: A pointer to a User model that will be populated with the found user's data
//
// Returns:
//   - error: An error if the user is not found, nil otherwise
func (userService *userService) FindUserByUsername(username string, user *model.User) error {
	return userService.userRepo.FindUserByUsername(username, user)
}

// IsUserExists checks if a user with the specified username exists.
// It delegates the check to the underlying repository.
//
// Parameters:
//   - username: The username to check for existence
//
// Returns:
//   - bool: true if a user with the given username exists, false otherwise
func (userService *userService) IsUserExists(username string, exceptId int) bool {
	return userService.userRepo.IsUserExists(username, exceptId)
}

// GetAllUsers retrieves all users stored in the system.
// It delegates the retrieval operation to the underlying repository.
//
// Parameters:
//   - users: A pointer to an array that will be populated with all user data
//
// Returns:
//   - error: An error if the retrieval fails, nil otherwise
func (userService *userService) GetAllUsers(users *[255]model.User) error {
	return userService.userRepo.GetAllUsers(users)
}

// SearchUsers finds users whose usernames contain the search string.
// It delegates the search operation to the underlying repository.
//
// Parameters:
//   - search: The substring to search for in usernames
//   - users: A pointer to an array that will be populated with matching users
//
// Returns:
//   - error: An error if the search fails, nil otherwise
func (userService *userService) SearchUsers(search string, users *[255]model.User) error {
	return userService.userRepo.SearchUsers(search, users)
}

// EditUser updates a user's information at the specified index.
// It delegates the update operation to the underlying repository.
// Only non-empty fields in data will overwrite existing values.
//
// Parameters:
//   - index: The index of the user to update
//   - data: User model containing the fields to update
//
// Returns:
//   - error: An error if the update fails or index is invalid, nil otherwise
func (userService *userService) EditUser(index int, data model.User) error {
	return userService.userRepo.EditUser(index, data)
}

// DeleteUser removes a user from the system.
// It delegates the deletion operation to the underlying repository.
//
// Parameters:
//   - id: The index of the user to remove
//
// Returns:
//   - error: An error if the deletion fails or id is invalid, nil otherwise
func (userService *userService) DeleteUser(id int) error {
	return userService.userRepo.DeleteUser(id)
}
