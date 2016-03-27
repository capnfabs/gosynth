package gosynth

import "math"

func Sine(period Clock) Synth {
	return func(clock Clock, out []Sample) {
		for i := 0; i < len(out); i++ {
			clk := clock + Clock(i)
			val := Sample(math.Sin(2 * math.Pi * float64(clk) / float64(period)))
			out[i] = val
		}
	}
}

func SineWithPeriod(period TimeFunc) Synth {
	return func(clock Clock, out []Sample) {
		for i := 0; i < len(out); i++ {
			clk := clock + Clock(i)
			val := Sample(math.Sin(2 * math.Pi * float64(clk) / float64(period(clk))))
			out[i] = val
		}
	}
}
