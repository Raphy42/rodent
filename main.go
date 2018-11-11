package main

import (
	"io/ioutil"
	"math/rand"
	"os"
	"runtime"
	"time"

	"github.com/go-gl/mathgl/mgl32"
	app "github.com/raphy42/rodent/core/application"
	"github.com/raphy42/rodent/core/application/input"
	"github.com/raphy42/rodent/core/graphic/gl"
	"github.com/raphy42/rodent/core/graphic/renderer"
	"github.com/raphy42/rodent/core/grid"
	"github.com/raphy42/rodent/core/logic"
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

	go func() {
		keys := bus.Subscribe(message.Keyboard.String())
		for {
			ev := (<-keys).(*input.KeyboardAction)
			switch ev.Key {
			case input.KeyEscape:
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

	testCube := grid.New(10, 10, 10)
	palette := make([]mgl32.Vec3, 1000)
	idx := 0
	testCube.Map(func(x, y, z int, in grid.Node) grid.Node {
		defer func() {
			idx += 1
		}()
		if idx%10 == 0 {
			return in
		}
		r, g, b := math.HSVtoRGB(float32(x)/10, float32(y)/10, float32(z)/10)
		palette[idx] = mgl32.Vec3{r, g, b}
		return grid.Node(idx + 1)
	})

	vertices := gl.NewVec3FBuffer(testCube.Vertices())
	colors := gl.NewVec3FBuffer(testCube.Colors(palette))

	// test := gl.NewVec3FBuffer(generateGarbage(36 * 6)).
	// 	SetAttributes(gl.Vec3, gl.Vec3)

	payload := gl.NewPayload(vertices, colors).
		First(0).
		Primitive(gl.Points, 1000).
		Build()

	keys := bus.Subscribe(message.Keyboard.String())
	go func() {
		wireframe := false
		for {
			ev := <-keys
			key := ev.(*input.KeyboardAction)
			if key.IsPressed() && key.Is(input.KeyLeftBracket) {
				payload.Wireframe(wireframe)
				wireframe = !wireframe
			}
		}
	}()

	camera := logic.NewCamera()
	camera.Register(bus)
	camera.Eye = mgl32.Vec3{15, 15, 15}

	renderer.Background(mgl32.Vec4{.5, .5, .5, 1})
	model := mgl32.Scale3D(.5, .5, .5)

	for !k.ShouldShutdown() {

		renderer.Clear()
		projection := camera.Projection()
		mvp := model.Mul4(projection.Mul4(camera.View()))
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
