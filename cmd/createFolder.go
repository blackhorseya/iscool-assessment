package cmd

import (
	"github.com/spf13/cobra"
)

// CreateFolderCmd represents the createFolder command
var CreateFolderCmd = &cobra.Command{
	Use:   "create-folder [username] [foldername] [description]?",
	Short: "create a new folder",
	Args:  cobra.RangeArgs(2, 3),
	Run: func(cmd *cobra.Command, args []string) {
		username := args[0]
		foldername := args[1]
		var description string
		if len(args) == 3 {
			description = args[2]
		}

		folder, err := fs.CreateFolder(username, foldername, description)
		if err != nil {
			cmd.PrintErrf("Error: %v\n", err)
			return
		}

		cmd.Printf("Create %v successfully.\n", folder.Name)
	},
}

func init() {
	rootCmd.AddCommand(CreateFolderCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// CreateFolderCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// CreateFolderCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
