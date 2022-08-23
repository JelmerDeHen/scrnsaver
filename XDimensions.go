package scrnsaver

// #cgo CFLAGS: -g -Wall -Wextra -Werror
// #cgo LDFLAGS: -lXss -lX11
// #include "XDimensions.h"
import "C"

import (
	"fmt"
)

func DisplayWidth() int {
	width := C.GetDisplayWidth()
	return int(width)
}

func DisplayHeight() int {
	height := C.GetDisplayHeight()
	return int(height)
}

func GetResolution() string {
	return fmt.Sprintf("%dx%d", DisplayWidth(), DisplayHeight())
}
