package main

import (
	. "github.com/OUCC/prism/logger"

	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"time"
)

func updateLog(memberID string, felicaIDm string) (string, string, bool, []string, error) {
	type post struct {
		Key       string `json:"key"`
		MemberID  string `json:"member_id"`
		FeliCaIDm string `json:"felica_idm"`
		When      int64  `json:"when"`
	}

	b, _ := json.Marshal(post{
		Key:       PRISM_KEY,
		MemberID:  memberID,
		FeliCaIDm: felicaIDm,
		When:      time.Now().Unix(),
	})

	resp, err := http.Post(LOG_POST_URL, "application/json", bytes.NewReader(b))
	if err != nil {
		return "", "", false, nil, err
	}
	Log.Debug(resp.Status)
	if resp.StatusCode != http.StatusCreated {
		return "", "", false, nil, errors.New(resp.Status)
	}

	type response struct {
		When       int64    `json:"when"`
		HandleName string   `json:"handle_name"`
		Event      string   `json:"event"`
		FirstLogin bool     `json:"first_login"`
		Occupants  []string `json:"occupants"`
	}

	data := response{}
	b, _ = ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()

	if err := json.Unmarshal(b, &data); err != nil {
		return "", "", false, nil, err
	}

	return data.Event, data.HandleName, data.FirstLogin, data.Occupants, nil
}

func registerFeliCa(memberID string, felicaIDm string) error {
	type post struct {
		Key       string `json:"key"`
		MemberID  string `json:"member_id"`
		FeliCaIDm string `json:"felica_idm"`
	}

	b, _ := json.Marshal(post{
		Key:       PRISM_KEY,
		MemberID:  memberID,
		FeliCaIDm: felicaIDm,
	})

	resp, err := http.Post(LOG_POST_URL, "application/json", bytes.NewReader(b))
	if err != nil {
		return err
	}
	Log.Debug(resp.Status)
	if resp.StatusCode != http.StatusCreated {
		return errors.New(resp.Status)
	}

	return nil
}
