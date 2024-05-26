package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// createFolderCmd represents the createFolder command
var createFolderCmd = &cobra.Command{
	Use:   "create-folder [username] [foldername] (description)",
	Short: "create a new folder",
	Run: func(cmd *cobra.Command, args []string) {
		// todo: 2024/5/26|sean|implement me
		fmt.Println("createFolder called")
	},
}

func init() {
	rootCmd.AddCommand(createFolderCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// createFolderCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// createFolderCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
