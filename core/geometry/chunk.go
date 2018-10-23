package geometry

import "fmt"

type Chunk struct {
	x, y, z, total int
	data           []int
	dirty bool
}

func NewChunk(x, y, z int) *Chunk {
	return &Chunk{
		x: x, y: y, z: z, total: x + y + z,
		data: make([]int, x+y+z),
		dirty: true,
	}
}

func (c Chunk) to1D(x, y, z int) int {
	return (z * c.x * c.y) + (y * c.x) + x;
}

func (c Chunk) to3D(idx int) (x, y, z int) {
	z = idx / (c.x * c.y)
	idx -= z * c.x * c.y
	y = idx / c.x
	x = idx % c.x
	return
}

func (c Chunk) SetBlock(x, y, z, block int) {
	index := c.to1D(x, y, z)
	if index > c.total {
		panic(fmt.Errorf("out of bounds"))
	}
	c.data[index] = block
	c.dirty = true
}

func (c Chunk) Block(x, y, z int) int {
	index := c.to1D(x, y, z)
	if index > c.total {
		panic(fmt.Errorf("out of bounds"))
	}
	return c.data[index]
}


