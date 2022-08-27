package scrnsaver

// #cgo CFLAGS: -g -Wall -Wextra -Werror
// #cgo LDFLAGS: -lXss -lX11
// #include "XOpenDisplay.h"
import "C"

func HasXorg() bool {
	if C.HasXorg() == 0 {
		return true
	}
	return false
}
