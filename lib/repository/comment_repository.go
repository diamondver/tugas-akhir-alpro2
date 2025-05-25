package repository

import (
	"fmt"
	"tugas-besar/lib/global"
	"tugas-besar/lib/model"
)

// commentRepository implements the CommentRepository interface using an in-memory
// storage mechanism for comment data.
type commentRepository struct {
}

// CommentRepository defines the interface for comment data operations.
// It provides methods to create new comments and retrieve existing comments by ID.
type CommentRepository interface {
	// Create adds a new comment to the repository.
	// Returns an error if the operation fails, nil otherwise.
	Create(comment *model.Comment) error

	// FindCommentByID retrieves a comment by its ID.
	// It populates the provided comment model with data if found.
	// Returns an error if the comment is not found, nil otherwise.
	FindCommentByID(id string, comment *model.Comment) error

	// IsCommentExists checks if a comment with the given ID exists in the repository.
	// Returns true if the comment exists, false otherwise.
	IsCommentExists(id string) bool
}

// NewCommentRepository creates and returns a new CommentRepository implementation.
//
// Returns:
//   - CommentRepository: A new instance of the commentRepository implementation
func NewCommentRepository() CommentRepository {
	return &commentRepository{}
}

// Create adds a new comment to the in-memory repository.
// The comment is assigned the next available index in the global comment storage.
//
// Parameters:
//   - comment: A pointer to the Comment model to be stored
//
// Returns:
//   - error: Always returns nil as this implementation doesn't have failure cases
func (repo *commentRepository) Create(comment *model.Comment) error {
	global.Comments[global.CommentCount] = *comment
	global.CommentCount++

	return nil
}

// FindCommentByID searches for a comment by its ID in the repository.
// If found, it populates the provided comment model with the comment's data.
//
// Parameters:
//   - id: The ID of the comment to search for
//   - comment: A pointer to a Comment model that will be populated with the found comment's data
//
// Returns:
//   - error: An error with a descriptive message if the comment is not found, nil otherwise
func (repo *commentRepository) FindCommentByID(id string, comment *model.Comment) error {
	for i := 0; i < global.CommentCount; i++ {
		if global.Comments[i].ID == id {
			*comment = global.Comments[i]
			return nil
		}
	}

	return fmt.Errorf("comment with ID %s not found", id)
}

// IsCommentExists checks if a comment with the specified ID exists in the repository.
// It iterates through all comments in the global storage and compares IDs.
//
// Parameters:
//   - id: The ID of the comment to search for
//
// Returns:
//   - bool: true if a comment with the given ID exists, false otherwise
func (repo *commentRepository) IsCommentExists(id string) bool {
	for i := 0; i < global.CommentCount; i++ {
		if global.Comments[i].ID == id {
			return true
		}
	}
	return false
}
