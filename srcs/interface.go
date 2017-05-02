package main

import (
	"fmt"
	"github.com/veandco/go-sdl2/sdl"
)

<<<<<<< HEAD:srcs/initInterphace.go
func eventHandel(tools *sdlTools) {
=======
func eventHandel(surface *sdl.Surface, win *sdl.Window, board*[SIZE][SIZE]int) {
>>>>>>> 1262a0e11aa90960ad10e68f541246644d3ea678:srcs/interface.go

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

<<<<<<< HEAD:srcs/initInterphace.go
func initSdl(tools *sdlTools) {
=======
	func initSdl(board*[SIZE][SIZE]int) {
		win, err := sdl.CreateWindow("Gomokou", sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED, W, H, sdl.WINDOW_SHOWN)

		if err != nil {
		fmt.Println("err CreateWindow", err)
			sdl.Quit()
		}
>>>>>>> 1262a0e11aa90960ad10e68f541246644d3ea678:srcs/interface.go

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
