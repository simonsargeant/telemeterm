package display

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"

	"github.com/rs/zerolog/log"
)

func Read(r io.Reader) ([]*SpanSnapshot, error) {
	scanner := bufio.NewScanner(r)

	log.Debug().Msg("Reading")

	var traces []*SpanSnapshot
	for scanner.Scan() {
		log.Debug().Msg("Reading line")

		var traceSlice []*SpanSnapshot

		err := json.Unmarshal(scanner.Bytes(), &traceSlice)
		if err != nil {
			return nil, fmt.Errorf("unmarshal traces: %w", err)
		}

		log.Debug().Int("traces", len(traceSlice)).Msg("Got traces")
		traces = append(traces, traceSlice...)
	}

	log.Debug().Msg("Finished reading")

	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("scan input: %w", err)
	}

	return traces, nil
}
