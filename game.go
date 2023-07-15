package main

import (
	"time"

	"github.com/gdamore/tcell/v2"
)

type Game struct {
	Screen tcell.Screen
}

func Run(screen tcell.Screen, defStyle tcell.Style) {
	y := 0
	x := 55
	for {
		screen.Clear()
		screen.SetContent(x, y+0, 'H', nil, defStyle)
		screen.SetContent(x+1, y+0, 'I', nil, defStyle)
		screen.SetContent(x+2, y+0, ' ', nil, defStyle)
		screen.SetContent(x+3, y+0, ' ', nil, defStyle)
		screen.SetContent(x+4, y+0, 'W', nil, defStyle)
		screen.SetContent(x+5, y+0, 'E', nil, defStyle)
		screen.SetContent(x+6, y+0, 'L', nil, defStyle)
		screen.SetContent(x+7, y+0, 'C', nil, defStyle)
		screen.SetContent(x+8, y+0, 'O', nil, defStyle)
		screen.SetContent(x+9, y+0, 'M', nil, defStyle)
		screen.SetContent(x+10, y+0, 'E', nil, defStyle)
		screen.SetContent(x+11, y+0, ' ', nil, defStyle)
		screen.SetContent(x+12, y+0, 'T', nil, defStyle)
		screen.SetContent(x+13, y+0, 'O', nil, defStyle)
		screen.SetContent(x+14, y+0, ' ', nil, defStyle)
		screen.SetContent(x+15, y+0, 'P', nil, defStyle)
		screen.SetContent(x+16, y+0, 'O', nil, defStyle)
		screen.SetContent(x+17, y+0, 'N', nil, defStyle)
		screen.SetContent(x+18, y+0, 'G', nil, defStyle)
		screen.SetContent(x+19, y+0, ' ', nil, defStyle)
		screen.SetContent(x+20, y+0, '!', nil, defStyle)

		screen.Show()

		time.Sleep(140 * time.Millisecond)

		y++
		if y == 15 {

			break
		}
	}
}
