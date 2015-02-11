package main

import (
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

func readAndPost() {
	for {
		time.Sleep(30 * time.Second)

		b, _ := json.Marshal(logData{
			Key:      PRISM_KEY,
			MemberID: DEBUG_MEMBER_ID,
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
