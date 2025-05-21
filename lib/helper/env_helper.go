package helper

import "os"

// GetEnv retrieves the value of an environment variable by its key.
// If the environment variable is not set or is empty, it returns
// the provided fallback value instead.
//
// Parameters:
//   - key: The name of the environment variable to retrieve
//   - fallback: The default value to return if the environment variable is not set
//
// Returns:
//   - The value of the environment variable if set and not empty,
//     otherwise the fallback value
func GetEnv(key, fallback string) string {
	value := os.Getenv(key)

	if value == "" {
		return fallback
	}

	return value
}
