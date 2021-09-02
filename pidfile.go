package pidfile

// #include <stdlib.h>
// #include <bsd/libutil.h>
// #cgo LDFLAGS: -lbsd
import "C"

import (
	"unsafe"
)

type PIDFile C.struct_pidfh

func Open(name string) (*PIDFile, error) {
	name_cstr := C.CString(name)
	p, e := C.pidfile_open(name_cstr, 0600, nil)
	C.free(unsafe.Pointer(name_cstr))
	return (*PIDFile)(p), e
}

func (p *PIDFile) Write() (error) {
	_, e := C.pidfile_write((*C.struct_pidfh)(p))
	return e
}

func (p *PIDFile) Close() (error) {
	_, e := C.pidfile_close((*C.struct_pidfh)(p))
	return e
}

func (p *PIDFile) Remove() (error) {
	_, e := C.pidfile_remove((*C.struct_pidfh)(p))
	return e
}
