package services

import (
	"fmt"
	"github.com/jedib0t/go-pretty/v6/table"
	"os"
	"strconv"

	"github.com/fatih/color"
	"github.com/manifoldco/promptui"

	"tugas-besar/lib/global"
	"tugas-besar/lib/helper"
	"tugas-besar/lib/model"
	"tugas-besar/lib/repository"
)

// CommentService defines the interface for comment management operations.
// It provides methods to create, find, and check the existence of comments.
type CommentService interface {
	// CreateCommentPage displays the comment creation interface for a user.
	// It shows a form where the user can input their comment text and select a category
	// (Positif, Netral, or Negatif). After submission, it creates the comment in the system.
	CreateCommentPage(user model.User) error

	// CreateComment adds a new comment to the system.
	// Returns an error if the creation fails, nil otherwise.
	CreateComment(comment *model.Comment, userId int) error

	// ShowComment displays all comments in the system in a tabular format.
	// After displaying the comments, it shows a menu with options for Search, Sorting, or Exit.
	// The user's selection is stored in the chose parameter.
	ShowComment(chose *string) error

	// SearchComment implements the comment search functionality.
	// It displays a search form, processes the search query against comment content,
	// and shows matching results in a tabular format. The function also handles
	// the option to search again or return to the previous menu.
	SearchComment() error

	// SortingComment handles the comment sorting functionality.
	// It presents options to sort comments by either comment text or category,
	// and in either ascending or descending order. The sorted results are
	// displayed in a tabular format.
	SortingComment() error

	// EditUserComment allows a user to edit their own comments.
	// It displays a list of the user's comments, prompts for the ID of the comment
	// to edit, and presents a form to update the comment text and category.
	EditUserComment(user model.User) error

	// DeleteUserComment allows a user to delete their own comments.
	// It displays a list of the user's comments, prompts for the ID of the comment
	// to delete, and removes the selected comment from the system.
	DeleteUserComment(user model.User) error

	// ShowTable retrieves and displays all comments in a formatted table.
	// It queries the repository for all comments and renders them in a table
	// with columns for comment number, ID, text content, and category.
	// The table is formatted with colored styling for better readability.
	ShowTable() error

	// CreateCommentForm displays interactive prompts for entering comment text and selecting a category.
	// It creates a text input prompt for the comment and a selection menu for the category
	// (Positif, Netral, Negatif) with custom styling. The user's inputs are stored in the provided
	// string pointers.
	CreateCommentForm(komentar, kategori *string) error

	// EditForm displays interactive prompts for editing comment text and selecting a category.
	// It creates a text input prompt for the comment and a selection menu for the category
	// (Positif, Netral, Negatif) with custom styling. The user's inputs are stored in the provided
	// string pointers.
	EditForm(komentar, kategori *string) error

	// EditComment updates a comment with the specified ID in the repository.
	// It delegates the update operation to the underlying repository implementation.
	EditComment(id int, komentar model.Comment) error
}

// commentService implements the commentService interface.
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

// CreateCommentPage displays a form for creating a new comment and processes the user's input.
// It clears the screen, shows a header for the comment input form, then prompts the user
// to enter comment text and select a category through the CreateCommentForm function.
// Upon successful input, it creates a new comment in the system with the provided information.
//
// Parameters:
//   - user: The model.User representing the currently logged-in user
//
// Returns:
//   - error: An error if the form display, user input, or comment creation fails, nil on success
func (c *commentService) CreateCommentPage(user model.User) error {
	helper.ClearScreen()
	color.Yellow("* MENU > USER > INPUT KOMENTAR")
	color.Yellow("========================================")
	color.Yellow("=           INPUT KOMENTAR             =")
	color.Yellow("========================================")

	var komentar, kategori string

	err := c.CreateCommentForm(&komentar, &kategori)
	if err != nil {
		return err
	}

	err = c.CreateComment(&model.Comment{
		Komentar: komentar,
		Kategori: kategori,
	}, user.Id)
	if err != nil {
		return err
	}

	return nil
}

// CreateCommentForm displays interactive prompts for entering comment text and selecting a category.
// It creates a text input prompt for the comment and a selection menu for the category
// (Positif, Netral, Negatif) with custom styling. The user's inputs are stored in the provided
// string pointers.
//
// Parameters:
//   - komentar: A pointer to a string where the comment text will be stored
//   - kategori: A pointer to a string where the selected category will be stored
//
// Returns:
//   - error: An error if any prompt operation fails, nil on success
func (c *commentService) CreateCommentForm(komentar, kategori *string) error {
	komentarPrompt := promptui.Prompt{Label: "Komentar"}
	kategoriPrompt := promptui.Select{
		Label: "Kategori",
		Items: []string{"Positif", "Netral", "Negatif"},
		Templates: &promptui.SelectTemplates{
			Label:    "{{ . | blue }}:",
			Active:   "\u27A1 {{ . | cyan }}",
			Inactive: "  {{ . | cyan }}",
			Selected: "\u2705 {{ . | blue | cyan }}",
		},
	}

	komentarInput, err := komentarPrompt.Run()
	if err != nil {
		return err
	}

	_, kategoriInput, err := kategoriPrompt.Run()
	if err != nil {
		return err
	}

	*komentar = komentarInput
	*kategori = kategoriInput

	return nil
}

// ShowComment displays all comments in the system in a tabular format.
// It first clears the screen and displays a header for the comment viewing section.
// Then it retrieves all comments from the repository, renders them in a table showing
// the comment number, text content, and category. After displaying the comments,
// it presents a menu with options for Search, Sorting, or Exit, and stores the
// user's selection in the chose parameter.
//
// Parameters:
//   - chose: A pointer to a string that will store the user's menu selection
//
// Returns:
//   - error: An error if retrieving comments or handling the menu fails, nil on success
func (c *commentService) ShowComment(chose *string) error {
	helper.ClearScreen()
	color.Yellow("* MENU > USER > LIHAT KOMENTAR")
	color.Yellow("========================================")
	color.Yellow("=           LIHAT KOMENTAR             =")
	color.Yellow("========================================")

	err := c.ShowTable()
	if err != nil {
		return err
	}

	prompt := promptui.Select{
		Label: "Pilih Menu",
		Items: []string{"Search", "Sorting", "Exit"},
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

// SearchComment implements the comment search functionality.
// It provides a user interface for searching comments by keyword and displays the results.
//
// The function follows these steps:
// 1. Clears the screen and displays the search interface header
// 2. Prompts the user to enter a keyword to search for in comments
// 3. Queries the repository for comments matching the keyword
// 4. Displays matching comments in a formatted table with numbering, comment text, and category
// 5. Asks the user if they want to search again
//
// Returns:
//   - error: Returns "continue" if the user wants to search again, "back" if the user wants
//     to return to the previous menu, or another error if any operation fails
func (c *commentService) SearchComment() error {
	helper.ClearScreen()
	color.Yellow("* MENU > USER > LIHAT KOMENTAR > CARI KOMENTAR")
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
	err = c.commentRepo.SearchComments(searchInput, &comments)
	if err != nil {
		return err
	}

	helper.ClearScreen()
	color.Yellow("* MENU > USER > LIHAT KOMENTAR > CARI KOMENTAR")
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

// SortingComment handles the comment sorting functionality.
// It provides a user interface for sorting comments by either comment text or category,
// in ascending or descending order.
//
// The function follows these steps:
// 1. Displays a header for the sorting interface
// 2. Prompts the user to select a field to sort by (Komentar or Kategori)
// 3. Prompts the user to select a sort direction (Ascending or Descending)
// 4. Converts the sort direction to an integer (0 for Ascending, 1 for Descending)
// 5. Calls the appropriate specialized sorting function based on user selections
//
// Returns:
//   - error: An error if any part of the sorting operation fails, nil on success
func (c *commentService) SortingComment() error {
	helper.ClearScreen()
	color.Yellow("* MENU > USER > LIHAT KOMENTAR > SORTING KOMENTAR")
	color.Yellow("========================================")
	color.Yellow("=           SORTING KOMENTAR           =")
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

	_, result, err := prompt.Run()
	if err != nil {
		return err
	}

	_, mode, err := promptMode.Run()
	if err != nil {
		return err
	}

	modeInt := 0
	if mode == "Descending" {
		modeInt = 1
	}

	switch result {
	case "Komentar":
		err := c.sortCommentByKomentar(modeInt)
		if err != nil {
			return err
		}
	case "Kategori":
		err := c.sortCommentByKategori(modeInt)
		if err != nil {
			return err
		}
	}

	return nil
}

// sortCommentByKomentar sorts and displays comments based on their content text.
// It retrieves comments from the repository sorted by their "Komentar" field,
// displays them in a formatted table, and waits for the user to press Enter before returning.
//
// The function follows these steps:
// 1. Retrieves comments from the repository sorted by comment text
// 2. Clears the screen and displays a header for the sorted comments
// 3. Creates and renders a table showing the sorted comments with numbering, text, and category
// 4. Waits for the user to press Enter (via fmt.Scanln()) before returning
//
// Parameters:
//   - mode: An integer indicating the sort direction (0 for ascending, 1 for descending)
//
// Returns:
//   - error: An error if retrieving or displaying the sorted comments fails, nil on success
func (c *commentService) sortCommentByKomentar(mode int) error {
	var comments [255]model.Comment

	err := c.commentRepo.SortCommentsByComment(&comments, mode)
	if err != nil {
		return err
	}

	helper.ClearScreen()
	color.Yellow("* MENU > USER > LIHAT KOMENTAR > SORTING KOMENTAR")
	color.Yellow("========================================")
	color.Yellow("=           SORTING KOMENTAR           =")
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
// It retrieves comments from the repository sorted by their "Kategori" field,
// displays them in a formatted table, and waits for the user to press Enter before returning.
//
// The function follows these steps:
// 1. Retrieves comments from the repository sorted by category
// 2. Clears the screen and displays a header for the sorted comments
// 3. Creates and renders a table showing the sorted comments with numbering, text, and category
// 4. Waits for the user to press Enter (via fmt.Scanln()) before returning
//
// Parameters:
//   - mode: An integer indicating the sort direction (0 for ascending, 1 for descending)
//
// Returns:
//   - error: An error if retrieving or displaying the sorted comments fails, nil on success
func (c *commentService) sortCommentByKategori(mode int) error {
	var comments [255]model.Comment

	err := c.commentRepo.SortCommentsByKategori(&comments, mode)
	if err != nil {
		return err
	}

	helper.ClearScreen()
	color.Yellow("* MENU > USER > LIHAT KOMENTAR > SORTING KOMENTAR")
	color.Yellow("========================================")
	color.Yellow("=           SORTING KOMENTAR           =")
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

// EditUserComment allows a user to edit their own comments.
// It provides a user interface for selecting and modifying an existing comment.
//
// The function follows these steps:
//  1. Clears the screen and displays a header for the comment editing interface
//  2. Retrieves and displays all comments created by the user in a formatted table
//     showing numbering, comment ID, text, and category
//  3. Prompts the user to enter the ID of the comment they want to edit
//  4. Validates the input to ensure it's a valid numeric ID
//  5. Displays a form for entering new comment text and selecting a new category
//  6. Updates the comment in the repository with the new information
//  7. If the update fails, displays an error and asks if the user wants to try again
//
// Parameters:
//   - user: The model.User representing the currently logged-in user
//
// Returns:
//   - error: Returns "continue" if the user wants to edit another comment after
//     an error, "back" if the user wants to return to the previous menu, nil on
//     successful update, or another error if any operation fails
func (c *commentService) EditUserComment(user model.User) error {
	helper.ClearScreen()
	color.Yellow("* MENU > USER > EDIT KOMENTAR")
	color.Yellow("========================================")
	color.Yellow("=            EDIT KOMENTAR             =")
	color.Yellow("========================================")

	err := c.showCommentByUserTable(user.Id)
	if err != nil {
		return err
	}

	prompt := promptui.Prompt{
		Label: "Masukkan id komentar yang ingin diedit",
		Validate: func(input string) error {
			if input == "" {
				return fmt.Errorf("id komentar tidak boleh kosong")
			}

			_, err := strconv.Atoi(input)
			if err != nil {
				return fmt.Errorf("id komentar harus berupa angka")
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
		return fmt.Errorf("id komentar harus berupa angka")
	}

	var komentar, kategori string
	err = c.EditForm(&komentar, &kategori)
	if err != nil {
		return err
	}

	err = c.commentRepo.EditUserComment(id, user.Id, model.Comment{
		Komentar: komentar,
		Kategori: kategori,
	})

	askPrompt := promptui.Prompt{
		Label:     "Edit Again?",
		IsConfirm: true,
	}

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

// EditForm displays interactive prompts for editing comment text and selecting a category.
// It creates a text input prompt for the comment and a selection menu for the category
// (Positif, Netral, Negatif) with custom styling. The user's inputs are stored in the provided
// string pointers.
//
// Parameters:
//   - komentar: A pointer to a string where the edited comment text will be stored
//   - kategori: A pointer to a string where the selected category will be stored
//
// Returns:
//   - error: An error if any prompt operation fails, nil on success
func (c *commentService) EditForm(komentar, kategori *string) error {
	komentarPrompt := promptui.Prompt{Label: "Komentar"}
	kategoriPrompt := promptui.Select{
		Label: "Kategori",
		Items: []string{"Positif", "Netral", "Negatif"},
		Templates: &promptui.SelectTemplates{
			Label:    "{{ . | blue }}:",
			Active:   "\u27A1 {{ . | cyan }}",
			Inactive: "  {{ . | cyan }}",
			Selected: "\u2705 {{ . | blue | cyan }}",
		},
	}

	komentarInput, err := komentarPrompt.Run()
	if err != nil {
		return err
	}

	_, kategoriInput, err := kategoriPrompt.Run()
	if err != nil {
		return err
	}

	*komentar = komentarInput
	*kategori = kategoriInput

	return nil
}

// DeleteUserComment allows a user to delete their own comments.
// It provides a user interface for selecting and removing an existing comment.
//
// The function follows these steps:
//  1. Clears the screen and displays a header for the comment deletion interface
//  2. Retrieves and displays all comments created by the user in a formatted table
//     showing numbering, comment ID, text, and category
//  3. Prompts the user to enter the ID of the comment they want to delete
//  4. Validates the input to ensure it's a valid numeric ID
//  5. Calls the repository to delete the comment with the specified ID
//  6. If the deletion fails, displays an error and asks if the user wants to try again
//
// Parameters:
//   - user: The model.User representing the currently logged-in user
//
// Returns:
//   - error: Returns "continue" if the user wants to delete another comment after
//     an error, "back" if the user wants to return to the previous menu, nil on
//     successful deletion, or another error if any operation fails
func (c *commentService) DeleteUserComment(user model.User) error {
	helper.ClearScreen()
	color.Yellow("* MENU > USER > HAPUS KOMENTAR")
	color.Yellow("========================================")
	color.Yellow("=            HAPUS KOMENTAR            =")
	color.Yellow("========================================")

	err := c.showCommentByUserTable(user.Id)
	if err != nil {
		return err
	}

	prompt := promptui.Prompt{
		Label: "Masukkan id komentar yang ingin dihapus",
		Validate: func(input string) error {
			if input == "" {
				return fmt.Errorf("id komentar tidak boleh kosong")
			}

			_, err := strconv.Atoi(input)
			if err != nil {
				return fmt.Errorf("id komentar harus berupa angka")
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

	err = c.commentRepo.DeleteUserComment(id, user.Id)
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

// CreateComment adds a new comment to the system.
// It delegates the creation operation to the underlying repository.
//
// Parameters:
//   - comment: A pointer to the Comment model to be created
//
// Returns:
//   - error: An error if the creation fails, nil otherwise
func (c *commentService) CreateComment(comment *model.Comment, userId int) error {
	return c.commentRepo.Create(comment, userId)
}

// CommentShowPage displays a menu for viewing different types of comments.
// It presents a selection interface with options to view all comments, positive comments,
// negative comments, search for comments, view comment statistics, or return to the previous menu.
//
// The function follows these steps:
// 1. Clears the screen and displays a header for the comment viewing section
// 2. Creates a selection menu with various comment viewing options
// 3. Captures the user's selection and stores it in the provided string pointer
//
// Parameters:
//   - chose: A pointer to a string that will store the user's menu selection
//
// Returns:
//   - error: An error if displaying the menu or capturing the selection fails, nil on success
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

// ShowTable retrieves and displays all comments in a formatted table.
// It creates a table with columns for comment number, text content, and category.
// The function queries the repository for all comments, adds each comment
// to the table (up to the global.CommentCount limit), and renders the table
// with colored formatting to standard output.
//
// Returns:
//   - error: An error if retrieving comments fails, nil on success
func (c *commentService) ShowTable() error {
	var comments [255]model.Comment

	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	t.AppendHeader(table.Row{"#", "Id", "Komentar", "Kategori"})

	err := c.commentRepo.GetAllComments(&comments)
	if err != nil {
		return err
	}

	for i := 0; i < global.CommentCount; i++ {
		t.AppendRow(table.Row{
			i + 1,
			comments[i].Id,
			comments[i].Komentar,
			comments[i].Kategori,
		})
	}

	t.SetStyle(table.StyleColoredBright)
	t.Render()

	return nil
}

// showCommentByUserTable retrieves and displays comments from a specific user in a formatted table.
// It creates a table with columns for row number, comment ID, text content, and category.
// The function queries the repository for comments belonging to the specified user,
// adds each non-empty comment to the table, and renders the table with colored formatting
// to standard output.
//
// Parameters:
//   - userId: An integer representing the ID of the user whose comments should be displayed
//
// Returns:
//   - error: An error if retrieving comments fails, nil on success
func (c *commentService) showCommentByUserTable(userId int) error {
	var comments [255]model.Comment

	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	t.AppendHeader(table.Row{"#", "Id", "Komentar", "Kategori"})
	err := c.commentRepo.GetCommentByUserId(userId, &comments)
	if err != nil {
		return err
	}
	var j int
	for i := 0; i < global.CommentCount; i++ {
		if comments[i].Komentar != "" {
			j++
			t.AppendRow(table.Row{
				j,
				comments[i].Id,
				comments[i].Komentar,
				comments[i].Kategori,
			})
		}
	}
	t.SetStyle(table.StyleColoredBright)
	t.Render()

	return nil
}

// EditComment updates a comment with the specified ID in the system.
// It delegates to the underlying repository implementation to perform the actual update.
// Only non-empty fields in the provided comment model will be updated.
//
// Parameters:
//   - id: The ID of the comment to edit
//   - komentar: The model.Comment containing fields to update
//
// Returns:
//   - error: An error if the comment is not found or update fails, nil on success
func (c *commentService) EditComment(id int, komentar model.Comment) error {
	err := c.commentRepo.EditComment(id, komentar)
	if err != nil {
		return err
	}

	return nil
}
