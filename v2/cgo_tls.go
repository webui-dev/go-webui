//go:build webui_tls

package webui

/*
#cgo CFLAGS: -DWEBUI_USE_TLS -DWEBUI_TLS -DNO_SSL_DL -DOPENSSL_API_1_1
#cgo LDFLAGS: -lssl -lcrypto
#cgo windows LDFLAGS: -lBcrypt

#cgo CFLAGS: -Iwebui/include
#include "webui.h"
*/
import "C"

import (
	"errors"
	"unsafe"
)

// SetTLSCertificate sets the SSL/TLS certificate and the private key content,
// both in PEM format. This works only with the `webui-2-secure` library.
// If set to empty, WebUI will generate a self-signed certificate.
func SetTLSCertificate(certificate_pem string, private_key_pem string) (err error) {
	ccertificate_pem := C.CString(certificate_pem)
	cprivate_key_pem := C.CString(private_key_pem)
	defer C.free(unsafe.Pointer(ccertificate_pem))
	defer C.free(unsafe.Pointer(cprivate_key_pem))
	if !C.webui_set_tls_certificate(ccertificate_pem, cprivate_key_pem) {
		err = errors.New("error: failed to set TLS certificate")
	}
	return
}
