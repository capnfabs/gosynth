package main

import (
	"crypto/md5"
	"flag"
	"fmt"
	"math"
	"strings"

	"github.com/capnfabs/gosynth"
	"github.com/gordonklaus/portaudio"
)

const sampleRate = 44100.0

// These are all in discrete period notation, not frequency.
// - Use [note] * 2 to go down an octave.
// - Use [note] / 2 to go up an octave.
var (
	// Middle C
	cRaw = sampleRate / 261.625565
	C    = clock(cRaw)
	Cs   = clock(cRaw / math.Exp2(1.0/12))
	Db   = Cs
	D    = clock(cRaw / math.Exp2(2.0/12))
	Ds   = clock(cRaw / math.Exp2(3.0/12))
	Eb   = Ds
	E    = clock(cRaw / math.Exp2(4.0/12))
	F    = clock(cRaw / math.Exp2(5.0/12))
	Fs   = clock(cRaw / math.Exp2(6.0/12))
	Gb   = Fs
	G    = clock(cRaw / math.Exp2(7.0/12))
	Gs   = clock(cRaw / math.Exp2(8.0/12))
	Ab   = Gs
	A    = clock(cRaw / math.Exp2(9.0/12))
	As   = clock(cRaw / math.Exp2(10.0/12))
	Bb   = As
	B    = clock(cRaw / math.Exp2(11.0/12))
)

var keys = []gosynth.Clock{
	C,
	Cs,
	D,
	Ds,
	E,
	F,
	Fs,
	G,
	Gs,
	A,
	As,
	B,
}

var Cmajor = []gosynth.Clock{
	C,
	D,
	E,
	F,
	G,
	A,
	B,
	C / 2,
}

var Cbass = []gosynth.Clock{
	C * 4,
	G * 4,
	F * 4,
	E * 4,
}

var Criff = []gosynth.Clock{
	C,
	C,
	F,
	E,
	E,
	B,
	C / 2,
	B,
}

func sine(out [][]float32) {
}

func clock(val float64) gosynth.Clock {
	return gosynth.Clock(val)
}

var quav = clock(sampleRate / 3)

func genSequence(palette []gosynth.Clock, data []byte, count int) []gosynth.Clock {
	ret := make([]gosynth.Clock, count)
	for i := 0; i < count; i++ {
		ret[i] = palette[int(data[i])%len(palette)]
	}
	return ret
}

func genPalette(baseNote gosynth.Clock, count int) []gosynth.Clock {
	ret := make([]gosynth.Clock, count)
	for i := 0; i < count; i++ {
		ret[i] = gosynth.Clock(float64(baseNote) / math.Exp2(float64(i)/12))
	}
	return ret
}

func argPlay(args []string) gosynth.Synth {
	// Take args and hash them.
	hash := md5.Sum([]byte(strings.Join(args, "")))
	// Choose a key.
	key := keys[int(hash[0])%len(keys)]
	// Generate a palette based on that key.
	// 1 octave at the moment.
	palette := genPalette(key, 13)
	// Now choose a length for each pattern. Min 3, max 8
	patternLen := int((hash[1] % 5)) + 3
	seq := genSequence(palette, hash[3:], patternLen)
	speedDivider := float32(hash[2] % 5)
	return gosynth.Mult(
		gosynth.Constant(0.1),
		gosynth.SawtoothWithPeriod(gosynth.StepSequencer(gosynth.Clock(sampleRate/speedDivider), seq)),
	)
}

func defaultPlay() gosynth.Synth {
	return gosynth.Mult(
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
}

func main() {
	flag.Parse()
	err := portaudio.Initialize()
	if err != nil {
		panic(err)
	}
	var master gosynth.Synth
	if len(flag.Args()) > 0 {
		fmt.Println("Playing using hashed input.")
		master = argPlay(flag.Args())
	} else {
		fmt.Println("Playing default track.")
		master = defaultPlay()
	}
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
