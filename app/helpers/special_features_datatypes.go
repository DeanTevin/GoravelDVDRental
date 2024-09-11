package helpers

import (
	"database/sql/driver"
	"errors"
	"strings"
)

type SpecialFeatures []string

// Implement driver.Valuer interface to convert the Tuple to a database value
func (t SpecialFeatures) Value() (driver.Value, error) {
	if t == nil {
		return nil, nil
	}
	return strings.Join(t, ","), nil
}

// Implement sql.Scanner to scan a database value into the tuple
func (t *SpecialFeatures) Scan(value interface{}) error {
	str, ok := value.(string)
	str = cleanString(str)
	if !ok {
		return errors.New("failed to scan multistring field - source is not a string")
	}
	*t = strings.Split(str, ",")
	return nil
}

func cleanString(input string) string {
	// Remove curly braces and quotes
	cleaned := strings.ReplaceAll(input, "{", "")
	cleaned = strings.ReplaceAll(cleaned, "}", "")
	cleaned = strings.ReplaceAll(cleaned, `"`, "")

	// Return the cleaned and trimmed string
	return strings.TrimSpace(cleaned)
}
