package main

import (
	"time"

	"github.com/capnfabs/gosynth"
	"github.com/gordonklaus/portaudio"
)

const sampleRate = 44100

func sine(out [][]float32) {
}

func main() {
	err := portaudio.Initialize()
	if err != nil {
		panic(err)
	}
	master := gosynth.Mult(
		gosynth.Constant(0.1),
		gosynth.Sine(sampleRate*4),
		gosynth.Sine(sampleRate/220),
	)
	stream, err := portaudio.OpenDefaultStream(0, 2, sampleRate, portaudio.FramesPerBufferUnspecified, master.PortAudioCallback())
	if err != nil {
		panic(err)
	}
	err = stream.Start()
	if err != nil {
		panic(err)
	}
	time.Sleep(3 * time.Second)
}
