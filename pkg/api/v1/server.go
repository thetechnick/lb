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

var serverConverters = []interface{}{
	// ServerList
	func(in *ServerList, out *api.ServerList, s runtime.Scope) (err error) {
		if err = s.Convert(in.Metadata, &out.ListMeta); err != nil {
			return
		}
		if err = s.Convert(&in.Items, &out.Items); err != nil {
			return
		}
		return
	},
	func(in *api.ServerList, out *ServerList, s runtime.Scope) (err error) {
		out.Metadata = &ListMeta{}
		if err = s.Convert(&in.ListMeta, out.Metadata); err != nil {
			return
		}
		if err = s.Convert(&in.Items, &out.Items); err != nil {
			return
		}
		return
	},

	// []Server
	func(in *[]*Server, out *[]api.Server, s runtime.Scope) (err error) {
		if in == nil {
			return
		}
		for _, f := range *in {
			outServer := &api.Server{}
			if err = s.Convert(f, outServer); err != nil {
				return
			}
			*out = append(*out, *outServer)
		}
		return
	},
	func(in *[]api.Server, out *[]*Server, s runtime.Scope) (err error) {
		if in == nil {
			return
		}
		for _, f := range *in {
			outServer := &Server{}
			if err = s.Convert(f, outServer); err != nil {
				return
			}
			*out = append(*out, outServer)
		}
		return
	},

	// Server
	func(in *Server, out *api.Server, s runtime.Scope) (err error) {
		if err = s.Convert(in.Metadata, &out.ObjectMeta); err != nil {
			return
		}
		if in.Spec != nil {
			if err = s.Convert(in.Spec, &out.Spec); err != nil {
				return
			}
		}
		if in.Status != nil {
			if err = s.Convert(in.Status, &out.Status); err != nil {
				return
			}
		}
		return
	},
	func(in *api.Server, out *Server, s runtime.Scope) (err error) {
		out.Metadata = &ObjectMeta{}
		if err = s.Convert(&in.ObjectMeta, out.Metadata); err != nil {
			return
		}
		out.Spec = &ServerSpec{}
		if err = s.Convert(&in.Spec, out.Spec); err != nil {
			return
		}
		out.Status = &ServerStatus{}
		if err = s.Convert(&in.Status, out.Status); err != nil {
			return
		}
		return
	},

	// ServerSpec
	func(in *ServerSpec, out *api.ServerSpec, s runtime.Scope) (err error) {
		out.Endpoints = in.Endpoints
		if err = s.Convert(&in.Ports, &out.Ports); err != nil {
			return
		}
		return
	},
	func(in *api.ServerSpec, out *ServerSpec, s runtime.Scope) (err error) {
		out.Endpoints = in.Endpoints
		if err = s.Convert(&in.Ports, &out.Ports); err != nil {
			return
		}
		return
	},

	// []ServerPort
	func(in *[]*ServerPort, out *[]api.ServerPort, s runtime.Scope) (err error) {
		if in == nil {
			return
		}
		for _, f := range *in {
			outServerPort := &api.ServerPort{}
			if err = s.Convert(f, outServerPort); err != nil {
				return
			}
			*out = append(*out, *outServerPort)
		}
		return
	},
	func(in *[]api.ServerPort, out *[]*ServerPort, s runtime.Scope) (err error) {
		if in == nil {
			return
		}
		for _, f := range *in {
			outServerPort := &ServerPort{}
			if err = s.Convert(&f, outServerPort); err != nil {
				return
			}
			*out = append(*out, outServerPort)
		}
		return
	},

	// ServerPort
	func(in *ServerPort, out *api.ServerPort, s runtime.Scope) (err error) {
		out.Port = in.Port
		out.Protocol = in.Protocol
		return
	},
	func(in *api.ServerPort, out *ServerPort, s runtime.Scope) (err error) {
		out.Port = in.Port
		out.Protocol = in.Protocol
		return
	},

	// ServerStatus
	func(in *ServerStatus, out *api.ServerStatus, s runtime.Scope) (err error) {
		if err = s.Convert(&in.Conditions, &out.Conditions); err != nil {
			return
		}
		if err = s.Convert(&in.Ingress, &out.Ingress); err != nil {
			return
		}
		return
	},
	func(in *api.ServerStatus, out *ServerStatus, s runtime.Scope) (err error) {
		if err = s.Convert(&in.Conditions, &out.Conditions); err != nil {
			return
		}
		if err = s.Convert(&in.Ingress, &out.Ingress); err != nil {
			return
		}
		return
	},

	// []ServerIngress
	func(in *[]*ServerIngress, out *[]api.ServerIngress, s runtime.Scope) (err error) {
		if in == nil {
			return
		}
		for _, f := range *in {
			outServerIngress := &api.ServerIngress{}
			if err = s.Convert(f, outServerIngress); err != nil {
				return
			}
			*out = append(*out, *outServerIngress)
		}
		return
	},
	func(in *[]api.ServerIngress, out *[]*ServerIngress, s runtime.Scope) (err error) {
		if in == nil {
			return
		}
		for _, f := range *in {
			outServerIngress := &ServerIngress{}
			if err = s.Convert(&f, outServerIngress); err != nil {
				return
			}
			*out = append(*out, outServerIngress)
		}
		return
	},

	// ServerIngress
	func(in *ServerIngress, out *api.ServerIngress, s runtime.Scope) (err error) {
		out.IP = in.Ip
		return
	},
	func(in *api.ServerIngress, out *ServerIngress, s runtime.Scope) (err error) {
		out.Ip = in.IP
		return
	},
}
