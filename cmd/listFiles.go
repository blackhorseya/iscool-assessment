package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// listFilesCmd represents the listFiles command
var listFilesCmd = &cobra.Command{
	Use:   "list-files [username] [foldername] [--sort-name|--sort-created] [asc|desc]",
	Short: "List all files in a folder",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("listFiles called")
	},
}

func init() {
	rootCmd.AddCommand(listFilesCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// listFilesCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// listFilesCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
