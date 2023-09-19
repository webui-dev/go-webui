package main

import "github.com/webui-dev/go-webui"

func main() {
	w := webui.NewWindow()
	w.Show("<html>Hello World</html>")
	webui.Wait()
}
