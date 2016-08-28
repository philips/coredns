package middleware

import (
	"net/http"
	"os"
	"path/filepath"
	"runtime"
)

// dir wraps http.Dir that restrict file access to a specific directory tree.
type dir http.Dir

var CoreDir dir = dir(fsPath())

// fsPath returns the path to the folder where the application may store data.
// If COREDNSPATH env variable is set, that value is used. Otherwise, the path is
// the result of evaluating "$HOME/.coredns".
func fsPath() string {
	if corePath := os.Getenv("COREDNSPATH"); corePath != "" {
		return corePath
	}
	return filepath.Join(userHomeDir(), ".coredns")
}

// userHomeDir returns the user's home directory according to environment
// variables.
//
// Credit: http://stackoverflow.com/a/7922977/1048862
func userHomeDir() string {
	if runtime.GOOS == "windows" {
		home := os.Getenv("HOMEDRIVE") + os.Getenv("HOMEPATH")
		if home == "" {
			home = os.Getenv("USERPROFILE")
		}
		return home
	}
	return os.Getenv("HOME")
}
