package main

import (
	"github.com/webui-dev/go-webui"
)

func Close(_ webui.Event) any {
	println("Exit.")

	webui.Exit()

	return nil

}

func main() {
	w := webui.NewWindow()

	w.Bind("Close", Close)

	w.Show("ui/MainWindow.html")

	webui.Wait()
}
