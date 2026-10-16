//go:build !cgo

package server

import (
	"errors"
	"github.com/ollama/ollama/api"
	"testing"
)

func TestNoCGOMLXUnsupported(t *testing.T) {
	for _, image := range []bool{false, true} {
		runner, err := newMLXServer("test-model", image)
		if runner != nil || !errors.Is(err, errMLXNotSupported) {
			t.Fatalf("image=%v: runner=%v, error=%v", image, runner, err)
		}
	}
	if err := populateImageModelDetails("test-model", &api.ModelDetails{}); !errors.Is(err, errMLXNotSupported) {
		t.Fatalf("manifest error: %v", err)
	}
}
