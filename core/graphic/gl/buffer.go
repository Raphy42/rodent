package gl

import "github.com/go-gl/gl/v4.1-core/gl"

type Buffer struct {
	handle     uint32
	xtype      int
	target     uint32
	count      int
	attributes func()
}

func (b Buffer) Bind() {
	gl.BindBuffer(b.target, b.handle)
}

type BufferBuilder struct {
	tasks []func()
	out   *Buffer
	size int32
	xtype uint32
	strided bool
}

func NewBufferBuilder() *BufferBuilder {
	b := &BufferBuilder{
		out: new(Buffer),
		tasks: make([]func(), 0),
	}
	b.Enqueue(func() {
		b.out.handle = genVBO()
	})
	return b
}

func (b *BufferBuilder) Target(target uint32) *BufferBuilder {
	b.out.target = target
	return b.Enqueue(func() {
		gl.BindBuffer(target, b.out.handle)
	})
}

func (b *BufferBuilder) Enqueue(fn func()) *BufferBuilder {
	b.tasks = append(b.tasks, fn)
	return b
}

func (b *BufferBuilder) SetAttribute(slot uint32) *BufferBuilder {
	return b.Enqueue(func() {
		gl.EnableVertexAttribArray(slot)
		gl.VertexAttribPointer(slot, b.size, b.xtype, false, 0, nil)
	})
}

func (b *BufferBuilder) SetAttributes(attributes ...Type) *BufferBuilder {
	b.strided = true
	return b.Enqueue(func() {
		stride := int32(0)
		for _, attribute := range attributes {
			stride += attribute.Count()
		}
		for index, xtype := range attributes {
			slot := uint32(index)
			gl.EnableVertexAttribArray(slot)
			gl.VertexAttribPointer(slot, xtype.Count(), xtype.Enum(), false, stride, nil)
		}
	})
	return b
}

func (b *BufferBuilder) Build() *Buffer {
	for _, task := range b.tasks {
		task()
	}
	return b.out
}

func genVBO() uint32 {
	var vbo uint32
	gl.GenBuffers(1, &vbo)
	return vbo
}

func NewVec3FBuffer(data []float32) *BufferBuilder {
	b := NewBufferBuilder()
	b.size = 3
	b.xtype = gl.FLOAT
	b.Target(gl.ARRAY_BUFFER).Enqueue(func() {
		gl.BufferData(gl.ARRAY_BUFFER, Float*len(data), gl.Ptr(data), gl.STATIC_DRAW)
	})
	return b
}

func NewInt32Buffer(data []int32) *BufferBuilder {
	b := NewBufferBuilder()
	b.size = 1
	b.xtype = gl.INT
	b.Target(gl.ARRAY_BUFFER).Enqueue(func() {
		gl.BufferData(gl.ARRAY_BUFFER, Int*len(data), gl.Ptr(data), gl.STATIC_DRAW)
	})
	return b
}
