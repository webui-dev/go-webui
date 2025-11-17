package webui

/*
#cgo CFLAGS: -Iwebui/include/
#cgo CFLAGS: -DNDEBUG -DNO_CACHING -DNO_CGI -DUSE_WEBSOCKET -DCGO
#cgo darwin CFLAGS: -x objective-c
#cgo windows CFLAGS: -DMUST_IMPLEMENT_CLOCK_GETTIME

#cgo darwin LDFLAGS: -framework WebKit -framework Cocoa
#cgo windows LDFLAGS: -lWs2_32 -lOle32 -lAdvapi32 -lShell32 -lUser32 -lOleAut32 -luuid

#include "webui/src/civetweb/civetweb.c"
// Prevent conflict with definition in `webui.c`.
#undef MG_BUF_LEN
// Allow webui.c to use standard vsnprintf
#undef vsnprintf

#ifdef __APPLE__
	#include "webui/src/webview/wkwebview.m"
#endif

#include "webui/src/webui.c"
#include "webui.h"
*/
import "C"
