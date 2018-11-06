package main

import (
	"io/ioutil"
	"math/rand"
	"os"
	"runtime"
	"time"

	"github.com/go-gl/glfw/v3.2/glfw"
	"github.com/go-gl/mathgl/mgl32"
	app "github.com/raphy42/rodent/core/application"
	"github.com/raphy42/rodent/core/graphic/gl"
	"github.com/raphy42/rodent/core/graphic/renderer"
	"github.com/raphy42/rodent/core/math"
	"github.com/raphy42/rodent/core/message"
)

func init() {
	runtime.LockOSThread()
}

func generateCube(w, h, d int) []float32 {
	out := make([]float32, 0)
	for x := 0; x < w; x++ {
		for y := 0; y < h; y++ {
			for z := 0; z < d; z++ {
				out = append(out,
					float32(x),
					float32(y),
					float32(z),
				)
			}
		}
	}
	return out
}

func generateGarbage(count int) []float32 {
	garbage := make([]float32, count)
	for i := 0; i < count; i += 3 {
		garbage[i] = math.Lerp(-5, 5, rand.Float32())
		garbage[i+1] = math.Lerp(-5, 5, rand.Float32())
		garbage[i+2] = math.Lerp(-5, 5, rand.Float32())
	}
	return garbage
}

func generateColors(count int) []float32 {
	garbage := make([]float32, count)
	for i := 0; i < count; i += 3 {
		garbage[i] = rand.Float32()
		garbage[i+1] = rand.Float32()
		garbage[i+2] = rand.Float32()
	}
	return garbage
}

func readFile(filename string) string {
	buf, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	return string(buf)
}

func test1() {
	bus := message.NewEventBus()

	k := app.New(app.Resizable(true), app.WindowSize(1200, 800), app.GLVersion(4, 1))
	if err := k.Init(); err != nil {
		panic(err)
	}

	k.RegisterEvents(bus)

	go func(){
		keys := bus.Channel("keyboard")
		for {
			ev := (<- keys).(*app.KeyboardEvent)
			switch ev.Key {
			case glfw.KeyEscape:
				os.Exit(0)
			}
		}
	}()

	shader := gl.NewShader(
		gl.Vertex(readFile("./assets/shaders/basic.vert.glsl")),
		gl.Geometry(readFile("./assets/shaders/basic.geom.glsl")),
		gl.Fragment(readFile("./assets/shaders/basic.frag.glsl")),
	)
	if err := shader.Build(); err != nil {
		panic(err)
	}
	vertices := gl.NewVec3FBuffer(generateCube(10, 10, 10))
	colors := gl.NewVec3FBuffer(generateColors(1000 * 3))

	// test := gl.NewVec3FBuffer(generateGarbage(36 * 6)).
	// 	SetAttributes(gl.Vec3, gl.Vec3)

	payload := gl.NewPayload(vertices, colors).
		First(0).
		Primitive(gl.Points, 1000).
		Build()

	projection := mgl32.Perspective(mgl32.DegToRad(60), 1200/800, 0.1, 10000.0)
	camera := mgl32.LookAtV(mgl32.Vec3{15, 15, 15}, mgl32.Vec3{0, 0, 0}, mgl32.Vec3{0, 1, 0})
	mvp := projection.Mul4(camera)

	renderer.Background(mgl32.Vec4{.5, .5, .5, 1})

	for !k.ShouldShutdown() {

		renderer.Clear()
		shader.SetMat4("mvp", mvp)
		payload.Bind()
		payload.DrawArrays()
		k.Tick(time.Now())
	}
	k.Dispose()
}

func main() {
	test1()
}
