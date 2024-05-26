package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// deleteFileCmd represents the deleteFile command
var deleteFileCmd = &cobra.Command{
	Use:   "delete-file [username] [foldername] [filename]",
	Short: "Delete a file from a folder",
	Run: func(cmd *cobra.Command, args []string) {
		// todo: 2024/5/26|sean|implement me
		fmt.Println("deleteFile called")
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
