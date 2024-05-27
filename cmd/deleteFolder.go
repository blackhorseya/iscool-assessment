package cmd

import (
	"github.com/spf13/cobra"
)

// DeleteFolderCmd represents the deleteFolder command
var DeleteFolderCmd = &cobra.Command{
	Use:   "delete-folder [username] [foldername]",
	Short: "delete a folder",
	Args:  cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		username := args[0]
		foldername := args[1]

		err := fs.DeleteFolder(username, foldername)
		if err != nil {
			cmd.Printf("Error: %v\n", err)
			return
		}

		cmd.Printf("Delete %v successfully.\n", foldername)
	},
}

func init() {
	rootCmd.AddCommand(DeleteFolderCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// DeleteFolderCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// DeleteFolderCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
