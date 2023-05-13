package main

import (
	b64 "encoding/base64"
	"fmt"
	"os"

	"github.com/webui-dev/go-webui/webui"
	"github.com/sqweek/dialog"  
)

var filePath string = ""

func Close(_ webui.Event) string {
	fmt.Println("Exit.")

	webui.Exit()

	return ""

}

func Save(e webui.Event) string {
	println("Save.")

	os.WriteFile(filePath, []byte(e.Data), 0644)

	return ""
}

func Open(e webui.Event) string {
	fmt.Println("Open.")

	filename, err := dialog.File().Load()

	if err == dialog.Cancelled{
		return ""
	}

	content, _ := os.ReadFile(filename)
	
	filePath = filename

	webui.Run(e.Window, fmt.Sprintf("addText('%s')", b64.StdEncoding.EncodeToString([]byte(content))))
	webui.Run(e.Window, fmt.Sprintf("SetFile('%s')", b64.StdEncoding.EncodeToString([]byte(filename))))

	return ""
}

func main () {
	window := webui.NewWindow()

	webui.Bind(window, "Open", Open)
	webui.Bind(window, "Save", Save)
	webui.Bind(window, "Close", Close)

	webui.Show(window, "ui/MainWindow.html")

	webui.Wait()
}