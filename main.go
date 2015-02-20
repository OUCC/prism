package main

import (
	"gopkg.in/qml.v1"
)

var (
	win *qml.Window
)

func main() {
	if err := qml.Run(run); err != nil {
		Log.Fatal(err)
	}
}

func run() error {
	engine := qml.NewEngine()
	engine.Context().SetVar("readerStatus", readerStatus)
	engine.Context().SetVar("occupants", occupants)
	component, err := engine.LoadFile("qml/main.qml")
	if err != nil {
		return err
	}
	win = component.CreateWindow(nil)
	win.Show()

	go readAndPost()
	go captureAndPost()

	win.Wait()
	return nil
}
