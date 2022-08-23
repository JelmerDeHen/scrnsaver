package scrnsaver

// #cgo CFLAGS: -g -Wall -Wextra -Werror
// #cgo LDFLAGS: -lXss -lX11
// #include "XScreenSaverQueryInfo.h"
import "C"

import (
	"fmt"
	"time"
)

type XID uint64

type ScreenSaverKind int

const (
	Blanked ScreenSaverKind = iota
	Internal
	External
)

type ScreenSaverState int

const (
	Off ScreenSaverState = iota
	On
	Disabled
)

type XScreenSaverInfo struct {
	// screen saver info
	Window XID
	// ScreenSaver{Off,On,Disabled}
	State ScreenSaverState
	// ScreenSaver{Blanked,Internal,External}
	Kind ScreenSaverKind
	// Milliseconds
	Til_or_since time.Duration // uint64
	// Milliseconds
	Idle time.Duration //uint64
	// Events
	EventMask uint64
}

func GetXScreenSaverInfo() (xscreensaverinfo *XScreenSaverInfo, err error) {
	// Handled by gc
	info := C.struct_XScreenSaverInfo_t{}

	errno := C.GetXScreenSaverQueryInfo(&info)
	if errno != 0 {
		return nil, fmt.Errorf("GetXScreenSaverQueryInfo failed errno=%d\n", errno)
	}

	// Copy over result
	xscreensaverinfo = &XScreenSaverInfo{
		Window:    XID(uint64(info.window)),
		State:     ScreenSaverState(int(info.state)),
		Kind:      ScreenSaverKind(int(info.kind)),
		EventMask: uint64(info.eventMask),
	}

	// Convert ms to time.Duration
	xscreensaverinfo.Til_or_since, err = time.ParseDuration(fmt.Sprintf("%dms", uint64(info.til_or_since)))
	xscreensaverinfo.Idle, err = time.ParseDuration(fmt.Sprintf("%dms", uint64(info.idle)))

	return xscreensaverinfo, nil
}
