package gosynth

func Sawtooth(period Clock) Synth {
	return func(clock Clock, out [][]Sample) {
		for i := 0; i < len(out[0]); i++ {
			val := Sample((clock+Clock(i))%period)*Sample(2) - Sample(1)
			out[0][i] = val
			out[1][i] = val
		}
	}
}
