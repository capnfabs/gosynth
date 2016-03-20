package gosynth

type Sample float32
type Clock uint64

// TODO: How to handle sample frequency?
type Synth func(clock Clock, out [][]Sample)

type PortAudioCallback func(out [][]Sample)

func (synth Synth) PortAudioCallback() PortAudioCallback {
	clock := Clock(0)
	return func(out [][]Sample) {
		synth(clock, out)
		clock += Clock(len(out[0]))
	}
}
