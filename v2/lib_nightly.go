//go:build webui_nightly

package webui

// #cgo CFLAGS: -Iwebui/include
// #include "webui.h"
import "C"

// SetProxy sets the web browser proxyServer to use. Need to be called before `Show()`.
func (w Window) SetProxy(name string, proxyServer string) {
	C.webui_set_proxy(C.size_t(w), C.CString(proxyServer))
}
