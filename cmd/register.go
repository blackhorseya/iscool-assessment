package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// registerCmd represents the register command
var registerCmd = &cobra.Command{
	Use:   "register [username]",
	Short: "register a new user",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		username := args[0]
		err := vfs.RegisterUser(username)
		if err != nil {
			_, _ = fmt.Fprintf(os.Stderr, "Error: %v\n", err)
			return
		}

		err = vfs.SaveToFile(dataFile)
		if err != nil {
			_, _ = fmt.Fprintf(os.Stderr, "Error: %v\n", err)
			return
		}
	},
}

func init() {
	rootCmd.AddCommand(registerCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// registerCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// registerCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
