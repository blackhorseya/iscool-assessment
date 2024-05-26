package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// renameFolderCmd represents the renameFolder command
var renameFolderCmd = &cobra.Command{
	Use:   "rename-folder [username] [foldername] [new-folder-name]",
	Short: "Rename a folder",
	Run: func(cmd *cobra.Command, args []string) {
		// todo: 2024/5/26|sean|implement me
		fmt.Println("renameFolder called")
	},
}

func init() {
	rootCmd.AddCommand(renameFolderCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// renameFolderCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// renameFolderCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
