package gosynth

func Constant(value Sample) Synth {
	return func(clock Clock, out []Sample) {
		for i := 0; i < len(out); i++ {
			out[i] = value
		}
	}
}
