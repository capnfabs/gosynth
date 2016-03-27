package gosynth

func addBuffers(out []Sample, a []Sample, b []Sample) {
	l := len(out)
	if l != len(a) || l != len(b) {
		panic("Expected equal buffer length")
	}
	for i := 0; i < l; i++ {
		out[i] = a[i] + b[i]
	}
}

func Add(synths ...Synth) Synth {
	return func(clock Clock, out []Sample) {
		synths[0](clock, out)
		for i := 1; i < len(synths); i++ {
			thisSynth := make([]Sample, len(out))
			synths[i](clock, thisSynth)
			addBuffers(out, out, thisSynth)
		}
		for i := 0; i < len(out); i++ {
			if out[i] > 1 {
				out[i] = 1
			} else if out[i] < -1 {
				out[i] = -1
			}
		}
	}
}
