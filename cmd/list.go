package cmd

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all available dog breeds",
	Long:  `List all available dog breeds from the Dog API.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("🐕 Fetching list of all breeds...")

		resp, err := http.Get("https://dog.ceo/api/breeds/list/all")
		if err != nil {
			fmt.Printf("❌ Error: %v\n", err)
			return
		}
		defer resp.Body.Close()

		body, err := io.ReadAll(resp.Body)
		if err != nil {
			fmt.Printf("❌ Error reading response: %v\n", err)
			return
		}

		var result struct {
			Message map[string][]string `json:"message"`
			Status  string              `json:"status"`
		}

		if err := json.Unmarshal(body, &result); err != nil {
			fmt.Printf("❌ Error parsing response: %v\n", err)
			return
		}

		if result.Status != "success" {
			fmt.Println("❌ API returned an error")
			return
		}

		fmt.Println("\n📋 Available Breeds:")
		fmt.Println(strings.Repeat("─", 40))

		count := 0
		for breed, subBreeds := range result.Message {
			count++
			if len(subBreeds) > 0 {
				fmt.Printf("%-20s (%d sub-breeds)\n", breed, len(subBreeds))
			} else {
				fmt.Printf("%-20s\n", breed)
			}
		}

		fmt.Println(strings.Repeat("─", 40))
		fmt.Printf("Total: %d breeds\n", count)
	},
}
