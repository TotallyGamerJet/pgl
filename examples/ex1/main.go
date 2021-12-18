package main

import (
	"fmt"
	"github.com/gotranspile/cxgo/runtime/libc"
	"github.com/veandco/go-sdl2/sdl"
	"pgl"
	"unsafe"
)

const (
	WIDTH  = 640
	HEIGHT = 480
)

var (
	Red   = pgl.Vec4{1.0, 0.0, 0.0, 0.0}
	Green = pgl.Vec4{0.0, 1.0, 0.0, 0.0}
	Blue  = pgl.Vec4{0.0, 0.0, 1.0, 0.0}
)

var (
	window *sdl.Window
	ren    *sdl.Renderer
	tex    *sdl.Texture
)

var bbufpix = (*pgl.U32)(nil)

var the_Context pgl.GlContext

type My_Uniforms struct {
	mvp_mat pgl.Mat4
	v_color pgl.Vec4
}

func main() {
	setup_context()

	var points = []float32{
		-0.5, -0.5, 0,
		0.5, -0.5, 0,
		0, 0.5, 0,
	}
	var the_uniforms My_Uniforms
	identity := pgl.Mat4{
		1, 0, 0, 0,
		0, 1, 0, 0,
		0, 0, 1, 0,
		0, 0, 0, 1,
	}

	var triangle pgl.GLuint
	pgl.GlGenBuffers(1, &triangle)
	pgl.GlBindBuffer(pgl.GL_ARRAY_BUFFER, triangle)
	pgl.GlBufferData(pgl.GL_ARRAY_BUFFER, pgl.GLsizei(unsafe.Sizeof(pgl.GLfloat(0))*9), unsafe.Pointer(&points[0]), pgl.GL_STATIC_DRAW)
	pgl.GlEnableVertexAttribArray(0)
	pgl.GlVertexAttribPointer(0, 3, pgl.GL_FLOAT, pgl.GL_FALSE, 0, 0)

	var myshader = pgl.PglCreateProgram(normal_vs, normal_fs, 0, nil, pgl.GL_FALSE)
	pgl.GlUseProgram(myshader)

	pgl.PglSetUniform(unsafe.Pointer(&the_uniforms))

	the_uniforms.v_color = Red

	libc.MemCpy(unsafe.Pointer(&the_uniforms.mvp_mat), unsafe.Pointer(&identity), int(unsafe.Sizeof(pgl.Mat4{})))

	var (
		e    sdl.Event
		quit = 0
	)

	var (
		old_time, new_time, counter,
		last_frame uint32
	)
	var frame_time float32
	_ = frame_time
	for quit == 0 {
		for {
			e = sdl.PollEvent()
			if e == nil {
				break
			}
			if e.GetType() == sdl.QUIT {
				quit = 1
			}
			/*if e.GetType() == sdl.KEYDOWN &&{
				quit = 1
			}*/
			if e.GetType() == sdl.MOUSEBUTTONDOWN {
				quit = 1
			}
		}

		new_time = sdl.GetTicks()
		frame_time = float32(new_time-last_frame) / 1000.0
		last_frame = new_time

		counter++
		if (counter % 300) == 0 {
			fmt.Printf("%d  %f FPS\n", new_time-old_time, 300000/((float32)(new_time-old_time)))
			old_time = new_time
			counter = 0
		}

		pgl.GlClearColor(0, 0, 0, 1)
		pgl.GlClear(pgl.GL_COLOR_BUFFER_BIT)
		pgl.GlDrawArrays(pgl.GL_TRIANGLES, 0, 3)

		tex.Update(nil, unsafe.Slice((*byte)(unsafe.Pointer(bbufpix)), int(HEIGHT*WIDTH*unsafe.Sizeof(pgl.U32(0)))), int(WIDTH*unsafe.Sizeof(pgl.U32(0))))
		//Render the scene
		ren.Copy(tex, nil, nil)
		ren.Present()
	}

	cleanup()
}

func normal_vs(vs_output *float32, vertex_attribs unsafe.Pointer, builtins *pgl.Shader_Builtins, uniforms unsafe.Pointer) {
	builtins.Gl_Position = pgl.Mult_mat4_vec4(*(*pgl.Mat4)(uniforms), *(*pgl.Vec4)(vertex_attribs))
}

func normal_fs(fs_input *float32, builtins *pgl.Shader_Builtins, uniforms unsafe.Pointer) {
	builtins.Gl_FragColor = (*My_Uniforms)(uniforms).v_color
}

func setup_context() {
	//SDL_SetMainReady();
	if err := sdl.Init(sdl.INIT_VIDEO); err != nil {
		panic(err)
	}
	var err error
	window, err = sdl.CreateWindow("c_ex1", 100, 100, WIDTH, HEIGHT, sdl.WINDOW_SHOWN)
	if err != nil {
		sdl.Quit()
		panic(err)
	}

	ren, _ = sdl.CreateRenderer(window, -1, sdl.RENDERER_SOFTWARE)
	tex, _ = ren.CreateTexture(sdl.PIXELFORMAT_ARGB8888, sdl.TEXTUREACCESS_STREAMING, WIDTH, HEIGHT)

	bbufpix = nil
	if pgl.Init_glContext(&the_Context, &bbufpix, WIDTH, HEIGHT, 32, 0x00FF0000, 0x0000FF00, 0x000000FF, 0xFF000000) == 0 {
		panic("failed to init context")
	}

	pgl.Set_glContext(&the_Context)
}

func cleanup() {
	pgl.Free_glContext(&the_Context)
	tex.Destroy()
	ren.Destroy()
	window.Destroy()

	sdl.Quit()
}
