package shader

import (
	"fmt"
	"strings"

	"github.com/go-gl/gl/v4.1-core/gl"
)

// Program struct for encapsulating openGL programs
type Program struct {
	ProgramID uint32
	Name      string
}

// NewProgram creates a new program struct
func NewProgram(name string, shaders []*Shader) (*Program, error) {
	program := new(Program)
	program.ProgramID = gl.CreateProgram()
	program.Name = name
	for _, shader := range shaders {
		gl.AttachShader(program.ProgramID, shader.ShaderID)
	}
	gl.LinkProgram(program.ProgramID)

	var status int32
	gl.GetProgramiv(program.ProgramID, gl.LINK_STATUS, &status)
	if status == gl.FALSE {
		var logLength int32
		gl.GetProgramiv(program.ProgramID, gl.INFO_LOG_LENGTH, &logLength)
		log := strings.Repeat("\x00", int(logLength+1))
		gl.GetProgramInfoLog(program.ProgramID, logLength, nil, gl.Str(log))
		return nil, fmt.Errorf("failed to link program: %v", log)
	}
	for _, shader := range shaders {
		gl.DeleteShader(shader.ShaderID)
	}
	return program, nil
}
