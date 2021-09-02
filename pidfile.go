package pidfile

// #include <bsd/libutil.h>
// #cgo LDFLAGS: -lbsd
import "C"

type PIDFile C.struct_pidfh

func Open() (*PIDFile, error) {
	p, e := C.pidfile_open(nil, 0600, nil)
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
