package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// deleteFolderCmd represents the deleteFolder command
var deleteFolderCmd = &cobra.Command{
	Use:   "delete-folder [username] [foldername]",
	Short: "delete a folder",
	Run: func(cmd *cobra.Command, args []string) {
		// todo: 2024/5/26|sean|implement me
		fmt.Println("deleteFolder called")
	},
}

func init() {
	rootCmd.AddCommand(deleteFolderCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// deleteFolderCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// deleteFolderCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
