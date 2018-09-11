// Get block device size

// +build linux

package local

import (
	"os"
	"syscall"
	"unsafe"
)

const (
	BLKGETSIZE64 = 0x80081272
)

func getDeviceSize(path string) (int64, error) {
	fd, err := syscall.Open(path, os.O_RDONLY, 0)
	if err != nil {
		return 0, err
	}
	defer syscall.Close(fd)
	var size int64
	if err := ioctl(uintptr(fd), BLKGETSIZE64, uintptr(unsafe.Pointer(&size))); err != nil {
		return 0, err
	}
	return size, nil
}
