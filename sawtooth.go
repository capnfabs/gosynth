package gosynth

func Sawtooth(period Clock) Synth {
	return func(clock Clock, out []Sample) {
		for i := 0; i < len(out); i++ {
			out[i] = Sample((clock+Clock(i))%period)/Sample(period)*Sample(2) - Sample(1)
		}
	}
}

func SawtoothWithPeriod(period TimeFunc) Synth {
	return func(clock Clock, out []Sample) {
		for i := 0; i < len(out); i++ {
			thisClock := clock + Clock(i)
			thisPeriod := period(clock)
			out[i] = Sample(thisClock%thisPeriod)/Sample(thisPeriod)*Sample(2) - Sample(1)
		}
	}
}
