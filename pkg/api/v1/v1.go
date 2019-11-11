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

package v1

//go:generate bash -c "protoc --go_out=plugins=grpc:. -I../../../protoc/include -I. *.proto"
//go:generate bash -c "for file in *.pb.go; do cat ../../../hack/boilerplate/boilerplate.generatego.txt | sed s/YEAR/2019/ | cat - $DOLLAR{file} > $DOLLAR{file}.tmp; mv $DOLLAR{file}.tmp $DOLLAR{file}; done"

import "github.com/thetechnick/lb/pkg/api/runtime"

var (
	Scheme *runtime.Scheme
)

func init() {
	Scheme = runtime.NewScheme()

	var converters []interface{}
	converters = append(converters, metaConverters...)
	converters = append(converters, conditionConverters...)
	converters = append(converters, serverConverters...)

	err := Scheme.AddConversionFuncs(converters...)
	if err != nil {
		panic(err)
	}
}
