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
