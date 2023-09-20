package main

import (
	"fmt"

	"github.com/webui-dev/go-webui"
)

const login_html string = `<!DOCTYPE html>
<html>
  <head>
    <title>WebUI 2 - Go Example</title>
    <script src="webui.js"></script>
    <style>
      body {
        color: white;
        background: #0F2027;
        background: -webkit-linear-gradient(to right, #4e99bb, #2c91b5, #07587a);
        background: linear-gradient(to right, #4e99bb, #2c91b5, #07587a);
        text-align: center;
        font-size: 18px;
        font-family: sans-serif;
      }
    </style>
  </head>
  <body>
    <h1>WebUI 2 - Go Example</h1>
    <br>
    <input type="password" id="MyInput" OnKeyUp="document.getElementById('err').innerHTML='&nbsp;';" autocomplete="off">
    <br>
    <h3 id="err" style="color: #dbdd52">&nbsp;</h3>
    <br>
    <button id="CheckPassword">Check Password</button> - <button id="Exit">Exit</button>
  </body>
</html>`

const dashboard_html string = `<!DOCTYPE html>
<html>
  <head>
    <title>Dashboard</title>
    <script src="/webui.js"></script>
    <style>
      body {
        color: white;
        background: #0F2027;
        background: -webkit-linear-gradient(to right, #4e99bb, #2c91b5, #07587a);
        background: linear-gradient(to right, #4e99bb, #2c91b5, #07587a);
        text-align: center;
        font-size: 18px;
        font-family: sans-serif;
      }
    </style>
  </head>
  <body>
    <h1>Welcome !</h1>
    <br>
    Call Secret() function and get the response
    <br>
    <br>
    <button OnClick="webui.call('Sec').then((response) => { alert('Response is ' + response) });">Secret</button>
    <br>
    <br>
    <button id="Exit">Exit</button>
  </body>
</html>`

func Exit(e webui.Event) any {
	webui.Exit()

	return nil
}

func Secret(e webui.Event) any {
	return "I Love Go!"
}

func Check(e webui.Event) any {
	// Run script
	resp, err := e.Window.Script("return document.getElementById('MyInput').value;", webui.ScriptOptions{BufferSize: 1024})
	if err != nil {
		fmt.Printf("JavaScript Error: %v\n", err)
		return nil
	}

	fmt.Printf("Password: [%s]\n", resp)

	// Check the password
	if resp == "123456" {
		e.Window.Show(dashboard_html)
	} else {
		e.Window.Script("document.getElementById('err').innerHTML = 'Sorry. Wrong password';", webui.ScriptOptions{})
	}

	return nil
}

func main() {
	// Create a new window.
	w := webui.NewWindow()

	// Bind go functions.
	w.Bind("CheckPassword", Check)
	w.Bind("Sec", Secret)
	w.Bind("Exit", Exit)

	// Show html UI.
	w.Show(login_html)

	// Wait until all windows get closed.
	webui.Wait()

	fmt.Println("Thank you.")
}
