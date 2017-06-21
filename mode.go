package main

import (
	"github.com/jroimartin/gocui"
)

// List of all modes
const (
	modeInsert = iota
	modeOverwrite
	modeEscape
	modeConnect
)

var notInsert = map[int]bool{
	modeEscape: true,
}

var modeChars = []struct {
	c  rune
	fg gocui.Attribute
	bg gocui.Attribute
}{
	{' ', 0, gocui.ColorGreen},
	{'R', gocui.ColorWhite | gocui.AttrBold, gocui.ColorGreen},
	{' ', 0, gocui.ColorRed},
	{'c', gocui.ColorWhite | gocui.AttrBold, gocui.ColorRed},
}

func modeBox(g *gocui.Gui) {
	maxX, maxY := g.Size()

	for i := 0; i < maxX; i++ {
		g.SetRune(i, maxY-2, '─', gocui.ColorWhite, gocui.ColorBlack)
	}

	ch := modeChars[state.Mode]
	g.SetRune(0, maxY-1, ch.c, ch.fg, ch.bg)
	g.SetRune(1, maxY-1, ' ', gocui.ColorBlack, 0)
}
