package internal

import "os"

// GetEnv return the value of the env or the fallback defined if the env does not exist
func GetEnv(name, fallback string) string {
	env, ok := os.LookupEnv(name)
	if !ok {
		env = fallback
	}

	return env
}
