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
	"errors"
	"fmt"
	"strconv"
	"time"

	"github.com/boltdb/bolt"
	"github.com/go-logr/logr"
	"github.com/google/uuid"

	"github.com/thetechnick/lb/pkg/api"
	db "github.com/thetechnick/lb/pkg/storage"
)

type toStorageFn func(obj api.Object) ([]byte, error)
type fromStorageFn func(b []byte) (obj api.Object, err error)
type updateHook func(old, new api.Object)
type keyFn func(obj api.Object) []byte

var noopUpdateHook = func(old, new api.Object) {}
var errBucketNotFound = errors.New("bucket not found")

type baseStorage struct {
	log logr.Logger

	db          *bolt.DB
	bucket      string
	toStorage   toStorageFn
	fromStorage fromStorageFn
	key         keyFn
	updateHook  updateHook
}

func (s *baseStorage) init() error {
	return s.db.Update(func(tx *bolt.Tx) error {
		if bucket := tx.Bucket([]byte(s.bucket)); bucket == nil {
			_, err := tx.CreateBucket([]byte(s.bucket))
			if err != nil && err != bolt.ErrBucketExists {
				return fmt.Errorf("create bucket: %s", err)
			}
		}
		return nil
	})
}

func (s *baseStorage) Delete(obj api.Object) (deleted api.Object, err error) {
	err = s.db.Update(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(s.bucket))
		if bucket == nil {
			return errBucketNotFound
		}

		key := s.key(obj)
		b := bucket.Get(key)
		if b == nil {
			return db.ErrNotFound
		}
		if deleted, err = s.fromStorage(b); err != nil {
			return err
		}
		return bucket.Delete(key)
	})
	s.log.WithValues("name", obj.GetName()).V(1).Info("deleted")
	return
}

func (s *baseStorage) Create(obj api.Object) (err error) {
	err = s.db.Update(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(s.bucket))
		if bucket == nil {
			return errBucketNotFound
		}

		if obj.GetName() == "" {
			return db.ErrIdentifierMissing
		}

		existingVolume, err := s.getKey(tx, s.key(obj))
		if err != nil {
			return err
		}
		if existingVolume != nil {
			return db.ErrNameConflict
		}

		uid, err := uuid.NewRandom()
		if err != nil {
			return err
		}
		now := time.Now()
		obj.SetResourceVersion("1")
		obj.SetUID(uid.String())
		obj.SetCreated(now)
		obj.SetUpdated(now)
		obj.SetGeneration(1)

		b, err := s.toStorage(obj)
		if err != nil {
			return err
		}
		if err = bucket.Put(s.key(obj), b); err != nil {
			return err
		}

		return nil
	})
	s.log.WithValues("name", obj.GetName()).V(1).Info("created")
	return
}

func (s *baseStorage) Update(obj api.Object, hook updateHook) (old, new api.Object, err error) {
	new = obj
	err = s.db.Update(func(tx *bolt.Tx) (err error) {
		bucket := tx.Bucket([]byte(s.bucket))
		if bucket == nil {
			return errBucketNotFound
		}

		if obj.GetName() == "" {
			return db.ErrIdentifierMissing
		}

		old, err = s.getKey(tx, s.key(obj))
		if err != nil {
			return err
		}
		if old == nil {
			return db.ErrNotFound
		}

		if obj.GetResourceVersion() != "" &&
			old.GetResourceVersion() != obj.GetResourceVersion() {
			return db.ErrResourceVersionChanged
		}

		if obj.GetUID() != "" && obj.GetUID() != old.GetUID() {
			return errors.New("can not change resource uid")
		}

		prevVersion, err := strconv.Atoi(old.GetResourceVersion())
		if err != nil {
			return err
		}

		hook(old, new)
		new.SetResourceVersion(strconv.Itoa(prevVersion + 1))
		new.SetUID(old.GetUID())
		new.SetUpdated(time.Now())
		new.SetCreated(old.GetCreated())
		b, err := s.toStorage(new)
		if err != nil {
			return err
		}
		if err = bucket.Put(s.key(new), b); err != nil {
			return err
		}

		return nil
	})
	s.log.WithValues("name", obj.GetName()).V(1).Info("updated")
	return
}

func (s *baseStorage) getKey(tx *bolt.Tx, key []byte) (obj api.Object, err error) {
	bucket := tx.Bucket([]byte(s.bucket))
	if bucket == nil {
		err = errBucketNotFound
		return
	}

	b := bucket.Get(key)
	return s.fromStorage(b)
}

func (s *baseStorage) Show(name string) (obj api.Object, err error) {
	err = s.db.View(func(tx *bolt.Tx) error {
		obj, err = s.getKey(tx, []byte(name))
		return err
	})
	return
}

func (s *baseStorage) All() (list []api.Object, err error) {
	err = s.db.View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(s.bucket))
		if bucket == nil {
			return errBucketNotFound
		}

		return bucket.ForEach(func(k []byte, v []byte) error {
			obj, err := s.fromStorage(v)
			if err != nil {
				return err
			}
			list = append(list, obj)
			return nil
		})
	})
	return
}
