package goretry

import "time"

func ConstantBackoff(d time.Duration) func(int) time.Duration {
	return func(_ int) time.Duration {
		return d
	}
}

func LinearBackoff(d time.Duration) func(int) time.Duration {
	return func(i int) time.Duration {
		if i <= 0 {
			return 0
		}
		return d * time.Duration(i)
	}
}

func ExponentialBackoff(base time.Duration, factor float64) func(int) time.Duration {
	return func(attempt int) time.Duration {
		return time.Duration(float64(base) * factor * float64(attempt))
	}
}
