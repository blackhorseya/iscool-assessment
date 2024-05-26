package vfs

import (
	"strings"
	"testing"
)

func TestNewFile(t *testing.T) {
	type args struct {
		name        string
		description string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "create file with valid name and description",
			args: args{
				name:        "validName",
				description: "validDescription",
			},
			wantErr: false,
		},
		{
			name: "create file with empty name",
			args: args{
				name:        "",
				description: "validDescription",
			},
			wantErr: true,
		},
		{
			name: "create file with name exceeding max length",
			args: args{
				name:        strings.Repeat("a", MaxInputLength+1),
				description: "validDescription",
			},
			wantErr: true,
		},
		{
			name: "create file with name containing invalid characters",
			args: args{
				name:        "invalidName!",
				description: "validDescription",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := NewFile(tt.args.name, tt.args.description)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewFile() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
