package xar

// #cgo !vendor LDFLAGS: -lxar
// #cgo vendor CFLAGS: -Ivendor/xar/include
// #cgo vendor LDFLAGS: -Lvendor/xar/lib -lxar
// #include <xar/xar.h>
import "C"
