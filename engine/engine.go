package engine

import (
	"diabgo/io"
	"diabgo/renderer"
	"diabgo/world"
	"github.com/veandco/go-sdl2/sdl"
)

type Engine struct {
	settings *Settings
	renderer *renderer.Renderer
	input *io.Input
	output *io.Output
	world *world.World
	tick uint64
	running bool
}


func (engine Engine) Run(){

	engine.running = true
	engine.renderer = renderer.NewRenderer("my game!", 300, 300)
	engine.world = world.NewWorld()
	for engine.running{

		for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
			switch event.(type) {
				case *sdl.QuitEvent:
					engine.Stop()
					break
			}
		}

		engine.world.Update()
		engine.renderer.Update()
		engine.tick++
	}


	engine.renderer.Destroy()

}
func (engine Engine) Stop(){
	engine.running = false
}
func (engine Engine) Pause(){

}
func (engine Engine) StartGame(){

}
func (engine Engine) StartFromSave(){

}
func (engine Engine) debug(msg string){

}