package display

import (
	"os"
	"sort"
	"time"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
	"github.com/rs/zerolog/log"
)

// Run reads in input, builds a tree of traces then flattens it to then display
func Run() {
	input, err := Read(os.Stdin)

	if err != nil {
		log.Error().Err(err).Msg("Read traces")
	}

	view := tview.NewFlex().SetDirection(tview.FlexRow)
	frame := tview.NewFrame(view)

	frame.SetBorders(1, 0, 1, 0, 0, 0)
	frame.AddText("telemeterm", true, tview.AlignCenter, tcell.ColorBlue)

	text := tview.NewTextView()
	text.SetBorder(true)

	traces := ToTraces(ToTree(input), text)

	traceGraph := NewTraceGraph(traces, text)
	traceGraph.SetBorder(true)

	view.AddItem(traceGraph, 0, 4, true).
		AddItem(text, 0, 1, true)

	if err := tview.NewApplication().SetRoot(frame, true).Run(); err != nil {
		log.Error().Err(err).Msg("Start screen")
	}
}

func ToTree(spans []*SpanSnapshot) *SnapshotNode {
	spanMap := make(map[string]*SnapshotNode)

	for _, span := range spans {
		node := getOrCreate(spanMap, span.SpanContext.SpanID)
		node.SpanSnapshot = span

		parent := getOrCreate(spanMap, span.Parent.SpanID)
		parent.Children = append(parent.Children, node)

		sort.Sort(parent.Children)
	}

	for _, node := range spanMap {
		if node.SpanSnapshot == nil {
			// Return the root node
			return node
		}
	}

	log.Error().Msg("No root node found")
	return nil
}

func getOrCreate(spanMap map[string]*SnapshotNode, spanID string) *SnapshotNode {
	node, ok := spanMap[spanID]
	if !ok {
		spanMap[spanID] = &SnapshotNode{}
		node = spanMap[spanID]
	}

	return node
}

func ToTraces(node *SnapshotNode, writer *tview.TextView) []*Trace {
	var traces []*Trace
	for _, child := range node.Children {
		traces = append(traces, NewTrace(
			child.SpanSnapshot.StartTime,
			child.SpanSnapshot.EndTime,
			buildLines(child, child.SpanSnapshot.StartTime, writer),
		))
	}

	traces[0].lines[0].selected = true

	return traces
}

func buildLines(node *SnapshotNode, rootStart time.Time, writer *tview.TextView) []*Line {
	lines := []*Line{NewLine(node, rootStart, writer)}

	for _, child := range node.Children {
		var childLines []*Line
		childLines = buildLines(child, rootStart, writer)
		lines = append(lines, childLines...)
	}

	return lines
}
