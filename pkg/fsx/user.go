//go:generate mockgen -destination=./mock_${GOFILE} -package=${GOPACKAGE} -source=${GOFILE}

package fsx

// UserManager defines the interface for user management
type UserManager interface {
	RegisterUser(username string) error
	DeleteUser(username string) error
	ListUsers() []string
}

// User represents a user with username and a list of folders
type User struct {
	Username string             `json:"username"`
	Folders  map[string]*Folder `json:"folders"`
}
