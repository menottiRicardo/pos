package common

import "strings"

// ParseValidationError parses the validation error string into a structured map
func ParseValidationError(errMsg string) map[string]string {
	errMap := make(map[string]string)
	lines := strings.Split(errMsg, "\n")
	for _, line := range lines {
		parts := strings.Split(line, " Error:")
		if len(parts) == 2 {
			field := strings.TrimPrefix(parts[0], "Key: 'MenuItem.")
			field = strings.TrimSuffix(field, "'")
			errMap[field] = parts[1]
		}
	}
	return errMap
}
