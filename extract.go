package xar

// #include <stdlib.h>
// #include <xar/xar.h>
import "C"
import "unsafe"

func (f *File) Extract() error {
	if C.xar_extract(f.archive.archive, f.file) != 0 {
		return ErrNonZero
	}

	return nil
}

func (f *File) ExtractToFile(path string) error {
	cPath := C.CString(path)
	defer C.free(unsafe.Pointer(cPath))

	if C.xar_extract_tofile(f.archive.archive, f.file, cPath) != 0 {
		return ErrNonZero
	}

	return nil
}

func (f *File) ExtractToBuffer() ([]byte, error) {
	//func C.GoBytes(unsafe.Pointer, C.int) []byte
	var size C.size_t
	var buffer *C.char

	if C.xar_extract_tobuffersz(f.archive.archive, f.file, &buffer, &size) != 0 {
		return nil, ErrNonZero
	}
	defer C.free(unsafe.Pointer(buffer))

	bte := C.GoBytes(unsafe.Pointer(buffer), C.int(size))
	return bte, nil
}

// TODO: xar_extract_tostream_init
// TODO: xar_extract_tostream
// TODO: xar_extract_tostream_end
