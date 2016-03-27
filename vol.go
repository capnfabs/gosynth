package gosynth

func Vol(k Sample, synth Synth) Synth {
	return func(clock Clock, out []Sample) {
		synth(clock, out)
		for i := 0; i < len(out); i++ {
			out[i] = k * out[i]
		}
	}
}
