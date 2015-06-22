package xar

// #include <stdlib.h>
// #include <xar/xar.h>
import "C"
import "unsafe"

func (a *Archive) GetOption(key string) string {
	cKey := C.CString(key)
	defer C.free(unsafe.Pointer(cKey))

	opt := C.xar_opt_get(a.archive, cKey)
	return C.GoString(opt)
}

func (a *Archive) SetOption(key, value string) error {
	cKey := C.CString(key)
	defer C.free(unsafe.Pointer(cKey))

	cValue := C.CString(value)
	defer C.free(unsafe.Pointer(cValue))

	if C.xar_opt_set(a.archive, cKey, cValue) != 0 {
		return ErrNonZero
	}
	return nil
}

func (a *Archive) UnsetOption(key string) error {
	cKey := C.CString(key)
	defer C.free(unsafe.Pointer(cKey))

	if C.xar_opt_unset(a.archive, cKey) != 0 {
		return ErrNonZero
	}
	return nil
}
