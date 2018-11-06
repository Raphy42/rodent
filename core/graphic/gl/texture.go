package gl

import (
	"image"
	"image/draw"
	"os"

	"github.com/go-gl/gl/v4.1-core/gl"
	"github.com/pkg/errors"
	"go.uber.org/zap"
)

type Texture struct {
	TextureOptions
	handle uint32
	build func() error
}

type TextureOptions struct {
	target uint32
	minFilter int32
	magFilter int32
	wrapS int32
	wrapT int32
}

type TextureOption func (*TextureOptions)

type TextureFilter int32

const (
	Nearest = TextureFilter(gl.NEAREST)
	Linear = TextureFilter(gl.LINEAR)
	NearestMipmapNearest = TextureFilter(gl.NEAREST_MIPMAP_NEAREST)
	LinearMipmapNearest = TextureFilter(gl.LINEAR_MIPMAP_NEAREST)
	NearestMipmapLinear = TextureFilter(gl.NEAREST_MIPMAP_LINEAR)
	LinearMipmapLinear = TextureFilter(gl.LINEAR_MIPMAP_LINEAR)
)

func Filter(min, mag TextureFilter) TextureOption {
	return func(options *TextureOptions) {
		options.minFilter = int32(min)
		options.magFilter = int32(mag)
	}
}

type TextureWrapper int32

const (
	Repeat = TextureWrapper(gl.REPEAT)
	MirroredRepeat = TextureWrapper(gl.MIRRORED_REPEAT)
	ClampToEdge = TextureWrapper(gl.CLAMP_TO_EDGE)
	ClampToBorder = TextureWrapper(gl.CLAMP_TO_BORDER)
)

func Wrap(width, height TextureWrapper) TextureOption {
	return func(options *TextureOptions) {
		options.wrapS = int32(width)
		options.wrapT = int32(height)
	}
}

func NewTexture(filename string) *Texture {
	t := new(Texture)

	t.build = func() error {
		file, err := os.Open(filename)
		if err != nil {
			return errors.Wrap(err, filename)
		}
		img, format, err := image.Decode(file)
		if err != nil {
			return errors.Wrap(err, filename)
		}
		log.Debug("decoded texture",
			zap.String("format", format),
			zap.String("filename", filename),
		)

		rgba := image.NewRGBA(img.Bounds())
		if rgba.Stride != rgba.Rect.Size().X * 4 {
			return errors.New("unsupported stride")
		}

		draw.Draw(rgba, rgba.Bounds(), img, image.Pt(0, 0), draw.Src)

		width, height := func(pt image.Point) (int32, int32){
			return int32(pt.X), int32(pt.Y)
		}(rgba.Rect.Size())

		gl.GenTextures(1, &t.handle)
		gl.ActiveTexture(gl.TEXTURE0)
		gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_MIN_FILTER, gl.LINEAR)
		gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_MAG_FILTER, gl.LINEAR)
		gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_WRAP_S, gl.CLAMP_TO_EDGE)
		gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_WRAP_T, gl.CLAMP_TO_EDGE)
		gl.TexImage2D(gl.TEXTURE_2D, 0, gl.RGBA, width, height, 0, gl.RGBA, gl.UNSIGNED_BYTE, gl.Ptr(rgba.Pix))

		return nil
	}
	return t
}

func (t *Texture) Build() error {
	return t.build()
}