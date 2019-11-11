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

package bolt

import (
	"context"

	"github.com/boltdb/bolt"
	"github.com/go-logr/logr"
	"github.com/golang/protobuf/proto"

	"github.com/thetechnick/lb/pkg/api"
	v1 "github.com/thetechnick/lb/pkg/api/v1"
	db "github.com/thetechnick/lb/pkg/storage"
)

const serverBucket = "lb:servers"

// NewServerStorage creates a new ServerStorage
func NewServerStorage(db *bolt.DB, log logr.Logger) (db.ServerStorage, error) {
	s := &serverStorage{}
	s.base = &baseStorage{
		log: log,

		db:          db,
		bucket:      serverBucket,
		toStorage:   toStorageFn(s.toStorage),
		fromStorage: fromStorageFn(s.fromStorage),
		key:         keyFn(s.key),
	}
	return s, s.base.init()
}

type serverStorage struct {
	base *baseStorage
}

func (s *serverStorage) Delete(ctx context.Context, obj *api.Server) (deleted *api.Server, err error) {
	var o api.Object
	if o, err = s.base.Delete(obj); err != nil {
		return
	}
	deleted = o.(*api.Server)
	return
}

func (s *serverStorage) Create(ctx context.Context, obj *api.Server) (err error) {
	err = s.base.Create(obj)
	if err != nil {
		return
	}
	return
}

func (s *serverStorage) Update(ctx context.Context, obj *api.Server) (err error) {
	_, _, err = s.base.Update(obj, func(old, new api.Object) {
		oldObj := old.(*api.Server)
		newObj := new.(*api.Server)

		// increment generation
		obj.SetGeneration(obj.GetGeneration() + 1)

		// do not update the status
		newObj.Status = oldObj.Status
	})
	return
}

func (s *serverStorage) UpdateStatus(ctx context.Context, obj *api.Server) (err error) {
	_, _, err = s.base.Update(obj, func(old, new api.Object) {
		oldObj := old.(*api.Server)
		newObj := new.(*api.Server)

		// do not update the spec or metadata
		newObj.Spec = oldObj.Spec
		newObj.ObjectMeta = oldObj.ObjectMeta
	})
	return
}

func (s *serverStorage) Show(ctx context.Context, name string) (volume *api.Server, err error) {
	var obj api.Object
	obj, err = s.base.Show(name)
	if err != nil {
		return
	}
	if obj != nil {
		volume = obj.(*api.Server)
	}
	return
}

func (s *serverStorage) All(ctx context.Context) (list []*api.Server, err error) {
	var objs []api.Object
	if objs, err = s.base.All(); err != nil {
		return
	}
	for _, obj := range objs {
		list = append(list, obj.(*api.Server))
	}
	return
}

func (s *serverStorage) key(obj api.Object) []byte {
	return []byte(obj.GetName())
}

func (s *serverStorage) fromStorage(b []byte) (o api.Object, err error) {
	if b == nil {
		return
	}

	so := &v1.Server{}
	if err = proto.Unmarshal(b, so); err != nil {
		return
	}

	printer := &api.Server{}
	err = v1.Scheme.Convert(so, printer)
	o = printer
	return
}

func (s *serverStorage) toStorage(obj api.Object) (b []byte, err error) {
	printer := obj.(*api.Server)
	so := &v1.Server{}
	if err = v1.Scheme.Convert(printer, so); err != nil {
		return
	}

	if b, err = proto.Marshal(so); err != nil {
		return
	}
	return
}
