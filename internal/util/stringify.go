package util

import (
	"encoding/json"
	"fmt"
	"strings"
)

func JSONPretty(v any) []byte {
	b, err := json.MarshalIndent(v, "", "  ")
	if err != nil {
		return []byte(fmt.Sprintf(`{"error":"json marshal failed: %v"}`, err))
	}
	return b
}

func StrAny(v any) string {
	if v == nil {
		return ""
	}
	switch t := v.(type) {
	case string:
		return strings.TrimSpace(t)
	default:
		return fmt.Sprintf("%v", v)
	}
}
