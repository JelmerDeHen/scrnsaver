package scrnsaver

import (
	"testing"
)

func TestGetXScreenSaverInfo(t *testing.T) {
	// info.Idle contains the idle time as time.Duration
	info, err := GetXScreenSaverInfo()
	if err != nil {
		t.Fatalf(`GetXScreenSaverInfo() = %+v, err=%q`, info, err)
	}
}
