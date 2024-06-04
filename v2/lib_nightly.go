//go:build webui_nightly

package webui

// #cgo CFLAGS: -Iwebui/include
// #include "webui.h"
import "C"
import "unsafe"

// SetProxy sets the web browser proxyServer to use. Need to be called before `Show()`.
func (w Window) SetProxy(name string, proxyServer string) {
	cproxyServer := C.CString(proxyServer)
	defer C.free(unsafe.Pointer(cproxyServer))
	C.webui_set_proxy(C.size_t(w), cproxyServer)
}
