package texture

import (
	"fmt"
	"image"
	"image/draw"
	"log"

	"../file"
	"github.com/go-gl/gl/v4.1-core/gl"
)

// Texture struct for holding information about texture assets
type Texture struct {
	TextureID uint32
	FileName  string
	Name      string
}

// NewTexture creates a texture
func NewTexture(filename string, name string) (*Texture, error) {
	texture := new(Texture)
	texture.FileName = filename
	texture.Name = name
	img, err := file.LoadTexture(filename)
	if err != nil {
		log.Fatalln(err)
		panic(err)
	}
	rgba := image.NewRGBA(img.Bounds())
	if rgba.Stride != rgba.Rect.Size().X*4 {
		return nil, fmt.Errorf("unsupported stride")
	}
	draw.Draw(rgba, rgba.Bounds(), img, image.Point{0, 0}, draw.Src)
	var textures uint32
	gl.GenTextures(1, &textures)
	texture.TextureID = textures
	gl.ActiveTexture(gl.TEXTURE0)
	gl.BindTexture(gl.TEXTURE_2D, textures)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_MIN_FILTER, gl.LINEAR)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_MAG_FILTER, gl.LINEAR)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_WRAP_S, gl.CLAMP_TO_EDGE)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_WRAP_T, gl.CLAMP_TO_EDGE)
	gl.TexImage2D(
		gl.TEXTURE_2D,
		0,
		gl.RGBA,
		int32(rgba.Rect.Size().X),
		int32(rgba.Rect.Size().Y),
		0,
		gl.RGBA,
		gl.UNSIGNED_BYTE,
		gl.Ptr(rgba.Pix))
	return texture, nil
}
