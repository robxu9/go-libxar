package xar

// #include <stdlib.h>
// #include <xar/xar.h>
import "C"

import (
	"unsafe"

	"golang.org/x/sys/unix"
)

func (a *Archive) Add(path string) (*File, error) {
	cStr := C.CString(path)
	defer C.free(unsafe.Pointer(cStr))

	handle := C.xar_add(a.archive, cStr)

	if handle == nil {
		return nil, ErrNil
	}

	return &File{
		archive: a,
		file:    handle,
	}, nil
}

func (a *Archive) AddFromBuffer(parent *File, name string, buffer []byte) (*File, error) {
	cName := C.CString(name)
	defer C.free(unsafe.Pointer(cName))

	var parentHandle C.xar_file_t
	if parent != nil {
		parentHandle = parent.file
	}

	handle := C.xar_add_frombuffer(a.archive, parentHandle, cName, (*C.char)(unsafe.Pointer(&buffer[0])), C.size_t(len(buffer)))

	if handle == nil {
		return nil, ErrNil
	}

	return &File{
		archive: a,
		file:    handle,
	}, nil
}

func (a *Archive) AddFolder(parent *File, name string, stat *unix.Stat_t) (*File, error) {
	cName := C.CString(name)
	defer C.free(unsafe.Pointer(cName))

	var parentHandle C.xar_file_t
	if parent != nil {
		parentHandle = parent.file
	}

	// translate our stat into c stat
	cStat := translateStatStruct(stat)

	handle := C.xar_add_folder(a.archive, parentHandle, cName, cStat)

	if handle == nil {
		return nil, ErrNil
	}

	return &File{
		archive: a,
		file:    handle,
	}, nil
}

func (a *Archive) AddFromPath(parent *File, name string, realpath string) (*File, error) {
	cName := C.CString(name)
	defer C.free(unsafe.Pointer(cName))

	cRealPath := C.CString(realpath)
	defer C.free(unsafe.Pointer(cRealPath))

	var parentHandle C.xar_file_t
	if parent != nil {
		parentHandle = parent.file
	}

	handle := C.xar_add_frompath(a.archive, parentHandle, cName, cRealPath)

	if handle == nil {
		return nil, ErrNil
	}

	return &File{
		archive: a,
		file:    handle,
	}, nil
}

func (a *Archive) AddFromArchive(parent *File, name string, source *File) (*File, error) {
	cName := C.CString(name)
	defer C.free(unsafe.Pointer(cName))

	var parentHandle C.xar_file_t
	if parent != nil {
		parentHandle = parent.file
	}

	handle := C.xar_add_from_archive(a.archive, parentHandle, cName, source.archive.archive, source.file)

	if handle == nil {
		return nil, ErrNil
	}

	return &File{
		archive: a,
		file:    handle,
	}, nil
}
