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
// Code generated by protoc-gen-go. DO NOT EDIT.
// source: meta.proto

package v1

import (
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type ListMeta struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListMeta) Reset()         { *m = ListMeta{} }
func (m *ListMeta) String() string { return proto.CompactTextString(m) }
func (*ListMeta) ProtoMessage()    {}
func (*ListMeta) Descriptor() ([]byte, []int) {
	return fileDescriptor_3b5ea8fe65782bcc, []int{0}
}

func (m *ListMeta) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListMeta.Unmarshal(m, b)
}
func (m *ListMeta) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListMeta.Marshal(b, m, deterministic)
}
func (m *ListMeta) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListMeta.Merge(m, src)
}
func (m *ListMeta) XXX_Size() int {
	return xxx_messageInfo_ListMeta.Size(m)
}
func (m *ListMeta) XXX_DiscardUnknown() {
	xxx_messageInfo_ListMeta.DiscardUnknown(m)
}

var xxx_messageInfo_ListMeta proto.InternalMessageInfo

type ObjectMeta struct {
	Uid                  string               `protobuf:"bytes,1,opt,name=uid,proto3" json:"uid,omitempty"`
	Name                 string               `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Created              *timestamp.Timestamp `protobuf:"bytes,3,opt,name=created,proto3" json:"created,omitempty"`
	Updated              *timestamp.Timestamp `protobuf:"bytes,4,opt,name=updated,proto3" json:"updated,omitempty"`
	ResourceVersion      string               `protobuf:"bytes,5,opt,name=resourceVersion,proto3" json:"resourceVersion,omitempty"`
	Generation           uint64               `protobuf:"varint,6,opt,name=generation,proto3" json:"generation,omitempty"`
	Annotations          map[string]string    `protobuf:"bytes,7,rep,name=annotations,proto3" json:"annotations,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *ObjectMeta) Reset()         { *m = ObjectMeta{} }
func (m *ObjectMeta) String() string { return proto.CompactTextString(m) }
func (*ObjectMeta) ProtoMessage()    {}
func (*ObjectMeta) Descriptor() ([]byte, []int) {
	return fileDescriptor_3b5ea8fe65782bcc, []int{1}
}

func (m *ObjectMeta) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ObjectMeta.Unmarshal(m, b)
}
func (m *ObjectMeta) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ObjectMeta.Marshal(b, m, deterministic)
}
func (m *ObjectMeta) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ObjectMeta.Merge(m, src)
}
func (m *ObjectMeta) XXX_Size() int {
	return xxx_messageInfo_ObjectMeta.Size(m)
}
func (m *ObjectMeta) XXX_DiscardUnknown() {
	xxx_messageInfo_ObjectMeta.DiscardUnknown(m)
}

var xxx_messageInfo_ObjectMeta proto.InternalMessageInfo

func (m *ObjectMeta) GetUid() string {
	if m != nil {
		return m.Uid
	}
	return ""
}

func (m *ObjectMeta) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *ObjectMeta) GetCreated() *timestamp.Timestamp {
	if m != nil {
		return m.Created
	}
	return nil
}

func (m *ObjectMeta) GetUpdated() *timestamp.Timestamp {
	if m != nil {
		return m.Updated
	}
	return nil
}

func (m *ObjectMeta) GetResourceVersion() string {
	if m != nil {
		return m.ResourceVersion
	}
	return ""
}

func (m *ObjectMeta) GetGeneration() uint64 {
	if m != nil {
		return m.Generation
	}
	return 0
}

func (m *ObjectMeta) GetAnnotations() map[string]string {
	if m != nil {
		return m.Annotations
	}
	return nil
}

func init() {
	proto.RegisterType((*ListMeta)(nil), "lb.api.v1.ListMeta")
	proto.RegisterType((*ObjectMeta)(nil), "lb.api.v1.ObjectMeta")
	proto.RegisterMapType((map[string]string)(nil), "lb.api.v1.ObjectMeta.AnnotationsEntry")
}

func init() { proto.RegisterFile("meta.proto", fileDescriptor_3b5ea8fe65782bcc) }

var fileDescriptor_3b5ea8fe65782bcc = []byte{
	// 289 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x90, 0xc1, 0x4b, 0x84, 0x40,
	0x18, 0xc5, 0xd1, 0x75, 0x77, 0xdb, 0xcf, 0x43, 0xcb, 0xd0, 0x41, 0x3c, 0x94, 0xec, 0x21, 0x3c,
	0xcd, 0xb2, 0x5b, 0x87, 0xe8, 0x10, 0x14, 0x04, 0x1d, 0x8a, 0x40, 0xa2, 0x43, 0xb7, 0x51, 0xbf,
	0xc4, 0x56, 0x67, 0x64, 0xfc, 0x14, 0xfc, 0xb7, 0xfb, 0x0b, 0xc2, 0x31, 0xdb, 0x65, 0x2f, 0xdd,
	0xbe, 0xf7, 0xfc, 0x3d, 0x9e, 0xf3, 0x00, 0x4a, 0x24, 0xc1, 0x2b, 0xad, 0x48, 0xb1, 0x45, 0x11,
	0x73, 0x51, 0xe5, 0xbc, 0xdd, 0xf8, 0x17, 0x99, 0x52, 0x59, 0x81, 0x6b, 0xf3, 0x21, 0x6e, 0x3e,
	0xd7, 0x94, 0x97, 0x58, 0x93, 0x28, 0xab, 0x81, 0x5d, 0x01, 0x9c, 0x3c, 0xe7, 0x35, 0xbd, 0x20,
	0x89, 0xd5, 0xb7, 0x0d, 0xf0, 0x1a, 0x7f, 0x61, 0x62, 0x24, 0x5b, 0xc2, 0xa4, 0xc9, 0x53, 0xcf,
	0x0a, 0xac, 0x70, 0x11, 0xf5, 0x27, 0x63, 0xe0, 0x48, 0x51, 0xa2, 0x67, 0x1b, 0xcb, 0xdc, 0xec,
	0x1a, 0xe6, 0x89, 0x46, 0x41, 0x98, 0x7a, 0x93, 0xc0, 0x0a, 0xdd, 0xad, 0xcf, 0x87, 0x4e, 0x3e,
	0x76, 0xf2, 0xb7, 0xb1, 0x33, 0x1a, 0xd1, 0x3e, 0xd5, 0x54, 0xa9, 0x49, 0x39, 0xff, 0xa7, 0x7e,
	0x51, 0x16, 0xc2, 0xa9, 0xc6, 0x5a, 0x35, 0x3a, 0xc1, 0x77, 0xd4, 0x75, 0xae, 0xa4, 0x37, 0x35,
	0xbf, 0x72, 0x6c, 0xb3, 0x73, 0x80, 0x0c, 0x25, 0x6a, 0x41, 0x3d, 0x34, 0x0b, 0xac, 0xd0, 0x89,
	0x0e, 0x1c, 0xf6, 0x04, 0xae, 0x90, 0x52, 0x91, 0x51, 0xb5, 0x37, 0x0f, 0x26, 0xa1, 0xbb, 0xbd,
	0xe4, 0x7f, 0xc3, 0xf1, 0xfd, 0x0e, 0xfc, 0x7e, 0x0f, 0x3e, 0x4a, 0xd2, 0x5d, 0x74, 0x18, 0xf5,
	0xef, 0x60, 0x79, 0x0c, 0xf4, 0xcb, 0xed, 0xb0, 0x1b, 0x97, 0xdb, 0x61, 0xc7, 0xce, 0x60, 0xda,
	0x8a, 0xa2, 0x19, 0xa7, 0x1b, 0xc4, 0xad, 0x7d, 0x63, 0x3d, 0x38, 0x1f, 0x76, 0xbb, 0x89, 0x67,
	0xe6, 0xd9, 0x57, 0x3f, 0x01, 0x00, 0x00, 0xff, 0xff, 0xc1, 0xf0, 0xe5, 0x90, 0xc7, 0x01, 0x00,
	0x00,
}
