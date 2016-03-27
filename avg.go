package gosynth

func Avg(synths ...Synth) Synth {
	return func(clock Clock, out []Sample) {
		synths[0](clock, out)
		for i := 1; i < len(synths); i++ {
			thisSynth := make([]Sample, len(out))
			synths[i](clock, thisSynth)
			addBuffers(out, out, thisSynth)
		}
		for i := 0; i < len(out); i++ {
			out[i] /= Sample(len(synths))
		}
	}
}
