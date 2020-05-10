package renderer

import (
	"github.com/veandco/go-sdl2/sdl"
)

type Renderer struct {
	Width, Height int32
	window *sdl.Window
	surface *sdl.Surface
}

func (renderer Renderer) Update(){
	renderer.window.UpdateSurface()
}

func (renderer Renderer) Destroy(){
	renderer.window.Destroy()
	sdl.Quit()
}

func NewRenderer(title string, width, height int32) *Renderer {

	if err := sdl.Init(sdl.INIT_EVERYTHING); err != nil {
		panic(err)
	}

	window, err := sdl.CreateWindow(title, sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED,
		width, height, sdl.WINDOW_SHOWN)
	if err != nil {
		panic(err)
	}

	surface, err := window.GetSurface()
	if err != nil {
		panic(err)
	}

	return &Renderer{width, height, window, surface}
}