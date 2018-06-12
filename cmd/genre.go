package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(genreCmd)
}

var genreCmd = &cobra.Command{
	Use:   "genre",
	Short: "Browse tabs by genre",
	Long:  "Browse tabs by genre",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Browse tabs by genre")
	},
}
