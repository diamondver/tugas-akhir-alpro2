package config

import (
	"github.com/fatih/color"
	"github.com/joho/godotenv"
)

// GetEnvConfig loads environment variables from the .env file at the project root.
// It uses the godotenv package to read the file and populate the environment.
// If the .env file cannot be loaded, it displays an error message in red text
// using the fatih/color package.
// No values are returned as this function modifies the environment directly.
func GetEnvConfig() {
	err := godotenv.Load()

	if err != nil {
		color.Red("Error loading .env file")
	}
}
