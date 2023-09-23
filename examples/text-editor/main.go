package main

import (
	b64 "encoding/base64"
	"fmt"
	"os"

	"github.com/sqweek/dialog"
	"github.com/webui-dev/go-webui"
)

var filePath string = ""

func Close(_ webui.Event) webui.Void {
	fmt.Println("Exit.")

	webui.Exit()

	return nil

}

func Save(e webui.Event) webui.Void {
	println("Save.")

	os.WriteFile(filePath, []byte(e.Data), 0644)

	return nil
}

func Open(e webui.Event) webui.Void {
	fmt.Println("Open.")

	filename, err := dialog.File().Load()

	if err == dialog.Cancelled {
		return nil
	}

	content, err := os.ReadFile(filename)

	if err != nil {
		fmt.Println("Error reading file ", filename)
		fmt.Println("Error: ", err)
		return nil
	}

	filePath = filename

	e.Window.Run(fmt.Sprintf("addText('%s')", b64.StdEncoding.EncodeToString([]byte(content))))
	e.Window.Run(fmt.Sprintf("SetFile('%s')", b64.StdEncoding.EncodeToString([]byte(filename))))

	return nil
}

func main() {
	w := webui.NewWindow()

	webui.Bind(w, "Open", Open)
	webui.Bind(w, "Save", Save)
	webui.Bind(w, "Close", Close)

	w.Show("ui/MainWindow.html")

	webui.Wait()
}
