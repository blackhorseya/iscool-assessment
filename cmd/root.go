package cmd

import (
	"fmt"
	"os"

	vfs2 "github.com/blackhorseya/iscool-assessment/pkg/legacy_vfs"
	"github.com/blackhorseya/iscool-assessment/pkg/utils"
	vfs3 "github.com/blackhorseya/iscool-assessment/pkg/vfs"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var vfs = vfs2.NewVFS()
var dataFile = "out/vfs.json"
var cfgFile string
var out string
var vfsV2 vfs3.VirtualFileSystem

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

	rootCmd.PersistentFlags().StringVar(
		&cfgFile,
		"config",
		"",
		"config file (default is $HOME/.config/iscool/.iscool.yaml)",
	)
	rootCmd.PersistentFlags().StringVar(&out, "out", "out/vfs.json", "output file or directory")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := os.UserHomeDir()
		cobra.CheckErr(err)

		// Search config in home directory with name ".iscool-assessment" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigType("yaml")
		viper.SetConfigName(".iscool")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Fprintln(os.Stderr, "Using config file:", viper.ConfigFileUsed())
	}

	// Load the virtual filesystem from the config file
	err := vfs.LoadFromFile(dataFile)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to load virtual filesystem: %v\n", err)
		os.Exit(1)
	}

	err = initVFS()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to init virtual filesystem: %v\n", err)
		return
	}
}

func initVFS() (err error) {
	pathType := utils.CheckPathType(out)
	if pathType == "json" {
		vfsV2, err = NewVFSWithJSON(out)
		if err != nil {
			return err
		}
	} else if pathType == "folder" {
		vfsV2, err = NewVFSWithSystem(out)
		if err != nil {
			return err
		}
	}

	return nil
}
