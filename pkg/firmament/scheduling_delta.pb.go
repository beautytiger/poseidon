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

// Code generated by protoc-gen-go.
// source: scheduling_delta.proto
// DO NOT EDIT!

package firmament

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type SchedulingDelta_ChangeType int32

const (
	SchedulingDelta_NOOP    SchedulingDelta_ChangeType = 0
	SchedulingDelta_PLACE   SchedulingDelta_ChangeType = 1
	SchedulingDelta_PREEMPT SchedulingDelta_ChangeType = 2
	SchedulingDelta_MIGRATE SchedulingDelta_ChangeType = 3
)

var SchedulingDelta_ChangeType_name = map[int32]string{
	0: "NOOP",
	1: "PLACE",
	2: "PREEMPT",
	3: "MIGRATE",
}
var SchedulingDelta_ChangeType_value = map[string]int32{
	"NOOP":    0,
	"PLACE":   1,
	"PREEMPT": 2,
	"MIGRATE": 3,
}

func (x SchedulingDelta_ChangeType) String() string {
	return proto.EnumName(SchedulingDelta_ChangeType_name, int32(x))
}
func (SchedulingDelta_ChangeType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor10, []int{0, 0}
}

type SchedulingDelta struct {
	TaskId     uint64                     `protobuf:"varint,1,opt,name=task_id,json=taskId" json:"task_id,omitempty"`
	ResourceId string                     `protobuf:"bytes,2,opt,name=resource_id,json=resourceId" json:"resource_id,omitempty"`
	Type       SchedulingDelta_ChangeType `protobuf:"varint,3,opt,name=type,enum=firmament.SchedulingDelta_ChangeType" json:"type,omitempty"`
}

func (m *SchedulingDelta) Reset()                    { *m = SchedulingDelta{} }
func (m *SchedulingDelta) String() string            { return proto.CompactTextString(m) }
func (*SchedulingDelta) ProtoMessage()               {}
func (*SchedulingDelta) Descriptor() ([]byte, []int) { return fileDescriptor10, []int{0} }

func (m *SchedulingDelta) GetTaskId() uint64 {
	if m != nil {
		return m.TaskId
	}
	return 0
}

func (m *SchedulingDelta) GetResourceId() string {
	if m != nil {
		return m.ResourceId
	}
	return ""
}

func (m *SchedulingDelta) GetType() SchedulingDelta_ChangeType {
	if m != nil {
		return m.Type
	}
	return SchedulingDelta_NOOP
}

func init() {
	proto.RegisterType((*SchedulingDelta)(nil), "firmament.SchedulingDelta")
	proto.RegisterEnum("firmament.SchedulingDelta_ChangeType", SchedulingDelta_ChangeType_name, SchedulingDelta_ChangeType_value)
}

func init() { proto.RegisterFile("scheduling_delta.proto", fileDescriptor10) }

var fileDescriptor10 = []byte{
	// 210 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xe2, 0x12, 0x2b, 0x4e, 0xce, 0x48,
	0x4d, 0x29, 0xcd, 0xc9, 0xcc, 0x4b, 0x8f, 0x4f, 0x49, 0xcd, 0x29, 0x49, 0xd4, 0x2b, 0x28, 0xca,
	0x2f, 0xc9, 0x17, 0xe2, 0x4c, 0xcb, 0x2c, 0xca, 0x4d, 0xcc, 0x4d, 0xcd, 0x2b, 0x51, 0x3a, 0xcc,
	0xc8, 0xc5, 0x1f, 0x0c, 0x57, 0xe5, 0x02, 0x52, 0x24, 0x24, 0xce, 0xc5, 0x5e, 0x92, 0x58, 0x9c,
	0x1d, 0x9f, 0x99, 0x22, 0xc1, 0xa8, 0xc0, 0xa8, 0xc1, 0x12, 0xc4, 0x06, 0xe2, 0x7a, 0xa6, 0x08,
	0xc9, 0x73, 0x71, 0x17, 0xa5, 0x16, 0xe7, 0x97, 0x16, 0x25, 0xa7, 0x82, 0x24, 0x99, 0x14, 0x18,
	0x35, 0x38, 0x83, 0xb8, 0x60, 0x42, 0x9e, 0x29, 0x42, 0x96, 0x5c, 0x2c, 0x25, 0x95, 0x05, 0xa9,
	0x12, 0xcc, 0x0a, 0x8c, 0x1a, 0x7c, 0x46, 0xaa, 0x7a, 0x70, 0x7b, 0xf4, 0xd0, 0xec, 0xd0, 0x73,
	0xce, 0x48, 0xcc, 0x4b, 0x4f, 0x0d, 0xa9, 0x2c, 0x48, 0x0d, 0x02, 0x6b, 0x51, 0xb2, 0xe6, 0xe2,
	0x42, 0x88, 0x09, 0x71, 0x70, 0xb1, 0xf8, 0xf9, 0xfb, 0x07, 0x08, 0x30, 0x08, 0x71, 0x72, 0xb1,
	0x06, 0xf8, 0x38, 0x3a, 0xbb, 0x0a, 0x30, 0x0a, 0x71, 0x73, 0xb1, 0x07, 0x04, 0xb9, 0xba, 0xfa,
	0x06, 0x84, 0x08, 0x30, 0x81, 0x38, 0xbe, 0x9e, 0xee, 0x41, 0x8e, 0x21, 0xae, 0x02, 0xcc, 0x49,
	0x6c, 0x60, 0x7f, 0x19, 0x03, 0x02, 0x00, 0x00, 0xff, 0xff, 0x8d, 0x3f, 0xf1, 0x44, 0xf1, 0x00,
	0x00, 0x00,
}
