// Get block device size

// +build !darwin,!linux

package local

import "fmt"

func getDeviceSize(path string) (int64, error) {
	return 0, fmt.Errorf("Not implemented")
}
