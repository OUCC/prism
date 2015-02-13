package main

import (
	"github.com/gvalkov/golang-evdev/evdev"

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

type eventData struct {
	MemberID   string   `json:"member_id"`
	When       int64    `json:"when"`
	HandleName string   `json:"handle_name"`
	Event      string   `json:"event"`
	FirstLogin bool     `json:"first_login"`
	Occupants  []string `json:"occupants"`
}

var (
	reader *evdev.InputDevice
)

func init() {
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

func readAndPost() {
	code := make(chan string, 0)

	go func() {
		ret := make([]rune, 0, 100)

		for {
			ev, err := reader.ReadOne()
			if err != nil {
				Log.Fatal(err.Error())
				continue
			}

			if ev.Type == evdev.EV_KEY && ev.Value == 1 {
				switch ev.Code {
				case evdev.KEY_ENTER:
					Log.Debug("Card reader input: %s", string(ret))
					code <- string(ret)
					ret = make([]rune, 0, 100)
				case evdev.KEY_SPACE:
					ret = append(ret, ' ')
				case evdev.KEY_0:
					ret = append(ret, '0')
				case evdev.KEY_1:
					ret = append(ret, '1')
				case evdev.KEY_2:
					ret = append(ret, '2')
				case evdev.KEY_3:
					ret = append(ret, '3')
				case evdev.KEY_4:
					ret = append(ret, '4')
				case evdev.KEY_5:
					ret = append(ret, '5')
				case evdev.KEY_6:
					ret = append(ret, '6')
				case evdev.KEY_7:
					ret = append(ret, '7')
				case evdev.KEY_8:
					ret = append(ret, '8')
				case evdev.KEY_9:
					ret = append(ret, '9')
				case evdev.KEY_A:
					ret = append(ret, 'A')
				case evdev.KEY_B:
					ret = append(ret, 'B')
				case evdev.KEY_C:
					ret = append(ret, 'C')
				case evdev.KEY_D:
					ret = append(ret, 'D')
				case evdev.KEY_E:
					ret = append(ret, 'E')
				case evdev.KEY_F:
					ret = append(ret, 'F')
				case evdev.KEY_G:
					ret = append(ret, 'G')
				case evdev.KEY_H:
					ret = append(ret, 'H')
				case evdev.KEY_I:
					ret = append(ret, 'I')
				case evdev.KEY_J:
					ret = append(ret, 'J')
				case evdev.KEY_K:
					ret = append(ret, 'K')
				case evdev.KEY_L:
					ret = append(ret, 'L')
				case evdev.KEY_M:
					ret = append(ret, 'M')
				case evdev.KEY_N:
					ret = append(ret, 'N')
				}
			}
		}
	}()

	for {
		id := <-code
		id = id[9:17]
		Log.Debug("Student ID: %s", id)

		b, _ := json.Marshal(logData{
			Key:      PRISM_KEY,
			MemberID: id,
			When:     time.Now().Unix(),
		})
		resp, err := http.Post(LOG_POST_URL, "application/json", bytes.NewReader(b))
		if err != nil {
			Log.Error(err.Error())
			continue
		}
		Log.Debug(resp.Status)
		if resp.StatusCode != http.StatusCreated {
			Log.Error("Error posting log data")
			continue
		}

		data := eventData{}
		b, _ = ioutil.ReadAll(resp.Body)
		defer resp.Body.Close()

		if err := json.Unmarshal(b, &data); err != nil {
			Log.Error(err.Error())
		}

		Log.Debug("HandleName: %s, Event: %s, FirstLogin: %t, Occupants: %v",
			data.HandleName, data.Event, data.FirstLogin, data.Occupants)
	}
}
