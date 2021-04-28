package display

import (
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

const (
	fullBlock  = "\u25A0"
	leftArrow  = "\u25C0"
	rightArrow = "\u25B6"
)

// TraceGraph implements the tview primative for displaying traces
type TraceGraph struct {
	*tview.Box
	traces        []*Trace
	text          *tview.TextView
	selectedTrace int
	selectedLine  int
}

func NewTraceGraph(traces []*Trace, text *tview.TextView) *TraceGraph {
	return &TraceGraph{
		Box:    tview.NewBox(),
		text:   text,
		traces: traces,
	}
}

func (t *TraceGraph) Draw(screen tcell.Screen) {
	t.Box.DrawForSubclass(screen, t)
	x, y, width, _ := t.GetInnerRect()

	t.text.SetText("All delighted traces raise their hands")

	for _, trace := range t.traces {
		y = trace.Draw(screen, x, y, width)
	}
}

func (t *TraceGraph) InputHandler() func(event *tcell.EventKey, setFocus func(p tview.Primitive)) {
	return t.WrapInputHandler(func(event *tcell.EventKey, setFocus func(p tview.Primitive)) {
		switch event.Key() {
		case tcell.KeyUp:
			t.dec()
		case tcell.KeyDown:
			t.inc()
		}
	})
}

func (t *TraceGraph) inc() {
	// At the top of the list, nothing to do
	if t.selectedTrace == len(t.traces)-1 && t.selectedLine == len(t.traces[t.selectedTrace].lines)-1 {
		return
	}

	// If at end of the lines in this trace, move on to the next trace
	if t.selectedLine == len(t.traces[t.selectedTrace].lines)-1 {
		t.traces[t.selectedTrace].lines[t.selectedLine].selected = false
		t.selectedTrace += 1
		t.selectedLine = 0
		t.traces[t.selectedTrace].lines[t.selectedLine].selected = true
	} else {
		t.traces[t.selectedTrace].lines[t.selectedLine].selected = false
		t.selectedLine += 1
		t.traces[t.selectedTrace].lines[t.selectedLine].selected = true
	}
}

func (t *TraceGraph) dec() {
	// At the bottom of the list, nothing to do
	if t.selectedTrace == 0 && t.selectedLine == 0 {
		return
	}

	// If at start of the lines in this trace, move on to the previous trace
	if t.selectedLine == 0 {
		t.traces[t.selectedTrace].lines[t.selectedLine].selected = false
		t.selectedTrace -= 1
		t.selectedLine = len(t.traces[t.selectedTrace].lines) - 1
		t.traces[t.selectedTrace].lines[t.selectedLine].selected = true
	} else {
		t.traces[t.selectedTrace].lines[t.selectedLine].selected = false
		t.selectedLine -= 1
		t.traces[t.selectedTrace].lines[t.selectedLine].selected = true
	}
}

func (t *TraceGraph) MouseHandler() func(action tview.MouseAction, event *tcell.EventMouse, setFocus func(p tview.Primitive)) (consumed bool, capture tview.Primitive) {
	return t.WrapMouseHandler(func(action tview.MouseAction, event *tcell.EventMouse, setFocus func(p tview.Primitive)) (consumed bool, capture tview.Primitive) {
		return
	})
}
