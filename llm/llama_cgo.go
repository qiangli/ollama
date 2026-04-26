//go:build cgo

package llm

import (
	"github.com/ollama/ollama/llama"
)

// llamaModelHandle wraps the CGo llama.Model for use across build configurations.
type llamaModelHandle struct {
	model *llama.Model
}

// Tokenize tokenizes the given content via CGo llama.cpp.
func (h *llamaModelHandle) Tokenize(content string, addSpecial bool, parseSpecial bool) ([]int, error) {
	return h.model.Tokenize(content, addSpecial, parseSpecial)
}

// TokenToPiece converts a token ID to its string representation via CGo.
func (h *llamaModelHandle) TokenToPiece(token int) string {
	return h.model.TokenToPiece(token)
}

// loadLlamaModel loads a model file for vocab-only access via CGo llama.cpp.
func loadLlamaModel(modelPath string) (*llamaModelHandle, error) {
	m, err := llama.LoadModelFromFile(modelPath, llama.ModelParams{VocabOnly: true})
	if err != nil {
		return nil, err
	}
	return &llamaModelHandle{model: m}, nil
}

// freeLlamaModel frees a CGo llama model. Safe to call with nil.
func freeLlamaModel(m *llamaModelHandle) {
	if m != nil && m.model != nil {
		llama.FreeModel(m.model)
	}
}

// schemaToGrammar converts a JSON Schema to a GBNF grammar string via CGo.
func schemaToGrammar(schema []byte) []byte {
	return llama.SchemaToGrammar(schema)
}
