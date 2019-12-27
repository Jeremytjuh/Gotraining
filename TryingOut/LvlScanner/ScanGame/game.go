package main

import (
	"bufio"
	tl "github.com/JoelOtter/termloop"
	"log"
	"os"
	"strings"
)

// Level bruv
type Level struct {
	*tl.Entity
}

// type Player struct {
// 	*tl.Entity
// 	P Coordinates
// }

// type Coordinates struct {
// 	X int
// 	Y int
// }

// var reserved = map[string]string{
// 	"#": "Wall",
// 	"G": "Golden door",
// 	"K": "Key",
// }

func main() {
	game := tl.NewGame()
	level := tl.NewBaseLevel(tl.Cell{
		Bg: tl.ColorCyan,
	})

	leveltje := new(Level)
	leveltje.Entity = tl.NewEntity(1, 1, 1, 1)

	level.AddEntity(leveltje)
	game.Screen().SetLevel(level)
	game.Start()
}

// Draw draw
func (level *Level) Draw(screen *tl.Screen) {
	file, err := os.Open("level.csv")
	if err != nil {
		log.Fatalln(file)
	}
	scanner := bufio.NewScanner(file)
	var startx = 0 // Determines at which coordinate the rendering starts
	var starty = 0 // Determines at which coordinate the rendering starts
	for scanner.Scan() {
		xline := strings.Split(scanner.Text(), "")
		for i, v := range xline {
			switch v {
			case "#":
				screen.RenderCell(startx+i+1, starty, &tl.Cell{
					Bg: tl.ColorWhite,
				})
			case " ":
				screen.RenderCell(startx+i+1, starty, &tl.Cell{
					Bg: tl.ColorBlack,
				})
			case "G":
				screen.RenderCell(startx+i+1, starty, &tl.Cell{
					Bg: tl.ColorYellow,
				})
			}
		}
		starty++
	}
}
