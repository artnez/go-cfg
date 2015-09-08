package internal

import "strings"

type Environ map[string]string

func NewEnviron(env []string) Environ {
	vars := map[string]string{}
	for _, entry := range env {
		if parts := strings.SplitN(entry, "=", 2); len(parts) >= 2 {
			name, value := parts[0], parts[1]
			if value == "" {
				value = "1"
			}
			vars[name] = value
		}
	}
	return vars
}
