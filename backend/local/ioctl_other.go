// IOCTL wrap functions

// +build !darwin,!dragonfly,!freebsd,!linux,!netbsd,!openbsd,!solaris

package local

import "fmt"

func ioctl(fd, cmd, ptr uintptr) error {
	return fmt.Errorf("Not implemented")
}
