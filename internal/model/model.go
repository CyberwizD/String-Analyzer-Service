package model

import "time"

// Properties represents the computed properties of a string.
type Properties struct {
	Length                int            `json:"length"`
	IsPalindrome          bool           `json:"is_palindrome"`
	UniqueCharacters      int            `json:"unique_characters"`
	WordCount             int            `json:"word_count"`
	SHA256Hash            string         `json:"sha256_hash"`
	CharacterFrequencyMap map[string]int `json:"character_frequency_map"`
}

// StringAnalysis represents the full analysis of a string, including its properties and metadata.
type StringAnalysis struct {
	ID         string     `json:"id"`
	Value      string     `json:"value"`
	Properties Properties `json:"properties"`
	CreatedAt  time.Time  `json:"created_at"`
}
