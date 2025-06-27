//go:build !windows
// +build !windows

package webview

func SetWindowIcon(hwnd uintptr, ico string) {
}

func SetWindowIconEmbed(hwnd uintptr, ico []byte) {}
