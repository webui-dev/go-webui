package main

import (
	"fmt"

	ui "github.com/webui-dev/go-webui/v2"
)

const doc = `<!DOCTYPE html>
<html>
	<head>
		<title>Call Go from JavaScript Example</title>
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
		<h1>WebUI - Call Go from JavaScript</h1>
		<br>
		<p>Call Go functions with arguments (<em>See the logs in your terminal</em>)</p>
		<button onclick="webui.handleStr('Hello', 'World');">Call handle_str()</button>
		<br>
		<button onclick="webui.handleInt(123, 456, 789);">Call handle_int()</button>
		<br>
		<button onclick="webui.handleBool(true, false);">Call handle_bool()</button>
		<br>
		<p>Call a Go function that returns a response</p>
		<button onclick="getRespFromGo();">Call get_response()</button>
		<div>Double: <input type="text" id="my-input" value="2"></div>
		<script>
			async function getRespFromGo() {
				const myInput = document.getElementById("my-input");
				const number = myInput.value;
				const result = await webui.handleResp(number);
				myInput.value = result;
			}
		</script>
	</body>
</html>`

// JavaScript: `webui.handleStr('Hello', 'World');`
func handleStr(e ui.Event) ui.Void {
	str1, err := ui.GetArg[string](e)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	// Omit error handling from here on for brevity.
	str2, _ := ui.GetArgAt[string](e, 1)

	fmt.Printf("handleStr 1: %s\n", str1) // Hello
	fmt.Printf("handleStr 2: %s\n", str2) // World

	return nil
}

// JavaScript: `webui.handleInt(123, 456, 789);`
func handleInt(e ui.Event) ui.Void {
	num1, _ := ui.GetArgAt[int](e, 0)
	num2, _ := ui.GetArgAt[int](e, 1)
	num3, _ := ui.GetArgAt[int](e, 2)

	fmt.Printf("handleInt 1: %d\n", num1) // 123
	fmt.Printf("handleInt 2: %d\n", num2) // 456
	fmt.Printf("handleInt 3: %d\n", num3) // 789

	return nil
}

// JavaScript: webui.handleBool(true, false);
func handleBool(e ui.Event) ui.Void {
	status1, _ := ui.GetArg[bool](e)
	status2, _ := ui.GetArgAt[bool](e, 1)

	fmt.Printf("handleBool 1: %t\n", status1) // true
	fmt.Printf("handleBool 2: %t\n", status2) // false

	return nil
}

// JavaScript: `const result = await webui.getResponse(number);`
func handleResp(e ui.Event) int {
	number, _ := ui.GetArg[int](e)

	response := number * 2
	fmt.Printf("handleResp: %d\n", response)

	return response
}

func main() {
	// Create a new window.
	w := ui.NewWindow()

	// Bind go functions.
	ui.Bind(w, "handleStr", handleStr)
	ui.Bind(w, "handleInt", handleInt)
	ui.Bind(w, "handleBool", handleBool)
	ui.Bind(w, "handleResp", handleResp)

	// Show html UI.
	w.Show(doc)

	// Wait until all windows get closed.
	ui.Wait()
}
