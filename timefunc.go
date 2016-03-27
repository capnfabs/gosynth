package gosynth

type TimeFunc func(clock Clock) Clock

func StepSequencer(stepLength Clock, values []Clock) TimeFunc {
	return func(clock Clock) Clock {
		idx := int(clock/stepLength) % len(values)
		return values[idx]
	}
}

// TODO: rename
func TimeConstant(value Clock) TimeFunc {
	return func(_ Clock) Clock {
		return value
	}
}
