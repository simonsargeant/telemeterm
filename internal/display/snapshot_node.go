package display

// SnapshotNode represents the tree of traces. Each trace may have multiple
// children but a single parent.
type SnapshotNode struct {
	SpanSnapshot *SpanSnapshot
	Children     Nodes
}

type Nodes []*SnapshotNode

func (n Nodes) Len() int      { return len(n) }
func (n Nodes) Swap(i, j int) { n[i], n[j] = n[j], n[i] }

func (n Nodes) Less(i, j int) bool {
	return n[i].SpanSnapshot.StartTime.Before(n[j].SpanSnapshot.StartTime)
}
