# scrnsaver

Calls XScreenSaverQueryInfo() (X11/extensions/scrnsaver.h) and provides XScreenSaverInfo (`man 3 XScreenSaverInfo`).

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

# Usage

```go
package main

import (
  "github.com/JelmerDeHen/scrnsaver"
  "fmt"
)

func main() {
  info, err := scrnsaver.GetXScreenSaverInfo()
  if err != nil {
      panic(err)
  }

  fmt.Printf("User is idle for %dms\n", info.Idle.Milliseconds())
}
```
