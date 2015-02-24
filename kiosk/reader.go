package main

import (
	. "github.com/OUCC/prism/logger"

	"github.com/gvalkov/golang-evdev/evdev"
	"gopkg.in/qml.v1"

	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"time"
)

type logData struct {
	Key      string `json:"key"`
	MemberID string `json:"member_id"`
	When     int64  `json:"when"`
}

type EventData struct {
	MemberID   string   `json:"member_id"`
	When       int64    `json:"when"`
	HandleName string   `json:"handle_name"`
	Event      string   `json:"event"`
	FirstLogin bool     `json:"first_login"`
	Occupants  []string `json:"occupants"`
}

type ReaderStatus struct {
	Status string
	Data   EventData
	Error  string
}

func (status *ReaderStatus) changed() {
	qml.Changed(status, &status.Error)
	qml.Changed(status, &status.Data)
	qml.Changed(status, &status.Status)
}

type Occupants struct {
	data []string
	Len  int
}

func (s *Occupants) set(data []string) {
	s.data = data
	s.Len = len(data)
	qml.Changed(s, &s.Len)
}

func (s *Occupants) Get(index int) string {
	return s.data[index]
}

var (
	reader       *evdev.InputDevice
	code         = make(chan string, 0)
	readerStatus = &ReaderStatus{}
	occupants    = &Occupants{}
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

func readDevice() {
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
				code <- string(ret)
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

func waitAndPost() {
	for {
		readerStatus.Status = "waiting"
		readerStatus.changed()

		id := <-code
		if len(id) == 0 {
			continue
		}

		id = id[12:20]
		Log.Debug("Student ID: %s", id)

		readerStatus.Status = "posting"
		readerStatus.changed()

		b, _ := json.Marshal(logData{
			Key:      PRISM_KEY,
			MemberID: id,
			When:     time.Now().Unix(),
		})
		resp, err := http.Post(LOG_POST_URL, "application/json", bytes.NewReader(b))
		if err != nil {
			Log.Error(err.Error())
			readerStatus.Status = "error"
			readerStatus.Error = err.Error()
			readerStatus.changed()
			time.Sleep(5 * time.Second)
			continue
		}
		Log.Debug(resp.Status)
		if resp.StatusCode != http.StatusCreated {
			Log.Error("Error posting log data")
			readerStatus.Status = "error"
			readerStatus.Error = resp.Status
			readerStatus.changed()
			time.Sleep(5 * time.Second)
			continue
		}

		data := EventData{}
		b, _ = ioutil.ReadAll(resp.Body)
		defer resp.Body.Close()

		if err := json.Unmarshal(b, &data); err != nil {
			Log.Error(err.Error())
		}

		Log.Debug("HandleName: %s, Event: %s, FirstLogin: %t, Occupants: %v",
			data.HandleName, data.Event, data.FirstLogin, data.Occupants)

		readerStatus.Status = "posted"
		readerStatus.Data = data
		readerStatus.changed()
		occupants.set(data.Occupants)

		time.Sleep(3 * time.Second)
	}
}
