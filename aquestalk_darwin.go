package aquestalk

import (
	"unsafe"
)

/*
#cgo CFLAGS: -F/Library/Frameworks
#cgo LDFLAGS: -F/Library/Frameworks
#cgo LDFLAGS: -framework AquesTalk
#import <AquesTalk/AquesTalk.h>
*/
import "C"

// Synthe synthesizes voice with an engine a.k.a. "Yukkuri".
func Synthe(koe string, speed int32) ([]byte, error) {
	param := C.AQTK_VOICE{
		bas: 0,
		spd: C.int(speed),
		vol: 100,
		pit: 100,
		acc: 100,
		lmd: 100,
		fsc: 100,
	}
	var size C.int
	r := C.AquesTalk_Synthe_Utf8(&param, C.CString(koe), &size)
	if r == nil {
		return nil, errno(size)
	}
	pWav := C.GoBytes(unsafe.Pointer(r), size)
	p := make([]byte, len(pWav))
	copy(p, pWav)
	C.AquesTalk_FreeWave(r)
	return p, nil
}
