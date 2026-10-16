//go:build !cgo

package server

import (
	"errors"
	"os"

	fsggml "github.com/ollama/ollama/fs/ggml"
)

// Quantization requires the native GGML backend. Serving existing GGUF models
// through the external runner does not use this path.
func quantize(in, out *os.File, orig *fsggml.GGML, newFileType fsggml.FileType, progressFn func(n uint64)) error {
	return errors.New("quantization not supported in this build")
}
