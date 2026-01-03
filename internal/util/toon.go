package util

import (
	"fmt"
	"os"

	toon "github.com/toon-format/toon-go"
)

func ParseTOONFile(path string) (map[string]any, error) {
	raw, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("read TOON file: %w", err)
	}
	var doc map[string]any
	if err := toon.Unmarshal(raw, &doc); err != nil {
		return nil, fmt.Errorf("parse TOON: %w", err)
	}
	return doc, nil
}
