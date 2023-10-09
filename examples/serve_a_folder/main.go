package main

import (
	"time"

	ui "github.com/webui-dev/go-webui"
)

const (
	w  = ui.Window(1)
	w2 = ui.Window(2)
)

func events(e ui.Event) any {
	if e.EventType == ui.Connected {
		println("Connected.")
	} else if e.EventType == ui.Disconnected {
		println("Disconnected.")
	} else if e.EventType == ui.MouseClick {
		println("Click.")
	} else if e.EventType == ui.Navigation {
		target, _ := ui.GetArg[string](e)
		println("Starting navigation to: ", target)
		// Since we bind all events, following `href` links is blocked by WebUI.
		// To control the navigation, we need to use `Navigate()`.
		e.Window.Navigate(target)
	}
	return nil
}

func switchToSecondPage(e ui.Event) any {
	e.Window.Show("second.html")
	return nil
}

func showSecondWindow(e ui.Event) any {
	w2.Show("second.html")
	// Remove the Go Back button when showing the second page in another window.
	// Wait max. 10 seconds until the window is recognized as shown.
	for i := 0; i < 1000; i++ {
		if w2.IsShown() {
			break
		}
		// Slow down check interval to reduce load.
		time.Sleep(10 * time.Millisecond)
	}
	if !w2.IsShown() {
		return nil
	}
	// Let the DOM load.
	time.Sleep(50 * time.Millisecond)
	// Remove `Go Back` button.
	w2.Run("document.getElementById('go-back').remove();")

	return nil
}

func exit(e ui.Event) any {
	ui.Exit()
	return nil
}

func main() {
	w.NewWindow()

	// Bind HTML elements to functions
	w.Bind("switch-to-second-page", switchToSecondPage)
	w.Bind("open-new-window", showSecondWindow)
	w.Bind("exit", exit)
	// Bind all events.
	w.Bind("", events)

	// Show the main window.
	w.Show("index.html")

	// Prepare the second window.
	w2.NewWindow()
	w2.Bind("exit", exit)

	ui.SetRootFolder("ui")
	ui.Wait()
}
