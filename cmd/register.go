package cmd

import (
	"github.com/spf13/cobra"
)

// RegisterCmd represents the register command
var RegisterCmd = &cobra.Command{
	Use:   "register [username]",
	Short: "register a new user",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		username := args[0]
		user, err := fs.RegisterUser(username)
		if err != nil {
			cmd.PrintErrf("Error: %v\n", err)
			return
		}

		cmd.Printf("Add %v successfully.\n", user.Username)
	},
}

func init() {
	rootCmd.AddCommand(RegisterCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// RegisterCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// RegisterCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
