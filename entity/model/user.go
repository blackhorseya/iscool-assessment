package model

// User represents a user with username and a list of folders.
type User struct {
	Username string             `json:"username"`
	Folders  map[string]*Folder `json:"folders"`
}

// NewUser creates a new User.
func NewUser(username string) (*User, error) {
	err := ValidateInput(username)
	if err != nil {
		return nil, err
	}

	return &User{
		Username: username,
		Folders:  make(map[string]*Folder),
	}, nil
}
