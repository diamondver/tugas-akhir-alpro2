package repository

import (
	"fmt"
	"strings"

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
	// GetAllComments retrieves all available comments from the repository.
	// It populates the provided comments array with all comments currently stored in the system.
	GetAllComments(comments *[255]model.Comment) error

	// Create adds a new comment to the repository.
	// Returns an error if the operation fails, nil otherwise.
	Create(comment *model.Comment, userId int) error

	// SearchComments searches for comments containing the specified search string.
	// It populates the provided comments array with matching comments.
	SearchComments(search string, comments *[255]model.Comment) error

	// SortCommentsByComment sorts the comments based on the length of the comment text.
	// The sorting can be done in either ascending or descending order.
	SortCommentsByComment(comments *[255]model.Comment, mode int) error

	// SortCommentsByKategori sorts the comments based on their category value.
	// Categories are ranked as: Positif (1), Netral (0), Negatif (-1).
	SortCommentsByKategori(comments *[255]model.Comment, mode int) error

	// EditUserComment updates a comment that belongs to a specific user.
	// Only allows editing if the comment exists and belongs to the specified user.
	EditUserComment(commentId int, userId int, comment model.Comment) error

	// DeleteUserComment removes a comment that belongs to a specific user.
	// Only allows deletion if the comment exists and belongs to the specified user.
	DeleteUserComment(commentId int, userId int) error

	// GetCommentByUserId retrieves all comments belonging to a specific user.
	// It populates the provided comments array with all comments from the specified user.
	GetCommentByUserId(userId int, comments *[255]model.Comment) error
}

// NewCommentRepository creates and returns a new CommentRepository implementation.
//
// Returns:
//   - CommentRepository: A new instance of the commentRepository implementation
func NewCommentRepository() CommentRepository {
	return &commentRepository{}
}

// GetAllComments retrieves all available comments from the repository.
// It directly assigns the global comment storage to the provided array pointer,
// which means the caller gets access to all comments currently in the system.
//
// Parameters:
//   - comments: A pointer to an array that will be filled with all comments
//
// Returns:
//   - error: Always returns nil as this implementation doesn't have failure cases
func (c *commentRepository) GetAllComments(comments *[255]model.Comment) error {
	*comments = global.Comments
	return nil
}

// Create adds a new comment to the in-memory repository.
// The comment is assigned the next available index in the global comment storage.
//
// Parameters:
//   - comment: A pointer to the Comment model to be stored
//
// Returns:
//   - error: Always returns nil as this implementation doesn't have failure cases
func (c *commentRepository) Create(comment *model.Comment, userId int) error {
	global.Comments[global.CommentCount] = model.Comment{
		Id:       global.IdCommentIncrement + 1,
		UserId:   userId,
		Komentar: comment.Komentar,
		Kategori: comment.Kategori,
	}
	global.CommentCount++
	global.IdCommentIncrement++

	return nil
}

// SearchComments searches for comments containing the specified search string.
// It implements a case-insensitive substring search by converting both the
// search term and comment text to lowercase before comparison.
//
// The method uses a manual substring matching algorithm that checks each position
// in the comment text as a potential starting point for a match.
//
// Parameters:
//   - search: The string to search for within comments
//   - comments: A pointer to an array that will be filled with matching comments
//
// Returns:
//   - error: Always returns nil as this implementation doesn't have failure cases
func (c *commentRepository) SearchComments(search string, comments *[255]model.Comment) error {
	searchLower := strings.ToLower(search)

	for i := 0; i < global.CommentCount; i++ {
		commentLower := strings.ToLower(global.Comments[i].Komentar)

		for j := 0; j <= len(commentLower)-len(searchLower); j++ {
			isMatch := true

			for k := 0; k < len(searchLower); k++ {
				if commentLower[j+k] != searchLower[k] {
					isMatch = false
					break
				}
			}

			if isMatch {
				(*comments)[i] = global.Comments[i]
				break
			}
		}
	}

	return nil
}

// SortCommentsByComment sorts the comments based on the length of the comment text.
// It first copies all global comments to the provided array, then sorts them using
// selection sort algorithm.
//
// The function implements a selection sort where:
// - For mode 0 (ascending): Comments with shorter text appear first
// - For mode 1 (descending): Comments with longer text appear first
//
// Parameters:
//   - comments: A pointer to an array that will be filled with sorted comments
//   - mode: The sorting mode (0 for ascending, 1 for descending)
//
// Returns:
//   - error: Always returns nil as this implementation doesn't have failure cases
func (c *commentRepository) SortCommentsByComment(comments *[255]model.Comment, mode int) error {
	for i := 0; i < global.CommentCount; i++ {
		(*comments)[i] = global.Comments[i]
	}

	for i := 0; i < global.CommentCount-1; i++ {
		index := i

		for j := i + 1; j < global.CommentCount; j++ {
			if mode == 0 { // Ascending
				if len((*comments)[j].Komentar) < len((*comments)[index].Komentar) {
					index = j
				}
			} else if mode == 1 { // Descending
				if len((*comments)[j].Komentar) > len((*comments)[index].Komentar) {
					index = j
				}
			}
		}

		if index != i {
			(*comments)[i], (*comments)[index] = (*comments)[index], (*comments)[i]
		}
	}

	return nil
}

// SortCommentsByKategori sorts the comments based on their category value.
// It first copies all global comments to the provided array, then sorts them using
// insertion sort algorithm.
//
// The function uses the following category values for sorting:
// - Positif: 1
// - Netral: 0
// - Negatif: -1
//
// The sorting behavior is determined by the mode parameter:
// - For mode 0 (ascending): Categories are sorted from Negatif to Positif
// - For mode 1 (descending): Categories are sorted from Positif to Negatif
//
// Parameters:
//   - comments: A pointer to an array that will be filled with sorted comments
//   - mode: The sorting mode (0 for ascending, 1 for descending)
//
// Returns:
//   - error: Always returns nil as this implementation doesn't have failure cases
func (c *commentRepository) SortCommentsByKategori(comments *[255]model.Comment, mode int) error {
	for i := 0; i < global.CommentCount; i++ {
		(*comments)[i] = global.Comments[i]
	}

	getCategoryValue := func(category string) int {
		switch category {
		case "Positif":
			return 1
		case "Netral":
			return 0
		case "Negatif":
			return -1
		default:
			return 0
		}
	}

	for i := 1; i < global.CommentCount; i++ {
		current := (*comments)[i]
		currentValue := getCategoryValue(current.Kategori)
		j := i - 1

		if mode == 0 {
			for j >= 0 && getCategoryValue((*comments)[j].Kategori) > currentValue {
				(*comments)[j+1] = (*comments)[j]
				j--
			}
		} else {
			for j >= 0 && getCategoryValue((*comments)[j].Kategori) < currentValue {
				(*comments)[j+1] = (*comments)[j]
				j--
			}
		}

		(*comments)[j+1] = current
	}

	return nil
}

// EditUserComment updates a comment that belongs to a specific user.
// It searches through all comments to find a match with both the specified commentId and userId.
// Only fields that contain values in the provided data will be updated (empty strings are ignored).
//
// Parameters:
//   - commentId: The ID of the comment to edit
//   - userId: The ID of the user who owns the comment
//   - data: The model.Comment containing fields to update
//
// Returns:
//   - error: An error if the comment is not found or doesn't belong to the user, nil on success
func (c *commentRepository) EditUserComment(commentId int, userId int, data model.Comment) error {
	for i := 0; i < global.CommentCount; i++ {
		if global.Comments[i].Id == commentId && global.Comments[i].UserId == userId {
			comment := &global.Comments[i]

			if data.Komentar != "" {
				comment.Komentar = data.Komentar
			}

			if data.Kategori != "" {
				comment.Kategori = data.Kategori
			}

			return nil
		}
	}

	return fmt.Errorf("comment with ID %d not found or does not belong to user with ID %d", commentId, userId)
}

// GetCommentByUserId retrieves all comments belonging to a specific user.
// It iterates through all comments in the global storage and copies those
// that match the specified user ID to the provided array, maintaining
// their original index positions.
//
// Note: This implementation preserves the original index positions of comments,
// which may result in sparse population of the results array if user comments
// are not contiguous in the global storage.
//
// Parameters:
//   - userId: The ID of the user whose comments to retrieve
//   - comments: A pointer to an array that will be filled with the user's comments
//
// Returns:
//   - error: Always returns nil as this implementation doesn't have failure cases
func (c *commentRepository) GetCommentByUserId(userId int, comments *[255]model.Comment) error {
	for i := 0; i < global.CommentCount; i++ {
		if global.Comments[i].UserId == userId {
			(*comments)[i] = global.Comments[i]
		}
	}

	return nil
}

// DeleteUserComment removes a comment that belongs to a specific user.
// It first searches for a comment with the matching commentId that also belongs to the specified userId.
// If found, it removes the comment by shifting all subsequent comments up by one position in the array
// and decrements the global comment count.
//
// Parameters:
//   - commentId: The ID of the comment to delete
//   - userId: The ID of the user who owns the comment
//
// Returns:
//   - error: An error if the comment is not found or doesn't belong to the user, nil on success
func (c *commentRepository) DeleteUserComment(commentId int, userId int) error {
	for i := 0; i < global.CommentCount; i++ {
		if global.Comments[i].Id == commentId && global.Comments[i].UserId == userId {
			for j := i; j < global.CommentCount-1; j++ {
				global.Comments[j] = global.Comments[j+1]
			}
			global.CommentCount--
			return nil
		}
	}

	return fmt.Errorf("comment with ID %d not found or does not belong to user with ID %d", commentId, userId)
}
