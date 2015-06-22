package xar

// #include <stdlib.h>
// #include <xar/xar.h>
import "C"
import "unsafe"

func (f *File) SetProperty(key, value string) error {
	cKey := C.CString(key)
	defer C.free(unsafe.Pointer(cKey))

	cValue := C.CString(value)
	defer C.free(unsafe.Pointer(cValue))

	if C.xar_prop_set(f.file, cKey, cValue) != 0 {
		return ErrNonZero
	}
	return nil
}

func (f *File) CreateProperty(key, value string) error {
	cKey := C.CString(key)
	defer C.free(unsafe.Pointer(cKey))

	cValue := C.CString(value)
	defer C.free(unsafe.Pointer(cValue))

	if C.xar_prop_create(f.file, cKey, cValue) != 0 {
		return ErrNonZero
	}
	return nil
}

func (f *File) GetProperty(key string) (string, error) {
	cKey := C.CString(key)
	defer C.free(unsafe.Pointer(cKey))

	var buffer *C.char

	if C.xar_prop_get(f.file, cKey, &buffer) != 0 {
		return "", ErrNonZero
	}

	return C.GoString(buffer), nil
}

func (f *File) ListProperties() []string {
	iter := C.xar_iter_new()
	if iter == nil {
		panic(ErrNil) // memory allocation failed
	}
	defer C.xar_iter_free(iter)

	list := []string{}
	next := C.xar_prop_first(f.file, iter)
	for next != nil {
		list = append(list, C.GoString(next))
		next = C.xar_prop_next(iter)
	}

	return list
}

func (f *File) RemoveProperty(key string) {
	cKey := C.CString(key)
	defer C.free(unsafe.Pointer(cKey))

	C.xar_prop_unset(f.file, cKey)
}

func (f *File) GetPropertyAttr(prop, attr string) string {
	cProp := C.CString(prop)
	defer C.free(unsafe.Pointer(cProp))

	cAttr := C.CString(attr)
	defer C.free(unsafe.Pointer(cAttr))

	result := C.xar_attr_get(f.file, cProp, cAttr)
	return C.GoString(result)
}

func (f *File) SetPropertyAttr(prop, attr, value string) error {
	cProp := C.CString(prop)
	defer C.free(unsafe.Pointer(cProp))

	cAttr := C.CString(attr)
	defer C.free(unsafe.Pointer(cAttr))

	cValue := C.CString(value)
	defer C.free(unsafe.Pointer(cValue))

	if C.xar_attr_set(f.file, cProp, cAttr, cValue) != 0 {
		return ErrNonZero
	}

	return nil
}

func (f *File) ListPropertyAttr(prop string) []string {
	cProp := C.CString(prop)
	defer C.free(unsafe.Pointer(cProp))

	iter := C.xar_iter_new()
	if iter == nil {
		panic(ErrNil) // memory allocation failed
	}
	defer C.xar_iter_free(iter)

	list := []string{}
	next := C.xar_attr_first(f.file, cProp, iter)
	for next != nil {
		list = append(list, C.GoString(next))
		next = C.xar_attr_next(iter)
	}

	return list
}
