package display

import (
	"fmt"
	"time"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

type Trace struct {
	start time.Time
	end   time.Time
	lines []*Line
}

func NewTrace(start time.Time, end time.Time, lines []*Line) *Trace {
	return &Trace{
		start: start,
		end:   end,
		lines: lines,
	}
}

func (t *Trace) Draw(screen tcell.Screen, x, y, width int) int {
	duration := t.end.Sub(t.start)

	// Calculate the scale of the trace line
	unit := int(duration) / width

	line := fmt.Sprintf("%s", t.start.String())
	tview.Print(screen, line, x, y, width, tview.AlignLeft, tcell.ColorWhite)
	line = fmt.Sprintf("%s", t.end.String())
	tview.Print(screen, line, x, y, width, tview.AlignRight, tcell.ColorWhite)

	y = y + 1

	for _, line := range t.lines {
		y = line.Draw(screen, x, y, width, unit)
	}

	return y + 1
}
