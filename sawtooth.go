package gosynth

func Sawtooth(period Clock) Synth {
	return func(clock Clock, out []Sample) {
		for i := 0; i < len(out); i++ {
			out[i] = Sample((clock+Clock(i))%period)/Sample(period)*Sample(2) - Sample(1)
		}
	}
}
