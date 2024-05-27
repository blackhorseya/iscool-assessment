package user

import (
	"context"
	"os"
	"testing"
)

func Test_system_Register(t *testing.T) {
	type args struct {
		ctx      context.Context
		username string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "register new user",
			args: args{
				ctx:      context.Background(),
				username: "newUser",
			},
			wantErr: false,
		},
		{
			name: "register with invalid username",
			args: args{
				ctx:      context.Background(),
				username: "invalidUsername!",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := &system{
				path: "out",
			}
			_, err := i.Register(tt.args.ctx, tt.args.username)
			if (err != nil) != tt.wantErr {
				t.Errorf("Register() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func Test_system_GetByUsername(t *testing.T) {
	type args struct {
		ctx      context.Context
		username string
		mock     func()
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "get user by valid username",
			args: args{
				ctx:      context.Background(),
				username: "existingUser",
				mock: func() {
					_ = os.MkdirAll("out/existingUser", os.ModePerm)
				},
			},
			wantErr: false,
		},
		{
			name: "get user by non-existing username",
			args: args{
				ctx:      context.Background(),
				username: "nonExistingUser",
			},
			wantErr: true,
		},
		{
			name: "get user by invalid username",
			args: args{
				ctx:      context.Background(),
				username: "invalidUsername!",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.args.mock != nil {
				tt.args.mock()
			}

			i := &system{
				path: "out",
			}
			_, err := i.GetByUsername(tt.args.ctx, tt.args.username)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetByUsername() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			_ = os.RemoveAll("out")
		})
	}
}
