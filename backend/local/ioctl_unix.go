// IOCTL wrap functions

// +build darwin dragonfly freebsd linux netbsd openbsd solaris

package local

import "syscall"

func ioctl(fd, cmd, ptr uintptr) error {
	_, _, e := syscall.Syscall(syscall.SYS_IOCTL, fd, cmd, ptr)
	if e != 0 {
		return e
	}
	return nil
}
