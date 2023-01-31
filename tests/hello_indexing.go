package tests

import (
	"github.com/totallygamerjet/pgl"
	"testing"
	"unsafe"
)

func hello_indexing(_ *testing.T) {
	var points = [...]float32{
		-0.5, 0.5, 0.0, // top left
		-0.5, -0.5, 0.0, // bottom left
		0.5, 0.5, 0.0, // top right
		0.5, -0.5, 0.0, // bottom right
	}

	// not that it matters here, but using CCW
	var indices = [...]uint32{
		0, 1, 2,
		2, 1, 3,
	}

	// using default VAO 0, already active (like compatibility profile)

	var square, elements pgl.GLuint
	pgl.GenBuffers(1, &square)
	pgl.BindBuffer(pgl.ARRAY_BUFFER, square)
	pgl.BufferData(pgl.ARRAY_BUFFER, pgl.GLsizei(unsafe.Sizeof(points)), unsafe.Pointer(&points[0]), pgl.STATIC_DRAW)

	pgl.GenBuffers(1, &elements)
	pgl.BindBuffer(pgl.ELEMENT_ARRAY_BUFFER, elements)
	pgl.BufferData(pgl.ELEMENT_ARRAY_BUFFER, pgl.GLsizei(unsafe.Sizeof(indices)), unsafe.Pointer(&indices[0]), pgl.STATIC_DRAW)

	pgl.EnableVertexAttribArray(0)
	pgl.VertexAttribPointer(0, 3, pgl.FLOAT, pgl.FALSE, 0, 0)

	// using default shader 0, already active

	pgl.ClearColor(0, 0, 0, 1)
	pgl.Clear(pgl.COLOR_BUFFER_BIT)
	pgl.DrawElements(pgl.TRIANGLES, 6, pgl.UNSIGNED_INT, 0)
}
