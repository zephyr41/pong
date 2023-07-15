package main

import (
	"log"
	"os"

	"github.com/gdamore/tcell/v2" // permet de crée une interface utlisateur
)

func main() {

	screen, err := tcell.NewScreen()
	if err != nil {
		log.Fatalf("could not create screen: %v", err) // gére les erreurs
	}
	if err := screen.Init(); err != nil {
		log.Fatalf("could not init screen: %v", err)
	}
	defStyle := tcell.StyleDefault.Background(tcell.ColorBlack).Foreground(tcell.ColorBlack)
	screen.SetStyle(defStyle)
	x := 0
	for {

		screen.SetContent(x, 10, 'H', nil, defStyle)
		screen.SetContent(x+1, 10, 'i', nil, defStyle)
		screen.SetContent(x+2, 10, '!', nil, defStyle)

		screen.Show()
		x++
		switch event := screen.PollEvent().(type) {
		case *tcell.EventResize:
			screen.Sync()
		case *tcell.EventKey:
			if event.Key() == tcell.KeyEscape || event.Key() == tcell.KeyCtrlC {
				screen.Fini()
				os.Exit(0)
			}
		}

	}

}
