package display

import (
	"fmt"
	"strings"
	"time"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

const (
	fullBlock  = "\u25A0"
	leftArrow  = "\u25C0"
	rightArrow = "\u25B6"
)

// Trace implements the tview primative for displaying traces
type Trace struct {
	*tview.Box
	root     *SnapshotNode
	selected int
}

func NewTrace(tree *SnapshotNode) *Trace {
	return &Trace{
		Box:  tview.NewBox(),
		root: tree,
	}
}

func (t *Trace) Draw(screen tcell.Screen) {
	t.Box.DrawForSubclass(screen, t)
	x, y, width, _ := t.GetInnerRect()

	for _, trace := range t.root.Children {
		duration := trace.SpanSnapshot.EndTime.Sub(trace.SpanSnapshot.StartTime)
		unit := int(duration) / width

		y = t.drawTrace(screen, trace, x, y, width, unit, trace.SpanSnapshot.StartTime)

		y += 1
	}
}

func (t *Trace) drawTrace(screen tcell.Screen, node *SnapshotNode, x, y, width, unit int, rootStart time.Time) int {
	startLag := node.SpanSnapshot.StartTime.Sub(rootStart)
	duration := node.SpanSnapshot.EndTime.Sub(node.SpanSnapshot.StartTime)

	startGap := strings.Repeat(" ", int(startLag)/unit)
	blocks := strings.Repeat(fullBlock, (int(duration)/unit)-2)

	line := fmt.Sprintf(` %s%s - %dus`, startGap, node.SpanSnapshot.Name, duration.Microseconds())
	tview.Print(screen, line, x, y, width, tview.AlignLeft, tcell.ColorWhite)

	line = fmt.Sprintf(`%s%s%s%s`, startGap, leftArrow, blocks, rightArrow)
	tview.Print(screen, line, x, y+1, width, tview.AlignLeft, tcell.ColorBlue)

	y += 2

	for _, child := range node.Children {
		y = t.drawTrace(screen, child, x, y, width, unit, rootStart)
	}

	return y
}

// TODO input handling doesn't really do anything yet
func (t *Trace) InputHandler() func(event *tcell.EventKey, setFocus func(p tview.Primitive)) {
	return t.WrapInputHandler(func(event *tcell.EventKey, setFocus func(p tview.Primitive)) {
	})
}

func (t *Trace) MouseHandler() func(action tview.MouseAction, event *tcell.EventMouse, setFocus func(p tview.Primitive)) (consumed bool, capture tview.Primitive) {
	return t.WrapMouseHandler(func(action tview.MouseAction, event *tcell.EventMouse, setFocus func(p tview.Primitive)) (consumed bool, capture tview.Primitive) {
		return
	})
}
