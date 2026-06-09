package cmd

import (
	"fmt"
	"os"
	"github.com/dogcli/dogcli/api"
	"github.com/spf13/cobra"
)

var (
	count      int
	outputDir string
)

var randomCmd = &cobra.Command{
	Use:   "random",
	Short: "Get random dog images",
	Long:  `Fetch random dog images from the Dog API.`,
	Run: func(cmd *cobra.Command, args []string) {
		if count < 1 {
			count = 1
		}
		if count > 50 {
			fmt.Println("⚠️  Maximum count is 50, setting to 50")
			count = 50
		}

		fmt.Printf("🐕 Fetching %d random dog image(s)...\n", count)

		images, err := api.GetRandomDogs(count)
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
	randomCmd.Flags().IntVarP(&count, "count", "c", 1, "Number of images to fetch (1-50)")
	randomCmd.Flags().StringVarP(&outputDir, "output", "o", ".", "Output directory for downloaded images")
}
