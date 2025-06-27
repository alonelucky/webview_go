//go:build windows
// +build windows

package webview

import (
	"syscall"
	"unsafe"
)

const (
	WindowIconTitle   = 0x01
	WindowIconTaskbar = 0x10
)

var (
	user32      = syscall.NewLazyDLL("user32.dll")
	modKernel32 = syscall.NewLazyDLL("kernel32.dll")

	sendMessage                  = user32.NewProc("SendMessageW")
	loadImage                    = user32.NewProc("LoadImageW")
	procCreateIconFromResourceEx = user32.NewProc("CreateIconFromResourceEx")
	pLookupIconIdFromDirectoryEx = user32.NewProc("LookupIconIdFromDirectoryEx")
)

func SetWindowIcon(hwnd uintptr, ico string) {
	// 加载图标文件
	icon, _, _ := loadImage.Call(
		0,
		uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr(ico))),
		1,    //IMAGE_ICON,
		0, 0, // 默认宽高
		0x0010, // LR_LOADFROMFILE
	)
	if icon == 0 {
		return
	}

	// 设置小图标（窗口标题栏）
	sendMessage.Call(
		hwnd,
		0x0080, // WM_SETICON,
		0,      // ICON_SMALL
		icon,
	)
}

func SetWindowIconEmbed(hwnd uintptr, ico []byte) {
	offset, _, _ := pLookupIconIdFromDirectoryEx.Call(
		uintptr(unsafe.Pointer(&ico[0])),
		uintptr(1),
		uintptr(0),
		uintptr(0),
		uintptr(0x00008000), /*LR_SHARED*/
	)

	icon, _, _ := procCreateIconFromResourceEx.Call(
		uintptr(unsafe.Pointer(&ico[offset])),
		uintptr(uint32(len(ico))),
		uintptr(1),
		uintptr(0x00030000),
		uintptr(0),
		uintptr(0),
		uintptr(0x00000040),
	)
	if icon == 0 {
		return
	}

	// 设置小图标（窗口标题栏）
	sendMessage.Call(
		hwnd,
		0x0080, // WM_SETICON,
		0,      // ICON_SMALL
		uintptr(unsafe.Pointer(icon)),
	)
}
