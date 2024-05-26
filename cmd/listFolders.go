package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// listFoldersCmd represents the listFolders command
var listFoldersCmd = &cobra.Command{
	Use:   "list-folders [username] [--sort-name|--sort-created] [asc|desc]",
	Short: "List all folders",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("listFolders called")
	},
}

func init() {
	rootCmd.AddCommand(listFoldersCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// listFoldersCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// listFoldersCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
