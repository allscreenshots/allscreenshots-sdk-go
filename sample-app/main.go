// Package main provides a sample web application demonstrating the Allscreenshots SDK.
package main

import (
	"context"
	"embed"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/allscreenshots/allscreenshots-sdk-go/pkg/allscreenshots"
)

//go:embed templates/*
var templatesFS embed.FS

//go:embed static/*
var staticFS embed.FS

var (
	client    *allscreenshots.Client
	templates *template.Template
)

// ScreenshotRequest represents the incoming screenshot request from the UI.
type ScreenshotRequest struct {
	URL      string `json:"url"`
	Device   string `json:"device"`
	FullPage bool   `json:"fullPage"`
}

// ScreenshotResponse represents the response sent back to the UI.
type ScreenshotResponse struct {
	Success bool   `json:"success"`
	Image   string `json:"image,omitempty"`
	Error   string `json:"error,omitempty"`
}

func main() {
	// Check for API key
	apiKey := os.Getenv("ALLSCREENSHOTS_API_KEY")
	if apiKey == "" {
		log.Fatal("ALLSCREENSHOTS_API_KEY environment variable is required")
	}

	// Initialize client
	client = allscreenshots.NewClient(
		allscreenshots.WithAPIKey(apiKey),
		allscreenshots.WithTimeout(120*time.Second),
	)

	// Parse templates
	var err error
	templates, err = template.ParseFS(templatesFS, "templates/*.html")
	if err != nil {
		log.Fatalf("Failed to parse templates: %v", err)
	}

	// Set up routes
	http.HandleFunc("/", handleIndex)
	http.HandleFunc("/api/screenshot", handleScreenshot)
	http.Handle("/static/", http.FileServer(http.FS(staticFS)))

	// Start server
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("Server starting on http://localhost:%s", port)
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatalf("Server failed: %v", err)
	}
}

func handleIndex(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	if err := templates.ExecuteTemplate(w, "index.html", nil); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func handleScreenshot(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var req ScreenshotRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		sendJSONResponse(w, ScreenshotResponse{
			Success: false,
			Error:   "Invalid request body",
		})
		return
	}

	// Validate URL
	if req.URL == "" {
		sendJSONResponse(w, ScreenshotResponse{
			Success: false,
			Error:   "URL is required",
		})
		return
	}

	// Create context with timeout
	ctx, cancel := context.WithTimeout(r.Context(), 120*time.Second)
	defer cancel()

	// Take screenshot
	imageData, err := client.Screenshot(ctx, &allscreenshots.ScreenshotRequest{
		URL:      req.URL,
		Device:   req.Device,
		FullPage: req.FullPage,
	})

	if err != nil {
		errorMsg := "Failed to capture screenshot"

		// Provide more specific error messages
		if allscreenshots.IsValidationError(err) {
			errorMsg = fmt.Sprintf("Invalid request: %v", err)
		} else if allscreenshots.IsUnauthorized(err) {
			errorMsg = "Invalid API key"
		} else if allscreenshots.IsRateLimited(err) {
			errorMsg = "Rate limit exceeded. Please try again later."
		} else if allscreenshots.IsServerError(err) {
			errorMsg = "Server error. Please try again."
		} else if apiErr, ok := allscreenshots.AsAPIError(err); ok {
			errorMsg = apiErr.Message
		}

		sendJSONResponse(w, ScreenshotResponse{
			Success: false,
			Error:   errorMsg,
		})
		return
	}

	// Encode image to base64
	base64Image := base64.StdEncoding.EncodeToString(imageData)

	sendJSONResponse(w, ScreenshotResponse{
		Success: true,
		Image:   base64Image,
	})
}

func sendJSONResponse(w http.ResponseWriter, resp ScreenshotResponse) {
	w.Header().Set("Content-Type", "application/json")
	if !resp.Success {
		w.WriteHeader(http.StatusBadRequest)
	}
	json.NewEncoder(w).Encode(resp)
}
