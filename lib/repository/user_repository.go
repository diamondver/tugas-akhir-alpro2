package repository

import (
	"fmt"
	"strings"
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
	IsUserExists(username string, exceptId int) bool

	// GetAllUsers retrieves all users stored in the repository.
	// It populates the provided users array with all user records
	// currently stored in the system.
	GetAllUsers(users *[255]model.User) error

	// SearchUsers finds users whose usernames contain the specified search string.
	// It performs a case-insensitive substring search on all usernames and
	// populates the provided array with matching user records.
	SearchUsers(search string, users *[255]model.User) error

	// EditUser updates a user's information at the specified index.
	// It allows partial updates - empty fields in the data parameter will not
	// overwrite existing values. Only non-empty fields will be updated.
	EditUser(index int, data model.User) error

	// DeleteUser removes a user from the repository.
	// It deletes the user at the specified index and shifts all subsequent users
	// to maintain contiguous storage, then decrements the global user count.
	DeleteUser(id int) error
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
func (repo *userRepository) IsUserExists(username string, exceptId int) bool {
	for i := 0; i < global.UserCount; i++ {
		if global.Users[i].Username == username && i != exceptId {
			return true
		}
	}
	return false
}

// GetAllUsers retrieves all users stored in the repository.
//
// This implementation simply copies all users from the global storage
// to the provided array. The function populates the users array with
// all user records currently stored in the system.
//
// Parameters:
//   - users: A pointer to a fixed-size array that will be populated with user data
//
// Returns:
//   - error: Always returns nil as this implementation doesn't have failure cases
func (repo *userRepository) GetAllUsers(users *[255]model.User) error {
	*users = global.Users

	return nil
}

// SearchUsers finds users whose usernames contain the specified search string.
//
// This implementation performs a manual case-insensitive substring search on usernames.
// The search algorithm works as follows:
// 1. Convert both the search term and each username to lowercase
// 2. For each possible position in the username, check if the search term matches
// 3. If a match is found, add the user to the results array
//
// The function uses a character-by-character comparison rather than built-in string
// functions like strings.Contains() to implement the substring search.
//
// Parameters:
//   - search: The substring to search for within usernames
//   - users: A pointer to a fixed-size array that will be populated with matching users
//
// Returns:
//   - error: Always returns nil as this implementation doesn't have failure cases
func (repo *userRepository) SearchUsers(search string, users *[255]model.User) error {
	searchLower := strings.ToLower(search)

	for i := 0; i < global.UserCount; i++ {
		usernameLower := strings.ToLower(global.Users[i].Username)

		for j := 0; j <= len(usernameLower)-len(searchLower); j++ {
			isMatch := true

			for k := 0; k < len(searchLower); k++ {
				if usernameLower[j+k] != searchLower[k] {
					isMatch = false
					break
				}
			}

			if isMatch {
				(*users)[i] = global.Users[i]
				break
			}
		}
	}

	return nil
}

// EditUser updates a user's information at the specified index.
//
// This implementation performs a partial update of the user data at the given index.
// Only non-empty fields in the data parameter will overwrite existing values.
// Currently, only Username and Password fields can be updated.
//
// Parameters:
//   - index: The array index of the user to be updated
//   - data: A User model containing the fields to update (empty fields are ignored)
//
// Returns:
//   - error: An error if the index is out of bounds, nil on success
func (repo *userRepository) EditUser(index int, data model.User) error {
	if index < 0 || index >= global.UserCount {
		return fmt.Errorf("index %d out of bounds", index)
	}

	user := &global.Users[index]

	if data.Username != "" {
		user.Username = data.Username
	}

	if data.Password != "" {
		user.Password = data.Password
	}

	return nil
}

// DeleteUser removes a user from the repository.
//
// This implementation deletes the user at the specified index by shifting all
// subsequent users one position back to maintain contiguous storage. After shifting,
// it clears the last user position to avoid duplicates and decrements the global user count.
//
// Parameters:
//   - id: The index of the user to remove
//
// Returns:
//   - error: An error if the id is out of bounds, nil on success
func (repo *userRepository) DeleteUser(id int) error {
	if id < 0 || id >= global.UserCount {
		return fmt.Errorf("id %d out of bounds", id)
	}

	for i := id; i < global.UserCount-1; i++ {
		global.Users[i] = global.Users[i+1]
	}

	global.Users[global.UserCount-1] = model.User{}

	global.UserCount--

	return nil
}
