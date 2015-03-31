// https://github.com/krig/go-sox/blob/master/examples/example0.go
package main

// Use this URL to import the go-sox library
import (
	. "github.com/OUCC/prism/logger"

	"github.com/krig/go-sox"
)

func playSound(filename string) {
	if !sox.Init() {
		Log.Error("Failed to initialize SoX")
		return
	}
	// Make sure to call Quit before terminating
	defer sox.Quit()

	// Open the input file
	in := sox.OpenRead(filename)
	if in == nil {
		Log.Error("Failed to open input file")
		return
	}
	// Close the file before exiting
	defer in.Release()

	// Open the output device: Specify the output signal characteristics.
	// Using "alsa" or "pulseaudio" should work for most files on Linux.
	// On other systems, other devices have to be used.
	out := sox.OpenWrite("default", in.Signal(), nil, SOUND_DEVICE)
	if out == nil {
		Log.Error("Failed to open output file")
		return
	}
	// Close the output device before exiting
	defer out.Release()

	// Create an effects chain: Some effects need to know about the
	// input or output encoding so we provide that information here.
	chain := sox.CreateEffectsChain(in.Encoding(), out.Encoding())
	// Make sure to clean up!
	defer chain.Release()

	// The first effect in the effect chain must be something that can
	// source samples; in this case, we use the built-in handler that
	// inputs data from an audio file.
	e := sox.CreateEffect(sox.FindEffect("input"))
	e.Options(in)
	// This becomes the first "effect" in the chain
	chain.Add(e, in.Signal(), in.Signal())
	e.Release()
	// The last effect in the effect chain must be something that only consumes
	// samples; in this case, we use the built-in handler that outputs data.
	e = sox.CreateEffect(sox.FindEffect("output"))
	e.Options(out)
	chain.Add(e, in.Signal(), in.Signal())
	e.Release()

	// Flow samples through the effects processing chain until EOF is reached.
	chain.Flow()
}
