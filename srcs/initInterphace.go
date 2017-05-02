package main

import (
	"fmt"
	"github.com/veandco/go-sdl2/sdl"
)

func eventHandel(surface *sdl.Surface, win *sdl.Window, board*[19][19]int) {

	var exit = false

	for !exit {
		for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
			switch t := event.(type) {
			case *sdl.QuitEvent:
				exit = true
				break
			case *sdl.MouseMotionEvent:
					//fmt.Println(t.X, " ", t.Y)
				break
			case *sdl.KeyDownEvent:
				if t.Keysym.Sym == 27 {
					exit = true
				}
				fmt.Print(t.Keysym)
				break
			case *sdl.MouseButtonEvent:
				if t.Type == 1025 {
					finDot(t, surface, win, board)
					break
				}
			}

		}
	}
}

	func initSdl(board*[19][19]int) {
		win, err := sdl.CreateWindow("Gomokou", sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED, W, H, sdl.WINDOW_SHOWN)

		if err != nil {
		fmt.Println("err CreateWindow", err)
			sdl.Quit()
		}

		surface, err1 := win.GetSurface()
		fmt.Println(err1)

		errCenter := putImageCenter("ressources/boardgo.bmp", surface)
		if errCenter != nil {
			fmt.Println("err imageCenter", errCenter)
		}

		win.UpdateSurface()
		eventHandel(surface, win, board)
		sdl.Quit();

		return;
	}
