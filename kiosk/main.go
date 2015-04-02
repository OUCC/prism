package main

import (
	. "github.com/OUCC/prism/logger"

	"gopkg.in/qml.v1"

	//"os/exec"
	"time"
)

// 部室滞在者のリスト
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
	win         *qml.Window
	readerModal qml.Object
	felicaModal qml.Object

	occupants = &Occupants{}
)

func main() {
	if CARDREADER_ENABLED {
		setupReader()
		go readerLoop()
	}
	if PASORI_ENABLED {
		setupPaSoRi()
		go pasoriLoop()
	}
	go waitAndPost()

	if err := qml.Run(run); err != nil {
		Log.Fatal(err)
	}
}

func run() error {
	engine := qml.NewEngine()
	engine.Context().SetVar("occupants", occupants)
	component, err := engine.LoadFile("qml/main.qml")
	if err != nil {
		return err
	}
	win = component.CreateWindow(nil)
	win.Show()

	readerModal = win.ObjectByName("readerModal")
	felicaModal = win.ObjectByName("felicaModal")
	win.Wait()
	return nil
}

func waitAndPost() {
	//	var s string
	for {
		select {
		case id := <-readerCode:
			if len(id) == 0 {
				continue
			}

			id = id[12:20]
			Log.Debug("Student ID: %s", id)

			readerModal.Call("showReaderPosting")

			info, handleName, isFirstLogin, occupantList, err := updateLog(id, "")
			if err != nil {
				readerModal.Call("showReaderError", err.Error())
				continue
			}

			occupants.set(occupantList)
			readerModal.Call("showReaderInfo", info, handleName, isFirstLogin)

		case id := <-pasoriCode:
			if len(id) == 0 {
				pasoriWait <- 0 * time.Second
				continue
			}

			if PASORI_SOUND {
				// play sound

				// SIGTRAP
				//cmd := exec.Command("play", "felica.mp3")
				//if err := cmd.Run(); err != nil {
				//Log.Error(err.Error())
				//}

				playSound("felica.mp3")
			}

			felicaModal.Call("showFeliCaPosting")

			info, handleName, isFirstLogin, occupantList, err := updateLog("", id)
			if err != nil {
				if err == ErrNotRegistered {
					// start registration
					felicaModal.Call("showFeliCaRegistration", id, "waiting")
					select {
					case id2 := <-readerCode:
						if err := registerFeliCa(id2[12:20], id); err != nil {
							felicaModal.Call("showFeliCaRegistration", id, err.Error())
						}
						felicaModal.Call("showFeliCaRegistration", id, "success")

					case <-time.After(30 * time.Second): // do nothing
					}
					pasoriWait <- 0 * time.Second // restart felica scan
				} else {
					felicaModal.Call("showFeliCaError", err.Error())
					pasoriWait <- 10 * time.Second
				}
				continue
			}

			occupants.set(occupantList)
			felicaModal.Call("showFeliCaInfo", info, handleName, isFirstLogin)
			pasoriWait <- 5 * time.Second
		}
	}
}
