package main

import (
	b64 "encoding/base64"
	"fmt"
	"os"

	"github.com/sqweek/dialog"
	"github.com/webui-dev/go-webui"
)

var filePath string = ""

func Close(_ webui.Event) any {
	fmt.Println("Exit.")

	webui.Exit()

	return nil

}

func Save(e webui.Event) any {
	println("Save.")

	os.WriteFile(filePath, []byte(e.Data), 0644)

	return nil
}

func Open(e webui.Event) any {
	fmt.Println("Open.")

	filename, err := dialog.File().Load()

	if err == dialog.Cancelled {
		return ""
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

	w.Bind("Open", Open)
	w.Bind("Save", Save)
	w.Bind("Close", Close)

	w.Show("ui/MainWindow.html")

	webui.Wait()
}
