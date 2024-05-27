package cmd_test

import (
	"bytes"
	"os"
	"testing"

	"github.com/blackhorseya/iscool-assessment/cmd"
	"github.com/spf13/cobra"
	"github.com/stretchr/testify/assert"
)

func executeCommand(root *cobra.Command, args ...string) (output string, err error) {
	_, output, err = executeCommandC(root, args...)
	return output, err
}

func executeCommandC(root *cobra.Command, args ...string) (c *cobra.Command, output string, err error) {
	buf := new(bytes.Buffer)
	root.SetOut(buf)
	root.SetErr(buf)
	root.SetArgs(args)

	c, err = root.ExecuteC()

	return c, buf.String(), err
}

func TestRegisterCmd(t *testing.T) {
	rootCmd := &cobra.Command{}
	rootCmd.AddCommand(cmd.RegisterCmd)

	testCases := []struct {
		name     string
		username string
		wantErr  bool
		wantMsg  string
	}{
		{
			name:     "register a new user",
			username: "test",
			wantErr:  false,
			wantMsg:  "Add test successfully.",
		},
		{
			name:     "register a new user with the same username",
			username: "test",
			wantErr:  false,
			wantMsg:  "Error: the test has already existed",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			output, err := executeCommand(rootCmd, "register", tc.username)
			if tc.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
			assert.Contains(t, output, tc.wantMsg)
		})
	}

	_ = os.Remove("out/vfs.json")
}

func TestCreateFolder(t *testing.T) {
	rootCmd := &cobra.Command{}
	rootCmd.AddCommand(cmd.RegisterCmd)
	rootCmd.AddCommand(cmd.CreateFolderCmd)

	testCases := []struct {
		name        string
		username    string
		foldername  string
		description string
		wantErr     bool
		wantMsg     string
	}{
		{
			name:        "create a new folder",
			username:    "test",
			foldername:  "test-folder",
			description: "test description",
			wantErr:     false,
			wantMsg:     "Create test-folder successfully.",
		},
		{
			name:        "create a new folder with the same foldername",
			username:    "test",
			foldername:  "test-folder",
			description: "test description",
			wantErr:     false,
			wantMsg:     "Error: the test-folder has already existed",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			_, _ = executeCommand(rootCmd, "register", tc.username)

			output, err := executeCommand(rootCmd, "create-folder", tc.username, tc.foldername, tc.description)
			if tc.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
			assert.Contains(t, output, tc.wantMsg)
		})
	}

	_ = os.Remove("out/vfs.json")
}

func TestListFilesCmd(t *testing.T) {
	rootCmd := &cobra.Command{}
	rootCmd.AddCommand(cmd.RegisterCmd)
	rootCmd.AddCommand(cmd.CreateFolderCmd)
	rootCmd.AddCommand(cmd.CreateFileCmd)
	rootCmd.AddCommand(cmd.ListFilesCmd)

	testCases := []struct {
		name        string
		username    string
		foldername  string
		sortName    string
		sortCreated string
		wantErr     bool
		wantMsg     string
		mock        func()
	}{
		{
			name:        "list files with valid username and foldername",
			username:    "user1",
			foldername:  "folder1",
			sortName:    "",
			sortCreated: "",
			wantErr:     false,
			wantMsg:     "folder1 user1",
			mock: func() {
				_, _ = executeCommand(rootCmd, "register", "user1")
				_, _ = executeCommand(rootCmd, "create-folder", "user1", "folder1", "test description")
				_, _ = executeCommand(rootCmd, "create-file", "user1", "folder1", "file1", "test description")
			},
		},
		{
			name:        "list files with invalid username",
			username:    "invalidUsername!",
			foldername:  "test-folder",
			sortName:    "",
			sortCreated: "",
			wantErr:     false,
			wantMsg:     "Error: the invalidUsername! doesn't exist",
		},
		{
			name:        "list files with invalid foldername",
			username:    "test",
			foldername:  "invalidFolder!",
			sortName:    "",
			sortCreated: "",
			wantErr:     false,
			wantMsg:     "Error: the invalidFolder! doesn't exist",
			mock: func() {
				_, _ = executeCommand(rootCmd, "register", "test")
			},
		},
		{
			name:        "list files with both sort-name and sort-created flags",
			username:    "test",
			foldername:  "test-folder",
			sortName:    "asc",
			sortCreated: "desc",
			wantErr:     false,
			wantMsg:     "Error: Cannot use both --sort-name and --sort-created flags together",
		},
		{
			name:        "list files with invalid sort-name value",
			username:    "test",
			foldername:  "test-folder",
			sortName:    "invalid",
			sortCreated: "",
			wantErr:     false,
			wantMsg:     "Error: Invalid value for --sort-name. Use 'asc' or 'desc'",
		},
		{
			name:        "list files with invalid sort-created value",
			username:    "test",
			foldername:  "test-folder",
			sortName:    "",
			sortCreated: "invalid",
			wantErr:     false,
			wantMsg:     "Error: Invalid value for --sort-created. Use 'asc' or 'desc'",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if tc.mock != nil {
				tc.mock()
			}

			_ = cmd.ListFilesCmd.Flags().Set("sort-name", tc.sortName)
			_ = cmd.ListFilesCmd.Flags().Set("sort-created", tc.sortCreated)
			output, err := executeCommand(rootCmd, "list-files", tc.username, tc.foldername)
			if tc.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
			assert.Contains(t, output, tc.wantMsg)

			// Clean up
			_ = os.Remove("out/vfs.json")
		})
	}
}

func TestListFoldersCmd(t *testing.T) {
	rootCmd := &cobra.Command{}
	rootCmd.AddCommand(cmd.RegisterCmd)
	rootCmd.AddCommand(cmd.CreateFolderCmd)
	rootCmd.AddCommand(cmd.ListFoldersCmd)

	testCases := []struct {
		name        string
		username    string
		sortName    string
		sortCreated string
		wantErr     bool
		wantMsg     string
		mock        func()
	}{
		{
			name:        "list folders with valid username",
			username:    "test",
			sortName:    "",
			sortCreated: "",
			wantErr:     false,
			wantMsg:     "test",
			mock: func() {
				_, _ = executeCommand(rootCmd, "register", "test")
				_, _ = executeCommand(rootCmd, "create-folder", "test", "folder1", "test description")
			},
		},
		{
			name:        "list folders with invalid username",
			username:    "invalidUsername!",
			sortName:    "",
			sortCreated: "",
			wantErr:     false,
			wantMsg:     "Error: the invalidUsername! doesn't exist",
		},
		{
			name:        "list folders with both sort-name and sort-created flags",
			username:    "test",
			sortName:    "asc",
			sortCreated: "desc",
			wantErr:     false,
			wantMsg:     "Error: Cannot use both --sort-name and --sort-created flags together",
		},
		{
			name:        "list folders with invalid sort-name value",
			username:    "test",
			sortName:    "invalid",
			sortCreated: "",
			wantErr:     false,
			wantMsg:     "Error: Invalid value for --sort-name. Use 'asc' or 'desc'",
		},
		{
			name:        "list folders with invalid sort-created value",
			username:    "test",
			sortName:    "",
			sortCreated: "invalid",
			wantErr:     false,
			wantMsg:     "Error: Invalid value for --sort-created. Use 'asc' or 'desc'",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if tc.mock != nil {
				tc.mock()
			}

			_ = cmd.ListFoldersCmd.Flags().Set("sort-name", tc.sortName)
			_ = cmd.ListFoldersCmd.Flags().Set("sort-created", tc.sortCreated)
			output, err := executeCommand(rootCmd, "list-folders", tc.username)
			if tc.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
			assert.Contains(t, output, tc.wantMsg)

			// Clean up
			_ = os.Remove("out/vfs.json")
		})
	}
}
