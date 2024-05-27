package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// createFileCmd represents the createFile command
var createFileCmd = &cobra.Command{
	Use:   "create-file [username] [foldername] [filename] [description]?",
	Short: "Create a file in a folder",
	Args:  cobra.RangeArgs(3, 4),
	Run: func(cmd *cobra.Command, args []string) {
		username := args[0]
		foldername := args[1]
		filename := args[2]
		description := ""
		if len(args) == 4 {
			description = args[3]
		}

		file, err := fs.CreateFile(username, foldername, filename, description)
		if err != nil {
			_, _ = fmt.Fprintf(os.Stderr, "Error: %v\n", err)
			return
		}

		// Create [filename]in[username]/[foldername] successfully.
		fmt.Printf("Create %v in %v/%v successfully.\n", file.Name, username, foldername)
	},
}

func init() {
	rootCmd.AddCommand(createFileCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// createFileCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// createFileCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
