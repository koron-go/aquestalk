// +build !windows
// +build !linux

package aquestalk

import "errors"

// Synthe synthesizes voice with an engine a.k.a. "Yukkuri".
func Synthe(koe string, speed int32) ([]byte, error) {
	return nil, errors.New("not supported platform")
}
