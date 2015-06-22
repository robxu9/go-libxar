package xar

// #include <stdlib.h>
// #include <xar/xar.h>
import "C"
import "unsafe"

//export golibxarSigCallback
func golibxarSigCallback(
	sig C.xar_signature_t,
	context unsafe.Pointer,
	data *C.uint8_t,
	length C.uint32_t,
	signedData **C.uint8_t,
	signedLength *C.uint32_t) C.int32_t {

	callback := (*sigCallback)(context)
	inBte := C.GoBytes(unsafe.Pointer(data), C.int(length))

	signed, err := callback.callback(callback.sig, inBte)
	if err != nil {
		return -1
	}

	outBte := C.malloc(C.size_t(len(signed)))

	for i := 0; i < len(signed); i++ {
		element := (*C.char)(unsafe.Pointer(uintptr(outBte) + uintptr(i)))
		*element = C.char(signed[i])
	}

	*signedData = (*C.uint8_t)(unsafe.Pointer(outBte))
	*signedLength = C.uint32_t(len(signed))

	return 0
}

// SignatureCallback defines the signature mechanism for signing data.
// It takes in unsigned data then returns the data, signed.
type SignatureCallback func(sig *Signature, data []byte) (signed []byte, err error)
