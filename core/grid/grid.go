package grid

import "github.com/go-gl/mathgl/mgl32"

type Node int

type Grid struct {
	total                            int
	width, height, depth             int
	halfWidth, halfHeight, halfDepth int
	data                             []Node
}

func New(w, h, d int) *Grid {
	total := w * h * d
	return &Grid{
		total: total,
		width: w, height: h, depth: d,
		halfWidth: w / 2, halfHeight: h / 2, halfDepth: d / 2,
		data: make([]Node, total),
	}
}

func (g *Grid) to3D(index int) (x, y, z int) {
	z = index / (g.width * g.height)
	index -= z * g.width * g.height
	y = index / g.width
	x = index % g.width
	return
}

func (g *Grid) to1D(x, y, z int) int {
	return z*g.width*g.height + y*g.width + x
}

func (g *Grid) Add(x, y, z int, node Node) bool {
	idx := g.to1D(x, y, z)
	if idx < 0 || idx > g.total {
		return false
	}
	g.data[idx] = node
	return true
}

func (g *Grid) At(x, y, z int) Node {
	idx := g.to1D(x, y, z)
	if idx < 0 || idx > g.total {
		return -1
	}
	return g.data[idx]
}

func (g *Grid) Vertices() []float32 {
	out := make([]float32, 0)
	for idx, node := range g.data {
		if node != 0 {
			x, y, z := g.to3D(int(idx))
			out = append(out, float32(x), float32(y), float32(z))
		}
	}
	return out
}

func (g *Grid) Colors(palette []mgl32.Vec3) []float32 {
	size := len(palette)
	out := make([]float32, 0)
	for _, node := range g.data {
		n := int(node)
		if n != 0 {
			if n > size+1 {
				out = append(out, 1, 1, 1)
			} else {
				color := palette[n-1]
				out = append(out, color.X(), color.Y(), color.Z())
			}
		}
	}
	return out
}

func (g *Grid) Map(fn func(x, y, z int, in Node) Node) {
	for idx, node := range g.data {
		x, y, z := g.to3D(idx)
		g.data[idx] = fn(x, y, z, node)
	}
}
