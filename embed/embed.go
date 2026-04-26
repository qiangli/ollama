// Package embed provides pure-Go components from the Ollama project for
// embedding in external Go applications.
//
// This package deliberately avoids importing Ollama's server/, discover/,
// llm/, llama/, and x/ packages which require CGo. Instead, it re-exports
// the pure-Go packages: api types, model management, manifest handling,
// template rendering, and configuration.
//
// The Ollama HTTP server and C++ inference runner are managed as external
// processes by the host application.
package embed

import (
	"fmt"
	"net/http"
	"net/url"

	"github.com/ollama/ollama/api"
	"github.com/ollama/ollama/envconfig"
	"github.com/ollama/ollama/manifest"
	"github.com/ollama/ollama/types/model"
)

// Re-export Ollama API types for use by the host application.
type (
	// GenerateRequest is the Ollama generate API request.
	GenerateRequest = api.GenerateRequest
	// GenerateResponse is the Ollama generate API response.
	GenerateResponse = api.GenerateResponse
	// ChatRequest is the Ollama chat API request.
	ChatRequest = api.ChatRequest
	// ChatResponse is the Ollama chat API response.
	ChatResponse = api.ChatResponse
	// PullRequest is the model pull API request.
	PullRequest = api.PullRequest
	// ProgressResponse tracks download/operation progress.
	ProgressResponse = api.ProgressResponse
	// ListResponse lists locally available models.
	ListResponse = api.ListResponse
	// ShowRequest gets model details.
	ShowRequest = api.ShowRequest
	// ShowResponse contains model details.
	ShowResponse = api.ShowResponse
	// DeleteRequest deletes a local model.
	DeleteRequest = api.DeleteRequest
	// CreateRequest creates a model from a file or other model.
	CreateRequest = api.CreateRequest
	// EmbedRequest creates embeddings.
	EmbedRequest = api.EmbedRequest
	// EmbedResponse contains embeddings.
	EmbedResponse = api.EmbedResponse
)

// ModelName is the parsed representation of an Ollama model name.
type ModelName = model.Name

// ParseModelName parses a model name string into its components.
var ParseModelName = model.ParseName

// DefaultOptions returns the default inference options.
var DefaultOptions = api.DefaultOptions

// ManifestsPath returns the base path for model manifests.
var ManifestsPath = manifest.Path

// BlobsPath returns the path for a blob by digest, or the blobs directory.
var BlobsPath = manifest.BlobsPath

// Manifests returns all locally stored model manifests.
var Manifests = manifest.Manifests

// AllowedOrigins returns configured allowed origins for CORS.
var AllowedOrigins = envconfig.AllowedOrigins

// NewClient creates an Ollama API client connected to the given base URL.
func NewClient(baseURL string) (*api.Client, error) {
	u, err := url.Parse(baseURL)
	if err != nil {
		return nil, fmt.Errorf("parse ollama base URL: %w", err)
	}
	return api.NewClient(u, http.DefaultClient), nil
}
