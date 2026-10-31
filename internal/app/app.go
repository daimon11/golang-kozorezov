// Package app contains the application logic shared by the commands.
package app

import (
	"fmt"
	"io"
)

// Run writes the startup message to w.
func Run(w io.Writer) error {
	_, err := fmt.Fprintln(w, "homeworks: ok")
	return err
}
