package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// deleteFileCmd represents the deleteFile command
var deleteFileCmd = &cobra.Command{
	Use:   "delete-file [username] [foldername] [filename]",
	Short: "Delete a file from a folder",
	Args:  cobra.ExactArgs(3),
	Run: func(cmd *cobra.Command, args []string) {
		username := args[0]
		foldername := args[1]
		filename := args[2]

		err := fs.DeleteFile(username, foldername, filename)
		if err != nil {
			_, _ = fmt.Fprintf(os.Stderr, "Error: %v\n", err)
			return
		}

		// Delete [filename]in[username]/[foldername] successfully.
		fmt.Printf("Delete %v in %v/%v successfully.\n", filename, username, foldername)
	},
}

func init() {
	rootCmd.AddCommand(deleteFileCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// deleteFileCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// deleteFileCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
