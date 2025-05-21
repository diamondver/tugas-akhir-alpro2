package main

import "tugas-besar/lib"

// main is the entry point of the application.
// It initializes the application by calling lib.Bootstrap(),
// which loads environment variables from the .env file,
// sets up application configuration, and prepares the
// necessary resources for the application to run.
func main() {
	lib.Bootstrap()
}
