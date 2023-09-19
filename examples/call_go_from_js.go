package main

import (
	"fmt"

	"github.com/webui-dev/go-webui"
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
		<button onclick="webui.call('MyID_One', 'Hello');">Call my_function_string()</button>
		<br>
		<button onclick="webui.call('MyID_Two', 123456789);">Call my_function_integer()</button>
		<br>
		<button onclick="webui.call('MyID_Three', true);">Call my_function_boolean()</button>
		<br>
		<p>Call a V function that returns a response</p>
		<button onclick="MyJS();">Call my_function_with_response()</button>
		<div>Double: <input type="text" id="MyInputID" value="2"></div>
		<script>
			async function MyJS() {
				const MyInput = document.getElementById("MyInputID");
				const number = MyInput.value;
				const result = await webui.call("MyID_Four", number);
				MyInput.value = result;
			}
		</script>
	</body>
</html>`

// JavaScript:
// webui.call('MyID_One', 'Hello');
func myFunctionString(e webui.Event) any {
	response := e.Data.String()

	fmt.Printf("myFunctionString: %s\n", response) // Hello

	// Need Multiple Arguments?
	//
	// WebUI supports only one argument. For multiple arguments,
	// send a JSON string from JavaScript and decode it.

	return nil
}

// JavaScript:
// webui.call('MyID_Two', 123456789);
func myFunctionInteger(e webui.Event) any {
	response := e.Data.Int()

	fmt.Printf("myFunctionInteger: %d\n", response) // 123456789

	return nil
}

// JavaScript:
// webui.call('MyID_Three', true);
func myFunctionBoolean(e webui.Event) any {
	response := e.Data.Bool()

	fmt.Printf("myFunctionBoolean: %t\n", response) // true

	return nil
}

// JavaScript:
// const result = webui.call('MyID_Four', number);
func myFunctionWithResponse(e webui.Event) any {
	number := e.Data.Int() * 2

	fmt.Printf("myFunctionWithResponse: %d\n", number)

	return number
}

func main() {
	// Create a new window.
	w := webui.NewWindow()

	// Bind go functions.
	w.Bind("MyID_One", myFunctionString)
	w.Bind("MyID_Two", myFunctionInteger)
	w.Bind("MyID_Three", myFunctionBoolean)
	w.Bind("MyID_Four", myFunctionWithResponse)

	// Show html UI.
	w.Show(doc)

	// Wait until all windows get closed.
	webui.Wait()
}
