//go:build !cgo

// This file provides stub implementations of the llama CGo functions
// for pure-Go builds. When CGO_ENABLED=0, model vocab loading and
// grammar conversion are not available — the Ollama engine path
// (tokenizer.Tokenizer) is used instead.

package llm

import "fmt"

// llamaModelHandle is an opaque placeholder for pure-Go builds.
type llamaModelHandle struct{}

// Tokenize is unavailable without CGo.
func (h *llamaModelHandle) Tokenize(content string, addSpecial bool, parseSpecial bool) ([]int, error) {
	return nil, fmt.Errorf("tokenize unavailable: built without CGo")
}

// TokenToPiece is unavailable without CGo.
func (h *llamaModelHandle) TokenToPiece(token int) string {
	return ""
}

// loadLlamaModel returns an error in pure-Go mode, causing the caller
// (NewLlamaServer) to fall through to the Ollama engine tokenizer.
func loadLlamaModel(modelPath string) (*llamaModelHandle, error) {
	return nil, fmt.Errorf("llama model loading unavailable: built without CGo (use Ollama engine tokenizer)")
}

// freeLlamaModel is a no-op in pure-Go builds. Safe to call with nil.
func freeLlamaModel(m *llamaModelHandle) {}

// schemaToGrammar is unavailable without CGo — returns nil to signal failure.
func schemaToGrammar(schema []byte) []byte {
	return nil
}
