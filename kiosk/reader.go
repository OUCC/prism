package main

import (
	. "github.com/OUCC/prism/logger"

	"github.com/gvalkov/golang-evdev/evdev"
)

var (
	reader     *evdev.InputDevice
	readerCode = make(chan string, 0)
)

func setupReader() {
	devices, err := evdev.ListInputDevices()
	if err != nil {
		Log.Fatal(err)
	}

	for _, dev := range devices {
		if dev.Name == CARDREADER_DEVICE {
			reader = dev
			break
		}
	}

	if reader == nil {
		Log.Fatal("card reader not found")
	}

	Log.Debug(reader.String())
}

func readerLoop() {
	ret := make([]rune, 0, 100)

	for {
		ev, err := reader.ReadOne()
		if err != nil {
			Log.Fatal(err.Error())
		}

		if ev.Type == evdev.EV_KEY && ev.Value == 1 {
			switch ev.Code {
			case evdev.KEY_ENTER:
				Log.Debug("Card reader input: %s", string(ret))
				readerCode <- string(ret)
				ret = make([]rune, 0, 100) // reset
			default:
				ret = append(ret, toRune(ev.Code))
			}
		}
	}
}

func toRune(key uint16) rune {
	switch key {
	case evdev.KEY_SPACE:
		return ' '
	case evdev.KEY_0:
		return '0'
	case evdev.KEY_1:
		return '1'
	case evdev.KEY_2:
		return '2'
	case evdev.KEY_3:
		return '3'
	case evdev.KEY_4:
		return '4'
	case evdev.KEY_5:
		return '5'
	case evdev.KEY_6:
		return '6'
	case evdev.KEY_7:
		return '7'
	case evdev.KEY_8:
		return '8'
	case evdev.KEY_9:
		return '9'
	case evdev.KEY_A:
		return 'A'
	case evdev.KEY_B:
		return 'B'
	case evdev.KEY_C:
		return 'C'
	case evdev.KEY_D:
		return 'D'
	case evdev.KEY_E:
		return 'E'
	case evdev.KEY_F:
		return 'F'
	case evdev.KEY_G:
		return 'G'
	case evdev.KEY_H:
		return 'H'
	case evdev.KEY_I:
		return 'I'
	case evdev.KEY_J:
		return 'J'
	case evdev.KEY_K:
		return 'K'
	case evdev.KEY_L:
		return 'L'
	case evdev.KEY_M:
		return 'M'
	case evdev.KEY_N:
		return 'N'
	default:
		return '?'
	}
}
