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

package runtime

import (
	"fmt"
	"reflect"
)

// Scope is passed to conversion funcs to allow them to continue an ongoing conversion.
type Scope interface {
	// Call Convert to convert sub-objects.
	Convert(src, dest interface{}) error
}

// scope contains information about an ongoing conversion.
type scope struct {
	converter *converter
	srcStack  scopeStack
	destStack scopeStack
}

func (s *scope) Convert(src, dest interface{}) error {
	return s.converter.Convert(src, dest)
}

// describe prints the path to get to the current (source, dest) values.
func (s *scope) describe() (src, dest string) {
	return s.srcStack.describe(), s.destStack.describe()
}

// error makes an error that includes information about where we were in the objects
// we were asked to convert.
func (s *scope) errorf(message string, args ...interface{}) error {
	srcPath, destPath := s.describe()
	where := fmt.Sprintf("converting %v to %v: ", srcPath, destPath)
	return fmt.Errorf(where+message, args...)
}

// scopeStackElement contains meta-information about a depth level.
type scopeStackElement struct {
	value reflect.Value
	key   string
}

// scopeStack contains an item for each depth level of object conversion.
type scopeStack []scopeStackElement

func (s *scopeStack) pop() {
	n := len(*s)
	*s = (*s)[:n-1]
}

func (s *scopeStack) push(e scopeStackElement) {
	*s = append(*s, e)
}

func (s *scopeStack) top() *scopeStackElement {
	return &(*s)[len(*s)-1]
}

func (s scopeStack) describe() string {
	desc := ""
	if len(s) > 1 {
		desc = "(" + s[1].value.Type().String() + ")"
	}
	for i, v := range s {
		if i < 2 {
			// First layer on stack is not real; second is handled specially above.
			continue
		}
		if v.key == "" {
			desc += fmt.Sprintf(".%v", v.value.Type())
		} else {
			desc += fmt.Sprintf(".%v", v.key)
		}
	}
	return desc
}
