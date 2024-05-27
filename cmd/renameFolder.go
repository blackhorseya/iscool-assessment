package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// renameFolderCmd represents the renameFolder command
var renameFolderCmd = &cobra.Command{
	Use:   "rename-folder [username] [foldername] [new-folder-name]",
	Short: "Rename a folder",
	Args:  cobra.ExactArgs(3),
	Run: func(cmd *cobra.Command, args []string) {
		username := args[0]
		foldername := args[1]
		newFolderName := args[2]

		_, err := fs.RenameFolder(username, foldername, newFolderName)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error: %v\n", err)
			return
		}

		fmt.Printf("Rename %v to %v successfully.\n", foldername, newFolderName)
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
