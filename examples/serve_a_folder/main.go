package main

import (
	"fmt"

	ui "github.com/webui-dev/go-webui/v2"
)

// Those constants are to avoid using global variables
// you can also use `const w ui.Window = 1`.
// A window ID can be between 1 and 512 (WEBUI_MAX_IDS)
const (
	MyWindow       = ui.Window(1)
	MySecondWindow = ui.Window(2)
)

func exitApp(e ui.Event) any {
	// Close all opened windows
	ui.Exit()
	return nil
}

func events(e ui.Event) any {
	// This function gets called every time
	// there is an event

	if e.EventType == ui.Connected {
		fmt.Println("Connected.")
	} else if e.EventType == ui.Disconnected {
		fmt.Println("Disconnected.")
	} else if e.EventType == ui.MouseClick {
		fmt.Println("Click.")
	} else if e.EventType == ui.Navigation {
		url, _ := ui.GetArg[string](e)
		fmt.Printf("Starting navigation to: %s\n", url)

		// Because we used `w.Bind("", events)`
		// WebUI will block all `href` link clicks and sent here instead.
		// We can then control the behaviour of links as needed.
		e.Window.Navigate(url)
	}
	return nil
}

func switchToSecondPage(e ui.Event) any {
	// This function gets called every
	// time the user clicks on "SwitchToSecondPage"

	// Switch to `/second.html` in the same opened window.
	e.Window.Show("second.html")
	return nil
}

func showSecondWindow(e ui.Event) any {
	// This function gets called every
	// time the user clicks on "OpenNewWindow"

	// Show a new window, and navigate to `/second.html`
	// if it's already open, then switch in the same window
	MySecondWindow.Show("second.html")
	return nil
}

// Counter for dynamic content
var count int

func myFilesHandler(filename string) ([]byte, int) {
	fmt.Printf("File: %s\n", filename)

	if filename == "/test.txt" {
		// Const static file example
		response := "HTTP/1.1 200 OK\r\n" +
			"Content-Type: text/html\r\n" +
			"Content-Length: 99\r\n\r\n" +
			"<html>" +
			"   This is a static embedded file content example." +
			"   <script src=\"webui.js\"></script>" + // To keep connection with WebUI
			"</html>"
		return []byte(response), len(response)
	} else if filename == "/dynamic.html" {
		// Dynamic file example

		// Generate body
		count++
		body := fmt.Sprintf(
			"<html>"+
				"   This is a dynamic file content example. <br>"+
				"   Count: %d <a href=\"dynamic.html\">[Refresh]</a><br>"+
				"   <script src=\"webui.js\"></script>"+ // To keep connection with WebUI
				"</html>",
			count,
		)

		// Generate header + body
		response := fmt.Sprintf(
			"HTTP/1.1 200 OK\r\n"+
				"Content-Type: text/html\r\n"+
				"Content-Length: %d\r\n\r\n"+
				"%s",
			len(body), body,
		)

		return []byte(response), len(response)
	}

	// Other files:
	// A nil return will make WebUI
	// look for the file locally.
	return nil, 0
}

func main() {
	ui.SetDefaultRootFolder("ui")
	
	// Create new windows
	MyWindow.NewWindow()
	MySecondWindow.NewWindow()

	// Bind HTML element IDs with Go functions
	MyWindow.Bind("SwitchToSecondPage", switchToSecondPage)
	MyWindow.Bind("OpenNewWindow", showSecondWindow)
	MyWindow.Bind("Exit", exitApp)
	MySecondWindow.Bind("Exit", exitApp)

	// Bind events
	MyWindow.Bind("", events)

	// Set the `.ts` and `.js` runtime
	// MyWindow.SetRuntime(ui.Nodejs)
	// MyWindow.SetRuntime(ui.Bun)
	MyWindow.SetRuntime(ui.Deno)

	// Set a custom files handler
	MyWindow.SetFileHandler(myFilesHandler)

	// Set window size
	MyWindow.SetSize(800, 800)

	// Set window position
	MyWindow.SetPosition(200, 200)

	// Show a new window
	// MyWindow.SetRootFolder("_MY_PATH_HERE_")
	// MyWindow.ShowBrowser("index.html", ui.Chrome)
	MyWindow.Show("index.html")

	// Wait until all windows get closed
	ui.Wait()

	// Free all memory resources (Optional)
	ui.Clean()
}
