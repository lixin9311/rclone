// Get block device size

// +build darwin

package local

import (
	"os"
	"syscall"
	"unsafe"
)

const (
	DKIOCGETBLOCKCOUNT = 0x40086419
	DKIOCGETBLOCKSIZE  = 0x40046418
)

func getDeviceSize(path string) (int64, error) {
	fd, err := syscall.Open(path, os.O_RDONLY, 0)
	if err != nil {
		return 0, err
	}
	defer syscall.Close(fd)
	var blksize int64
	var blkcnt int64
	if err := ioctl(uintptr(fd), DKIOCGETBLOCKCOUNT, uintptr(unsafe.Pointer(&blkcnt))); err != nil {
		return 0, err
	}
	if err := ioctl(uintptr(fd), DKIOCGETBLOCKSIZE, uintptr(unsafe.Pointer(&blksize))); err != nil {
		return 0, err
	}
	return blkcnt * blksize, nil
}
