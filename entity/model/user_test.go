package model

import (
	"strings"
	"testing"
)

func TestNewUser(t *testing.T) {
	tests := []struct {
		name     string
		username string
		wantErr  bool
	}{
		{
			name:     "Valid input",
			username: "valid-username",
			wantErr:  false,
		},
		{
			name:     "Empty username",
			username: "",
			wantErr:  true,
		},
		{
			name:     "Username exceeding max length",
			username: strings.Repeat("a", 256),
			wantErr:  true,
		},
		{
			name:     "Invalid characters in username",
			username: "invalid!",
			wantErr:  true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			user, err := NewUser(tt.username)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewUser() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if err == nil && user.Username != tt.username {
				t.Errorf("NewUser() user.Username = %v, want %v", user.Username, tt.username)
			}
		})
	}
}
