//go:build !webui_tls

package webui

// #cgo CFLAGS: -DNO_SSL
import "C"
