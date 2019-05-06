// Code generated by 'go generate'; DO NOT EDIT.

package conf

import (
	"syscall"
	"unsafe"

	"golang.org/x/sys/windows"
)

var _ unsafe.Pointer

// Do the interface allocations only once for common
// Errno values.
const (
	errnoERROR_IO_PENDING = 997
)

var (
	errERROR_IO_PENDING error = syscall.Errno(errnoERROR_IO_PENDING)
)

// errnoErr returns common boxed Errno values, to prevent
// allocations at runtime.
func errnoErr(e syscall.Errno) error {
	switch e {
	case 0:
		return nil
	case errnoERROR_IO_PENDING:
		return errERROR_IO_PENDING
	}
	// TODO: add more here, after collecting data on the common
	// error values see on Windows. (perhaps when running
	// all.bat?)
	return e
}

var (
	modole32    = windows.NewLazySystemDLL("ole32.dll")
	modshell32  = windows.NewLazySystemDLL("shell32.dll")
	modkernel32 = windows.NewLazySystemDLL("kernel32.dll")

	procCoTaskMemFree                = modole32.NewProc("CoTaskMemFree")
	procSHGetKnownFolderPath         = modshell32.NewProc("SHGetKnownFolderPath")
	procFindFirstChangeNotificationW = modkernel32.NewProc("FindFirstChangeNotificationW")
	procFindNextChangeNotification   = modkernel32.NewProc("FindNextChangeNotification")
)

func coTaskMemFree(pointer uintptr) {
	syscall.Syscall(procCoTaskMemFree.Addr(), 1, uintptr(pointer), 0, 0)
	return
}

func shGetKnownFolderPath(id *windows.GUID, flags uint32, token windows.Handle, path **uint16) (err error) {
	r1, _, e1 := syscall.Syscall6(procSHGetKnownFolderPath.Addr(), 4, uintptr(unsafe.Pointer(id)), uintptr(flags), uintptr(token), uintptr(unsafe.Pointer(path)), 0, 0)
	if r1 != 0 {
		if e1 != 0 {
			err = errnoErr(e1)
		} else {
			err = syscall.EINVAL
		}
	}
	return
}

func findFirstChangeNotification(path *uint16, watchSubtree bool, filter uint32) (handle windows.Handle, err error) {
	var _p0 uint32
	if watchSubtree {
		_p0 = 1
	} else {
		_p0 = 0
	}
	r0, _, e1 := syscall.Syscall(procFindFirstChangeNotificationW.Addr(), 3, uintptr(unsafe.Pointer(path)), uintptr(_p0), uintptr(filter))
	handle = windows.Handle(r0)
	if handle == 0 {
		if e1 != 0 {
			err = errnoErr(e1)
		} else {
			err = syscall.EINVAL
		}
	}
	return
}

func findNextChangeNotification(handle windows.Handle) (err error) {
	r1, _, e1 := syscall.Syscall(procFindNextChangeNotification.Addr(), 1, uintptr(handle), 0, 0)
	if r1 == 0 {
		if e1 != 0 {
			err = errnoErr(e1)
		} else {
			err = syscall.EINVAL
		}
	}
	return
}