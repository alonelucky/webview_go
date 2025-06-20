package main

import webview "github.com/webview/webview_go"

func main() {
	w := webview.New(false, "G:\\golang\\webview_go\\examples\\basic\\test")
	defer w.Destroy()
	w.SetTitle("Basic Example")
	w.SetSize(480, 320, webview.HintNone)
	w.SetHtml("Thanks for using webview!")
	w.Run()
}
