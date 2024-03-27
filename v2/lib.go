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
#cgo CFLAGS: -Iwebui/include
#include "webui.h"

typedef bool bool_t;

// for webui_bind
extern void goWebuiEventHandler(webui_event_t* e);
static size_t go_webui_bind(size_t win, const char* element) {
	return webui_bind(win, element, goWebuiEventHandler);
}
*/
import "C"

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"reflect"
	"unsafe"
)

type Window uint

type Browser uint8

const (
	NoBrowser Browser = iota
	AnyBrowser
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
	MouseClick
	Navigation
	Callback
)

type Event struct {
	Window      Window
	EventType   EventType
	Element     string
	eventNumber uint
	bindId      uint
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

type returnError struct {
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

// Private function that receives and handles webui events as go events.
//
//export goWebuiEventHandler
func goWebuiEventHandler(e *C.webui_event_t) {
	// Create Go event from C event.
	goEvent := Event{
		Window:      Window(e.window),
		EventType:   EventType(e.event_type),
		Element:     C.GoString(e.element),
		eventNumber: uint(e.event_number),
		bindId:      uint(e.bind_id),
	}
	// Call user callback function.
	result := funcList[goEvent.Window][goEvent.bindId](goEvent)
	if result == nil {
		return
	}
	response, err := json.Marshal(result)
	if err != nil {
		log.Println("Failed encoding JS result into JSON", err)
	}
	C.webui_interface_set_response(e.window, e.event_number, C.CString(string(response)))
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
}

// SetRootFolder sets the web-server root folder path for the window.
func (w Window) SetRootFolder(path string) {
	C.webui_set_root_folder(C.size_t(w), C.CString(path))
}

// SetRootFolder sets the web-server root folder path for all windows.
// Deprecated: use SetDefaultRootFolder instead
func SetRootFolder(path string) {
	C.webui_set_default_root_folder(C.CString(path))
}

// SetDefaultRootFolder sets the web-server root folder path for all windows.
func SetDefaultRootFolder(path string) {
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

// SetProxy sets the web browser proxyServer to use. Need to be called before `Show()`.
func (w Window) SetProxy(name string, proxyServer string) {
	C.webui_set_proxy(C.size_t(w), C.CString(proxyServer))
}

// GetUrl returns the full current URL
func (w Window) GetUrl() string {
	return C.GoString(C.webui_get_url(C.size_t(w)))
}

// SetPublic allows a specific window address to be accessible from a public network
func (w Window) SetPublic(name string, status bool) {
	C.webui_set_public(C.size_t(w), C.bool_t(status))
}

// Navigate navigates to a specific URL
func (w Window) Navigate(url string) {
	C.webui_navigate(C.size_t(w), C.CString(url))
}

// Clean free all memory resources. Should be called only at the end.
func Clean() {
	C.webui_clean()
}

// DeleteAllProfiles deletes all local web-browser profiles folder. It should called at the end.
func DeleteAllProfiles() {
	C.webui_delete_all_profiles()
}

// DeleteProfile deletes a specific window web-browser local folder profile.
func (w Window) DeleteProfile() {
	C.webui_delete_profile(C.size_t(w))
}

// GetParentProcessID returns the ID of the parent process (The web browser may re-create another new process).
func (w Window) GetParentProcessID() uint64 {
	return uint64(C.webui_get_parent_process_id(C.size_t(w)))
}

// GetParentProcessID returns the ID of the last child process.
func (w Window) GetChildProcessID() uint64 {
	return uint64(C.webui_get_child_process_id(C.size_t(w)))
}

// SetPort sets a custom web-server network port to be used by WebUI.
func (w Window) SetPort(port uint) bool {
	return bool(C.webui_set_port(C.size_t(w), C.size_t(port)))
}

// -- SSL/TLS -------------------------

// Run sets the SSL/TLS certificate and the private key content, both in PEM
// format. This works only with `webui-2-secure` library. If set empty WebUI
// will generate a self-signed certificate.
func SetTLSCertificate(certificate_pem string, private_key_pem string) {
	C.webui_set_tls_certificate(C.CString(certificate_pem), C.CString(private_key_pem))
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

func (e *returnError) Error() string {
	return fmt.Sprintf("Failed returning the response of type `%s` for `%s`.", e.typ, e.element)
}

func (e Event) cStruct() *C.webui_event_t {
	return &C.webui_event_t{
		window:       C.size_t(e.Window),
		event_type:   C.size_t(e.EventType),
		element:      C.CString(e.Element),
		event_number: C.size_t(e.eventNumber),
		bind_id:      C.size_t(e.bindId),
	}
}

// GetSize returns the size of the first JavaScript argument.
func (e Event) GetSize() uint {
	return uint(C.webui_get_size(e.cStruct()))
}

// GetSize returns the size of the JavaScript at the specified index.
func (e Event) GetSizeAt(idx uint) uint {
	return uint(C.webui_get_size_at(e.cStruct(), C.size_t(idx)))
}

// GetArg parses the JavaScript argument into a Go data type.
func GetArg[T any](e Event) (arg T, err error) {
	cEvent := e.cStruct()
	if uint(C.webui_get_size(cEvent)) == 0 {
		err = &noArgError{e.Element}
	}
	var ret T
	switch p := any(&ret).(type) {
	case *string:
		*p = C.GoString(C.webui_get_string(cEvent))
	case *int:
		*p = int(C.webui_get_int(cEvent))
	case *bool:
		*p = bool(C.webui_get_bool(cEvent))
	default:
		if jsonErr := json.Unmarshal([]byte(C.GoString(C.webui_get_string(cEvent))), p); err != nil {
			err = &getArgError{jsonErr, e.Element, reflect.TypeOf(ret).String()}
		}
	}
	arg = ret
	return
}

// GetArgAt parses the JavaScript argument with the specified index into a Go data type.
func GetArgAt[T any](e Event, idx uint) (arg T, err error) {
	cEvent := e.cStruct()
	cIdx := C.size_t(idx)
	if uint(C.webui_get_size_at(cEvent, cIdx)) == 0 {
		err = &noArgError{e.Element}
	}
	var ret T
	switch p := any(&ret).(type) {
	case *string:
		*p = C.GoString(C.webui_get_string_at(cEvent, cIdx))
	case *int:
		*p = int(C.webui_get_int_at(cEvent, cIdx))
	case *bool:
		*p = bool(C.webui_get_bool_at(cEvent, cIdx))
	default:
		if jsonErr := json.Unmarshal([]byte(C.GoString(C.webui_get_string_at(cEvent, cIdx))), p); err != nil {
			err = &getArgError{jsonErr, e.Element, reflect.TypeOf(ret).String()}
		}
	}
	arg = ret
	return
}

// Return return the response to JavaScript.
func Return[T any](e Event, value T) (err error) {
	cEvent := e.cStruct()

	switch v := any(&value).(type) {
	case string:
		C.webui_return_string(cEvent, C.CString(v))
	case int:
		C.webui_return_int(cEvent, C.longlong(v))
	case int32:
		C.webui_return_int(cEvent, C.longlong(v))
	case int64:
		C.webui_return_int(cEvent, C.longlong(v))
	case bool:
		C.webui_return_bool(cEvent, C.bool_t(v))
	default:
		err = &returnError{e.Element, reflect.TypeOf(value).String()}
	}
	return
}
