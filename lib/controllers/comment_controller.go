package controllers

import (
	"fmt"
	"tugas-besar/lib/services"

	"github.com/fatih/color"
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

// MainMenu displays the main menu of the application and captures the user's choice.
// It delegates to the commentService to handle menu display and selection logic.
//
// Parameters:
//   - result: A pointer to a string that will store the user's menu selection
//
// The function displays errors in red if any occur during menu operations
// and waits for user acknowledgment by pressing Enter before returning.
func (c *CommentController) CommentInputPage(result *string) {
	err := c.commentService.CommentInputPage(result)

	if err != nil {
		if e, ok := err.(error); ok {
			color.Red(e.Error())
		} else {
			color.Red(fmt.Sprintf("%v", err))
		}
		fmt.Scanln()
		return
	}
}
