package gosynth

func Constant(value Sample) Synth {
	return func(clock Clock, out [][]Sample) {
		for i := 0; i < len(out[0]); i++ {
			out[0][i] = value
			out[1][i] = value
		}
	}
}
