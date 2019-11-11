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
	"context"

	"github.com/go-logr/logr"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/thetechnick/lb/pkg/api"
	v1 "github.com/thetechnick/lb/pkg/api/v1"
	db "github.com/thetechnick/lb/pkg/storage"
)

type serverService struct {
	log     logr.Logger
	storage db.ServerStorage
}

func NewServerService(
	log logr.Logger,
	storage db.ServerStorage,
) v1.ServerServiceServer {
	return &serverService{
		log:     log,
		storage: storage,
	}
}

func (s *serverService) checkError(err error) error {
	return checkError(err, s.log)
}

func (s *serverService) Create(ctx context.Context, req *v1.Server) (res *v1.Server, err error) {
	defer func() {
		err = s.checkError(err)
	}()
	res = &v1.Server{}

	// validation

	obj := &api.Server{}
	if err = v1.Scheme.Convert(req, obj); err != nil {
		return
	}
	if err = s.storage.Create(ctx, obj); err != nil {
		return
	}
	if err = v1.Scheme.Convert(obj, res); err != nil {
		return
	}
	return
}

func (s *serverService) Delete(ctx context.Context, req *v1.ServerDeleteRequest) (res *v1.Server, err error) {
	defer func() {
		err = s.checkError(err)
	}()
	res = &v1.Server{}

	// validation

	var p *api.Server
	p, err = s.storage.Show(ctx, req.Name)
	if err != nil {
		return
	}
	if p == nil {
		err = status.Error(codes.NotFound, "printer not found")
		return
	}

	if p, err = s.storage.Delete(ctx, p); err != nil {
		return
	}

	if err = v1.Scheme.Convert(p, res); err != nil {
		return
	}
	return
}

func (s *serverService) List(ctx context.Context, req *v1.ServerListRequest) (res *v1.ServerList, err error) {
	defer func() {
		err = s.checkError(err)
	}()
	res = &v1.ServerList{}

	all, err := s.storage.All(ctx)
	if err != nil {
		return
	}

	res = &v1.ServerList{}
	printerListItems := []*v1.Server{}
	if err = v1.Scheme.Convert(&all, &printerListItems); err != nil {
		return
	}
	res.Items = printerListItems
	return
}

func (s *serverService) Show(ctx context.Context, req *v1.ServerShowRequest) (res *v1.Server, err error) {
	defer func() {
		err = s.checkError(err)
	}()
	res = &v1.Server{}

	// validate

	obj, err := s.storage.Show(ctx, req.Name)
	if err != nil {
		return
	}
	if obj == nil {
		err = status.Error(codes.NotFound, "server not found")
		return
	}
	if err = v1.Scheme.Convert(obj, res); err != nil {
		return
	}
	return
}

func (s *serverService) Update(ctx context.Context, req *v1.Server) (res *v1.Server, err error) {
	defer func() {
		err = s.checkError(err)
	}()
	res = &v1.Server{}

	// validate

	obj := &api.Server{}
	if err = v1.Scheme.Convert(req, obj); err != nil {
		return
	}
	if err = s.storage.Update(ctx, obj); err != nil {
		return
	}
	if err = v1.Scheme.Convert(obj, res); err != nil {
		return
	}
	return
}

func (s *serverService) UpdateStatus(ctx context.Context, req *v1.Server) (res *v1.Server, err error) {
	defer func() {
		err = s.checkError(err)
	}()
	res = &v1.Server{}

	// validate

	obj := &api.Server{}
	if err = v1.Scheme.Convert(req, obj); err != nil {
		return
	}
	if err = s.storage.UpdateStatus(ctx, obj); err != nil {
		return
	}
	if err = v1.Scheme.Convert(obj, res); err != nil {
		return
	}
	return
}
