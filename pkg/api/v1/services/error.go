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

package services

import (
	"github.com/go-logr/logr"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	db "github.com/thetechnick/lb/pkg/storage"
)

// api.ErrorFieldViolation descriptions
const (
	// NotEmpty is returned when the field cannot be empty
	NotEmpty = "Cannot be empty"
)

func checkError(err error, log logr.Logger) error {
	if _, ok := status.FromError(err); ok {
		return err
	}
	switch err {
	case db.ErrNotFound:
		return status.New(codes.NotFound, "not found").Err()
	case db.ErrNameConflict:
		return status.New(codes.AlreadyExists, err.Error()).Err()
	case db.ErrIdentifierMissing:
		return status.New(codes.InvalidArgument, err.Error()).Err()
	case db.ErrResourceVersionChanged:
		return status.New(codes.Aborted, err.Error()).Err()
	}
	log.Error(err, "internal error")
	return status.New(codes.Internal, "internal error").Err()
}
