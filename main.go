package main

import (
	"io/ioutil"
	"os"
	"runtime"
	"time"

	"github.com/go-gl/mathgl/mgl32"
	app "github.com/raphy42/rodent/core/application"
	"github.com/raphy42/rodent/core/application/input"
	"github.com/raphy42/rodent/core/graphic/gl"
	"github.com/raphy42/rodent/core/graphic/renderer"
	"github.com/raphy42/rodent/core/grid"
	"github.com/raphy42/rodent/core/logger"
	"github.com/raphy42/rodent/core/logic"
	"github.com/raphy42/rodent/core/math"
	"github.com/raphy42/rodent/core/message"
)

var (
	log = logger.New()
)

func init() {
	runtime.LockOSThread()
}

func readFile(filename string) string {
	buf, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	return string(buf)
}

func main() {
	bus := message.NewEventBus()

	k := app.New(app.Resizable(true), app.WindowSize(800, 800), app.GLVersion(4, 1))
	if err := k.Init(); err != nil {
		panic(err)
	}

	k.RegisterEvents(bus)

	shader := gl.NewShader(
		gl.Vertex(readFile("./assets/shaders/basic.vert.glsl")),
		gl.Geometry(readFile("./assets/shaders/basic.geom.glsl")),
		gl.Fragment(readFile("./assets/shaders/basic.frag.glsl")),
	)
	if err := shader.Build(); err != nil {
		panic(err)
	}

	const (
		W = 30
		H = 30
		D = 30
	)

	testCube := grid.New(W, H, D)
	palette := make([]mgl32.Vec3, W*H*D)
	idx := 0
	const radius = 15
	testCube.Map(func(x, y, z int, in grid.Node) grid.Node {
		defer func() {
			idx += 1
		}()

		dx := x - testCube.HalfWidth
		dy := y - testCube.HalfHeight
		dz := z - testCube.HalfDepth

		distance := math.Sqrt(float32(dx*dx + dy*dy + dz*dz))
		if distance > radius {
			return in
		}
		var r, g, b float32
		if x%3 == 0 || y%3 == 0 || z%3 == 0 {
			if x%3 == y%3 && x%3 == z%3 {
				r, g, b = math.HSVtoRGB(float32(x)/W, 60, float32(z)/D)
				goto setPalette
			}
			return in
		} else {
			r, g, b = math.HSVtoRGB(float32(x)/W, float32(y)/H, float32(z)/D)
		}
	setPalette:
		palette[idx] = mgl32.Vec3{r, g, b}
		// value returned must correspond to a valid index + 1
		// 1 < idx + 1 < len(palette) + 1
		return grid.Node(idx + 1)
	})

	vertices := gl.NewVec3FBuffer(testCube.Vertices())
	colors := gl.NewVec3FBuffer(testCube.Colors(palette))

	payload := gl.NewPayload(vertices, colors).
		First(0).
		Primitive(gl.Points, W*H*D).
		Build()

	keys := bus.Subscribe(message.Keyboard)
	go func() {
		wireframe := false
		for {
			ev := <-keys
			key := ev.(*input.KeyboardAction)

			switch {
			case key.IsPressed() && key.Is(input.KeyLeftBracket):
				{
					wireframe = !wireframe
					payload.Wireframe(wireframe)
				}
			case key.Is(input.KeyEscape):
				os.Exit(0)
			}
		}
	}()

	camera := logic.NewCamera()
	camera.Register(bus)

	// camera position
	camera.Eye = mgl32.Vec3{15, 15, 15}

	// background color
	renderer.Background(mgl32.Vec4{.5, .5, .5, 1})

	// model scale
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
