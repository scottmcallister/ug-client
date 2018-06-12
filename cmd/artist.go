package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(artistCmd)
}

var artistCmd = &cobra.Command{
	Use:   "artist",
	Short: "Show a list of artists",
	Long:  "Show a list of artists",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Show a list of artists")
	},
}
