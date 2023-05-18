package main

import "github.com/webui-dev/go-webui/v2"

func main() {
	var my_window = webui.NewWindow()

	webui.Show(my_window, "<html>Hello World</html>")
	webui.Wait()
}
