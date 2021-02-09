package main

import (
	"fmt"
	"github.com/jroimartin/gocui"
	"log"
)

func main() {
	g, err := gocui.NewGui(gocui.OutputNormal)
	if err != nil {
		log.Panicln(err)
	}
	defer g.Close()

	g.SetManagerFunc(layout)

	if err := g.SetKeybinding("", gocui.KeyCtrlC, gocui.ModNone, quit); err != nil {
		log.Panicln(err)
	}

	if err := g.MainLoop(); err != nil && err != gocui.ErrQuit {
		log.Panicln(err)
	}

}

func layout(g *gocui.Gui) error {
	g.BgColor = gocui.ColorBlack
	g.FgColor = gocui.ColorGreen
	maxX, maxY := g.Size()
	if nodesView, err := g.SetView("nodes", 0, 0, maxX-1, maxY/3-1); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		nodesView.Title = "Nodes"
		nodesView.FgColor = gocui.ColorWhite
		//nodesView.Frame = false // get rid of border
		// Print out read text
		fmt.Fprintln(nodesView, "This is node view")
	}

	if namespacesView, err := g.SetView("namespaces", 0, maxY/3+1, maxX-1, maxY*2/3-1); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		namespacesView.Title = "NameSpaces"
		// Print out read text
		fmt.Fprintln(namespacesView, "This is ns view")
	}

	if podsView, err := g.SetView("pods", 0, maxY*2/3+1, maxX-1, maxY-1); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		podsView.Title = "Pods"
		// Print out read text
		fmt.Fprintln(podsView, "This is pod view")
	}
	return nil
}

func quit(g *gocui.Gui, v *gocui.View) error {
	return gocui.ErrQuit
}