//go:build !cgo

package server

import "testing"

func TestQuantizationUnavailableWithoutCGO(t *testing.T) {
	err := quantize(nil, nil, nil, 0, func(uint64) { t.Fatal("unexpected progress callback") })
	if err == nil || err.Error() != "quantization not supported in this build" {
		t.Fatalf("quantize error = %v", err)
	}
}
