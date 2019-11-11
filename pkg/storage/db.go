/*
Copyright 2019 The LB Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package storage

import (
	"context"
	"errors"

	"github.com/thetechnick/lb/pkg/api"
)

var (
	// ErrIdentifierMissing is returned if the given object is missing its name or uid
	ErrIdentifierMissing = errors.New("missing name or uid")

	// ErrResourceVersionChanged is returned if the resource version of the object has changed
	ErrResourceVersionChanged = errors.New("resource version changed")

	// ErrNameConflict is returned if a item with the name already exists
	ErrNameConflict = errors.New("a resource with this name already exists")

	// ErrNotFound is returned when the requested item was not found
	ErrNotFound = errors.New("not found")
)

// ServerStorage persists volumes
type ServerStorage interface {
	Create(ctx context.Context, server *api.Server) error
	Update(ctx context.Context, server *api.Server) (err error)
	UpdateStatus(ctx context.Context, server *api.Server) (err error)
	All(ctx context.Context) (list []*api.Server, err error)
	Delete(ctx context.Context, server *api.Server) (deleted *api.Server, err error)
	Show(ctx context.Context, name string) (server *api.Server, err error)
}
