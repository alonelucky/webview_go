//go:build !windows
// +build !windows

package webview

func SetWindowIcon(hwnd uintptr, ico string, showicon uint) {
}

func SetWindowIconEmbed(hwnd uintptr, ico []byte, showicon uint) {}
