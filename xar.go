package xar

// #cgo !vendor LDFLAGS: -lxar
// #cgo vendor CFLAGS: -Ivendor/xar/include
// #cgo vendor LDFLAGS: -Lvendor/xar/lib -lxar
// #include <stdlib.h>
// #include <xar/xar.h>
import "C"

import (
	"errors"
	"unsafe"
)

const (
	readFlag  = 0
	writeFlag = 1
)

var (
	ErrNotOpened = errors.New("xar: couldn't open the archive")
	ErrNil       = errors.New("xar: libxar returned NULL on fail")
	ErrNonZero   = errors.New("xar: libxar returned non-zero code")
)

type Archive struct {
	archive C.xar_t
}

func OpenArchive(file string) (*Archive, error) {
	return open(file, readFlag)
}

func CreateArchive(file string) (*Archive, error) {
	return open(file, writeFlag)
}

func open(file string, mode int) (*Archive, error) {
	cStr := C.CString(file)
	defer C.free(unsafe.Pointer(cStr))

	cArchive := C.xar_open(cStr, C.int32_t(mode))
	if cArchive == nil {
		return nil, ErrNotOpened
	}

	return &Archive{
		archive: cArchive,
	}, nil
}

func (a *Archive) Close() error {
	ret := C.xar_close(a.archive)
	if ret != 0 {
		return ErrNil
	}
	return nil
}

// WriteMetadata is comparable to xar_serialize.
// It writes out the table of contents to an XML file.
func (a *Archive) WriteMetadata(xmlFile string) {
	cXmlFile := C.CString(xmlFile)
	defer C.free(unsafe.Pointer(cXmlFile))

	C.xar_serialize(a.archive, cXmlFile)
}
