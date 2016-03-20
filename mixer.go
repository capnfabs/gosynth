package gosynth

func multBuffers(out []Sample, a []Sample, b []Sample) {
	l := len(out)
	if l != len(a) || l != len(b) {
		panic("Expected equal buffer length")
	}
	for i := 0; i < l; i++ {
		out[i] = a[i] * b[i]
	}
}

func multStreams(out [][]Sample, a [][]Sample, b [][]Sample) {
	if len(out) != len(a) || len(out) != len(b) {
		panic("Expected equal number of channels")
	}
	for i := 0; i < len(out); i++ {
		multBuffers(out[i], a[i], b[i])
	}
}

func allocateSameSizeAs(a [][]Sample) [][]Sample {
	out := make([][]Sample, len(a))
	for i := 0; i < len(a); i++ {
		out[i] = make([]Sample, len(a[i]))
	}
	return out
}

func Mult(synths ...Synth) Synth {
	return func(clock Clock, out [][]Sample) {
		synths[0](clock, out)
		for i := 1; i < len(synths); i++ {
			thisSynth := allocateSameSizeAs(out)
			synths[i](clock, thisSynth)
			multStreams(out, out, thisSynth)
		}
	}
}
