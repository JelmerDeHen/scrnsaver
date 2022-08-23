# scrnsaver

This package calls XScreenSaverQueryInfo() (X11/extensions/scrnsaver.h) and provides content of XScreenSaverInfo (`man 3 XScreenSaverInfo`) in a Go struct. This information contains the amount of time the user was idle.

```go
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
```

The `XScreenSaverInfo.Idle` field contains the idle time as `time.Duration`.

