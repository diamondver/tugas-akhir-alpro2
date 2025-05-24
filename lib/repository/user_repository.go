package repository

import (
	"fmt"
	"tugas-besar/lib/global"
	"tugas-besar/lib/model"
)

// userRepository implements the UserRepository interface using an in-memory
// storage mechanism for user data.
type userRepository struct {
}

// UserRepository defines the interface for user data operations.
// It provides methods to create new users and retrieve existing users by username.
type UserRepository interface {
	// Create adds a new user to the repository.
	// Returns an error if the operation fails, nil otherwise.
	Create(user *model.User) error

	// FindUserByUsername retrieves a user by their username.
	// It populates the provided user model with data if found.
	// Returns an error if the user is not found, nil otherwise.
	FindUserByUsername(username string, user *model.User) error

	// IsUserExists checks if a user with the given username exists in the repository.
	// Returns true if the user exists, false otherwise.
	IsUserExists(username string) bool
}

// NewUserRepository creates and returns a new UserRepository implementation.
//
// Returns:
//   - UserRepository: A new instance of the userRepository implementation
func NewUserRepository() UserRepository {
	return &userRepository{}
}

// Create adds a new user to the in-memory repository.
// The user is assigned the next available index in the global user storage.
//
// Parameters:
//   - user: A pointer to the User model to be stored
//
// Returns:
//   - error: Always returns nil as this implementation doesn't have failure cases
func (repo *userRepository) Create(user *model.User) error {
	global.Users[global.UserCount] = *user
	global.UserCount++

	return nil
}

// FindUserByUsername searches for a user by their username in the repository.
// If found, it populates the provided user model with the user's data.
//
// Parameters:
//   - username: The username to search for
//   - user: A pointer to a User model that will be populated with the found user's data
//
// Returns:
//   - error: An error with a descriptive message if the user is not found, nil otherwise
func (repo *userRepository) FindUserByUsername(username string, user *model.User) error {
	for i := 0; i < global.UserCount; i++ {
		if global.Users[i].Username == username {
			*user = global.Users[i]
			return nil
		}
	}

	return fmt.Errorf("user with username %s not found", username)
}

// IsUserExists checks if a user with the specified username exists in the repository.
// It iterates through all users in the global storage and compares usernames.
//
// Parameters:
//   - username: The username to search for
//
// Returns:
//   - bool: true if a user with the given username exists, false otherwise
func (repo *userRepository) IsUserExists(username string) bool {
	for i := 0; i < global.UserCount; i++ {
		if global.Users[i].Username == username {
			return true
		}
	}
	return false
}
