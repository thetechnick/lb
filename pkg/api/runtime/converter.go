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

// converter converts api object to other versions
type converter struct {
	conversionFuncs conversionFuncs
}

// newConverter creates a new Converter
func newConverter() *converter {
	c := &converter{
		conversionFuncs: newConversionFuncs(),
	}
	for _, dc := range defaultConversions {
		if err := c.RegisterConversionFunc(dc); err != nil {
			panic(err)
		}
	}
	return c
}

// RegisterConversionFunc adds a new conversion function to the lookup table.
func (c *converter) RegisterConversionFunc(conversionFunc interface{}) error {
	return c.conversionFuncs.Add(conversionFunc)
}

// Convert runs the registered type converter for the given type pair.
func (c *converter) Convert(src, dest interface{}) error {
	return c.doConvert(src, dest, c.convert)
}

func (c *converter) convert(sv, dv reflect.Value, scope *scope) error {
	// Convert sv to dv.
	dt, st := dv.Type(), sv.Type()
	pair := typePair{st, dt}
	if fv, ok := c.conversionFuncs.fns[pair]; ok {
		return c.callConverter(sv, dv, fv, scope)
	}

	return fmt.Errorf("no converter registered for conversion of %v to %v", dt, st)
}

type conversionFunc func(sv, dv reflect.Value, scope *scope) error

func (c *converter) doConvert(src, dest interface{}, f conversionFunc) error {
	dv, err := enforcePtr(dest)
	if err != nil {
		return fmt.Errorf("%v: for (src: %v) (dest: %v)", err, src, dest)
	}
	if !dv.CanAddr() && !dv.CanSet() {
		return fmt.Errorf("can't write to dest")
	}
	sv, err := enforcePtr(src)
	if err != nil {
		if err == ErrIsNil {
			return nil
		}
		return err
	}
	s := &scope{
		converter: c,
	}
	s.srcStack.push(scopeStackElement{})
	s.destStack.push(scopeStackElement{})
	return f(sv, dv, s)
}

func (c *converter) callConverter(sv, dv, custom reflect.Value, scope *scope) error {
	if !sv.CanAddr() {
		sv2 := reflect.New(sv.Type())
		sv2.Elem().Set(sv)
		sv = sv2
	} else {
		sv = sv.Addr()
	}
	if !dv.CanAddr() {
		if !dv.CanSet() {
			return scope.errorf("can't addr or set dest")
		}
		dvOrig := dv
		dvv := reflect.New(dvOrig.Type())
		defer func() { dvOrig.Set(dvv) }()
	} else {
		dv = dv.Addr()
	}
	args := []reflect.Value{sv, dv, reflect.ValueOf(scope)}
	ret := custom.Call(args)[0].Interface()
	// This convolution is necessary because nil interfaces won't convert
	// to errors.
	if ret == nil {
		return nil
	}
	return ret.(error)
}

type typePair struct {
	source reflect.Type
	dest   reflect.Type
}

// conversionFuncs holds the typePair to conversionFunc mapping.
type conversionFuncs struct {
	fns map[typePair]reflect.Value
}

// newConversionFuncs creates a new conversionFuncs mapping.
func newConversionFuncs() conversionFuncs {
	return conversionFuncs{fns: make(map[typePair]reflect.Value)}
}

// Add adds the provided conversion functions to the lookup table - they must have the signature
// `func(type1, type2, Scope) error`. Functions are added in the order passed and will override
// previously registered pairs.
func (c conversionFuncs) Add(fns ...interface{}) error {
	for _, fn := range fns {
		fv := reflect.ValueOf(fn)
		ft := fv.Type()
		if err := verifyConversionFunctionSignature(ft); err != nil {
			return err
		}
		c.fns[typePair{ft.In(0).Elem(), ft.In(1).Elem()}] = fv
	}
	return nil
}
