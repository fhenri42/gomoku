package main

import (
	"fmt"
	//"github.com/banthar/Go-SDL/sdl"
	"github.com/veandco/go-sdl2/sdl"
	"github.com/veandco/go-sdl2/sdl_image"
)

const pushBack = 17
const pushBackB = 28
var turn = true

func getXY(src *sdl.Surface, dst *sdl.Surface, x int32, y int32) error {
	var centerRect sdl.Rect
	centerRect.X = x
	centerRect.Y = y
	return src.Blit(nil, dst, &centerRect)
}


func putImageXY(file string, dst *sdl.Surface, x int32, y int32) error {
	imgSurface, err := img.Load(file)
	if err != nil {
		return err
	}
	defer imgSurface.Free()
	return getXY(imgSurface, dst, x, y)
}

func getCenter(src *sdl.Surface, dst *sdl.Surface) error {
	var centerRect sdl.Rect
	centerRect.X = dst.ClipRect.H/2 - src.ClipRect.H/2
	centerRect.Y = dst.ClipRect.W/2 - src.ClipRect.W/2
	return src.Blit(nil, dst, &centerRect)
}

func putImageCenter(file string, dst *sdl.Surface) error {
	imgSurface, err := img.Load(file)
	if err != nil {
		return err
	}
	defer imgSurface.Free()
	return getCenter(imgSurface, dst)
}

func eventHandel(surface *sdl.Surface, win *sdl.Window) {

	var exit = false

	for !exit {
		for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
			switch t := event.(type) {
			case *sdl.QuitEvent:
				exit = true
				break
			case *sdl.MouseMotionEvent:
				//	fmt.Println(t.X, " ", t.Y)
				break
			case *sdl.KeyDownEvent:
				if t.Keysym.Sym == 27 {
					exit = true
				}
				fmt.Print(t.Keysym)
				break
			case *sdl.MouseButtonEvent:
				if t.Type == 1025 {
					finDot(t, surface, win)
					break
				}
			}

		}
	}
}

func  finDot(t *sdl.MouseButtonEvent, surface *sdl.Surface, win *sdl.Window)  {
	var nbrX int32 = 166
	var x = 0
	var y = 0
	for x < 19 {
		y = 0
		var nbrY int32 = 166

		for y < 19 {
			if (t.X - 4 <= nbrX && t.Y - 4 <= nbrY && t.Y + 4 >= nbrY && t.X + 4 >= nbrX) {
				putAdot(nbrX, nbrY, surface, win)
				turn = !turn
				break
			}
			nbrY+=40
			y++
		}
		nbrX+= 40
		x++
	}
}
func putAdot(x int32, y int32, surface *sdl.Surface,win *sdl.Window) {

	if turn {
		errXY := putImageXY("ressources/red-dot.bmp",surface, x - pushBack, y - pushBack)
		if errXY != nil {
			fmt.Println("errXY", errXY)
		}
		} else  {
			errXY := putImageXY("ressources/black-dot.bmp",surface, x - pushBackB, y - pushBackB)
			if errXY != nil {
				fmt.Println("errXY", errXY)
			}
		}
		win.UpdateSurface()

	}

	func intGraphics() {
		fmt.Print("DEEEEEE")
		win, err := sdl.CreateWindow("Gomokou", sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED, 800, 800, sdl.WINDOW_SHOWN)

		if err != nil {
			panic("Video mode set failed: ")
			sdl.Quit()
		}

		surface, err1 := win.GetSurface()
		fmt.Println(err1)

		errCenter := putImageCenter("ressources/board.bmp", surface)
		if errCenter != nil {
			fmt.Println("err imageCenter", errCenter)
		}
		
		win.UpdateSurface()
		eventHandel(surface, win)
		sdl.Quit();

		return;
	}
