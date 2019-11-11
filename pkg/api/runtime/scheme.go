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

// Scheme defines methods for serializing and deserializing objects, a type
// registry for converting group, version, and kind information to and from Go
// schemas, and mappings between Go schemas of different versions.
type Scheme struct {
	// converter stores all registered conversion functions. It also has
	// default coverting behavior.
	converter *converter
}

// NewScheme creates a new Scheme. This scheme is pluggable by default.
func NewScheme() *Scheme {
	s := &Scheme{
		converter: newConverter(),
	}
	return s
}

// Convert will attempt to convert in into out. Both must be pointers. For easy
// testing of conversion functions. Returns an error if the conversion isn't
// possible. You can call this with types that haven't been registered (for example,
// a to test conversion of types that are nested within registered types). The
// context interface is passed to the convertor.
func (s *Scheme) Convert(in, out interface{}) error {
	return s.converter.Convert(in, out)
}

// AddConversionFuncs adds functions to the list of conversion functions. The given
// functions should know how to convert between two of your API objects, or their
// sub-objects.
func (s *Scheme) AddConversionFuncs(conversionFuncs ...interface{}) error {
	for _, f := range conversionFuncs {
		if err := s.converter.RegisterConversionFunc(f); err != nil {
			return err
		}
	}
	return nil
}
