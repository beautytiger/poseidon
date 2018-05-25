/*
Copyright 2018 The Kubernetes Authors.

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
// source: label_selector.proto

package firmament

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type LabelSelector_SelectorType int32

const (
	LabelSelector_IN_SET         LabelSelector_SelectorType = 0
	LabelSelector_NOT_IN_SET     LabelSelector_SelectorType = 1
	LabelSelector_EXISTS_KEY     LabelSelector_SelectorType = 2
	LabelSelector_NOT_EXISTS_KEY LabelSelector_SelectorType = 3
)

var LabelSelector_SelectorType_name = map[int32]string{
	0: "IN_SET",
	1: "NOT_IN_SET",
	2: "EXISTS_KEY",
	3: "NOT_EXISTS_KEY",
}
var LabelSelector_SelectorType_value = map[string]int32{
	"IN_SET":         0,
	"NOT_IN_SET":     1,
	"EXISTS_KEY":     2,
	"NOT_EXISTS_KEY": 3,
}

func (x LabelSelector_SelectorType) String() string {
	return proto.EnumName(LabelSelector_SelectorType_name, int32(x))
}
func (LabelSelector_SelectorType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor5, []int{0, 0}
}

// LabelSelector can be used by client/user to identify a set of objects if their Labels match the LabelSelector.
// Set-based label requirements allow filtering keys according to a set of values. Four kinds of operators are supported:
// in, notin, exists (only the key identifier) and notexist.
type LabelSelector struct {
	Type   LabelSelector_SelectorType `protobuf:"varint,1,opt,name=type,enum=firmament.LabelSelector_SelectorType" json:"type,omitempty"`
	Key    string                     `protobuf:"bytes,2,opt,name=key" json:"key,omitempty"`
	Values []string                   `protobuf:"bytes,3,rep,name=values" json:"values,omitempty"`
}

func (m *LabelSelector) Reset()                    { *m = LabelSelector{} }
func (m *LabelSelector) String() string            { return proto.CompactTextString(m) }
func (*LabelSelector) ProtoMessage()               {}
func (*LabelSelector) Descriptor() ([]byte, []int) { return fileDescriptor5, []int{0} }

func (m *LabelSelector) GetType() LabelSelector_SelectorType {
	if m != nil {
		return m.Type
	}
	return LabelSelector_IN_SET
}

func (m *LabelSelector) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *LabelSelector) GetValues() []string {
	if m != nil {
		return m.Values
	}
	return nil
}

func init() {
	proto.RegisterType((*LabelSelector)(nil), "firmament.LabelSelector")
	proto.RegisterEnum("firmament.LabelSelector_SelectorType", LabelSelector_SelectorType_name, LabelSelector_SelectorType_value)
}

func init() { proto.RegisterFile("label_selector.proto", fileDescriptor5) }

var fileDescriptor5 = []byte{
	// 191 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0xc9, 0x49, 0x4c, 0x4a,
	0xcd, 0x89, 0x2f, 0x4e, 0xcd, 0x49, 0x4d, 0x2e, 0xc9, 0x2f, 0xd2, 0x2b, 0x28, 0xca, 0x2f, 0xc9,
	0x17, 0xe2, 0x4c, 0xcb, 0x2c, 0xca, 0x4d, 0xcc, 0x4d, 0xcd, 0x2b, 0x51, 0x3a, 0xc2, 0xc8, 0xc5,
	0xeb, 0x03, 0x52, 0x13, 0x0c, 0x55, 0x22, 0x64, 0xc9, 0xc5, 0x52, 0x52, 0x59, 0x90, 0x2a, 0xc1,
	0xa8, 0xc0, 0xa8, 0xc1, 0x67, 0xa4, 0xaa, 0x07, 0x57, 0xab, 0x87, 0xa2, 0x4e, 0x0f, 0xc6, 0x08,
	0xa9, 0x2c, 0x48, 0x0d, 0x02, 0x6b, 0x11, 0x12, 0xe0, 0x62, 0xce, 0x4e, 0xad, 0x94, 0x60, 0x52,
	0x60, 0xd4, 0xe0, 0x0c, 0x02, 0x31, 0x85, 0xc4, 0xb8, 0xd8, 0xca, 0x12, 0x73, 0x4a, 0x53, 0x8b,
	0x25, 0x98, 0x15, 0x98, 0x35, 0x38, 0x83, 0xa0, 0x3c, 0x25, 0x3f, 0x2e, 0x1e, 0x64, 0xfd, 0x42,
	0x5c, 0x5c, 0x6c, 0x9e, 0x7e, 0xf1, 0xc1, 0xae, 0x21, 0x02, 0x0c, 0x42, 0x7c, 0x5c, 0x5c, 0x7e,
	0xfe, 0x21, 0xf1, 0x50, 0x3e, 0x23, 0x88, 0xef, 0x1a, 0xe1, 0x19, 0x1c, 0x12, 0x1c, 0xef, 0xed,
	0x1a, 0x29, 0xc0, 0x24, 0x24, 0xc4, 0xc5, 0x07, 0x92, 0x47, 0x12, 0x63, 0x4e, 0x62, 0x03, 0x7b,
	0xcc, 0x18, 0x10, 0x00, 0x00, 0xff, 0xff, 0x10, 0xf2, 0x09, 0xce, 0xf0, 0x00, 0x00, 0x00,
}
