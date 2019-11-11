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

import (
	"github.com/thetechnick/lb/pkg/api"
	"github.com/thetechnick/lb/pkg/api/runtime"
)

var metaConverters = []interface{}{
	// ListMeta
	func(in *ListMeta, out *api.ListMeta, s runtime.Scope) (err error) {
		return
	},
	func(in *api.ListMeta, out *ListMeta, s runtime.Scope) (err error) {
		return
	},

	// ObjectMeta
	func(in *ObjectMeta, out *api.ObjectMeta, s runtime.Scope) (err error) {
		out.UID = in.Uid
		out.Name = in.Name
		if out.Created, err = Timestamp(in.Created); err != nil {
			return
		}
		if out.Updated, err = Timestamp(in.Updated); err != nil {
			return
		}
		out.ResourceVersion = in.ResourceVersion
		out.Generation = in.Generation
		out.Annotations = in.Annotations
		return
	},
	func(in *api.ObjectMeta, out *ObjectMeta, s runtime.Scope) (err error) {
		out.Uid = in.UID
		out.Name = in.Name
		if out.Created, err = TimestampProto(in.Created); err != nil {
			return
		}
		if out.Updated, err = TimestampProto(in.Updated); err != nil {
			return
		}
		out.ResourceVersion = in.ResourceVersion
		out.Generation = in.Generation
		out.Annotations = in.Annotations
		return
	},
}
