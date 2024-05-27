package cmd

import (
	"github.com/spf13/cobra"
)

// RenameFolderCmd represents the renameFolder command
var RenameFolderCmd = &cobra.Command{
	Use:   "rename-folder [username] [foldername] [new-folder-name]",
	Short: "Rename a folder",
	Args:  cobra.ExactArgs(3),
	Run: func(cmd *cobra.Command, args []string) {
		username := args[0]
		foldername := args[1]
		newFolderName := args[2]

		_, err := fs.RenameFolder(username, foldername, newFolderName)
		if err != nil {
			cmd.Printf("Error: %v\n", err)
			return
		}

		cmd.Printf("Rename %v to %v successfully.\n", foldername, newFolderName)
	},
}

func init() {
	rootCmd.AddCommand(RenameFolderCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// RenameFolderCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// RenameFolderCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
