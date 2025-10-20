package nlp

import (
	"regexp"
	"strconv"
	"strings"
)

// ParseQuery interprets a natural language query and returns a map of filters.
func ParseQuery(query string) (map[string]string, error) {
	filters := make(map[string]string)
	query = strings.ToLower(query)

	// Palindrome
	if strings.Contains(query, "palindromic") || strings.Contains(query, "palindrome") {
		filters["is_palindrome"] = "true"
	}

	// Word count
	if strings.Contains(query, "single word") {
		filters["word_count"] = "1"
	}

	// Length
	re := regexp.MustCompile(`longer than (\d+) characters`)
	if matches := re.FindStringSubmatch(query); len(matches) > 1 {
		length, _ := strconv.Atoi(matches[1])
		filters["min_length"] = strconv.Itoa(length + 1)
	}

	// Contains character
	re = regexp.MustCompile(`containing the letter ([a-z])`)
	if matches := re.FindStringSubmatch(query); len(matches) > 1 {
		filters["contains_character"] = matches[1]
	}

	// Contains first vowel
	if strings.Contains(query, "first vowel") {
		filters["contains_character"] = "a"
	}

	return filters, nil
}
