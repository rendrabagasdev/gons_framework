package env

import "os"

// Get retrieves the value of the environment variable named by the key.
// It returns the value, or the defaultValue if the variable is not present.
func Get(key string, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}
