//go:build webui_log

package webui

/*
#cgo windows,amd64 LDFLAGS: -L../webui/debug -lwebui-2-static -lws2_32
#cgo linux darwin LDFLAGS: -L../webui/debug -lwebui-2-static -lpthread -lm
*/
import "C"
