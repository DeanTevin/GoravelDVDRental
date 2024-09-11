package helpers

import (
	"database/sql/driver"
	"errors"
	"regexp"
	"strconv"
	"strings"
)

type Fulltext []stringArray

type stringArray struct {
	Key    string
	Values int
}

// Implement driver.Valuer interface to convert the Tuple to a database value
func (t Fulltext) Value() (driver.Value, error) {
	if t == nil {
		return nil, nil
	}

	// Serialize the Fulltext to a string format: 'key1':16, 'key2':12, 'key3':10
	var sb strings.Builder
	for i, item := range t {
		if i > 0 {
			sb.WriteString(", ")
		}
		sb.WriteString("'" + item.Key + "':" + strconv.Itoa(item.Values))
	}
	return sb.String(), nil
}

// Implement sql.Scanner to scan a database value into the tuple
func (t *Fulltext) Scan(value interface{}) error {
	str, ok := value.(string)
	str = cleanString(str)
	if !ok {
		return errors.New("failed to scan multistring field - source is not a string")
	}
	*t = parseStringToMap(str)
	return nil
}

func parseStringToMap(input string) []stringArray {
	// Create an array to store the key-value pairs
	var result []stringArray

	// Use a regular expression to find all occurrences of 'key':value
	re := regexp.MustCompile(`'(\w+)':(\d+)`)
	matches := re.FindAllStringSubmatch(input, -1)

	// Iterate over the matches and store them in the array
	for _, match := range matches {
		key := match[1]
		value, _ := strconv.Atoi(match[2]) // Convert value from string to int
		result = append(result, stringArray{Key: key, Values: value})
	}

	return result
}
