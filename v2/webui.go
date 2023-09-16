package webui

/*
  WebUI Library 2.4.0
  http://webui.me
  https://github.com/webui-dev/webui
  Copyright (c) 2020-2023 Hassan Draga.
  Licensed under MIT License.
  All rights reserved.
  Canada.
*/

/*
#cgo CFLAGS: -I ./
#cgo windows LDFLAGS: -L ./webui-windows-gcc-x64 -lwebui-2-static -lws2_32
#cgo darwin LDFLAGS: -L ./webui-macos-clang-x64 -lwebui-2-static -lpthread -lm
#cgo linux LDFLAGS: -L ./webui-linux-gcc-x64 -lwebui-2-static -lpthread -lm
#include <webui.h>
extern void goWebuiEvent(size_t _window, size_t _event_type, char* _element, char* _data, size_t _event_number);
static void go_webui_event_handler(webui_event_t* e) {
	goWebuiEvent(e->window, e->event_type, e->element, e->data, e->event_number);
}
static size_t go_webui_bind(size_t win, const char* element) {
	return webui_bind(win, element, go_webui_event_handler);
}
*/
import "C"

import (
	"bytes"
	"encoding/json"
	"log"
	"strconv"
	"unsafe"
)

// Web browsers enum
const AnyBrowser uint = 0 // 0. Default recommended web browser
const Chrome uint = 1     // 1. Google Chrome
const Firefox uint = 2    // 2. Mozilla Firefox
const Edge uint = 3       // 3. Microsoft Edge
const Safari uint = 4     // 4. Apple Safari
const Chromium uint = 5   // 5. The Chromium Project
const Opera uint = 6      // 6. Opera Browser
const Brave uint = 7      // 7. The Brave Browser
const Vivaldi uint = 8    // 8. The Vivaldi Browser
const Epic uint = 9       // 9. The Epic Browser
const Yandex uint = 10    // 10. The Yandex Browser

// Runtimes enum
const None uint = 0
const Deno uint = 1
const Nodejs uint = 2

// Events enum
const WEBUI_EVENT_DISCONNECTED uint = 0        // 0. Window disconnection event
const WEBUI_EVENT_CONNECTED uint = 1           // 1. Window connection event
const WEBUI_EVENT_MULTI_CONNECTION uint = 2    // 2. New window connection event
const WEBUI_EVENT_UNWANTED_CONNECTION uint = 3 // 3. New unwanted window connection event
const WEBUI_EVENT_MOUSE_CLICK uint = 4         // 4. Mouse click event
const WEBUI_EVENT_NAVIGATION uint = 5          // 5. Window navigation event
const WEBUI_EVENT_CALLBACK uint = 6            // 6. Function call event

type Window uint

type Data string

type Event struct {
	Window    Window
	EventType uint
	Element   string
	Data      Data
}

type JavaScript struct {
	Timeout    uint
	BufferSize uint
	Response   string
}

// User Go Callback Functions list
var funcList = make(map[Window]map[uint]func(Event) any)

// This private function receives all events
//
//export goWebuiEvent
func goWebuiEvent(window C.size_t, _event_type C.size_t, _element *C.char, _data *C.char, _event_number C.size_t) {
	// Create a new event struct
	e := Event{
		Window:    Window(window),
		EventType: uint(_event_type),
		Element:   C.GoString(_element),
		Data:      Data(C.GoString(_data)),
	}
	// Call user callback function
	funcId := uint(C.webui_interface_get_bind_id(window, _element))
	result := funcList[Window(window)][funcId](e)
	if result == nil {
		return
	}
	jsonRes, err := json.Marshal(result)
	if err != nil {
		log.Println("Failed encoding JS result into JSON", err)
	}
	C.webui_interface_set_response(window, _event_number, C.CString(string(jsonRes)))
}

// -- Public APIs --

// JavaScript object constructor
func NewJavaScript() JavaScript {
	js := JavaScript{
		Timeout:    0,
		BufferSize: (1024 * 8),
		Response:   "",
	}
	return js
}

// Run a JavaScript, and get the response back (Make sure your local buffer can hold the response).
func (w Window) Script(js *JavaScript, script string) bool {
	// Create a local buffer to hold the response
	ResponseBuffer := make([]byte, uint64(js.BufferSize))

	// Create a pointer to the local buffer
	ptr := (*C.char)(unsafe.Pointer(&ResponseBuffer[0]))

	// Run the JavaScript and wait for response
	status := C.webui_script(C.size_t(w), C.CString(script), C.size_t(js.Timeout), ptr, C.size_t(uint64(js.BufferSize)))

	// Copy the response to the users struct
	ResponseLen := bytes.IndexByte(ResponseBuffer[:], 0)
	js.Response = string(ResponseBuffer[:ResponseLen])

	// return the status of the JavaScript execution
	// True: No JavaScript error.
	// False: JavaScript error.
	return bool(status)
}

// Run JavaScript quickly with no waiting for the response.
func (w Window) Run(script string) {
	C.webui_run(C.size_t(w), C.CString(script))
}

func Encode(str string) string {
	return C.GoString(C.webui_encode(C.CString(str)))
}

func Decode(str string) string {
	return C.GoString(C.webui_decode(C.CString(str)))
}

// Chose between Deno and Nodejs runtime for .js and .ts files.
func (w Window) SetRuntime(runtime uint) {
	C.webui_set_runtime(C.size_t(w), C.size_t(runtime))
}

// Create a new window object
func NewWindow() Window {
	w := Window(C.size_t(C.webui_new_window()))
	funcList[w] = make(map[uint]func(Event) any)
	return w
}

// Check a specific window if it's still running
func (w Window) IsShown() bool {
	status := C.webui_is_shown(C.size_t(w))
	return bool(status)
}

// Close a specific window.
func (w Window) Close() {
	C.webui_close(C.size_t(w))
}

// Set the maximum time in seconds to wait for browser to start
func SetTimeout(seconds uint) {
	C.webui_set_timeout(C.size_t(seconds))
}

// Allow the window URL to be re-used in normal web browsers
func (w Window) SetMultiAccess(access bool) {
	C.webui_set_multi_access(C.size_t(w), C._Bool(access))
}

// Close all opened windows
func Exit() {
	C.webui_exit()
}

// Show a window using a embedded HTML, or a file. If the window is already opened then it will be refreshed.
func (w Window) Show(content string) {
	C.webui_show(C.size_t(w), C.CString(content))
}

// Same as Show(). But with a specific web browser.
func (w Window) ShowBrowser(content string, browser uint) {
	C.webui_show_browser(C.size_t(w), C.CString(content), C.size_t(browser))
}

// Wait until all opened windows get closed.
func Wait() {
	C.webui_wait()
}

// Bind a specific html element click event with a function. Empty element means all events.
func (w Window) Bind(element string, callback func(Event) any) {
	funcId := uint(C.go_webui_bind(C.size_t(w), C.CString(element)))
	funcList[w][funcId] = callback
}

func (d Data) String() string {
	return string(d)
}

func (d Data) Int() int {
	num, err := strconv.Atoi(string(d))
	if err != nil {
		log.Println("Failed getting event int argument", err)
	}
	return num
}

func (d Data) Bool() bool {
	boolVal, err := strconv.ParseBool(string(d))
	if err != nil {
		log.Println("Failed getting event bool argument", err)
	}
	return boolVal
}
