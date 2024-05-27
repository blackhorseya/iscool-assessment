package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// deleteFolderCmd represents the deleteFolder command
var deleteFolderCmd = &cobra.Command{
	Use:   "delete-folder [username] [foldername]",
	Short: "delete a folder",
	Args:  cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		username := args[0]
		foldername := args[1]

		err := fs.DeleteFolder(username, foldername)
		if err != nil {
			_, _ = fmt.Fprintf(os.Stderr, "Error: %v\n", err)
			return
		}

		fmt.Printf("Delete %v successfully.\n", foldername)
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
