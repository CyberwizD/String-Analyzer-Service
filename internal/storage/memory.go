package storage

import (
	"fmt"
	"sync"

	"github.com/CyberwizD/String-Analyzer-Service/internal/analyzer"
	"github.com/CyberwizD/String-Analyzer-Service/internal/model"
)

// MemoryStore is an in-memory storage for string analysis data.
type MemoryStore struct {
	mu   sync.RWMutex
	data map[string]model.StringAnalysis
}

// NewMemoryStore creates and returns a new MemoryStore.
func NewMemoryStore() *MemoryStore {
	return &MemoryStore{
		data: make(map[string]model.StringAnalysis),
	}
}

// Create adds a new string analysis to the store.
func (s *MemoryStore) Create(analysis model.StringAnalysis) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	if _, exists := s.data[analysis.Value]; exists {
		return fmt.Errorf("string already exists")
	}

	s.data[analysis.Value] = analysis
	return nil
}

// Get retrieves a string analysis from the store by its value.
func (s *MemoryStore) Get(value string) (model.StringAnalysis, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	analysis, exists := s.data[value]
	if !exists {
		return model.StringAnalysis{}, fmt.Errorf("string not found")
	}

	return analysis, nil
}

// GetAll retrieves all string analyses from the store.
func (s *MemoryStore) GetAll() []model.StringAnalysis {
	s.mu.RLock()
	defer s.mu.RUnlock()

	analyses := make([]model.StringAnalysis, 0, len(s.data))
	for _, analysis := range s.data {
		analyses = append(analyses, analysis)
	}

	return analyses
}

// Delete removes a string analysis from the store by its value.
func (s *MemoryStore) Delete(value string) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	if _, exists := s.data[value]; !exists {
		return fmt.Errorf("string not found")
	}

	delete(s.data, value)
	return nil
}

// GetByHash retrieves a string analysis from the store by its SHA256 hash.
func (s *MemoryStore) GetByHash(hash string) (model.StringAnalysis, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	for _, analysis := range s.data {
		if analysis.ID == hash {
			return analysis, nil
		}
	}

	return model.StringAnalysis{}, fmt.Errorf("string not found")
}

// GetByValue retrieves a string analysis from the store by its value.
func (s *MemoryStore) GetByValue(value string) (model.StringAnalysis, error) {
    s.mu.RLock()
    defer s.mu.RUnlock()

    // To ensure consistency, we always calculate the hash of the input value
    // and use that for lookup, as the ID is the hash.
    hash := analyzer.AnalyzeString(value).ID
    
    for _, analysis := range s.data {
        if analysis.ID == hash {
            return analysis, nil
        }
    }

    return model.StringAnalysis{}, fmt.Errorf("string not found")
}
