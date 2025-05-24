package services

import (
	"github.com/fatih/color"
	"github.com/manifoldco/promptui"
	"tugas-besar/lib/helper"
)

// MainService defines the interface for the main operations of the application.
// It abstracts the core business logic to allow for better testing and modularity.
type MainService interface {
	MainMenu(chose *string) error
}

// mainServiceImpl implements the MainService interface with concrete business logic.
type mainServiceImpl struct {
}

// NewMainService creates and returns a new instance of MainService.
// This factory function follows the dependency injection pattern to create
// properly initialized service objects.
//
// Returns:
//   - A concrete implementation of the MainService interface
func NewMainService() MainService {
	return &mainServiceImpl{}
}

func (*mainServiceImpl) MainMenu(chose *string) error {
	helper.ClearScreen()
	color.Yellow("=========================================")
	color.Yellow("=  Selamat datang di Tugas Besar Alpro  =")
	color.Yellow("=       Aplikasi Analisis Sentimen      =")
	color.Yellow("=            Kelompok 2                 =")
	color.Yellow("=========================================")

	prompt := promptui.Select{
		Label: "Pilih Menu",
		Items: []string{"Login", "Register", "Admin", "Exit"},
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
