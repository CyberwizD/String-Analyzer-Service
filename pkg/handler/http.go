package handler

import (
	"net/http"
	"strconv"
	"strings"

	"github.com/CyberwizD/String-Analyzer-Service/internal/analyzer"
	"github.com/CyberwizD/String-Analyzer-Service/internal/model"
	"github.com/CyberwizD/String-Analyzer-Service/internal/nlp"
	"github.com/CyberwizD/String-Analyzer-Service/internal/storage"
	"github.com/gin-gonic/gin"
)

// GinHandler handles HTTP requests for the string analysis service using Gin.
type GinHandler struct {
	store *storage.MemoryStore
}

// NewGinHandler creates a new GinHandler.
func NewGinHandler(store *storage.MemoryStore) *GinHandler {
	return &GinHandler{store: store}
}

// CreateString handles the creation of a new string analysis.
func (h *GinHandler) CreateString(c *gin.Context) {
	var req map[string]interface{}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	value, ok := req["value"]
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{"error": `Missing "value" field`})
		return
	}

	stringValue, ok := value.(string)
	if !ok {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error": "Invalid data type for \"value\" (must be string)"})
		return
	}

	if stringValue == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": `The "value" field cannot be empty`})
		return
	}

	analysis := analyzer.AnalyzeString(stringValue)
	if err := h.store.Create(analysis); err != nil {
		c.JSON(http.StatusConflict, gin.H{"error": "String already exists"})
		return
	}

	c.JSON(http.StatusCreated, analysis)
}

// GetString handles the retrieval of a specific string analysis.
func (h *GinHandler) GetString(c *gin.Context) {
	value := c.Param("string_value")

	analysis, err := h.store.Get(value)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "String not found"})
		return
	}

	c.JSON(http.StatusOK, analysis)
}

// GetAllStrings handles the retrieval of all string analyses with optional filtering.
func (h *GinHandler) GetAllStrings(c *gin.Context) {
	filters := c.Request.URL.Query()
	analyses := h.store.GetAll()

	filteredAnalyses := filterAnalyses(analyses, filters)

	response := gin.H{
		"data":            filteredAnalyses,
		"count":           len(filteredAnalyses),
		"filters_applied": parseFilters(filters),
	}

	c.JSON(http.StatusOK, response)
}

// filterAnalyses filters a slice of StringAnalysis based on query parameters.
func filterAnalyses(analyses []model.StringAnalysis, filters map[string][]string) []model.StringAnalysis {
	var filtered []model.StringAnalysis

	for _, analysis := range analyses {
		if matchFilters(analysis, filters) {
			filtered = append(filtered, analysis)
		}
	}

	return filtered
}

// filterAnalysesWithMap filters a slice of StringAnalysis based on a map of filters.
func filterAnalysesWithMap(analyses []model.StringAnalysis, filters map[string]string) []model.StringAnalysis {
	var filtered []model.StringAnalysis

	for _, analysis := range analyses {
		if matchFiltersWithMap(analysis, filters) {
			filtered = append(filtered, analysis)
		}
	}

	return filtered
}

// matchFilters checks if a StringAnalysis matches the provided filters.
func matchFilters(analysis model.StringAnalysis, filters map[string][]string) bool {
	for key, values := range filters {
		value := values[0]
		switch key {
		case "is_palindrome":
			isPalindrome, err := strconv.ParseBool(value)
			if err != nil || analysis.Properties.IsPalindrome != isPalindrome {
				return false
			}
		case "min_length":
			minLength, err := strconv.Atoi(value)
			if err != nil || analysis.Properties.Length < minLength {
				return false
			}
		case "max_length":
			maxLength, err := strconv.Atoi(value)
			if err != nil || analysis.Properties.Length > maxLength {
				return false
			}
		case "word_count":
			wordCount, err := strconv.Atoi(value)
			if err != nil || analysis.Properties.WordCount != wordCount {
				return false
			}
		case "contains_character":
			if !strings.Contains(analysis.Value, value) {
				return false
			}
		}
	}
	return true
}

// parseFilters parses the query parameters into a map for the response.
func parseFilters(filters map[string][]string) map[string]interface{} {
	parsed := make(map[string]interface{})
	for key, values := range filters {
		value := values[0]
		switch key {
		case "is_palindrome":
			parsed[key], _ = strconv.ParseBool(value)
		case "min_length", "max_length", "word_count":
			parsed[key], _ = strconv.Atoi(value)
		default:
			parsed[key] = value
		}
	}
	return parsed
}

// matchFiltersWithMap checks if a StringAnalysis matches the provided filters from a map.
func matchFiltersWithMap(analysis model.StringAnalysis, filters map[string]string) bool {
	for key, value := range filters {
		switch key {
		case "is_palindrome":
			isPalindrome, err := strconv.ParseBool(value)
			if err != nil || analysis.Properties.IsPalindrome != isPalindrome {
				return false
			}
		case "min_length":
			minLength, err := strconv.Atoi(value)
			if err != nil || analysis.Properties.Length < minLength {
				return false
			}
		case "max_length":
			maxLength, err := strconv.Atoi(value)
			if err != nil || analysis.Properties.Length > maxLength {
				return false
			}
		case "word_count":
			wordCount, err := strconv.Atoi(value)
			if err != nil || analysis.Properties.WordCount != wordCount {
				return false
			}
		case "contains_character":
			if !strings.Contains(analysis.Value, value) {
				return false
			}
		}
	}
	return true
}

// DeleteString handles the deletion of a specific string analysis.
func (h *GinHandler) DeleteString(c *gin.Context) {
	value := c.Param("string_value")

	if err := h.store.Delete(value); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "String not found"})
		return
	}

	c.Status(http.StatusNoContent)
}

// FilterByNaturalLanguage handles filtering by natural language query.
func (h *GinHandler) FilterByNaturalLanguage(c *gin.Context) {
	query := c.Query("query")
	if query == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Missing query parameter"})
		return
	}

	parsedFilters, err := nlp.ParseQuery(query)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Unable to parse natural language query"})
		return
	}

	analyses := h.store.GetAll()
	filteredAnalyses := filterAnalysesWithMap(analyses, parsedFilters)

	response := gin.H{
		"data":  filteredAnalyses,
		"count": len(filteredAnalyses),
		"interpreted_query": gin.H{
			"original":       query,
			"parsed_filters": parsedFilters,
		},
	}

	c.JSON(http.StatusOK, response)
}
