package main

import (
	"fmt"
	"math"

	"github.com/capnfabs/gosynth"
	"github.com/gordonklaus/portaudio"
)

const sampleRate = 44100.0

// These are all in discrete period notation, not frequency.
// - Use [note] * 2 to go down an octave.
// - Use [note] / 2 to go up an octave.
var (
	C  = sampleRate / 261.625565
	Cs = C / math.Exp2(1.0/12)
	Db = Cs
	D  = C / math.Exp2(2.0/12)
	Ds = C / math.Exp2(3.0/12)
	Eb = Ds
	E  = C / math.Exp2(4.0/12)
	F  = C / math.Exp2(5.0/12)
	Fs = C / math.Exp2(6.0/12)
	Gb = Fs
	G  = C / math.Exp2(7.0/12)
	Gs = C / math.Exp2(8.0/12)
	Ab = Gs
	A  = C / math.Exp2(9.0/12)
	As = C / math.Exp2(10.0/12)
	Bb = As
	B  = C / math.Exp2(11.0/12)
)

var Cmajor = []gosynth.Clock{
	clock(C),
	clock(D),
	clock(E),
	clock(F),
	clock(G),
	clock(A),
	clock(B),
	clock(C),
}

var Cbass = []gosynth.Clock{
	clock(C) * 4,
	clock(G) * 4,
	clock(F) * 4,
	clock(E) * 4,
}

var Criff = []gosynth.Clock{
	clock(C),
	clock(C),
	clock(F),
	clock(E),
	clock(E),
	clock(B),
	clock(C / 2),
	clock(B),
}

func sine(out [][]float32) {
}

func clock(val float64) gosynth.Clock {
	return gosynth.Clock(val)
}

var quav = clock(sampleRate / 3)

func main() {
	err := portaudio.Initialize()
	if err != nil {
		panic(err)
	}
	master := gosynth.Mult(
		gosynth.Constant(0.1),
		gosynth.Avg(
			// C chord
			//gosynth.Sawtooth(clock(E)),
			//gosynth.Square(clock(G*2)),
			// End C Chord
			// Step Sequencer
			gosynth.SawtoothWithPeriod(gosynth.StepSequencer(quav*8, Cbass)),
			gosynth.SawtoothWithPeriod(gosynth.StepSequencer(quav, Criff)),
			//gosynth.Sine(clock(sampleRate/32.70/7)),
			//gosynth.Sawtooth(sampleRate/660),
		),
	)
	stream, err := portaudio.OpenDefaultStream(0, 1, sampleRate, portaudio.FramesPerBufferUnspecified, master.PortAudioCallback())
	if err != nil {
		panic(err)
	}
	defer stream.Close()
	err = stream.Start()
	if err != nil {
		panic(err)
	}
	fmt.Println("Press `Enter` to stop")
	fmt.Scanln()
	stream.Stop()
}
