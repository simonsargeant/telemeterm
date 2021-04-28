package display

import (
	"fmt"
	"strings"
	"time"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
	"github.com/rs/zerolog/log"
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
	// missing var is height
	x, y, width, _ := t.GetInnerRect()

	index := 0
	for _, trace := range t.root.Children {
		start, err := time.Parse(time.RFC3339Nano, trace.SpanSnapshot.StartTime)
		if err != nil {
			log.Error().Err(err).Msg("Invalid start timestamp")
			continue
		}

		end, err := time.Parse(time.RFC3339Nano, trace.SpanSnapshot.EndTime)
		if err != nil {
			log.Error().Err(err).Msg("Invalid end timestamp")
			continue
		}

		duration := end.Sub(start)
		blocks := strings.Repeat(fullBlock, width-2)

		unit := int(duration) / width

		line := fmt.Sprintf(` %s - %dus`, trace.SpanSnapshot.Name, duration.Microseconds())
		tview.Print(screen, line, x, y+index, width, tview.AlignLeft, tcell.ColorWhite)
		line = fmt.Sprintf(`%s%s%s`, leftArrow, blocks, rightArrow)
		tview.Print(screen, line, x, y+index+1, width, tview.AlignLeft, tcell.ColorBlue)

		index += 2
		for _, child := range trace.Children {
			index, err = t.drawChild(screen, child, x, y, width, index, unit, start)
			if err != nil {
				log.Error().Err(err).Msg("Draw child nodes")
			}
		}

		index += 1
	}
}

func (t *Trace) drawChild(screen tcell.Screen, node *SnapshotNode, x, y, width, index, unit int, parentStart time.Time) (int, error) {
	start, err := time.Parse(time.RFC3339Nano, node.SpanSnapshot.StartTime)
	if err != nil {
		return index, fmt.Errorf("start timestamp: %w", err)
	}

	end, err := time.Parse(time.RFC3339Nano, node.SpanSnapshot.EndTime)
	if err != nil {
		return index, fmt.Errorf("end timestamp: %w", err)
	}

	startLag := start.Sub(parentStart)
	duration := end.Sub(start)

	startGap := strings.Repeat(" ", int(startLag)/unit)
	blocks := strings.Repeat(fullBlock, int(duration)/unit)

	line := fmt.Sprintf(` %s%s - %dus`, startGap, node.SpanSnapshot.Name, duration.Microseconds())
	tview.Print(screen, line, x, y+index, width, tview.AlignLeft, tcell.ColorWhite)
	line = fmt.Sprintf(`%s%s%s%s`, startGap, leftArrow, blocks, rightArrow)
	tview.Print(screen, line, x, y+index+1, width, tview.AlignLeft, tcell.ColorBlue)

	index += 2
	for _, child := range node.Children {
		index, err = t.drawChild(screen, child, x, y, width, index, unit, parentStart)
		if err != nil {
			log.Error().Err(err).Msg("Draw child nodes")
		}
	}

	return index, nil
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
