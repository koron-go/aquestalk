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

type errno int

func (nr errno) Error() string {
	switch nr {
	case 100:
		return "misc error (error code: 100)"
	case 101:
		return "out of memory"
	case 102:
		return "undefined reading symbols (error code: 102)"
	case 103:
		return "negative prosody data"
	case 104:
		return "undefined delimiters"
	case 105:
		return "undefined reading symbols (error code: 105)"
	case 106:
		return "illegal tags"
	case 107:
		return "too long tags"
	case 108:
		return "invalid tag values"
	case 109:
		return "failed to play wave (error code: 109)"
	case 110:
		return "failed to play wave (error code: 110)"
	case 111:
		return "no sound data to play"
	case 200:
		return "too long phonetic string (error code: 200)"
	case 201:
		return "too many reading symbols in a phrase"
	case 202:
		return "too long phonetic string (error code: 202)"
	case 203:
		return "heap memory exhaust"
	case 204:
		return "too long phonetic string (error code: 204)"
	default:
		return fmt.Sprintf("undefined error code: %d", nr)
	}
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
