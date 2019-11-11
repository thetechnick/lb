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

import (
	"sort"
	"time"
)

// Condition represents a state that a object is in
type Condition struct {
	Status                ConditionStatus
	Type, Reason, Message string
	LastTransitionTime    time.Time
}

// ConditionStatus is the status a condition is in
type ConditionStatus string

// ConditionStatus values
const (
	ConditionStatusUnknown ConditionStatus = "Unknown"
	ConditionStatusTrue    ConditionStatus = "True"
	ConditionStatusFalse   ConditionStatus = "False"
)

// ConditionCollection represents a collection of conditions
type ConditionCollection interface {
	Get(t string) (c Condition, ok bool)
	Set(c Condition)
	SetMultiple(conditions ...Condition)
	List() (l []Condition)
}

// ConditionMapCollection implements ConditionCollection by using a map
type ConditionMapCollection map[string]Condition

// NewConditionMapCollection creates a new ConditionMapCollection
func NewConditionMapCollection(conditions []Condition) ConditionMapCollection {
	c := ConditionMapCollection{}
	c.SetMultiple(conditions...)
	return c
}

// Get returns a condition of the given type if it exists in the collection
func (m ConditionMapCollection) Get(t string) (c Condition, ok bool) {
	c, ok = m[t]
	return
}

// Set overrides a existing condition with a new one
func (m ConditionMapCollection) Set(c Condition) {
	m[c.Type] = c
}

// List returns a slice of conditions
func (m ConditionMapCollection) List() (l []Condition) {
	for _, c := range m {
		l = append(l, c)
	}
	sort.Slice(l, func(i, j int) bool {
		return l[i].Type < l[j].Type
	})
	return
}

// SetMultiple sets multiple conditions at once
func (m ConditionMapCollection) SetMultiple(conditions ...Condition) {
	for _, c := range conditions {
		m.Set(c)
	}
}
