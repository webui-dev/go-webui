package main

import (
	"fmt"

	ui "github.com/webui-dev/go-webui"
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
		<button onclick="webui.call('MyID_One', 'Hello', 'World');">Call my_function_string()</button>
		<br>
		<button onclick="webui.call('MyID_Two', 123, 456, 789);">Call my_function_integer()</button>
		<br>
		<button onclick="webui.call('MyID_Three', true, false);">Call my_function_boolean()</button>
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
func myFunctionString(e ui.Event) ui.Void {
	str1, err := ui.GetArg[string](e)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	// Omit error handling from here on for brevity.
	str2, _ := ui.GetArgAt[string](e, 1)

	fmt.Printf("myFunctionString 1: %s\n", str1) // Hello
	fmt.Printf("myFunctionString 2: %s\n", str2) // World

	return nil
}

// JavaScript:
// webui.call('MyID_Two', 123456789);
func myFunctionInteger(e ui.Event) ui.Void {
	num1, _ := ui.GetArgAt[int](e, 0)
	num2, _ := ui.GetArgAt[int](e, 1)
	num3, _ := ui.GetArgAt[int](e, 2)

	fmt.Printf("myFunctionInteger 1: %d\n", num1) // 123
	fmt.Printf("myFunctionInteger 2: %d\n", num2) // 456
	fmt.Printf("myFunctionInteger 3: %d\n", num3) // 789

	return nil
}

// JavaScript:
// webui.call('MyID_Three', true);
func myFunctionBoolean(e ui.Event) ui.Void {
	status1, _ := ui.GetArg[bool](e)
	status2, _ := ui.GetArgAt[bool](e, 1)

	fmt.Printf("myFunctionBoolean 1: %t\n", status1) // true
	fmt.Printf("myFunctionBoolean 2: %t\n", status2) // false

	return nil
}

// JavaScript:
// const result = webui.call('MyID_Four', number);
func myFunctionWithResponse(e ui.Event) int {
	number, _ := ui.GetArg[int](e)

	response := number * 2
	fmt.Printf("myFunctionWithResponse: %d\n", response)

	return response
}

func main() {
	// Create a new window.
	w := ui.NewWindow()

	// Bind go functions.
	ui.Bind(w, "MyID_One", myFunctionString)
	ui.Bind(w, "MyID_Two", myFunctionInteger)
	ui.Bind(w, "MyID_Three", myFunctionBoolean)
	ui.Bind(w, "MyID_Four", myFunctionWithResponse)

	// Show html UI.
	w.Show(doc)

	// Wait until all windows get closed.
	ui.Wait()
}
