package cmd

import (
	"fmt"
	"os"

	"github.com/blackhorseya/iscool-assessment/pkg/utils"
	"github.com/blackhorseya/iscool-assessment/pkg/vfs"
	"github.com/spf13/cobra"
)

var cfgFile string
var out string
var fs vfs.VirtualFileSystem

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "iscool",
	Short: "iscool is a CLI tool to virtual file system",
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.
	rootCmd.PersistentFlags().StringVar(&out, "out", "out/vfs.json", "output file or directory")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	err := initVFS()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to init virtual filesystem: %v\n", err)
		return
	}
}

func initVFS() (err error) {
	pathType := utils.CheckPathType(out)
	if pathType == "json" {
		fs, err = NewVFSWithJSON(out)
		if err != nil {
			return err
		}
	} else if pathType == "folder" {
		fs, err = NewVFSWithSystem(out)
		if err != nil {
			return err
		}
	}

	return nil
}
