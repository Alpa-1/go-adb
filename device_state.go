package adb

import "github.com/Alpa-1/go-adb/internal/errors"

// DeviceState represents one of the 3 possible states adb will report devices.
// A device can be communicated with when it's in StateOnline.
// A USB device will make the following state transitions:
// 	Plugged in: StateDisconnected->StateOffline->StateOnline
// 	Unplugged:  StateOnline->StateDisconnected
//go:generate stringer -type=DeviceState
type DeviceState int8

// https://android.googlesource.com/platform/packages/modules/adb/+/refs/heads/master/adb.cpp - line 115
const (
	StateInvalid DeviceState = iota
	StateAuthorizing
	StateUnauthorized
	StateDisconnected
	StateOffline
	StateOnline
	StateRecovery
	StateBootloader
	StateHost
	StateRescue
	StateSideload
	StateConnecting
	StateUnknown
)

var deviceStateStrings = map[string]DeviceState{
	"":             StateDisconnected,
	"offline":      StateOffline,
	"device":       StateOnline,
	"unauthorized": StateUnauthorized,
	"authorizing":  StateAuthorizing,
	"bootloader":   StateBootloader,
	"host":         StateHost,
	"rescue":       StateRescue,
	"sideload":     StateSideload,
	"connecting":   StateConnecting,
	"recovery":     StateRecovery,
	"unknown":      StateUnknown,
}

func parseDeviceState(str string) (DeviceState, error) {
	state, ok := deviceStateStrings[str]
	if !ok {
		return StateInvalid, errors.Errorf(errors.ParseError, "invalid device state: %q", state)
	}
	return state, nil
}
