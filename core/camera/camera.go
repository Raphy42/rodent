package camera

import (
	"sync"

	"github.com/go-gl/mathgl/mgl32"
	"github.com/raphy42/rodent/core/math"
)

type Camera struct {
	sync.RWMutex
	Eye                    mgl32.Vec3
	Target                 mgl32.Vec3
	Up                     mgl32.Vec3
	Near, Far, FOV, Aspect float32

	Speed, yaw, pitch, sensitivity, zoom float32
	dirty                                bool
	front, right, worldUp                mgl32.Vec3
	view                                 mgl32.Mat4
	projection                           mgl32.Mat4
}

func NewPerspective() *Camera {
	return &Camera{
		Eye:     mgl32.Vec3{1, 1, 1},
		Target:  mgl32.Vec3{0, 0, 0},
		Up:      mgl32.Vec3{0, 1, 0},
		worldUp: mgl32.Vec3{0, 1, 0},
		Near:    0.1, Far: 100,
		FOV: 60, Aspect: 4 / 3,
		Speed: mgl32.DegToRad(1),
		front: mgl32.Vec3{0, 0, -1},
		right: mgl32.Vec3{1, 0, 0},
		yaw:   -90, sensitivity: 0.25, zoom: 68,
		dirty: true,
	}
}

func (c *Camera) View() mgl32.Mat4 {
	c.Lock()
	defer c.Unlock()

	if c.dirty {
		c.rebuild()
		c.view = mgl32.LookAtV(c.Eye, c.Eye.Add(c.front), c.Up)
		c.dirty = true
	}
	return c.view
}

func (c *Camera) rebuild() {
	c.front = mgl32.Vec3{
		math.Cos(mgl32.DegToRad(c.yaw)) * math.Cos(mgl32.DegToRad(c.pitch)),
		math.Sin(mgl32.DegToRad(c.pitch)),
		math.Sin(mgl32.DegToRad(c.yaw)) * math.Cos(mgl32.DegToRad(c.pitch)),
	}.Normalize()
	c.right = c.front.Cross(c.worldUp).Normalize()
	c.Up = c.right.Cross(c.front).Normalize()
}

func (c *Camera) Projection() mgl32.Mat4 {
	c.Lock()
	defer c.Unlock()

	if c.dirty {
		c.projection = mgl32.Perspective(mgl32.DegToRad(c.zoom), c.Aspect, c.Near, c.Far)
		c.dirty = true
	}
	return c.projection
}

func (c *Camera) Zoom(amount float32) {
	c.Lock()
	defer c.Unlock()

	c.zoom = mgl32.Clamp(c.zoom-amount, 20, 120)
	c.dirty = true
}

func (c *Camera) Move(action Action, delta float32) {
	c.Lock()
	defer c.Unlock()

	vel := c.Speed * delta
	switch action {
	case Forward:
		c.Eye = c.Eye.Add(c.front.Mul(vel))
	case Backward:
		c.Eye = c.Eye.Sub(c.front.Mul(vel))
	case Left:
		c.Eye = c.Eye.Sub(c.right.Mul(vel))
	case Right:
		c.Eye = c.Eye.Add(c.right.Mul(vel))
	}
	c.dirty = true
}

func (c *Camera) Center(x, y float32) {
	c.Lock()
	defer c.Unlock()

	x *= c.sensitivity
	y *= c.sensitivity

	c.yaw += x
	c.pitch += y
	// mgl32.Clamp(c.pitch, -89, 89)
	c.dirty = true
}
