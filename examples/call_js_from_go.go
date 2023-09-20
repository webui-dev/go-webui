package main

import (
	"fmt"
	"strconv"

	ui "github.com/webui-dev/go-webui"
)

// UI HTML
const doc = `<!DOCTYPE html>
<html>
	<head>
		<title>Call JavaScript from Go Example</title>
		<script src="webui.js"></script>
		<style>
			body {
				background: linear-gradient(to left, #36265a, #654da9);
				color: AliceBlue;
				font: 16px sans-serif;
				text-align: center;
				margin-top: 30px;
			}
			button {
				margin: 5px 0 10px;
			}
		</style>
	</head>
	<body>
		<h1>WebUI - Call JavaScript from Go</h1>
		<br>
		<button id="MyButton1">Count <span id="count">0<span></button>
		<br>
		<button id="MyButton2">Exit</button>
		<script>
			let count = document.getElementById("count").innerHTML;
			function SetCount(number) {
				document.getElementById("count").innerHTML = number;
				count = number;
			}
		</script>
	</body>
</html>`

func myCountFunc(e ui.Event) any {
	// Error handling omitted for brevity
	count, _ := e.Window.Script("return count;", ui.ScriptOptions{})
	i, _ := strconv.Atoi(count)
	e.Window.Run(fmt.Sprintf("SetCount(%v);", i+1))
	return nil
}

// Close all opened windows
func myExitFunc(e ui.Event) any {
	ui.Exit()
	return nil
}

func main() {
	// Create a window
	w := ui.NewWindow()

	// Bind HTML elements to functions
	w.Bind("MyButton1", myCountFunc)
	w.Bind("MyButton2", myExitFunc)

	// Show the window.
	w.Show(doc)

	// Wait until all windows get closed
	ui.Wait()
}
