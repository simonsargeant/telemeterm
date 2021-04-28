package display

import (
	"os"
	"sort"

	"github.com/rs/zerolog/log"
)

func Run() {
	traces, err := Read(os.Stdin)

	if err != nil {
		log.Error().Err(err).Msg("Read traces")
	}

	err = Display(ToTree(traces))

	if err != nil {
		log.Error().Err(err).Msg("Display traces")
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
