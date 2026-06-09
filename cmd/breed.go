package cmd

import (
	"fmt"
	"os"
	"github.com/dogcli/dogcli/api"
	"github.com/spf13/cobra"
)

var breedCmd = &cobra.Command{
	Use:   "breed [breed-name]",
	Short: "Get dog images by breed",
	Long:  `Fetch dog images for a specific breed from the Dog API.`,
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		breed := args[0]

		if count < 1 {
			count = 1
		}
		if count > 50 {
			fmt.Println("⚠️  Maximum count is 50, setting to 50")
			count = 50
		}

		fmt.Printf("🐕 Fetching %d image(s) for breed '%s'...\n", count, breed)

		images, err := api.GetBreedDogs(breed, count)
		if err != nil {
			fmt.Printf("❌ Error: %v\n", err)
			os.Exit(1)
		}

		if err := api.DownloadImages(images, outputDir); err != nil {
			fmt.Printf("❌ Error downloading images: %v\n", err)
			os.Exit(1)
		}

		fmt.Printf("✅ Successfully downloaded %d image(s) to %s\n", len(images), outputDir)
	},
}

func init() {
	breedCmd.Flags().IntVarP(&count, "count", "c", 1, "Number of images to fetch (1-50)")
	breedCmd.Flags().StringVarP(&outputDir, "output", "o", ".", "Output directory for downloaded images")
}
