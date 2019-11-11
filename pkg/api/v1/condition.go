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

var conditionConverters = []interface{}{
	// ConditionCollection
	func(in *[]*Condition, out *api.ConditionCollection, s runtime.Scope) (err error) {
		m := api.ConditionMapCollection{}
		*out = m
		if in == nil {
			return
		}
		for _, f := range *in {
			outCondition := api.Condition{}
			if err = s.Convert(f, &outCondition); err != nil {
				return
			}
			m[outCondition.Type] = outCondition
		}
		return
	},
	func(in *api.ConditionCollection, out *[]*Condition, s runtime.Scope) (err error) {
		if in == nil || *in == nil {
			return
		}
		for _, f := range (*in).List() {
			outCondition := &Condition{}
			if err = s.Convert(&f, outCondition); err != nil {
				return
			}
			*out = append(*out, outCondition)
		}
		return
	},

	// []Condition
	func(in *[]*Condition, out *[]api.Condition, s runtime.Scope) (err error) {
		if in == nil {
			return
		}
		for _, f := range *in {
			outCondition := api.Condition{}
			if err = s.Convert(f, &outCondition); err != nil {
				return
			}
			*out = append(*out, outCondition)
		}
		return
	},
	func(in *[]api.Condition, out *[]*Condition, s runtime.Scope) (err error) {
		if in == nil {
			return
		}
		for _, f := range *in {
			outCondition := &Condition{}
			if err = s.Convert(&f, outCondition); err != nil {
				return
			}
			*out = append(*out, outCondition)
		}
		return
	},

	// Condition
	func(in *Condition, out *api.Condition, s runtime.Scope) (err error) {
		out.Type = in.Type
		out.Reason = in.Reason
		out.Message = in.Message
		if out.LastTransitionTime, err = Timestamp(in.LastTransitionTime); err != nil {
			return
		}
		if err = s.Convert(&in.Status, &out.Status); err != nil {
			return
		}
		return
	},
	func(in *api.Condition, out *Condition, s runtime.Scope) (err error) {
		out.Type = in.Type
		out.Reason = in.Reason
		out.Message = in.Message
		if out.LastTransitionTime, err = TimestampProto(in.LastTransitionTime); err != nil {
			return
		}
		if err = s.Convert(&in.Status, &out.Status); err != nil {
			return
		}
		return
	},

	// ConditionStatus
	func(in *Condition_Status, out *api.ConditionStatus, s runtime.Scope) (err error) {
		n, ok := Condition_Status_name[int32(*in)]
		if !ok {
			*out = api.ConditionStatusUnknown
			return
		}
		*out = api.ConditionStatus(n)
		return
	},
	func(in *api.ConditionStatus, out *Condition_Status, s runtime.Scope) (err error) {
		v, ok := Condition_Status_value[string(*in)]
		if !ok {
			*out = Condition_Unknown
			return
		}
		*out = Condition_Status(v)
		return
	},
}
