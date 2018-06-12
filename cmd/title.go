package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(titleCmd)
}

var titleCmd = &cobra.Command{
	Use:   "title",
	Short: "Display a list of tabs based on a search query",
	Long:  "Display a list of tabs based on a search query",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Display a list of tabs based on a search query")
	},
}
