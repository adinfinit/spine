package spine

import (
	"math"
)

type Color struct{ R, G, B, A float32 }

func (c Color) WithAlpha(a float32) Color {
	c.A = a
	return c
}

func u32(v float32) uint32 {
	if v <= 0 {
		return 0
	} else if v >= 1 {
		return 0xFFFF
	}
	return uint32(v * 0xFFFF)
}

func (c Color) RGBA() (r, g, b, a uint32) {
	return u32(c.R), u32(c.G), u32(c.B), u32(c.A)
}
func (c Color) Float32() (r, g, b, a float32) { return c.R, c.G, c.B, c.A }
func (c Color) Float64() (r, g, b, a float64) {
	return float64(c.R), float64(c.G), float64(c.B), float64(c.A)
}

func (c Color) RGB64() (r, g, b float64) {
	return float64(c.R), float64(c.G), float64(c.B)
}
func (c Color) RGBA64() (r, g, b, a float64) {
	return float64(c.R), float64(c.G), float64(c.B), float64(c.A)
}

func abs(x float32) float32 {
	if x < 0 {
		return -x
	}
	return x
}
func atan2(y, x float32) float32 { return float32(math.Atan2(float64(y), float64(x))) }
func mod(x, y float32) float32   { return float32(math.Mod(float64(x), float64(y))) }
func sqrt(v float32) float32     { return float32(math.Sqrt(float64(v))) }
func sin(v float32) float32      { return float32(math.Sin(float64(v))) }
func cos(v float32) float32      { return float32(math.Cos(float64(v))) }
func sincos(v float32) (float32, float32) {
	sn, cs := math.Sincos(float64(v))
	return float32(sn), float32(cs)
}

func clamp01(v float32) float32 {
	if v <= 0 {
		return 0
	} else if v >= 1 {
		return 1
	}
	return v
}

func lerp(a, b float32, p float32) float32 { return a*(1-p) + b*p }

func lerpAngle(a, b, p float32) float32 {
	delta := b - a
	for delta > math.Pi {
		delta -= 2 * math.Pi
	}
	for delta < -math.Pi {
		delta += 2 * math.Pi
	}
	return a + delta*p
}

func lerpVector(a, b Vector, p float32) Vector {
	return Vector{
		a.X*(1-p) + b.X*p,
		a.Y*(1-p) + b.Y*p,
	}
}

func lerpAngleVector(a, b Vector, p float32) Vector {
	return Vector{
		lerpAngle(a.X, b.X, p),
		lerpAngle(a.Y, b.Y, p),
	}
}

func lerpColor(a, b Color, p float32) Color {
	return Color{
		a.R*(1-p) + b.R*p,
		a.G*(1-p) + b.G*p,
		a.B*(1-p) + b.B*p,
		a.A*(1-p) + b.A*p,
	}
}
