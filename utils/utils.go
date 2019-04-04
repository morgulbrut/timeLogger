/*
Copyright © 2019 morgulbrut
This work is free. You can redistribute it and/or modify it under the
terms of the Do What The Fuck You Want To Public License, Version 2,
as published by Sam Hocevar. See the LICENSE file or
 http://www.wtfpl.net/ for more details.
*/
package utils

import (
	"fmt"
	"os"

	"github.com/morgulbrut/color"
)

// Error prints an error and quits the programm
func Error(err string) {
	color.Red(fmt.Sprintf("ERROR: %s", err))
	os.Exit(1)
}

// AppendToFile appends a string as a new line to a file
func AppendToFile(s, filename string) (err error) {
	// If the file doesn't exist, create it, or append to the file
	f, err := os.OpenFile(filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	_, err = f.Write([]byte(s + "\n"))
	err = f.Close()
	return err
}
