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
	defStyle := tcell.StyleDefault.Background(tcell.ColorPaleTurquoise).Foreground(tcell.ColorBlack)
	screen.SetStyle(defStyle)

	go Run(screen, defStyle)

	for {
		switch ev := screen.PollEvent().(type) {
		case *tcell.EventResize:
			screen.Sync()
		case *tcell.EventKey:
			if ev.Key() == tcell.KeyEscape || ev.Key() == tcell.KeyCtrlC {
				screen.Fini()
				os.Exit(0)
			}

		}
	}
}
