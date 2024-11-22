// Copyright 2023 Canonical Ltd.
// Licensed under the AGPLv3, see LICENCE file for details.

package service

import (
	"context"

	"github.com/juju/juju/core/changestream"
	coreobjectstore "github.com/juju/juju/core/objectstore"
	"github.com/juju/juju/core/watcher"
	"github.com/juju/juju/core/watcher/eventsource"
	interrors "github.com/juju/juju/internal/errors"
)

// State describes retrieval and persistence methods for the coreobjectstore.
type State interface {
	// GetMetadata returns the persistence metadata for the specified path.
	GetMetadata(ctx context.Context, path string) (coreobjectstore.Metadata, error)
	// PutMetadata adds a new specified path for the persistence metadata.
	PutMetadata(ctx context.Context, metadata coreobjectstore.Metadata) error
	// ListMetadata returns the persistence metadata for all paths.
	ListMetadata(ctx context.Context) ([]coreobjectstore.Metadata, error)
	// RemoveMetadata removes the specified path for the persistence metadata.
	RemoveMetadata(ctx context.Context, path string) error
	// InitialWatchStatement returns the table and the initial watch statement
	// for the persistence metadata.
	InitialWatchStatement() (string, string)
}

// WatcherFactory describes methods for creating watchers.
type WatcherFactory interface {
	// NewNamespaceWatcher returns a new namespace watcher
	// for events based on the input change mask.
	NewNamespaceWatcher(string, changestream.ChangeType, eventsource.NamespaceQuery) (watcher.StringsWatcher, error)
}

// Service provides the API for working with the coreobjectstore.
type Service struct {
	st State
}

// NewService returns a new service reference wrapping the input state.
func NewService(st State) *Service {
	return &Service{
		st: st,
	}
}

// GetMetadata returns the persistence metadata for the specified path.
func (s *Service) GetMetadata(ctx context.Context, path string) (coreobjectstore.Metadata, error) {
	metadata, err := s.st.GetMetadata(ctx, path)
	if err != nil {
		return coreobjectstore.Metadata{}, interrors.Errorf("retrieving metadata %s %w", path, err)
	}
	return coreobjectstore.Metadata{
		Path: metadata.Path,
		Hash: metadata.Hash,
		Size: metadata.Size,
	}, nil
}

// ListMetadata returns the persistence metadata for all paths.
func (s *Service) ListMetadata(ctx context.Context) ([]coreobjectstore.Metadata, error) {
	metadata, err := s.st.ListMetadata(ctx)
	if err != nil {
		return nil, interrors.Errorf("retrieving metadata: %w", err)
	}
	m := make([]coreobjectstore.Metadata, len(metadata))
	for i, v := range metadata {
		m[i] = coreobjectstore.Metadata{
			Path: v.Path,
			Hash: v.Hash,
			Size: v.Size,
		}
	}
	return m, nil
}

// PutMetadata adds a new specified path for the persistence metadata.
func (s *Service) PutMetadata(ctx context.Context, metadata coreobjectstore.Metadata) error {
	err := s.st.PutMetadata(ctx, coreobjectstore.Metadata{
		Hash: metadata.Hash,
		Path: metadata.Path,
		Size: metadata.Size,
	})
	if err != nil {
		return interrors.Errorf("adding path %s %w", metadata.Path, err)
	}
	return nil
}

// RemoveMetadata removes the specified path for the persistence metadata.
func (s *Service) RemoveMetadata(ctx context.Context, path string) error {
	err := s.st.RemoveMetadata(ctx, path)
	if err != nil {
		return interrors.Errorf("removing path %s %w", path, err)
	}
	return nil
}

// WatchableService provides the API for working with the objectstore
// and the ability to create watchers.
type WatchableService struct {
	Service
	watcherFactory WatcherFactory
}

// NewWatchableService returns a new service reference wrapping the input state.
func NewWatchableService(st State, watcherFactory WatcherFactory) *WatchableService {
	return &WatchableService{
		Service: Service{
			st: st,
		},
		watcherFactory: watcherFactory,
	}
}

// Watch returns a watcher that emits the path changes that either have been
// added or removed.
func (s *WatchableService) Watch() (watcher.StringsWatcher, error) {
	table, stmt := s.st.InitialWatchStatement()
	return s.watcherFactory.NewNamespaceWatcher(
		table,
		changestream.All,
		eventsource.InitialNamespaceChanges(stmt),
	)
}
