package main

import (
	"fmt"
	"github.com/veandco/go-sdl2/sdl"
	"github.com/veandco/go-sdl2/sdl_image"
)

func printBoard(tools *sdlTools) {
  tools.surface.Free()
  loadMap(tools)
  var i int32 = 0
  var j int32 = 0

  for i < SIZE {
    j = 0
    for j < SIZE {
      if (tools.board[i][j] != 0) {
        putAdot(i, j, tools, tools.board[i][j])
      }
      j++
    }
    i++
  }

  tools.win.UpdateSurface()
}

func loadMap(tools *sdlTools)  {
	errCenter := putImageCenter("ressources/boardgo.bmp", tools.surface)
	if errCenter != nil {
		fmt.Println("err imageCenter", errCenter)
	}
	tools.win.UpdateSurface()
}

func putAdot(i int32, j int32, tools *sdlTools, player int) {
  var file string

	if player == PLAYER1 {
    file = "ressources/red-dot.bmp"
  } else {
    file = "ressources/black-dot.bmp"
  }
	errXY := putImageXY(file, tools.surface, j * SQUARE + OFFSET - SQUARE / 4, i * SQUARE + OFFSET - SQUARE / 4)
  if errXY != nil {
    fmt.Println("errXY", errXY)
  }
}

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
