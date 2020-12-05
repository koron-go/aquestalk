package aquestalk

import (
	"errors"
	"fmt"
	"unsafe"
)

/*
#cgo LDFLAGS: -ldl
#include <dlfcn.h>

unsigned char* callSynthe(void *p, const char *koe, int speed, int *size) {
	unsigned char* (*synthe)(const char*, int, int*) = p;
	return synthe(koe, speed, size);
}

void callFree(void *p, unsigned char *wav) {
	void (*freewav)(unsigned char *) = p;
	freewav(wav);
}
*/
import "C"

// DLLName declares name of shared object. You can change this only before call Synthe().
var DLLName = "./libAquesTalk.so"

// Synthe synthesizes voice with an engine a.k.a. "Yukkuri".
func Synthe(koe string, speed int32) ([]byte, error) {
	h := C.dlopen(C.CString(DLLName), C.RTLD_LAZY)
	if h == nil {
		return nil, fmt.Errorf("failed to load %s", DLLName)
	}
	pSynthe := C.dlsym(h, C.CString("AquesTalk_Synthe_Utf8"))
	if pSynthe == nil {
		return nil, errors.New("not found symbol: AquesTalk_Synthe_Utf8")
	}
	pFree := C.dlsym(h, C.CString("AquesTalk_FreeWave"))
	if pFree == nil {
		return nil, errors.New("not found symbol: AquesTalk_FreeWave")
	}

	var size C.int
	r := C.callSynthe(pSynthe, C.CString(koe), C.int(speed), &size)
	if r == nil {
		return nil, errno(size)
	}
	pWav := C.GoBytes(unsafe.Pointer(r), size)
	p := make([]byte, len(pWav))
	copy(p, pWav)
	C.callFree(pFree, r)
	return p, nil
}
