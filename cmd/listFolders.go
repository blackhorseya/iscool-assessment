package cmd

import (
	"github.com/spf13/cobra"
)

const orderAsc = "asc"

// ListFoldersCmd represents the listFolders command
var ListFoldersCmd = &cobra.Command{
	Use:   "list-folders [username]",
	Short: "List all folders",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		username := args[0]
		sortName, _ := cmd.Flags().GetString("sort-name")
		sortCreated, _ := cmd.Flags().GetString("sort-created")

		if sortName != "" && sortCreated != "" {
			cmd.Println("Error: Cannot use both --sort-name and --sort-created flags together")
			return
		}

		if sortName != "" && sortName != orderAsc && sortName != "desc" {
			cmd.Println("Error: Invalid value for --sort-name. Use 'asc' or 'desc'")
			return
		}

		if sortCreated != "" && sortCreated != orderAsc && sortCreated != "desc" {
			cmd.Println("Error: Invalid value for --sort-created. Use 'asc' or 'desc'")
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

		folders, err := fs.ListFolders(username, sortCriteria, order)
		if err != nil {
			cmd.Printf("Error: %v\n", err)
			return
		}

		// Warning: The [username] doesn't have any folders.
		if len(folders) == 0 {
			cmd.Printf("Warning: The %s doesn't have any folders.\n", username)
			return
		}

		// List all the folders within the [username] scope in following formats: [foldername] [description]
		// [created at] [username]
		for _, folder := range folders {
			createdAt := folder.CreatedAt.Format("2006-01-02 15:04:05")
			cmd.Printf("%s %s %s %s\n", folder.Name, folder.Description, createdAt, username)
		}
	},
}

func init() {
	rootCmd.AddCommand(ListFoldersCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// ListFoldersCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// ListFoldersCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	ListFoldersCmd.Flags().String("sort-name", "", "Sort folders by name (asc or desc)")
	ListFoldersCmd.Flags().String("sort-created", "", "Sort folders by created time (asc or desc)")
}
