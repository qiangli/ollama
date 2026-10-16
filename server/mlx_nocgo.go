//go:build !cgo

package server

import (
	"errors"
	"github.com/ollama/ollama/api"
	"github.com/ollama/ollama/llm"
)

var errMLXNotSupported = errors.New("MLX models and image generation are not supported in this build")

func newMLXServer(string, bool) (llm.LlamaServer, error) {
	return nil, errMLXNotSupported
}

func populateImageModelDetails(string, *api.ModelDetails) error {
	return errMLXNotSupported
}
