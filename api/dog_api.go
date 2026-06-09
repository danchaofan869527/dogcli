package api

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

// DogAPIResponse represents the API response structure
type DogAPIResponse struct {
	Message []string `json:"message"`
	Status  string   `json:"status"`
}

// GetRandomDogs fetches random dog images
func GetRandomDogs(count int) ([]string, error) {
	url := fmt.Sprintf("https://dog.ceo/api/breeds/image/random/%d", count)

	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch images: %w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response: %w", err)
	}

	var result DogAPIResponse
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, fmt.Errorf("failed to parse response: %w", err)
	}

	if result.Status != "success" {
		return nil, fmt.Errorf("API returned error status")
	}

	return result.Message, nil
}

// GetBreedDogs fetches dog images for a specific breed
func GetBreedDogs(breed string, count int) ([]string, error) {
	// Normalize breed name (replace spaces with slashes)
	breed = strings.ReplaceAll(breed, " ", "/")
	breed = strings.ToLower(breed)

	url := fmt.Sprintf("https://dog.ceo/api/breed/%s/images/random/%d", breed, count)

	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch images: %w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response: %w", err)
	}

	var result DogAPIResponse
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, fmt.Errorf("failed to parse response: %w", err)
	}

	if result.Status != "success" {
		return nil, fmt.Errorf("breed '%s' not found or API error", breed)
	}

	return result.Message, nil
}

// DownloadImages downloads images from URLs to specified directory
func DownloadImages(urls []string, outputDir string) error {
	// Create output directory if it doesn't exist
	if err := os.MkdirAll(outputDir, 0755); err != nil {
		return fmt.Errorf("failed to create output directory: %w", err)
	}

	fmt.Printf("⏳ Downloading images to %s...\n", outputDir)

	for i, url := range urls {
		// Extract filename from URL
		parts := strings.Split(url, "/")
		filename := parts[len(parts)-1]

		// Add prefix if multiple images
		if len(urls) > 1 {
			filename = fmt.Sprintf("%03d_%s", i+1, filename)
		}

		outputPath := filepath.Join(outputDir, filename)

		// Download the image
		if err := downloadImage(url, outputPath); err != nil {
			fmt.Printf("⚠️  Failed to download %s: %v\n", filename, err)
			continue
		}

		fmt.Printf("   Downloaded: %s\n", filename)
	}

	return nil
}

// downloadImage downloads a single image
func downloadImage(url, outputPath string) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("HTTP status: %d", resp.StatusCode)
	}

	out, err := os.Create(outputPath)
	if err != nil {
		return err
	}
	defer out.Close()

	_, err = io.Copy(out, resp.Body)
	return err
}
