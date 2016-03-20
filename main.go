package main

import (
	"time"

	"github.com/gordonklaus/portaudio"
)

const sampleRate = 44100

var left = float32(-1)
var right = float32(-1)

func sawTooth(out [][]float32) {
	for i := 0; i < len(out[0]); i++ {
		out[0][i] = left
		out[1][i] = right
		left += 0.025
		right += 0.03
		if left > 1 {
			left -= 2
		}
		if right > 1 {
			right -= 2
		}
	}
}

func main() {
	err := portaudio.Initialize()
	if err != nil {
		panic(err)
	}
	stream, err := portaudio.OpenDefaultStream(0, 2, sampleRate, portaudio.FramesPerBufferUnspecified, sawTooth)
	if err != nil {
		panic(err)
	}
	err = stream.Start()
	if err != nil {
		panic(err)
	}
	time.Sleep(3 * time.Second)
}
