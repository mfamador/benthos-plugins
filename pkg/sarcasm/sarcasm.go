// plugin package for test application
package sarcasm

import (
	"bytes"
	"strconv"
	"time"

	"github.com/Jeffail/benthos/v3/lib/types"
)

// HowSarcastic implements de sarcasm algorithm
func HowSarcastic(content []byte) float64 {
	if bytes.Contains(bytes.ToLower(content), []byte("/s")) {
		const fullSarcastic = 100
		return fullSarcastic
	}
	return 0
}

// SarcasmProc structure
type SarcasmProc struct {
	MetadataKey string `json:"metadata_key" yaml:"metadata_key"`
}

// ProcessMessage returns messages mutated with their sarcasm level.
func (s *SarcasmProc) ProcessMessage(msg types.Message) ([]types.Message, types.Response) {
	newMsg := msg.Copy()
	if err := newMsg.Iter(func(i int, p types.Part) error {
		sarcasm := HowSarcastic(p.Get())
		sarcasmStr := strconv.FormatFloat(sarcasm, 'f', -1, 64)
		if len(s.MetadataKey) > 0 {
			p.Metadata().Set(s.MetadataKey, sarcasmStr)
		} else {
			p.Set([]byte(sarcasmStr))
		}
		return nil
	}); err != nil {
		panic("Invalid config")
	}
	return []types.Message{newMsg}, nil
}

// CloseAsync does nothing.
func (s *SarcasmProc) CloseAsync() {}

// WaitForClose does nothing.
func (s *SarcasmProc) WaitForClose(timeout time.Duration) error {
	return nil
}
