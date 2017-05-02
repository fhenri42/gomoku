package main

import (
  "fmt"
	"github.com/veandco/go-sdl2/sdl"
	"github.com/veandco/go-sdl2/sdl_image"
)

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


func  finDot(t *sdl.MouseButtonEvent, tools *sdlTools)  {
	var nbrX int32 = 21
	var x = 0
	var y = 0

	for x < 19 {
		y = 0
		var nbrY int32 = 22

		for y < 19 {
			if ((t.X - AVREGEREPAIR <= nbrX && t.Y - AVREGEREPAIR <= nbrY && t.Y + AVREGEREPAIR >= nbrY && t.X + AVREGEREPAIR >= nbrX) && tools.board[x][y] == 0) {

        putAdot(nbrX, nbrY, tools)
				if tools.turn {
					tools.board[x][y] = 1
				} else  {
					tools.board[x][y] = 2
				}
			tools.turn = !tools.turn
				break
			}
			nbrY+=53
			y++
		}
		nbrX+= 53
		x++
	}
}

func putAdot(x int32, y int32, tools *sdlTools) {

	if tools.turn {
		errXY := putImageXY("ressources/red-dot.bmp",tools.surface, x - PUSHBACK, y - PUSHBACK)
		if errXY != nil {
			fmt.Println("errXY", errXY)
		}
		} else  {
			errXY := putImageXY("ressources/black-dot.bmp",tools.surface, x - PUSHBACKB, y - PUSHBACKB)
			if errXY != nil {
				fmt.Println("errXY", errXY)
			}
		}
		tools.win.UpdateSurface()
	}
