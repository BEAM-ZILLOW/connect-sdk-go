package metadata

import (
	"errors"
	"strings"
)

// ShoppingCartExtension represents the shopping extension present in the MetaDataProvider
type ShoppingCartExtension struct {
	Creator string `json:"creator,omitempty"`
	Name    string `json:"name,omitempty"`
	Version string `json:"version,omitempty"`
}

// NewShoppingCartExtension creates a ShoppingCartExtension with the given creator, name and version
func NewShoppingCartExtension(creator string, name string, version string) (*ShoppingCartExtension, error) {
	if strings.TrimSpace(creator) == "" {
		return nil, errors.New("creator is required")
	}
	if strings.TrimSpace(creator) == "" {
		return nil, errors.New("name is required")
	}
	if strings.TrimSpace(creator) == "" {
		return nil, errors.New("version is required")
	}
	return &ShoppingCartExtension{creator, name, version}, nil
}