package kutils

import (
	"os"
	"os/user"

	"github.com/kumose/kprinter"
)

// CurrentUser returns current login user
func CurrentUser() string {
	user, err := user.Current()
	if err != nil {
		kprinter.Errorf("Get current user: %s", err)
		return "root"
	}
	return user.Username
}

// UserHome returns home directory of current user
func UserHome() string {
	homedir, err := os.UserHomeDir()
	if err != nil {
		kprinter.Errorf("Get current user home: %s", err)
		return "root"
	}
	return homedir
}
