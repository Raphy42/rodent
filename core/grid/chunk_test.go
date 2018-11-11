package grid

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	assert.New(t).NotNil(New(1, 1, 1))
}

func TestGrid_Colors(t *testing.T) {
	a := assert.New(t)

	grid := New(1, 1, 1)
	a.NotNil(grid)

	for i := 0; i < grid.total; i++ {
		x, y, z := grid.to3D(i)
		grid.Add(x, y, z, Node(10))
	}

	a.Len(grid.Colors(nil), grid.total*3)
}
