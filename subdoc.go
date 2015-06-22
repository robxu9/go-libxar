package xar

// #include <stdlib.h>
// #include <xar/xar.h>
import "C"
import "unsafe"

type Subdoc struct {
	subdoc  C.xar_subdoc_t
	archive *Archive
}

func (a *Archive) NewSubdoc(name string) *Subdoc {
	cName := C.CString(name)
	defer C.free(unsafe.Pointer(cName))

	subdoc := C.xar_subdoc_new(a.archive, cName)
	if subdoc == nil {
		panic(ErrNil)
	}

	return &Subdoc{
		subdoc:  subdoc,
		archive: a,
	}
}

func (a *Archive) ListSubdocs() []*Subdoc {
	iter := C.xar_iter_new()
	if iter == nil {
		panic(ErrNil) // memory allocation failed
	}
	defer C.xar_iter_free(iter)

	var list []*Subdoc
	next := C.xar_subdoc_first(a.archive)
	for next != nil {
		list = append(list, &Subdoc{
			subdoc:  next,
			archive: a,
		})
		next = C.xar_subdoc_next(next)
	}

	return list
}

func (s *Subdoc) SetProperty(key, value string) error {
	cKey := C.CString(key)
	defer C.free(unsafe.Pointer(cKey))

	cValue := C.CString(value)
	defer C.free(unsafe.Pointer(cValue))

	if C.xar_subdoc_prop_set(s.subdoc, cKey, cValue) != 0 {
		return ErrNonZero
	}
	return nil
}

func (s *Subdoc) GetProperty(key string) (string, error) {
	cKey := C.CString(key)
	defer C.free(unsafe.Pointer(cKey))

	var buffer *C.char

	if C.xar_subdoc_prop_get(s.subdoc, cKey, &buffer) != 0 {
		return "", ErrNonZero
	}

	return C.GoString(buffer), nil
}

func (s *Subdoc) SetPropertyAttr(prop, attr, value string) error {
	cProp := C.CString(prop)
	defer C.free(unsafe.Pointer(cProp))

	cAttr := C.CString(attr)
	defer C.free(unsafe.Pointer(cAttr))

	cValue := C.CString(value)
	defer C.free(unsafe.Pointer(cValue))

	if C.xar_subdoc_attr_set(s.subdoc, cProp, cAttr, cValue) != 0 {
		return ErrNonZero
	}

	return nil
}

func (s *Subdoc) GetPropertyAttr(prop, attr string) string {
	cProp := C.CString(prop)
	defer C.free(unsafe.Pointer(cProp))

	cAttr := C.CString(attr)
	defer C.free(unsafe.Pointer(cAttr))

	result := C.xar_subdoc_attr_get(s.subdoc, cProp, cAttr)

	return C.GoString(result)
}

func (s *Subdoc) Name() string {
	return C.GoString(C.xar_subdoc_name(s.subdoc))
}

// Export gets the byte output of the subdocument in XML format.
// Synonymous to xar_subdoc_copyout in the libxar API.
func (s *Subdoc) Export() ([]byte, error) {
	var buffer *C.uchar
	var size C.uint

	if C.xar_subdoc_copyout(s.subdoc, &buffer, &size) != 0 {
		return nil, ErrNonZero
	}
	defer C.free(unsafe.Pointer(buffer))

	bte := C.GoBytes(unsafe.Pointer(buffer), C.int(size))
	return bte, nil
}

// Import fills the subdoc with the information from the byte output.
// Synonymous to xar_subdoc_copyin in the libxar API.
func (s *Subdoc) Import(b []byte) error {
	if C.xar_subdoc_copyin(s.subdoc, (*C.uchar)(unsafe.Pointer(&b[0])), C.uint(len(b))) != 0 {
		return ErrNonZero
	}

	return nil
}

// Remove removes the subdoc from the archive. Once it is removed,
// it becomes unusable. Do not touch or errors will occur.
func (s *Subdoc) Remove() {
	C.xar_subdoc_remove(s.subdoc)
}
