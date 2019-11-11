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
	"time"

	"github.com/golang/protobuf/ptypes"
	"github.com/golang/protobuf/ptypes/timestamp"
)

// TimestampProto converts a time.Time into google.protobuf.Timestamp
func TimestampProto(t time.Time) (*timestamp.Timestamp, error) {
	if t.IsZero() {
		return nil, nil
	}
	return ptypes.TimestampProto(t)
}

// Timestamp converts google.protobuf.Timestamp into time.Time
func Timestamp(ts *timestamp.Timestamp) (time.Time, error) {
	if ts == nil {
		return time.Time{}, nil
	}
	return ptypes.Timestamp(ts)
}
