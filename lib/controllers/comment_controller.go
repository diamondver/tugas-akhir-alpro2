package controllers

import (
	"fmt"

	"github.com/fatih/color"

	"tugas-besar/lib/model"
	"tugas-besar/lib/services"
)

// CommentController handles application requests and delegates operations to the comment service.
// It implements the controller logic for comment functionality of the application.
type CommentController struct {
	commentService services.CommentService
}

// NewCommentController creates a new CommentController instance with the provided service dependency.
// It follows the dependency injection pattern, allowing for better testability and modular design.
//
// Parameters:
//   - service: An implementation of the CommentService interface
//
// Returns:
//   - A pointer to the newly created CommentController
func NewCommentController(service services.CommentService) *CommentController {
	return &CommentController{
		commentService: service,
	}
}

// CommentInputPage handles the user interface flow for adding a new comment.
// It calls the comment service to display the comment input form and process the submission.
//
// The function handles several control flow paths:
// - On successful comment creation, it displays a success message and returns
// - If the service returns "back" error, it exits the input flow
// - If the service returns "continue" error, it restarts the input flow
// - For other errors, it displays the error message and exits
//
// Parameters:
//   - user: The model.User who is creating the comment
func (c *CommentController) CommentInputPage(user model.User) {
	for {
		err := c.commentService.CreateCommentPage(user)
		if err != nil {
			if err.Error() == "back" {
				break
			}

			if err.Error() == "continue" {
				continue
			}

			color.Red(err.Error())
			break
		}

		color.Green("Komentar berhasil ditambahkan!")
		fmt.Scanln()
		break
	}
}

// CommentView handles the user interface flow for viewing, searching, and sorting comments.
// It continuously calls the comment service to display comments and process user actions.
//
// The function handles several control flow paths based on user selection:
// - If the service returns an error, it displays the error message and exits
// - If the user selects "Exit", it breaks out of the viewing loop
// - If the user selects "Search", it invokes the search comments functionality
// - If the user selects "Sorting", it calls the comment sorting functionality
//
// The function does not take any parameters and does not return any values.
func (c *CommentController) CommentView() {
	var result string

	for {
		err := c.commentService.ShowComment(&result)
		if err != nil {
			color.Red(err.Error())
			fmt.Scanln()
			return
		}

		if result == "Exit" {
			break
		}

		switch result {
		case "Search":
			c.searchComment()
		case "Sorting":
			err := c.commentService.SortingComment()
			if err != nil {
				return
			}
		}
	}
}

// searchComment handles the user interface flow for searching comments.
// It continuously calls the comment service's search functionality until exited.
//
// The function handles several control flow paths:
// - If the service returns "back" error, it exits the search flow
// - If the service returns "continue" error, it restarts the search flow
// - For other errors, it displays the error message and exits
//
// This is an internal method with no parameters and no return values.
func (c *CommentController) searchComment() {
	for {
		err := c.commentService.SearchComment()
		if err != nil {
			if err.Error() == "back" {
				break
			}

			if err.Error() == "continue" {
				continue
			}

			color.Red(err.Error())
			fmt.Scanln()
			return
		}
	}
}

// EditComment handles the user interface flow for editing a user's comment.
// It calls the comment service to display the comment edit form and process the submission.
//
// The function handles several control flow paths:
// - On successful comment edit, it displays a success message and returns
// - If the service returns "back" error, it exits the edit flow
// - If the service returns "continue" error, it restarts the edit flow
// - For other errors, it displays the error message and exits
//
// Parameters:
//   - user: The model.User whose comments are being edited
func (c *CommentController) EditComment(user model.User) {
	for {
		err := c.commentService.EditUserComment(user)
		if err != nil {
			if err.Error() == "back" {
				break
			}

			if err.Error() == "continue" {
				continue
			}

			color.Red(err.Error())
			fmt.Scanln()
			return
		}

		color.Green("Komentar berhasil diubah!")
		fmt.Scanln()
		break
	}
}

// DeleteComment handles the user interface flow for deleting a user's comment.
// It calls the comment service to display the comment deletion interface and process the request.
//
// The function handles several control flow paths:
// - On successful comment deletion, it displays a success message and returns
// - If the service returns "back" error, it exits the deletion flow
// - If the service returns "continue" error, it restarts the deletion flow
// - For other errors, it displays the error message and exits
//
// Parameters:
//   - user: The model.User whose comments are being deleted
func (c *CommentController) DeleteComment(user model.User) {
	for {
		err := c.commentService.DeleteUserComment(user)
		if err != nil {
			if err.Error() == "back" {
				break
			}

			if err.Error() == "continue" {
				continue
			}

			color.Red(err.Error())
			fmt.Scanln()
			return
		}

		color.Green("Komentar berhasil dihapus!")
		fmt.Scanln()
		break
	}
}
