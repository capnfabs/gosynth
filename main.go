package main

import (
	"time"

	"github.com/gordonklaus/portaudio"
)

const sampleRate = 44100

var signal = float32(-1)

func sawTooth(out []float32) {
	for i := 0; i < len(out); i++ {
		out[i] = signal
		signal += 0.05
		if signal > 1 {
			signal -= 2
		}
	}
}

func main() {
	err := portaudio.Initialize()
	if err != nil {
		panic(err)
	}
	stream, err := portaudio.OpenDefaultStream(0, 1, sampleRate, portaudio.FramesPerBufferUnspecified, sawTooth)
	if err != nil {
		panic(err)
	}
	err = stream.Start()
	if err != nil {
		panic(err)
	}
	time.Sleep(5 * time.Second)
}
