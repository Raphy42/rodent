package gl

import (
	"github.com/go-gl/gl/v4.1-core/gl"
)

type Payload struct {
	handle    uint32
	mode      uint32
	first     int32
	count     int32
	wireframe bool
}

type Primitive uint32

const (
	Triangles = gl.TRIANGLES
	Points    = gl.POINTS
)

type PayloadBuilder struct {
	out      *Payload
	builders []*BufferBuilder
}

func genVAO() uint32 {
	var vao uint32
	gl.GenVertexArrays(1, &vao)
	return vao
}

func NewPayload(builders ...*BufferBuilder) *PayloadBuilder {
	return &PayloadBuilder{out: new(Payload), builders: builders}
}

func (p *PayloadBuilder) Primitive(mode Primitive, count int32) *PayloadBuilder {
	p.out.mode = uint32(mode)
	p.out.count = count
	return p
}

func (p *PayloadBuilder) First(first int32) *PayloadBuilder {
	p.out.first = first
	return p
}

func (p *PayloadBuilder) Build() *Payload {
	p.out.handle = genVAO()
	p.out.Bind()

	for slot, builder := range p.builders {
		if !builder.strided {
			builder.SetAttribute(uint32(slot))
		}
		builder.Build()
	}
	return p.out
}

func (p *Payload) Bind() {
	gl.BindVertexArray(p.handle)
}

func (p *Payload) DrawArrays() {
	if p.wireframe {
		gl.PolygonMode(gl.FRONT_AND_BACK, gl.LINE)
	} else {
		gl.PolygonMode(gl.FRONT_AND_BACK, gl.FILL)
	}
	gl.DrawArrays(p.mode, p.first, p.count)
}

func (p *Payload) Wireframe(value bool) {
	p.wireframe = value
}
