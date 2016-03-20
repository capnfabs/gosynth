package gosynth

func Square(period Clock) Synth {
	return SquareMinMax(period, -1, 1)
}

func SquareMinMax(period Clock, min, max Sample) Synth {
	halfPeriod := period / 2
	return func(clock Clock, out []Sample) {
		for i := 0; i < len(out); i++ {
			if (clock+Clock(i))%period < halfPeriod {
				out[i] = min
			} else {
				out[i] = max
			}
		}
	}
}
