package model

import (
	"strings"
	"testing"
)

func TestValidateInput(t *testing.T) {
	type args struct {
		input string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "valid input",
			args: args{
				input: "valid-input",
			},
			wantErr: false,
		},
		{
			name: "empty input",
			args: args{
				input: "",
			},
			wantErr: true,
		},
		{
			name: "input exceeding max length",
			args: args{
				input: strings.Repeat("a", MaxInputLength+1),
			},
			wantErr: true,
		},
		{
			name: "input with invalid characters",
			args: args{
				input: "invalid!",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := ValidateInput(tt.args.input); (err != nil) != tt.wantErr {
				t.Errorf("ValidateInput() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
