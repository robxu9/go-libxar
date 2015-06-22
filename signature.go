package xar

// #include <stdlib.h>
// #include <xar/xar.h>
//
// int32_t golibxarSigCallback(xar_signature_t sig, void *context, uint8_t *data, uint32_t length, uint8_t **signed_data, uint32_t *signed_len);
//
// int32_t golibxarSigCallbackCgo(xar_signature_t sig, void *context, uint8_t *data, uint32_t length, uint8_t **signed_data, uint32_t *signed_len) {
//     return golibxarSigCallback(sig, context, data, length, signed_data, signed_len);
// }
import "C"

import (
	"crypto/x509"
	"unsafe"
)

type Signature struct {
	signature C.xar_signature_t
	archive   *Archive
}

type sigCallback struct {
	sig      *Signature
	callback SignatureCallback
}

func (a *Archive) NewSignature(sigType string, length int32, callback SignatureCallback) (*Signature, error) {
	sig := &Signature{
		archive: a,
	} // set signature later
	cb := &sigCallback{
		sig:      sig,
		callback: callback,
	}

	cSigType := C.CString(sigType)
	defer C.free(unsafe.Pointer(cSigType))

	xarSig := C.xar_signature_new(a.archive, cSigType, C.int32_t(length), (C.xar_signer_callback)(unsafe.Pointer(C.golibxarSigCallbackCgo)), unsafe.Pointer(cb))
	if xarSig == nil {
		return nil, ErrNil
	}

	sig.signature = xarSig
	return sig, nil
}

func (a *Archive) ListSignatures() []*Signature {
	list := []*Signature{}
	next := C.xar_signature_first(a.archive)
	for next != nil {
		list = append(list, &Signature{
			signature: next,
			archive:   a,
		})
		next = C.xar_signature_next(next)
	}

	return list
}

func (s *Signature) Type() string {
	return C.GoString(C.xar_signature_type(s.signature))
}

func (s *Signature) AddX509Certificate(cert *x509.Certificate) error {
	raw := cert.Raw

	if C.xar_signature_add_x509certificate(s.signature, (*C.uint8_t)(unsafe.Pointer(&raw[0])), C.uint32_t(len(raw))) != 0 {
		return ErrNonZero
	}

	return nil
}

func (s *Signature) NumX509Certifiates() int {
	return int(C.xar_signature_get_x509certificate_count(s.signature))
}

func (s *Signature) GetX509Certificate(index int) (*x509.Certificate, error) {

	var dataptr *C.uint8_t
	var length C.uint32_t

	if C.xar_signature_get_x509certificate_data(s.signature, C.int32_t(index), &dataptr, &length) != 0 {
		return nil, ErrNonZero
	}
	defer C.free(unsafe.Pointer(dataptr))

	bte := C.GoBytes(unsafe.Pointer(dataptr), C.int(length))
	return x509.ParseCertificate(bte)
}

// TODO: xar_signature_copy_signed_data
