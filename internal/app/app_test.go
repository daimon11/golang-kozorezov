package app

import (
	"bytes"
	"testing"
)

func TestRun(t *testing.T) {
	var buf bytes.Buffer

	if err := Run(&buf); err != nil {
		t.Fatalf("Run() returned error: %v", err)
	}

	if got, want := buf.String(), "homeworks: ok\n"; got != want {
		t.Errorf("Run() wrote %q, want %q", got, want)
	}
}
