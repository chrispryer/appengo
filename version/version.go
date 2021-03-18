package version

import "fmt"

var Version = "0.0.1"

func FullVersion() (string, error) {
	return fmt.Sprintf("appengo version: %s", Version)
}
