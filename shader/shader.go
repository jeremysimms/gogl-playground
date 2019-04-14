package shader

import (
	"fmt"
	"regexp"
	"strings"

	"../file"
	"github.com/go-gl/gl/v4.1-core/gl"
)

func compileShader(source string, shaderType uint32) (uint32, error) {
	shader := gl.CreateShader(shaderType)
	sources, free := gl.Strs(source + "\x00")
	gl.ShaderSource(shader, 1, sources, nil)
	gl.CompileShader(shader)
	free()
	var status int32
	gl.GetShaderiv(shader, gl.COMPILE_STATUS, &status)
	if status == gl.FALSE {
		var logLength int32
		gl.GetShaderiv(shader, gl.INFO_LOG_LENGTH, &logLength)
		log := strings.Repeat("\x00", int(logLength+1))
		gl.GetShaderInfoLog(shader, logLength, nil, gl.Str(log))
		return 0, fmt.Errorf("Failed to compile shader %v: %v", source, log)
	}
	return shader, nil
}

func loadShader(filename string, name string) (*Shader, error) {
	isFrag, err := regexp.MatchString(".frag", filename)
	if err != nil {
		return nil, err
	}
	isVert, err := regexp.MatchString(".vert", filename)
	if err != nil {
		return nil, err
	}
	shaderProgram := file.LoadFileAsString(filename)
	var shaderType Type
	if isFrag {
		shaderType = Fragment
	} else if isVert {
		shaderType = Vertex
	} else {
		return nil, fmt.Errorf("Unrecognized shader format in file %v", filename)
	}
	shaderID, error := compileShader(shaderProgram, uint32(shaderType))
	shader := new(Shader)
	shader.Name = name
	shader.ShaderID = shaderID
	shader.ShaderType = shaderType
	return shader, error
}

// Type enum for determining the type of the shader
type Type uint32

const (
	//Vertex shader type
	Vertex Type = gl.VERTEX_SHADER
	// Fragment shader type
	Fragment Type = gl.FRAGMENT_SHADER
)

// Shader struct for holding data about shaders
type Shader struct {
	ShaderID   uint32
	Name       string
	FileName   string
	ShaderType Type
}

// NewShader loads a shader by filename
func NewShader(filename string, name string) (*Shader, error) {
	return loadShader(filename, name)
}
