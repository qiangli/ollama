//go:build cgo

package server

import (
	"github.com/ollama/ollama/api"
	"github.com/ollama/ollama/format"
	"github.com/ollama/ollama/llm"
	"github.com/ollama/ollama/x/imagegen"
	imagegenmanifest "github.com/ollama/ollama/x/imagegen/manifest"
	"github.com/ollama/ollama/x/mlxrunner"
)

func newMLXServer(name string, image bool) (llm.LlamaServer, error) {
	if image {
		return imagegen.NewServer(name)
	}
	return mlxrunner.NewClient(name)
}

func populateImageModelDetails(name string, details *api.ModelDetails) error {
	if info, err := imagegenmanifest.GetModelInfo(name); err == nil {
		details.Family = info.Architecture
		details.ParameterSize = format.HumanNumber(uint64(info.ParameterCount))
		details.QuantizationLevel = info.Quantization
	}
	return nil
}
