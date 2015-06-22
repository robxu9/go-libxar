package xar

// #include <stdlib.h>
// #include <xar/xar.h>
import "C"
import "unsafe"

type File struct {
	archive *Archive
	file    C.xar_file_t
}

func (f *File) Verify() bool {
	return C.xar_verify(f.archive.archive, f.file) == 0
}

func (f *File) Size() string {
	cStr := C.xar_get_size(f.archive.archive, f.file)
	defer C.free(unsafe.Pointer(cStr))

	return C.GoString(cStr)
}

func (f *File) Type() string {
	cStr := C.xar_get_type(f.archive.archive, f.file)
	defer C.free(unsafe.Pointer(cStr))

	return C.GoString(cStr)
}

func (f *File) Mode() string {
	cStr := C.xar_get_mode(f.archive.archive, f.file)
	defer C.free(unsafe.Pointer(cStr))

	return C.GoString(cStr)
}

func (f *File) Owner() string {
	cStr := C.xar_get_owner(f.archive.archive, f.file)
	defer C.free(unsafe.Pointer(cStr))

	return C.GoString(cStr)
}

func (f *File) Group() string {
	cStr := C.xar_get_group(f.archive.archive, f.file)
	defer C.free(unsafe.Pointer(cStr))

	return C.GoString(cStr)
}

func (f *File) Mtime() string {
	cStr := C.xar_get_mtime(f.archive.archive, f.file)
	defer C.free(unsafe.Pointer(cStr))

	return C.GoString(cStr)
}

func (f *File) Path() string {
	cStr := C.xar_get_path(f.file)
	defer C.free(unsafe.Pointer(cStr))

	return C.GoString(cStr)
}
