package main

import (
	"fmt"
	"github.com/veandco/go-sdl2/sdl"
)

func eventHandel(tools *sdlTools) {

	for !tools.exit {
		for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
			switch t := event.(type) {
			case *sdl.QuitEvent:
				tools.exit = true
				break
			case *sdl.MouseMotionEvent:
				//fmt.Println(t.X, " ", t.Y)
				break
			case *sdl.KeyDownEvent:
				if t.Keysym.Sym == 27 {
					tools.exit = true
				}
				fmt.Print(t.Keysym)
				break
			case *sdl.MouseButtonEvent:
				if t.Type == 1025 {
					finDot(t, tools)
					break
				}
			}
		}
	}
}


func initSdl(tools *sdlTools) {

	tools.win,_ = sdl.CreateWindow("Gomokou", sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED, W, H, sdl.WINDOW_SHOWN)
	tools.surface, _ = tools.win.GetSurface()

	errCenter := putImageCenter("ressources/boardgo.bmp", tools.surface)
	if errCenter != nil {
		fmt.Println("err imageCenter", errCenter)
	}

	tools.win.UpdateSurface()
	eventHandel(tools)
	sdl.Quit();

	return;
}
