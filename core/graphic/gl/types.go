package gl

import (
	"github.com/go-gl/gl/v4.1-core/gl"
)

type Attribute struct {
	Type Type
}

type Type int

const (
	Int = 4
	Float = 4
	Vec2 = 2 * Float
	Vec3 = 3 * Float
	Vec4 = 4 * Float
	Mat2 = 2 * 2 * Float
	Mat3 = 3 * 3 * Float
	Mat4 = 4 * 4 * Float
)

func (t Type) Count() int32 {
	switch t {
	case Vec2:
		return 2
	case Vec3:
		return 3
	case Mat2:
		return 4
	case Mat3:
		return 9
	case Mat4:
		return 16
	}
	return 0
}

func (t Type) Enum() uint32 {
	switch t {
	case Float, Vec2, Vec3, Vec4 /*Mat2*/, Mat3, Mat4: return gl.FLOAT
	}
	return 0
}
