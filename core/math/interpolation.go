package math

import "math"

func Lerp(min, max, value float32) float32 {
	return (1 - value) * min + value * max
}

func Sin(v float32) float32 {
	return float32(math.Sin(float64(v)))
}

func Cos(v float32) float32 {
	return float32(math.Cos(float64(v)))
}