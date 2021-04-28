package display

import (
	"fmt"

	"github.com/rivo/tview"
)

func Display(traceTree *SnapshotNode) error {
	view := NewTrace(traceTree)

	view.SetBorder(true).
		SetTitle("telemeterm")

	if err := tview.NewApplication().SetRoot(view, true).Run(); err != nil {
		return fmt.Errorf("start screen: %w", err)
	}

	return nil
}
