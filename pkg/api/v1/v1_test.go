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
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/stretchr/testify/require"

	"github.com/thetechnick/lb/pkg/api"
)

func TestServer(t *testing.T) {
	apiObj := &api.Server{
		Spec: api.ServerSpec{
			Endpoints: []string{"1.2.3.4"},
			Ports: []api.ServerPort{
				{Port: 22, Protocol: "tcp"},
			},
		},
		Status: api.ServerStatus{
			Conditions: api.NewConditionMapCollection([]api.Condition{
				{Status: api.ConditionStatusTrue},
			}),
			Ingress: []api.ServerIngress{
				{IP: "1.2.3.4"},
			},
		},
	}
	v1Obj := &Server{
		Metadata: &ObjectMeta{},
		Spec: &ServerSpec{
			Endpoints: []string{"1.2.3.4"},
			Ports: []*ServerPort{
				{Port: 22, Protocol: "tcp"},
			},
		},
		Status: &ServerStatus{
			Conditions: []*Condition{{
				Status: Condition_True,
			}},
			Ingress: []*ServerIngress{
				{Ip: "1.2.3.4"},
			},
		},
	}

	t.Run("from api.Server to v1.Server", func(t *testing.T) {
		v1ObjOut := &Server{}

		err := Scheme.Convert(apiObj, v1ObjOut)
		require.NoError(t, err)
		assert.Equal(t, v1Obj, v1ObjOut)
	})

	t.Run("from v1.Server to api.Server", func(t *testing.T) {
		apiObjOut := &api.Server{}

		err := Scheme.Convert(v1Obj, apiObjOut)
		require.NoError(t, err)
		assert.Equal(t, apiObj, apiObjOut)
	})
}
