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
	// Node View
	if nodesView, err := g.SetView("nodes", 0, 0, maxX-1, maxY/3-1); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		nodesView.Title = "nodes"
		nodesView.FgColor = gocui.ColorWhite
	}

	if nodesName, err := g.SetView("nodeName", 0, 0, maxX/4, maxY/3-1); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		nodesName.Title = "NAME"
		nodesName.FgColor = gocui.ColorWhite
		nodesName.Frame = false
		// Print out read text
		fmt.Fprintln(nodesName, "NAME")
		fmt.Fprintln(nodesName, "ip-10-0-26-43.ec2.internal")
		fmt.Fprintln(nodesName, "ip-10-0-26-43.ec2.internal")
		fmt.Fprintln(nodesName, "ip-10-0-26-43.ec2.internal")
		fmt.Fprintln(nodesName, "ip-10-0-26-43.ec2.internal")
	}

	if nodesStatus, err := g.SetView("nodeStatus", maxX/4, 0, maxX*2/4, maxY/3-1); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		nodesStatus.Title = "STATUS"
		nodesStatus.FgColor = gocui.ColorWhite
		nodesStatus.Frame = false
		// Print out read text
		fmt.Fprintln(nodesStatus, "STATUS")
		fmt.Fprintln(nodesStatus, "Ready")
		fmt.Fprintln(nodesStatus, "Ready")
		fmt.Fprintln(nodesStatus, "Ready")
		fmt.Fprintln(nodesStatus, "Ready")
	}

	if nodesCpu, err := g.SetView("nodeCpu", maxX*2/4, 0, maxX*3/4, maxY/3-1); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		nodesCpu.Title = "CPU"
		nodesCpu.FgColor = gocui.ColorWhite
		nodesCpu.Frame = false
		// Print out read text
		fmt.Fprintln(nodesCpu, "CPU")
		fmt.Fprintln(nodesCpu, "100%")
		fmt.Fprintln(nodesCpu, "100%")
		fmt.Fprintln(nodesCpu, "100%")
		fmt.Fprintln(nodesCpu, "100%")
	}

	if nodesMem, err := g.SetView("nodeMem", maxX*3/4, 0, maxX-1, maxY/3-1); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		nodesMem.Title = "MEM"
		nodesMem.FgColor = gocui.ColorWhite
		nodesMem.Frame = false
		// Print out read text
		fmt.Fprintln(nodesMem, "MEM")
		fmt.Fprintln(nodesMem, "100%")
		fmt.Fprintln(nodesMem, "100%")
		fmt.Fprintln(nodesMem, "100%")
		fmt.Fprintln(nodesMem, "100%")
	}

	// Namespace View

	if nsView, err := g.SetView("namespaces", 0, maxY/3, maxX-1, maxY*2/3-1); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		nsView.Title = "namespaces"
		nsView.FgColor = gocui.ColorWhite
	}

	if nsName, err := g.SetView("nsName", 0, maxY/3, maxX/4, maxY*2/3-1); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		nsName.Title = "NAMESPACE"
		nsName.FgColor = gocui.ColorWhite
		nsName.Frame = false
		// Print out read text
		fmt.Fprintln(nsName, "NAME")
		fmt.Fprintln(nsName, "atom8-prod")
		fmt.Fprintln(nsName, "saas-prod")
		fmt.Fprintln(nsName, "backorder-prod")
		fmt.Fprintln(nsName, "thirdparty-prod")
	}

	if nsStatus, err := g.SetView("nsStatus", maxX/4, maxY/3, maxX*2/4, maxY*2/3-1); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		nsStatus.Title = "STATUS"
		nsStatus.FgColor = gocui.ColorWhite
		nsStatus.Frame = false
		// Print out read text
		fmt.Fprintln(nsStatus, "STATUS")
		fmt.Fprintln(nsStatus, "Active")
		fmt.Fprintln(nsStatus, "Active")
		fmt.Fprintln(nsStatus, "Active")
		fmt.Fprintln(nsStatus, "Active")
	}

	if nsCpu, err := g.SetView("nsCpu", maxX*2/4, maxY/3, maxX*3/4, maxY*2/3-1); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		nsCpu.Title = "CPU"
		nsCpu.FgColor = gocui.ColorWhite
		nsCpu.Frame = false
		// Print out read text
		fmt.Fprintln(nsCpu, "CPU")
		fmt.Fprintln(nsCpu, "100%")
		fmt.Fprintln(nsCpu, "100%")
		fmt.Fprintln(nsCpu, "100%")
		fmt.Fprintln(nsCpu, "100%")
	}

	if nsMem, err := g.SetView("nsMem", maxX*3/4, maxY/3, maxX-1, maxY*2/3-1); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		nsMem.Title = "MEM"
		nsMem.FgColor = gocui.ColorWhite
		nsMem.Frame = false
		// Print out read text
		fmt.Fprintln(nsMem, "MEM")
		fmt.Fprintln(nsMem, "100%")
		fmt.Fprintln(nsMem, "100%")
		fmt.Fprintln(nsMem, "100%")
		fmt.Fprintln(nsMem, "100%")
	}

	// Pod View

	if podsView, err := g.SetView("pods", 0, maxY*2/3, maxX-1, maxY-1); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		podsView.Title = "pods"
		// Print out read text
	}

	if podName, err := g.SetView("podName", 0, maxY*2/3, maxX/4, maxY-1); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		podName.Title = "NAME"
		podName.FgColor = gocui.ColorWhite
		podName.Frame = false
		// Print out read text
		fmt.Fprintln(podName, "NAME")
		fmt.Fprintln(podName, "pod123")
		fmt.Fprintln(podName, "pod123")
		fmt.Fprintln(podName, "pod123")
		fmt.Fprintln(podName, "pod123")
	}

	if podNs, err := g.SetView("podNs", maxX/4, maxY*2/3, maxX*2/4-1, maxY-1); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		podNs.Title = "NAMESPACE"
		podNs.FgColor = gocui.ColorWhite
		podNs.Frame = false
		// Print out read text
		fmt.Fprintln(podNs, "NAMESPACE")
		fmt.Fprintln(podNs, "atom-prod")
		fmt.Fprintln(podNs, "atom-prod")
		fmt.Fprintln(podNs, "saas-prod")
		fmt.Fprintln(podNs, "saas-prod")
	}

	if podCpu, err := g.SetView("podCpu", maxX*2/4, maxY*2/3, maxX*3/4-1, maxY-1); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		podCpu.Title = "CPU"
		podCpu.FgColor = gocui.ColorWhite
		podCpu.Frame = false
		// Print out read text
		fmt.Fprintln(podCpu, "CPU")
		fmt.Fprintln(podCpu, "100%")
		fmt.Fprintln(podCpu, "100%")
		fmt.Fprintln(podCpu, "100%")
		fmt.Fprintln(podCpu, "100%")
	}

	if podMem, err := g.SetView("nsMem", maxX*3/4, maxY*2/3, maxX-1, maxY-1); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		podMem.Title = "MEM"
		podMem.FgColor = gocui.ColorWhite
		podMem.Frame = false
		// Print out read text
		fmt.Fprintln(podMem, "MEM")
		fmt.Fprintln(podMem, "100%")
		fmt.Fprintln(podMem, "100%")
		fmt.Fprintln(podMem, "100%")
		fmt.Fprintln(podMem, "100%")
	}

	return nil
}

func quit(g *gocui.Gui, v *gocui.View) error {
	return gocui.ErrQuit
}