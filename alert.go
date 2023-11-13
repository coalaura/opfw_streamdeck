package main

import (
	"syscall"
	"unsafe"
)

var (
	user32         = syscall.NewLazyDLL("user32.dll")
	procMessageBox = user32.NewProc("MessageBoxW")
)

const (
	MB_OK          = 0x00000000
	MB_SYSTEMMODAL = 0x00001000
	MB_ICONWARNING = 0x00000030
)

func alert(msg string) {
	lpCaption, _ := syscall.UTF16PtrFromString("OP-FW Streamdeck")
	lpText, _ := syscall.UTF16PtrFromString(msg)

	syscall.SyscallN(procMessageBox.Addr(),
		0,
		uintptr(unsafe.Pointer(lpText)),
		uintptr(unsafe.Pointer(lpCaption)),
		MB_OK|MB_ICONWARNING|MB_SYSTEMMODAL,
	)
}
