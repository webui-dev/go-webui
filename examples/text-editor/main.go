package main

import "github.com/webui-dev/go-webui/v2"

func Close(_ webui.Event) any {
	println("Exit.")

	webui.Exit()

	return nil

}

func main() {
	w := webui.NewWindow()

	w.SetRootFolder("ui")

	w.Bind("__close-btn", Close)

	if err := w.ShowBrowser("index.html", webui.ChromiumBased); err != nil {
		println("Warning: Install a Chromium-based web browser for an optimized experience.")
		w.Show("index.html")
	}

	webui.Wait()
}
