package main

import (
	"fmt"
	"os"
	"time"
	"github.com/veandco/go-sdl2/sdl"
	"github.com/veandco/go-sdl2/sdl_ttf"
)
func  displayTime(tools *sdlTools)  {
	ttf.Init()
	var clr sdl.Color
	clr.R = 0
	clr.G = 0
	clr.B = 0

	var rect  sdl.Rect
	font, _:= ttf.OpenFont("ressources/Zalight.ttf", 42)
	text, _:= font.RenderUTF8_Solid(tools.time.String(),clr)
	defer text.Free()

	rect.X = 100
	rect.Y = 120
	text.Blit(nil, tools.surface, &rect);
	tools.win.UpdateSurface()
}

func initMessage(tools *sdlTools) sdl.MessageBoxData  {

	var msg sdl.MessageBoxData
	var buttonArray = make([]sdl.MessageBoxButtonData, 2)
	buttonArray[0].Flags = sdl.MESSAGEBOX_BUTTON_RETURNKEY_DEFAULT
	buttonArray[0].ButtonId = 1
	buttonArray[0].Text = "Replay"
	buttonArray[1].Flags = sdl.MESSAGEBOX_BUTTON_ESCAPEKEY_DEFAULT
	buttonArray[1].ButtonId = 2
	buttonArray[1].Text = "Quit"
	var colorBox *sdl.MessageBoxColorScheme = new(sdl.MessageBoxColorScheme)

	colorBox.Colors[0].R = 255
	colorBox.Colors[0].G = 0
	colorBox.Colors[0].B = 0

	colorBox.Colors[1].R = 255
	colorBox.Colors[1].G = 0
	colorBox.Colors[1].B = 0

	colorBox.Colors[2].R = 255
	colorBox.Colors[2].G = 0
	colorBox.Colors[2].B = 0

	colorBox.Colors[3].R = 255
	colorBox.Colors[3].G = 0
	colorBox.Colors[3].B = 0

	colorBox.Colors[4].R = 255
	colorBox.Colors[4].G = 0
	colorBox.Colors[4].B = 0

	msg.Flags = sdl.MESSAGEBOX_INFORMATION
	msg.Window = tools.win
	msg.Title = "End of game"
	msg.NumButtons = 2
	msg.Buttons = buttonArray
	msg.ColorScheme = colorBox
	return msg
}

func playAgain(tools *sdlTools, winner int) {

		msg := initMessage(tools)
	if winner == 1 {
		msg.Message = "PLAYER 1 win"
	} else {
		msg.Message = "PLAYER 2 win"
	}

	_, key := sdl.ShowMessageBox(&msg)


	if (key == 2) {
		sdl.Quit()
		os.Exit(1)
	}

	err := putImageCenter("ressources/menu.bmp", tools.surface)
	if err != nil {
		fmt.Println("err", err)
	}
	tools.win.UpdateSurface()
	initSdlTools(tools)
}

func play(tools *sdlTools, i int, j int) {
	var isEnd bool

  tools.board, isEnd = moveAndEat(tools.board, i, j, tools.player, &(tools.scorePlayer2), &(tools.scorePlayer1))
	tools.player = tools.player % 2 + 1
  printBoard(tools)
	if (isEnd) {
		playAgain(tools, tools.player % 2 + 1)
	}
}

func  onClic(t *sdl.MouseButtonEvent, tools *sdlTools)  {
	var j int = (int(t.X) + SQUARE / 2 - OFFSET_X) / (SQUARE + SPACING)
	var moduloj = (int(t.X) + SQUARE / 2 - OFFSET_X) % (SQUARE + SPACING)
	var i int = (int(t.Y) + SQUARE / 2 - OFFSET_Y) / (SQUARE + SPACING)
	var moduloi = (int(t.Y) + SQUARE / 2 - OFFSET_Y) % (SQUARE + SPACING)

	if (moduloj > 0 && moduloi > 0 && isPlayable(tools, i, j)) {
    play(tools, i, j)
		if (tools.gameType == SOLO) {
			tools.wait = true
			timeBfore := time.Now()
			bestMove := getNextMove(tools.board, tools.scorePlayer1, tools.scorePlayer2, tools.player, 0)
			timeAfter := time.Now()
			tools.time = timeAfter.Sub(timeBfore)
			play(tools, bestMove.x, bestMove.y)
			displayTime(tools)
			tools.wait = false
		}
	}
}

func handleEvent(tools *sdlTools) {

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
				break
			case *sdl.MouseButtonEvent:
				if t.Type == 1025 && tools.gameState && !tools.wait {
					onClic(t, tools)
					break
				} else if (t.Type == 1025 && t.X  <= 740 && t.Y  <= 625 && t.Y  >= 492 && t.X >= 466 && !tools.gameState) {
					tools.gameState = true
					tools.gameType = 2
					loadMap(tools, "ressources/board.bmp")
				} else if (t.Type == 1025 && t.X  <= 740 && t.Y  <= 822 && t.Y  >= 688 && t.X >= 466 && !tools.gameState) {
					tools.gameState = true
					tools.gameType = 1
					loadMap(tools, "ressources/board.bmp")
				} else if (t.Type == 1025 && t.X  <= 510 && t.Y  <= 921 && t.Y  >= 898 && t.X >= 485 && !tools.gameState) {
					if (tools.iaStart) {
						loadMenu(tools, "ressources/menu.bmp")
					} else {
						loadMenu(tools, "ressources/menu1.bmp")
					}
					tools.iaStart = !tools.iaStart
			 }
			break
			}
		}
	}
}
