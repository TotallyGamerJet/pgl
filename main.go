package main

import (
	"fmt"
	"github.com/gotranspile/cxgo/runtime/libc"
	"github.com/veandco/go-sdl2/sdl"
	"unsafe"
)

const (
	WIDTH  = 640
	HEIGHT = 480
)

var (
	Red   = vec4{1.0, 0.0, 0.0, 0.0}
	Green = vec4{0.0, 1.0, 0.0, 0.0}
	Blue  = vec4{0.0, 0.0, 1.0, 0.0}
)

var (
	window *sdl.Window
	ren    *sdl.Renderer
	tex    *sdl.Texture
)

var bbufpix = (*u32)(nil)

var the_Context glContext

type My_Uniforms struct {
	mvp_mat mat4
	v_color vec4
}

func main() {
	setup_context()

	var points = []float32{
		-0.5, -0.5, 0,
		0.5, -0.5, 0,
		0, 0.5, 0,
	}
	var the_uniforms My_Uniforms
	identity := mat4{1, 0, 0, 0, 0, 1, 0, 0, 0, 0, 1, 0, 0, 0, 0, 1}

	var triangle GLuint
	glGenBuffers(1, &triangle)
	glBindBuffer(GL_ARRAY_BUFFER, triangle)
	glBufferData(GL_ARRAY_BUFFER, GLsizei(unsafe.Sizeof(GLfloat(0))*9), unsafe.Pointer(&points[0]), GL_STATIC_DRAW)
	glEnableVertexAttribArray(0)
	glVertexAttribPointer(0, 3, GL_FLOAT, GL_FALSE, 0, 0)

	var myshader = pglCreateProgram(normal_vs, normal_fs, 0, nil, GL_FALSE)
	glUseProgram(myshader)

	pglSetUniform(unsafe.Pointer(&the_uniforms))

	the_uniforms.v_color = Red

	libc.MemCpy(unsafe.Pointer(&the_uniforms.mvp_mat), unsafe.Pointer(&identity), int(unsafe.Sizeof(mat4{})))

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

		glClearColor(0, 0, 0, 1)
		glClear(GL_COLOR_BUFFER_BIT)
		glDrawArrays(GL_TRIANGLES, 0, 3)

		tex.Update(nil, unsafe.Slice((*byte)(unsafe.Pointer(bbufpix)), int(HEIGHT*WIDTH*unsafe.Sizeof(u32(0)))), int(WIDTH*unsafe.Sizeof(u32(0))))
		//Render the scene
		ren.Copy(tex, nil, nil)
		ren.Present()
	}

	cleanup()
}

func normal_vs(vs_output *float32, vertex_attribs unsafe.Pointer, builtins *Shader_Builtins, uniforms unsafe.Pointer) {
	builtins.Gl_Position = mult_mat4_vec4(*(*mat4)(uniforms), unsafe.Slice((*vec4)(vertex_attribs), 1)[0])
}

func normal_fs(fs_input *float32, builtins *Shader_Builtins, uniforms unsafe.Pointer) {
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
	if init_glContext(&the_Context, &bbufpix, WIDTH, HEIGHT, 32, 0x00FF0000, 0x0000FF00, 0x000000FF, 0xFF000000) == 0 {
		panic("failed to init context")
	}

	set_glContext(&the_Context)
}

func cleanup() {
	free_glContext(&the_Context)
	tex.Destroy()
	ren.Destroy()
	window.Destroy()

	sdl.Quit()
}
