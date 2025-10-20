package analyzer

import (
	"crypto/sha256"
	"encoding/hex"
	"strings"
	"time"
	"unicode"

	"github.com/CyberwizD/String-Analyzer-Service/internal/model"
)

// AnalyzeString computes the properties of a given string and returns a StringAnalysis object.
func AnalyzeString(value string) model.StringAnalysis {
	hash := calculateSHA256(value)
	properties := model.Properties{
		Length:                calculateLength(value),
		IsPalindrome:          isPalindrome(value),
		UniqueCharacters:      countUniqueCharacters(value),
		WordCount:             countWords(value),
		SHA256Hash:            hash,
		CharacterFrequencyMap: calculateCharacterFrequency(value),
	}

	return model.StringAnalysis{
		ID:         hash,
		Value:      value,
		Properties: properties,
		CreatedAt:  time.Now().UTC(),
	}
}

// calculateLength returns the number of characters in a string.
func calculateLength(s string) int {
	return len(s)
}

// isPalindrome checks if a string is a palindrome (case-insensitive).
func isPalindrome(s string) bool {
	s = strings.ToLower(s)
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		if runes[i] != runes[j] {
			return false
		}
	}
	return true
}

// countUniqueCharacters returns the number of unique characters in a string.
func countUniqueCharacters(s string) int {
	charSet := make(map[rune]bool)
	for _, char := range s {
		charSet[char] = true
	}
	return len(charSet)
}

// countWords returns the number of words in a string, separated by whitespace.
func countWords(s string) int {
	return len(strings.Fields(s))
}

// calculateSHA256 returns the SHA-256 hash of a string.
func calculateSHA256(s string) string {
	hash := sha256.Sum256([]byte(s))
	return hex.EncodeToString(hash[:])
}

// calculateCharacterFrequency returns a map of each character to its occurrence count.
func calculateCharacterFrequency(s string) map[string]int {
	freqMap := make(map[string]int)
	for _, char := range s {
		if !unicode.IsSpace(char) {
			freqMap[string(char)]++
		}
	}
	return freqMap
}
