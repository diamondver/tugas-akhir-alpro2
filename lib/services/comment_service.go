package services

import (
	"tugas-besar/lib/helper"
	"tugas-besar/lib/model"
	"tugas-besar/lib/repository"

	"github.com/fatih/color"
	"github.com/manifoldco/promptui"
)

// CommentService defines the interface for comment management operations.
// It provides methods to create, find, and check the existence of comments.
type CommentService interface {
	// CreateComment adds a new comment to the system.
	// Returns an error if the creation fails, nil otherwise.
	CreateComment(comment *model.Comment) error
}

func (c CommentService) CommentInputPage(result *string) any {
	panic("unimplemented")
}

// komentarService implements the KomentarService interface.
// It acts as a service layer between the application and the repository.
type commentService struct {
	commentRepo repository.CommentRepository
}

// NewCommentService creates and returns a new CommentService implementation.
//
// Parameters:
//   - commentRepo: The comment repository implementation to use for data operations
//
// Returns:
//   - CommentService: A new instance of the commentService implementation
func NewCommentService(commentRepo repository.CommentRepository) CommentService {
	return &commentService{
		commentRepo: commentRepo,
	}
}

// CreateComment adds a new comment to the system.
// It delegates the creation operation to the underlying repository.
//
// Parameters:
//   - comment: A pointer to the Comment model to be created
//
// Returns:
//   - error: An error if the creation fails, nil otherwise
func (commentService *commentService) CreateComment(comment *model.Comment) error {
	return commentService.commentRepo.Create(comment)
}

// FindCommentByID retrieves a comment by its ID.
// It delegates the search operation to the underlying repository.
//
// Parameters:
//   - id: The ID of the comment to search for
//   - comment: A pointer to a Comment model that will be populated with the found comment's data
//
// Returns:
//   - error: An error if the comment is not found, nil otherwise
//func (commentService *commentService) FindCommentByID(id string, comment *model.Comment) error {
//	return commentService.commentRepo.FindByID(id, comment)
//}

// IsCommentExists checks if a comment with the specified ID exists.
// It delegates the check to the underlying repository.
//
// Parameters:
//   - id: The ID of the comment to check for existence
//
// Returns:
//   - bool: true if a comment with the given ID exists, false otherwise
func (commentService *commentService) IsCommentExists(id string) bool {
	return commentService.commentRepo.IsCommentExists(id)
}

func (commentService *commentService) commentInputPage() error {
	var komentar, kategori string

	helper.ClearScreen()
	color.Yellow("* MENU > INPUT KOMENTAR")
	color.Yellow("========================================")
	color.Yellow("=           INPUT KOMENTAR             =")
	color.Yellow("========================================")

	err := commentForm(&komentar, &kategori)
	if err != nil {
		return err
	}

	// Process the komentar and kategori as needed
	// For example, you might want to save them to a database or perform some other action

	return nil
}

func commentForm(komentar, kategori *string) error {
	komentarPrompt := promptui.Prompt{Label: "Komentar"}
	kategoriPrompt := promptui.Prompt{Label: "Kategori"}

	komentarInput, err := komentarPrompt.Run()
	if err != nil {
		return err
	}

	kategoriInput, err := kategoriPrompt.Run()
	if err != nil {
		return err
	}

	*komentar = komentarInput
	*kategori = kategoriInput

	return nil
}

func (*commentService) CommentShowPage(chose *string) error {
	helper.ClearScreen()
	color.Yellow("* MENU > LIHAT KOMENTAR")
	color.Yellow("========================================")
	color.Yellow("=           LIHAT KOMENTAR             =")
	color.Yellow("========================================")

	prompt := promptui.Select{
		Label: "Pilih Menu",
		Items: []string{"Lihat Semua Komentar", "Lihat Komentar Positif", "Lihat Komentar Negatif", "Cari Komentar", "Statistik Komentar", "Kembali"},
		Templates: &promptui.SelectTemplates{
			Label:    "{{ . | blue }}:",
			Active:   "\u27A1 {{ . | cyan }}",
			Inactive: "  {{ . | cyan }}",
			Selected: "\u2705 {{ . | blue | cyan }}",
		},
	}

	_, result, err := prompt.Run()

	if err != nil {
		return err
	}

	*chose = result

	return nil
}
