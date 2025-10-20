package main

import (
	"log"

	"github.com/CyberwizD/String-Analyzer-Service/internal/storage"
	"github.com/CyberwizD/String-Analyzer-Service/pkg/handler"
	"github.com/gin-gonic/gin"
)

func main() {
	store := storage.NewMemoryStore()
	h := handler.NewGinHandler(store)

	r := gin.Default()

	r.POST("/strings", h.CreateString)
	r.GET("/strings/:string_value", h.GetString)
	r.GET("/strings", h.GetAllStrings)
	r.GET("/strings/filter-by-natural-language", h.FilterByNaturalLanguage)
	r.DELETE("/strings/:string_value", h.DeleteString)

	log.Println("Server starting on :8080")
	if err := r.Run(":8080"); err != nil {
		log.Fatalf("could not start server: %v", err)
	}
}
