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

package api

import "time"

// ListMeta holds common metadata for lists.
type ListMeta struct{}

// ObjectMeta holds common metadata for different objects.
type ObjectMeta struct {
	UID, Name, ResourceVersion string
	Created, Updated           time.Time
	Generation                 uint64
	Annotations                map[string]string
}

func (m *ObjectMeta) GetUID() string {
	return m.UID
}

func (m *ObjectMeta) SetUID(uid string) {
	m.UID = uid
}

func (m *ObjectMeta) GetName() string {
	return m.Name
}

func (m *ObjectMeta) SetName(name string) {
	m.Name = name
}

func (m *ObjectMeta) GetResourceVersion() string {
	return m.ResourceVersion
}

func (m *ObjectMeta) SetResourceVersion(resourceVersion string) {
	m.ResourceVersion = resourceVersion
}

func (m *ObjectMeta) GetCreated() time.Time {
	return m.Created
}

func (m *ObjectMeta) SetCreated(created time.Time) {
	m.Created = created
}

func (m *ObjectMeta) GetUpdated() time.Time {
	return m.Updated
}

func (m *ObjectMeta) SetUpdated(updated time.Time) {
	m.Updated = updated
}

func (m *ObjectMeta) GetGeneration() uint64 {
	return m.Generation
}

func (m *ObjectMeta) SetGeneration(generation uint64) {
	m.Generation = generation
}

func (m *ObjectMeta) GetAnnotations() map[string]string {
	return m.Annotations
}

func (m *ObjectMeta) SetAnnotations(annotations map[string]string) {
	m.Annotations = annotations
}

// Object is a generic metadata interface for all object.
type Object interface {
	GetUID() string
	SetUID(string)
	GetName() string
	SetName(string)
	GetResourceVersion() string
	SetResourceVersion(string)
	GetCreated() time.Time
	SetCreated(time.Time)
	GetUpdated() time.Time
	SetUpdated(time.Time)
	GetGeneration() uint64
	SetGeneration(uint64)
	GetAnnotations() map[string]string
	SetAnnotations(map[string]string)
}
