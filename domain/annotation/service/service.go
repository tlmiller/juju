// Copyright 2023 Canonical Ltd.
// Licensed under the AGPLv3, see LICENCE file for details.

package service

import (
	"context"

	"github.com/juju/juju/core/annotations"
	"github.com/juju/juju/internal/errors"
)

// State describes retrieval and persistence methods for annotations.
type State interface {
	// GetAnnotations retrieves all the annotations associated with a given ID.
	// If no annotations are found, an empty map is returned.
	GetAnnotations(ctx context.Context, ID annotations.ID) (map[string]string, error)

	// SetAnnotations associates key/value annotation pairs with a given ID.
	// If an annotation already exists for the given ID, then it will be updated
	// with the given value. First all annotations are deleted, then the given
	// pairs are inserted, so unsetting an annotation is implicit.
	SetAnnotations(ctx context.Context, ID annotations.ID, annotations map[string]string) error
}

// Service provides the API for working with annotations.
type Service struct {
	st State
}

// NewService returns a new service reference wrapping the given annotations state.
func NewService(st State) *Service {
	return &Service{
		st: st,
	}
}

// GetAnnotations retrieves all the annotations associated with a given ID. If
// no annotations are found, an empty map is returned.
func (s *Service) GetAnnotations(ctx context.Context, id annotations.ID) (map[string]string, error) {
	annotations, err := s.st.GetAnnotations(ctx, id)
	return annotations, errors.Capture(err)
}

// SetAnnotations associates key/value annotation pairs with a given ID. If
// an annotation already exists for the given ID, then it will be updated with
// the given value.
func (s *Service) SetAnnotations(ctx context.Context, id annotations.ID, annotations map[string]string) error {
	err := s.st.SetAnnotations(ctx, id, annotations)
	return errors.Errorf("updating annotations for ID: %q %w", id.Name, err)
}
