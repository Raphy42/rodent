package math

import (
	"math"
)

func Lerp(min, max, value float32) float32 {
	return (1-value)*min + value*max
}

func Sin(v float32) float32 {
	return float32(math.Sin(float64(v)))
}

func Cos(v float32) float32 {
	return float32(math.Cos(float64(v)))
}

func HSVtoRGB(h, s, v float32) (float32, float32, float32) {
	if s == 0.0 {
		return v, v, v
	}
	i := int(h * 6.)
	f := (h * 6.) - float32(i)
	p, q, t := v*(1.-s), v*(1.-s*f), v*(1.-s*(1.-f))
	i %= 6
	switch i {
	case 0:
		return v, t, p
	case 1:
		return q, v, p
	case 2:
		return p, v, t
	case 3:
		return p, q, v
	case 4:
		return t, p, v
	case 5:
		return v, p, q
	}
	return 0, 0, 0
}
