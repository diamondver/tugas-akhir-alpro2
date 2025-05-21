package lib

import "tugas-besar/lib/config"

// Bootstrap initializes the application by loading environment configurations.
// It calls config.GetEnvConfig() to load environment variables from the .env file.
// After initializing configurations, it enters an infinite loop to keep the
// application running. This function is called from the main function to start
// the application processes.
//
// The function does not accept any parameters and does not return any values.
func Bootstrap() {
	// Configuration
	config.GetEnvConfig()

	for true {

	}
}
