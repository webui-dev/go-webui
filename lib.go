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
#cgo CFLAGS: -Iwebui
#cgo windows LDFLAGS: -Lwebui/webui-windows-gcc-x64 -lwebui-2-static -lws2_32
#cgo darwin,amd64 LDFLAGS: -Lwebui/webui-macos-clang-x64 -lwebui-2-static -lpthread -lm
#cgo darwin,arm64 LDFLAGS: -Lwebui/webui-macos-clang-arm64 -lwebui-2-static -lpthread -lm
#cgo linux LDFLAGS: -Lwebui/webui-linux-gcc-x64 -lwebui-2-static -lpthread -lm

#include <webui.h>
extern void goWebuiEvent(size_t _window, size_t _event_type, char* _element, char* _data, size_t _size, size_t _event_number);
static void go_webui_event_handler(webui_event_t* e) {
	goWebuiEvent(e->window, e->event_type, e->element, e->data, e->size, e->event_number);
}
static size_t go_webui_bind(size_t win, const char* element) {
	return webui_bind(win, element, go_webui_event_handler);
}
*/
import "C"

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"strconv"
	"time"
	"unsafe"
)

type Browser uint8

const (
	AnyBrowser Browser = iota
	Chrome
	Firefox
	Edge
	Safari
	Chromium
	Opera
	Brave
	Vivaldi
	Epic
	Yandex
	ChromiumBased
)

type Runtime uint8

const (
	None Runtime = iota
	Deno
	Nodejs
)

type EventType uint8

const (
	Disconnected EventType = iota
	Connected
	MultiConnection
	UnwantedConnection
	MouseClick
	Navigation
	Callback
)

type Window uint

type Data string

type Event struct {
	Window    Window
	EventType EventType
	Element   string
	Data      Data
	Size      uint
}

type ScriptOptions struct {
	Timeout    uint
	BufferSize uint
}

type Void *struct{}

type noArgError struct {
	element string
}

type getArgError struct {
	err     error
	element string
	typ     string
}

// User Go Callback Functions list
var funcList = make(map[Window]map[uint]func(Event) any)

// == Definitions =============================================================

// NewWindow creates a new WebUI window object and returns the window number.
func NewWindow() Window {
	w := Window(C.size_t(C.webui_new_window()))
	funcList[w] = make(map[uint]func(Event) any)
	return w
}

// NewWindow creates a new webui window object using a specified window number.
func (w Window) NewWindow() {
	funcList[w] = make(map[uint]func(Event) any)
	C.webui_new_window_id(C.size_t(w))
}

// NewWindowId returns a free window number that can be used with `NewWindow`.
func NewWindowId() Window {
	return Window(C.webui_get_new_window_id())
}

// Private function that receives and handles webui events as go events
//
//export goWebuiEvent
func goWebuiEvent(_window C.size_t, _event_type C.size_t, _element *C.char, _data *C.char, _size C.size_t, _event_number C.size_t) {
	// Create a new event struct
	w := Window(_window)
	e := Event{
		Window:    w,
		EventType: EventType(_event_type),
		Element:   C.GoString(_element),
		Data:      Data(C.GoString(_data)),
		Size:      uint(_size),
	}
	// Call user callback function
	funcId := uint(C.webui_interface_get_bind_id(_window, _element))
	result := funcList[w][funcId](e)
	if result == nil {
		return
	}
	response, err := json.Marshal(result)
	if err != nil {
		log.Println("Failed encoding JS result into JSON", err)
	}
	C.webui_interface_set_response(_window, _event_number, C.CString(string(response)))
}

// Bind binds a specific html element click event with a function. Empty element means all events.
func (w Window) Bind(element string, callback func(Event) any) {
	funcId := uint(C.go_webui_bind(C.size_t(w), C.CString(element)))
	funcList[w][funcId] = callback
}

// Bind binds a specific html element click event with a function. Empty element means all events.
func Bind[T any](w Window, element string, callback func(Event) T) {
	funcId := uint(C.go_webui_bind(C.size_t(w), C.CString(element)))
	funcList[w][funcId] = func(e Event) any {
		return callback(e)
	}
}

// Show opens a window using embedded HTML, or a file. If the window is already open, it will be refreshed.
func (w Window) Show(content string) (err error) {
	if !C.webui_show(C.size_t(w), C.CString(content)) {
		err = errors.New("Failed showing window.")
	}
	return
}

// ShowBrowser opens a window using embedded HTML, or a file in a specific web browser.
// If the window is already open, it will be refreshed.
func (w Window) ShowBrowser(content string, browser Browser) (err error) {
	if !C.webui_show_browser(C.size_t(w), C.CString(content), C.size_t(browser)) {
		err = errors.New("Failed showing window.")
	}
	return
}

// SetKiosk determines whether Kiosk mode (full screen) is enabled for the window.
func (w Window) SetKiosk(enable bool) {
	C.webui_set_kiosk(C.size_t(w), C._Bool(enable))
}

// Wait waits until all opened windows get closed.
func Wait() {
	C.webui_wait()
}

// Close closes the window. The window object will still exist.
func (w Window) Close() {
	C.webui_close(C.size_t(w))
}

// Destroy closes the window and free all memory resources.
func (w Window) Destroy() {
	C.webui_destroy(C.size_t(w))
}

// Exit closes all open windows. `Wait()` will return (Break).
func Exit() {
	C.webui_exit()
	// Quickfix segfault on `Exit` call when all Events where bound. E.g.: `w.Bind("", events)`
	time.Sleep(1 * time.Millisecond)
}

// SetRootFolder sets the web-server root folder path for the window.
func (w Window) SetRootFolder(path string) {
	C.webui_set_root_folder(C.size_t(w), C.CString(path))
}

// SetRootFolder sets the web-server root folder path for all windows.
func SetRootFolder(path string) {
	C.webui_set_default_root_folder(C.CString(path))
}

// IsShown checks if the window it's still running.
func (w Window) IsShown() bool {
	status := C.webui_is_shown(C.size_t(w))
	return bool(status)
}

// SetTimeout sets the maximum time in seconds to wait for the browser to start.
func SetTimeout(seconds uint) {
	C.webui_set_timeout(C.size_t(seconds))
}

// SetIcon sets the default embedded HTML favicon.
func (w Window) SetIcon(icon string, icon_type string) {
	C.webui_set_icon(C.size_t(w), C.CString(icon), C.CString(icon_type))
}

// SetMultiAccess determines whether the window URL can be reused in normal web browsers.
func (w Window) SetMultiAccess(access bool) {
	C.webui_set_multi_access(C.size_t(w), C._Bool(access))
}

// Encode sends text based data to the UI using base64 encoding.
func Encode(str string) string {
	return C.GoString(C.webui_encode(C.CString(str)))
}

// Decode decodes Base64 encoded text received from the the UI.
func Decode(str string) string {
	return C.GoString(C.webui_decode(C.CString(str)))
}

// SetHide determines whether the window is run in hidden mode.
func (w Window) SetHide(status bool) {
	C.webui_set_hide(C.size_t(w), C._Bool(status))
}

// SetSize sets the window size.
func (w Window) SetSize(width uint, height uint) {
	C.webui_set_size(C.size_t(w), C.uint(width), C.uint(height))
}

// SetPosition sets the window position.
func (w Window) SetPosition(x uint, y uint) {
	C.webui_set_position(C.size_t(w), C.uint(x), C.uint(y))
}

// SetProfile sets the web browser profile to use.
// An empty `name` and `path` means the default user profile.
// Needs to be called before `webui_show()`.
func (w Window) SetProfile(name string, path string) {
	C.webui_set_profile(C.size_t(w), C.CString(name), C.CString(path))
}

// GetUrl returns the full current URL
func (w Window) GetUrl() string {
	return C.GoString(C.webui_get_url(C.size_t(w)))
}

// Navigate navigates to a specific URL
func (w Window) Navigate(url string) {
	C.webui_navigate(C.size_t(w), C.CString(url))
}

// == Javascript ==============================================================

// Run executres JavaScript without waiting for the response.
func (w Window) Run(script string) {
	C.webui_run(C.size_t(w), C.CString(script))
}

// Script executes JavaScript and returns the response (Make sure the response buffer can hold the response).
// The default BufferSize is 8KiB.
func (w Window) Script(script string, options ScriptOptions) (resp string, err error) {
	opts := ScriptOptions{
		Timeout:    options.Timeout,
		BufferSize: options.BufferSize,
	}
	if options.BufferSize == 0 {
		opts.BufferSize = (1024 * 8)
	}

	// Create a local buffer to hold the response
	buffer := make([]byte, uint64(opts.BufferSize))

	// Create a pointer to the local buffer
	ptr := (*C.char)(unsafe.Pointer(&buffer[0]))

	// Run the script and wait for the response
	ok := C.webui_script(C.size_t(w), C.CString(script), C.size_t(opts.Timeout), ptr, C.size_t(uint64(opts.BufferSize)))
	if !ok {
		err = errors.New(fmt.Sprintf("Failed running script: %s.\n", script))
	}
	respLen := bytes.IndexByte(buffer[:], 0)
	resp = string(buffer[:respLen])

	return
}

// SetRuntime sets the runtime for .js and .ts files to Deno and Nodejs.
func (w Window) SetRuntime(runtime Runtime) {
	C.webui_set_runtime(C.size_t(w), C.size_t(runtime))
}

func (e *noArgError) Error() string {
	return fmt.Sprintf("`%s` did not receive an argument.", e.element)
}

func (e *getArgError) Error() string {
	return fmt.Sprintf("Failed getting argument of type `%s` for `%s`. %v", e.typ, e.element, e.err)
}

// Int parses the JavaScript argument as integer.
func (e Event) Int() (arg int, err error) {
	if e.Size == 0 {
		err = &noArgError{e.Element}
	}
	arg, err = strconv.Atoi(string(e.Data))
	if err != nil {
		err = &getArgError{err, "int", e.Element}
	}
	return
}

// String parses the JavaScript argument as integer.
func (e Event) String() (arg string, err error) {
	if e.Size == 0 {
		err = &noArgError{e.Element}
	}
	arg = string(e.Data)
	return
}

// Bool parses the JavaScript argument as integer.
func (e Event) Bool() (arg bool, err error) {
	if e.Size == 0 {
		err = &noArgError{e.Element}
	}
	arg, err = strconv.ParseBool(string(e.Data))
	if err != nil {
		err = &getArgError{err, "bool", e.Element}
	}
	return
}

// GetArg parses the JavaScript argument into a Go data type.
func GetArg[T any](e Event) (arg T, err error) {
	if e.Size == 0 {
		err = &noArgError{e.Element}
	}
	var ret T
	switch p := any(&ret).(type) {
	case *string:
		*p, err = e.String()
	case *int:
		*p, err = e.Int()
	case *bool:
		*p, err = e.Bool()
	default:
		err = json.Unmarshal([]byte(e.Data), p)
	}
	arg = ret
	return
}
