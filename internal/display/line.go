package display

import (
	"fmt"
	"strings"
	"time"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

type Line struct {
	node      *SnapshotNode
	rootStart time.Time
	selected  bool
	writer    *tview.TextView
}

func NewLine(node *SnapshotNode, rootStart time.Time, writer *tview.TextView) *Line {
	return &Line{
		node:      node,
		rootStart: rootStart,
		writer:    writer,
		selected:  false,
	}
}

func (l *Line) Draw(screen tcell.Screen, x, y, width, unit int) int {
	if l.selected {
		l.writeText(screen)
	}

	startLag := l.node.SpanSnapshot.StartTime.Sub(l.rootStart)
	duration := l.node.SpanSnapshot.EndTime.Sub(l.node.SpanSnapshot.StartTime)

	startGap := strings.Repeat(" ", int(startLag)/unit)

	blocks := strings.Repeat(fullBlock, (int(duration)/unit)-2)

	var bg string
	if l.selected {
		bg = "[:white]"
	}
	line := fmt.Sprintf(`%s%s%s%s%s`, startGap, bg, leftArrow, blocks, rightArrow)
	tview.Print(screen, line, x, y, width, tview.AlignLeft, tcell.ColorBlue)

	line = fmt.Sprintf(` %s%s%s - %dus`, startGap, bg, l.node.SpanSnapshot.Name, duration.Microseconds())
	color := tcell.ColorWhite
	if l.selected {
		color = tcell.ColorBlack
	}
	tview.Print(screen, line, x, y+1, width, tview.AlignLeft, color)

	return y + 2
}

func (l *Line) writeText(screen tcell.Screen) {
	s := l.node.SpanSnapshot.Name + " - " + l.node.SpanSnapshot.SpanContext.TraceID + ": " + l.node.SpanSnapshot.SpanContext.SpanID + "\n"

	for _, attribute := range l.node.SpanSnapshot.Attributes {
		s += fmt.Sprintf("%s = %v\n", attribute.Key, attribute.Value.Value)
	}

	l.writer.SetText(s)
	l.writer.ScrollToBeginning()
	l.writer.Draw(screen)
}
