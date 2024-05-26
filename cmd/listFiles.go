package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

const (
	orderByName = "name"
	orderDesc   = "desc"
)

// listFilesCmd represents the listFiles command
var listFilesCmd = &cobra.Command{
	Use:   "list-files [username] [foldername]",
	Short: "List all files in a folder",
	Args:  cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		username := args[0]
		foldername := args[1]
		sortName, _ := cmd.Flags().GetString("sort-name")
		sortCreated, _ := cmd.Flags().GetString("sort-created")

		if sortName != "" && sortCreated != "" {
			fmt.Fprintln(os.Stderr, "Error: Cannot use both --sort-name and --sort-created flags together")
			return
		}

		if sortName != "" && sortName != orderAsc && sortName != orderDesc {
			fmt.Fprintln(os.Stderr, "Error: Invalid value for --sort-name. Use 'asc' or 'desc'")
			return
		}

		if sortCreated != "" && sortCreated != orderAsc && sortCreated != orderDesc {
			fmt.Fprintln(os.Stderr, "Error: Invalid value for --sort-created. Use 'asc' or 'desc'")
			return
		}

		// Default sorting by name in ascending order
		sortCriteria := orderByName
		order := orderAsc

		if sortName != "" {
			sortCriteria = orderByName
			order = sortName
		} else if sortCreated != "" {
			sortCriteria = "created"
			order = sortCreated
		}

		files, err := vfs.ListFiles(username, foldername, sortCriteria, order)
		if err != nil {
			_, _ = fmt.Fprintf(os.Stderr, "Error: %v\n", err)
			return
		}

		// Warning: The folder is empty.
		if len(files) == 0 {
			fmt.Fprintln(os.Stderr, "Warning: The folder is empty.")
			return
		}

		// List files with the following fields: [filename] [description] [created at] [foldername] [username]
		for _, file := range files {
			createdAt := file.CreatedAt.Format("2006-01-02 15:04:05")
			fmt.Printf("%s %s %s %s %s\n", file.Name, file.Description, createdAt, foldername, username)
		}
	},
}

func init() {
	rootCmd.AddCommand(listFilesCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// listFilesCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// listFilesCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	listFilesCmd.Flags().String("sort-name", "", "Sort folders by name (asc or desc)")
	listFilesCmd.Flags().String("sort-created", "", "Sort folders by created time (asc or desc)")
}
