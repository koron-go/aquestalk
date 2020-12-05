package aquestalk

import (
	"fmt"
	"reflect"
	"sync"
	"syscall"
	"unsafe"
)

// Synthe synthesizes voice with an engine a.k.a. "Yukkuri".
func Synthe(koe string, speed int32) ([]byte, error) {
	err := dllInit()
	if err != nil {
		return nil, err
	}
	pKoe := append([]byte(koe), 0)
	var size int32
	r1, _, err := dllSynthe.Call(uintptr(unsafe.Pointer(&pKoe[0])), uintptr(speed), uintptr(unsafe.Pointer(&size)))
	if nr, ok := err.(syscall.Errno); !ok || nr != 0 {
		return nil, nr
	}
	if r1 == 0 {
		return nil, errno(size)
	}
	pWav := toBytes(r1, size)
	p := make([]byte, len(pWav))
	copy(p, pWav)
	dllFree.Call(r1)
	return p, nil
}

func toBytes(p uintptr, size int32) []byte {
	return *(*[]byte)(unsafe.Pointer(&reflect.SliceHeader{
		Data: p,
		Len:  int(size),
		Cap:  int(size),
	}))
}

// DLLName declares name of DLL. You can change this only before call Synthe().
var DLLName = "AquesTalk.dll"

var (
	dllOnce   sync.Once
	dllPtr    *syscall.DLL
	dllSynthe *syscall.Proc
	dllFree   *syscall.Proc
	dllErr    error
)

func dllInit() error {
	dllOnce.Do(func() {
		if dllErr != nil || dllPtr != nil {
			return
		}
		dll, err := syscall.LoadDLL(DLLName)
		if err != nil {
			dllErr = err
			return
		}
		pSynthe, err := dll.FindProc("AquesTalk_Synthe_Utf8")
		if err != nil {
			dll.Release()
			dllErr = err
			return
		}
		pFreeWave, err := dll.FindProc("AquesTalk_FreeWave")
		if err != nil {
			dll.Release()
			dllErr = err
			return
		}
		dllPtr, dllSynthe, dllFree, dllErr = dll, pSynthe, pFreeWave, nil
	})
	return dllErr
}
