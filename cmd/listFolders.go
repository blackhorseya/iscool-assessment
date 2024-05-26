package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

const orderAsc = "asc"

// listFoldersCmd represents the listFolders command
var listFoldersCmd = &cobra.Command{
	Use:   "list-folders [username]",
	Short: "List all folders",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		username := args[0]
		sortName, _ := cmd.Flags().GetString("sort-name")
		sortCreated, _ := cmd.Flags().GetString("sort-created")

		if sortName != "" && sortCreated != "" {
			fmt.Fprintln(os.Stderr, "Error: Cannot use both --sort-name and --sort-created flags together")
			return
		}

		if sortName != "" && sortName != orderAsc && sortName != "desc" {
			fmt.Fprintln(os.Stderr, "Error: Invalid value for --sort-name. Use 'asc' or 'desc'")
			return
		}

		if sortCreated != "" && sortCreated != orderAsc && sortCreated != "desc" {
			fmt.Fprintln(os.Stderr, "Error: Invalid value for --sort-created. Use 'asc' or 'desc'")
			return
		}

		// Default sorting by name in ascending order
		sortCriteria := "name"
		order := orderAsc

		if sortName != "" {
			sortCriteria = "name"
			order = sortName
		} else if sortCreated != "" {
			sortCriteria = "created"
			order = sortCreated
		}

		folders, err := vfs.ListFolders(username, sortCriteria, order)
		if err != nil {
			_, _ = fmt.Fprintf(os.Stderr, "Error: %v\n", err)
			return
		}

		for _, folder := range folders {
			createdAt := folder.CreatedAt.Format("2006-01-02 15:04:05")
			fmt.Printf("%s %s %s %s\n", folder.Name, folder.Description, createdAt, username)
		}
	},
}

func init() {
	rootCmd.AddCommand(listFoldersCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// listFoldersCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// listFoldersCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	listFoldersCmd.Flags().String("sort-name", "", "Sort folders by name (asc or desc)")
	listFoldersCmd.Flags().String("sort-created", "", "Sort folders by created time (asc or desc)")
}
