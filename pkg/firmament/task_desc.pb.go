// Code generated by protoc-gen-go.
// source: task_desc.proto
// DO NOT EDIT!

package firmament

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type TaskDescriptor_TaskState int32

const (
	TaskDescriptor_CREATED   TaskDescriptor_TaskState = 0
	TaskDescriptor_BLOCKING  TaskDescriptor_TaskState = 1
	TaskDescriptor_RUNNABLE  TaskDescriptor_TaskState = 2
	TaskDescriptor_ASSIGNED  TaskDescriptor_TaskState = 3
	TaskDescriptor_RUNNING   TaskDescriptor_TaskState = 4
	TaskDescriptor_COMPLETED TaskDescriptor_TaskState = 5
	TaskDescriptor_FAILED    TaskDescriptor_TaskState = 6
	TaskDescriptor_ABORTED   TaskDescriptor_TaskState = 7
	TaskDescriptor_DELEGATED TaskDescriptor_TaskState = 8
	TaskDescriptor_UNKNOWN   TaskDescriptor_TaskState = 9
)

var TaskDescriptor_TaskState_name = map[int32]string{
	0: "CREATED",
	1: "BLOCKING",
	2: "RUNNABLE",
	3: "ASSIGNED",
	4: "RUNNING",
	5: "COMPLETED",
	6: "FAILED",
	7: "ABORTED",
	8: "DELEGATED",
	9: "UNKNOWN",
}
var TaskDescriptor_TaskState_value = map[string]int32{
	"CREATED":   0,
	"BLOCKING":  1,
	"RUNNABLE":  2,
	"ASSIGNED":  3,
	"RUNNING":   4,
	"COMPLETED": 5,
	"FAILED":    6,
	"ABORTED":   7,
	"DELEGATED": 8,
	"UNKNOWN":   9,
}

func (x TaskDescriptor_TaskState) String() string {
	return proto.EnumName(TaskDescriptor_TaskState_name, int32(x))
}
func (TaskDescriptor_TaskState) EnumDescriptor() ([]byte, []int) { return fileDescriptor11, []int{0, 0} }

type TaskDescriptor_TaskType int32

const (
	TaskDescriptor_SHEEP  TaskDescriptor_TaskType = 0
	TaskDescriptor_RABBIT TaskDescriptor_TaskType = 1
	TaskDescriptor_DEVIL  TaskDescriptor_TaskType = 2
	TaskDescriptor_TURTLE TaskDescriptor_TaskType = 3
)

var TaskDescriptor_TaskType_name = map[int32]string{
	0: "SHEEP",
	1: "RABBIT",
	2: "DEVIL",
	3: "TURTLE",
}
var TaskDescriptor_TaskType_value = map[string]int32{
	"SHEEP":  0,
	"RABBIT": 1,
	"DEVIL":  2,
	"TURTLE": 3,
}

func (x TaskDescriptor_TaskType) String() string {
	return proto.EnumName(TaskDescriptor_TaskType_name, int32(x))
}
func (TaskDescriptor_TaskType) EnumDescriptor() ([]byte, []int) { return fileDescriptor11, []int{0, 1} }

type TaskDescriptor struct {
	Uid   uint64                   `protobuf:"varint,1,opt,name=uid" json:"uid,omitempty"`
	Name  string                   `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
	State TaskDescriptor_TaskState `protobuf:"varint,3,opt,name=state,enum=firmament.TaskDescriptor_TaskState" json:"state,omitempty"`
	JobId string                   `protobuf:"bytes,4,opt,name=job_id,json=jobId" json:"job_id,omitempty"`
	Index uint64                   `protobuf:"varint,5,opt,name=index" json:"index,omitempty"`
	// Inputs/outputs
	Dependencies []*ReferenceDescriptor `protobuf:"bytes,6,rep,name=dependencies" json:"dependencies,omitempty"`
	Outputs      []*ReferenceDescriptor `protobuf:"bytes,7,rep,name=outputs" json:"outputs,omitempty"`
	// Command and arguments
	Binary string   `protobuf:"bytes,8,opt,name=binary" json:"binary,omitempty"`
	Args   []string `protobuf:"bytes,9,rep,name=args" json:"args,omitempty"`
	// Children
	Spawned []*TaskDescriptor `protobuf:"bytes,10,rep,name=spawned" json:"spawned,omitempty"`
	// Runtime meta-data
	ScheduledToResource   string `protobuf:"bytes,11,opt,name=scheduled_to_resource,json=scheduledToResource" json:"scheduled_to_resource,omitempty"`
	LastHeartbeatLocation string `protobuf:"bytes,12,opt,name=last_heartbeat_location,json=lastHeartbeatLocation" json:"last_heartbeat_location,omitempty"`
	LastHeartbeatTime     uint64 `protobuf:"varint,13,opt,name=last_heartbeat_time,json=lastHeartbeatTime" json:"last_heartbeat_time,omitempty"`
	// Delegation info
	DelegatedTo   string `protobuf:"bytes,14,opt,name=delegated_to,json=delegatedTo" json:"delegated_to,omitempty"`
	DelegatedFrom string `protobuf:"bytes,15,opt,name=delegated_from,json=delegatedFrom" json:"delegated_from,omitempty"`
	// Timestamps
	SubmitTime uint64 `protobuf:"varint,16,opt,name=submit_time,json=submitTime" json:"submit_time,omitempty"`
	StartTime  uint64 `protobuf:"varint,17,opt,name=start_time,json=startTime" json:"start_time,omitempty"`
	FinishTime uint64 `protobuf:"varint,18,opt,name=finish_time,json=finishTime" json:"finish_time,omitempty"`
	// The total time spent unscheduled before previous runs.
	TotalUnscheduledTime uint64 `protobuf:"varint,19,opt,name=total_unscheduled_time,json=totalUnscheduledTime" json:"total_unscheduled_time,omitempty"`
	// The total time spent in previous runs. This field only gets updated when
	// the task finishes running.
	TotalRunTime uint64 `protobuf:"varint,20,opt,name=total_run_time,json=totalRunTime" json:"total_run_time,omitempty"`
	// Deadlines
	RelativeDeadline uint64 `protobuf:"varint,21,opt,name=relative_deadline,json=relativeDeadline" json:"relative_deadline,omitempty"`
	AbsoluteDeadline uint64 `protobuf:"varint,22,opt,name=absolute_deadline,json=absoluteDeadline" json:"absolute_deadline,omitempty"`
	// TODO(malte): move these to sub-messages
	Port      uint64 `protobuf:"varint,23,opt,name=port" json:"port,omitempty"`
	InputSize uint64 `protobuf:"varint,24,opt,name=input_size,json=inputSize" json:"input_size,omitempty"`
	// TaskLib related stuff
	InjectTaskLib bool `protobuf:"varint,25,opt,name=inject_task_lib,json=injectTaskLib" json:"inject_task_lib,omitempty"`
	// Task resource request and priority
	ResourceRequest *ResourceVector `protobuf:"bytes,26,opt,name=resource_request,json=resourceRequest" json:"resource_request,omitempty"`
	Priority        uint32          `protobuf:"varint,27,opt,name=priority" json:"priority,omitempty"`
	// TODO(malte): move this to a policy-specific sub-message
	TaskType TaskDescriptor_TaskType `protobuf:"varint,28,opt,name=task_type,json=taskType,enum=firmament.TaskDescriptor_TaskType" json:"task_type,omitempty"`
	// Final report after successful execution
	FinalReport *TaskFinalReport `protobuf:"bytes,29,opt,name=final_report,json=finalReport" json:"final_report,omitempty"`
	// Simulation related fields
	TraceJobId  uint64 `protobuf:"varint,30,opt,name=trace_job_id,json=traceJobId" json:"trace_job_id,omitempty"`
	TraceTaskId uint64 `protobuf:"varint,31,opt,name=trace_task_id,json=traceTaskId" json:"trace_task_id,omitempty"`
	// Task labels
	Labels []*Label `protobuf:"bytes,32,rep,name=labels" json:"labels,omitempty"`
	// Resource label selectors
	LabelSelectors []*LabelSelector `protobuf:"bytes,33,rep,name=label_selectors,json=labelSelectors" json:"label_selectors,omitempty"`
}

func (m *TaskDescriptor) Reset()                    { *m = TaskDescriptor{} }
func (m *TaskDescriptor) String() string            { return proto.CompactTextString(m) }
func (*TaskDescriptor) ProtoMessage()               {}
func (*TaskDescriptor) Descriptor() ([]byte, []int) { return fileDescriptor11, []int{0} }

func (m *TaskDescriptor) GetUid() uint64 {
	if m != nil {
		return m.Uid
	}
	return 0
}

func (m *TaskDescriptor) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *TaskDescriptor) GetState() TaskDescriptor_TaskState {
	if m != nil {
		return m.State
	}
	return TaskDescriptor_CREATED
}

func (m *TaskDescriptor) GetJobId() string {
	if m != nil {
		return m.JobId
	}
	return ""
}

func (m *TaskDescriptor) GetIndex() uint64 {
	if m != nil {
		return m.Index
	}
	return 0
}

func (m *TaskDescriptor) GetDependencies() []*ReferenceDescriptor {
	if m != nil {
		return m.Dependencies
	}
	return nil
}

func (m *TaskDescriptor) GetOutputs() []*ReferenceDescriptor {
	if m != nil {
		return m.Outputs
	}
	return nil
}

func (m *TaskDescriptor) GetBinary() string {
	if m != nil {
		return m.Binary
	}
	return ""
}

func (m *TaskDescriptor) GetArgs() []string {
	if m != nil {
		return m.Args
	}
	return nil
}

func (m *TaskDescriptor) GetSpawned() []*TaskDescriptor {
	if m != nil {
		return m.Spawned
	}
	return nil
}

func (m *TaskDescriptor) GetScheduledToResource() string {
	if m != nil {
		return m.ScheduledToResource
	}
	return ""
}

func (m *TaskDescriptor) GetLastHeartbeatLocation() string {
	if m != nil {
		return m.LastHeartbeatLocation
	}
	return ""
}

func (m *TaskDescriptor) GetLastHeartbeatTime() uint64 {
	if m != nil {
		return m.LastHeartbeatTime
	}
	return 0
}

func (m *TaskDescriptor) GetDelegatedTo() string {
	if m != nil {
		return m.DelegatedTo
	}
	return ""
}

func (m *TaskDescriptor) GetDelegatedFrom() string {
	if m != nil {
		return m.DelegatedFrom
	}
	return ""
}

func (m *TaskDescriptor) GetSubmitTime() uint64 {
	if m != nil {
		return m.SubmitTime
	}
	return 0
}

func (m *TaskDescriptor) GetStartTime() uint64 {
	if m != nil {
		return m.StartTime
	}
	return 0
}

func (m *TaskDescriptor) GetFinishTime() uint64 {
	if m != nil {
		return m.FinishTime
	}
	return 0
}

func (m *TaskDescriptor) GetTotalUnscheduledTime() uint64 {
	if m != nil {
		return m.TotalUnscheduledTime
	}
	return 0
}

func (m *TaskDescriptor) GetTotalRunTime() uint64 {
	if m != nil {
		return m.TotalRunTime
	}
	return 0
}

func (m *TaskDescriptor) GetRelativeDeadline() uint64 {
	if m != nil {
		return m.RelativeDeadline
	}
	return 0
}

func (m *TaskDescriptor) GetAbsoluteDeadline() uint64 {
	if m != nil {
		return m.AbsoluteDeadline
	}
	return 0
}

func (m *TaskDescriptor) GetPort() uint64 {
	if m != nil {
		return m.Port
	}
	return 0
}

func (m *TaskDescriptor) GetInputSize() uint64 {
	if m != nil {
		return m.InputSize
	}
	return 0
}

func (m *TaskDescriptor) GetInjectTaskLib() bool {
	if m != nil {
		return m.InjectTaskLib
	}
	return false
}

func (m *TaskDescriptor) GetResourceRequest() *ResourceVector {
	if m != nil {
		return m.ResourceRequest
	}
	return nil
}

func (m *TaskDescriptor) GetPriority() uint32 {
	if m != nil {
		return m.Priority
	}
	return 0
}

func (m *TaskDescriptor) GetTaskType() TaskDescriptor_TaskType {
	if m != nil {
		return m.TaskType
	}
	return TaskDescriptor_SHEEP
}

func (m *TaskDescriptor) GetFinalReport() *TaskFinalReport {
	if m != nil {
		return m.FinalReport
	}
	return nil
}

func (m *TaskDescriptor) GetTraceJobId() uint64 {
	if m != nil {
		return m.TraceJobId
	}
	return 0
}

func (m *TaskDescriptor) GetTraceTaskId() uint64 {
	if m != nil {
		return m.TraceTaskId
	}
	return 0
}

func (m *TaskDescriptor) GetLabels() []*Label {
	if m != nil {
		return m.Labels
	}
	return nil
}

func (m *TaskDescriptor) GetLabelSelectors() []*LabelSelector {
	if m != nil {
		return m.LabelSelectors
	}
	return nil
}

func init() {
	proto.RegisterType((*TaskDescriptor)(nil), "firmament.TaskDescriptor")
	proto.RegisterEnum("firmament.TaskDescriptor_TaskState", TaskDescriptor_TaskState_name, TaskDescriptor_TaskState_value)
	proto.RegisterEnum("firmament.TaskDescriptor_TaskType", TaskDescriptor_TaskType_name, TaskDescriptor_TaskType_value)
}

func init() { proto.RegisterFile("task_desc.proto", fileDescriptor11) }

var fileDescriptor11 = []byte{
	// 913 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x8c, 0x55, 0xdd, 0x6f, 0xdb, 0x36,
	0x10, 0xaf, 0xeb, 0xd8, 0xb1, 0xcf, 0x1f, 0x51, 0x98, 0x2f, 0xd6, 0x5b, 0x1b, 0x37, 0xfb, 0x80,
	0x81, 0x01, 0x79, 0x48, 0x87, 0xa1, 0x7b, 0x18, 0x06, 0x3b, 0x56, 0x52, 0xaf, 0x9e, 0x53, 0xd0,
	0x4e, 0xf7, 0x28, 0x50, 0x12, 0xdd, 0x30, 0x95, 0x25, 0x8d, 0xa4, 0xba, 0xa5, 0xff, 0xc6, 0xfe,
	0x9b, 0xfd, 0x75, 0x03, 0x4f, 0x92, 0xe3, 0x04, 0xd8, 0xb0, 0x37, 0xde, 0xef, 0xe3, 0xee, 0x70,
	0x22, 0x4f, 0xb0, 0x63, 0xb8, 0xfe, 0xe8, 0x85, 0x42, 0x07, 0xa7, 0xa9, 0x4a, 0x4c, 0x42, 0x9a,
	0x4b, 0xa9, 0x56, 0x7c, 0x25, 0x62, 0xd3, 0x6b, 0x45, 0xdc, 0x17, 0x51, 0x8e, 0xf7, 0xf6, 0x31,
	0xf0, 0xb4, 0x88, 0x44, 0x60, 0x12, 0x55, 0xa2, 0x4a, 0x2c, 0x85, 0x12, 0x71, 0x20, 0x36, 0x72,
	0xf4, 0x0e, 0x94, 0xd0, 0x49, 0xa6, 0x02, 0xe1, 0x7d, 0xda, 0x14, 0x1f, 0x61, 0xad, 0xa5, 0x8c,
	0x79, 0xe4, 0x29, 0x91, 0x26, 0xca, 0xe4, 0xc4, 0xc9, 0xdf, 0x6d, 0xe8, 0x2e, 0xb8, 0xfe, 0x38,
	0x16, 0x3a, 0x50, 0x32, 0x35, 0x89, 0x22, 0x0e, 0x54, 0x33, 0x19, 0xd2, 0x4a, 0xbf, 0x32, 0xd8,
	0x62, 0xf6, 0x48, 0x08, 0x6c, 0xc5, 0x7c, 0x25, 0xe8, 0xd3, 0x7e, 0x65, 0xd0, 0x64, 0x78, 0x26,
	0x3f, 0x42, 0x4d, 0x1b, 0x6e, 0x04, 0xad, 0xf6, 0x2b, 0x83, 0xee, 0xd9, 0x57, 0xa7, 0xeb, 0xe6,
	0x4f, 0x1f, 0xe6, 0xc3, 0x70, 0x6e, 0xa5, 0x2c, 0x77, 0x90, 0x03, 0xa8, 0xdf, 0x26, 0xbe, 0x27,
	0x43, 0xba, 0x85, 0x09, 0x6b, 0xb7, 0x89, 0x3f, 0x09, 0xc9, 0x3e, 0xd4, 0x64, 0x1c, 0x8a, 0x3f,
	0x69, 0x0d, 0x2b, 0xe7, 0x01, 0x19, 0x41, 0x3b, 0x14, 0xa9, 0x88, 0x43, 0x11, 0x07, 0x52, 0x68,
	0x5a, 0xef, 0x57, 0x07, 0xad, 0xb3, 0x17, 0x1b, 0xe5, 0x58, 0x39, 0x87, 0xfb, 0x9a, 0xec, 0x81,
	0x87, 0xbc, 0x86, 0xed, 0x24, 0x33, 0x69, 0x66, 0x34, 0xdd, 0xfe, 0x5f, 0xf6, 0x52, 0x4e, 0x0e,
	0xa1, 0xee, 0xcb, 0x98, 0xab, 0x3b, 0xda, 0xc0, 0x56, 0x8b, 0xc8, 0x4e, 0x84, 0xab, 0x0f, 0x9a,
	0x36, 0xfb, 0x55, 0x3b, 0x11, 0x7b, 0x26, 0xaf, 0x60, 0x5b, 0xa7, 0xfc, 0x8f, 0x58, 0x84, 0x14,
	0xb0, 0xca, 0xb3, 0x7f, 0x9d, 0x09, 0x2b, 0x95, 0xe4, 0x0c, 0x0e, 0x74, 0x70, 0x23, 0xc2, 0x2c,
	0x12, 0xa1, 0x67, 0x12, 0xaf, 0xfc, 0x7c, 0xb4, 0x85, 0xf5, 0xf6, 0xd6, 0xe4, 0x22, 0x61, 0x05,
	0x45, 0x7e, 0x80, 0xa3, 0x88, 0x6b, 0xe3, 0xdd, 0x08, 0xae, 0x8c, 0x2f, 0xb8, 0xf1, 0xa2, 0x24,
	0xe0, 0x46, 0x26, 0x31, 0x6d, 0xa3, 0xeb, 0xc0, 0xd2, 0x6f, 0x4a, 0x76, 0x5a, 0x90, 0xe4, 0x14,
	0xf6, 0x1e, 0xf9, 0x8c, 0x5c, 0x09, 0xda, 0xc1, 0x71, 0xef, 0x3e, 0xf0, 0x2c, 0xe4, 0x4a, 0x90,
	0x97, 0x76, 0xf4, 0x91, 0xf8, 0xc0, 0x0d, 0xf6, 0x46, 0xbb, 0x98, 0xbc, 0xb5, 0xc6, 0x16, 0x09,
	0xf9, 0x06, 0xba, 0xf7, 0x92, 0xa5, 0x4a, 0x56, 0x74, 0x07, 0x45, 0x9d, 0x35, 0x7a, 0xa1, 0x92,
	0x15, 0x39, 0x86, 0x96, 0xce, 0xfc, 0x95, 0x2c, 0x2a, 0x3a, 0x58, 0x11, 0x72, 0x08, 0x4b, 0x3d,
	0x07, 0xd0, 0x86, 0xab, 0x82, 0xdf, 0x45, 0xbe, 0x89, 0x08, 0xd2, 0xc7, 0xd0, 0x5a, 0xca, 0x58,
	0xea, 0x9b, 0x9c, 0x27, 0xb9, 0x3f, 0x87, 0x50, 0xf0, 0x3d, 0x1c, 0x9a, 0xc4, 0xf0, 0xc8, 0xcb,
	0xe2, 0x8d, 0x71, 0x5a, 0xed, 0x1e, 0x6a, 0xf7, 0x91, 0xbd, 0xbe, 0x27, 0xd1, 0xf5, 0x35, 0x74,
	0x73, 0x97, 0xca, 0xe2, 0x5c, 0xbd, 0x8f, 0xea, 0x36, 0xa2, 0x2c, 0x8b, 0x51, 0xf5, 0x1d, 0xec,
	0x2a, 0x11, 0x71, 0x23, 0x3f, 0xd9, 0x97, 0xc6, 0xc3, 0x48, 0xc6, 0x82, 0x1e, 0xa0, 0xd0, 0x29,
	0x89, 0x71, 0x81, 0x5b, 0x31, 0xf7, 0x75, 0x12, 0x65, 0x66, 0x43, 0x7c, 0x98, 0x8b, 0x4b, 0x62,
	0x2d, 0x26, 0xb0, 0x65, 0x9f, 0x22, 0x3d, 0x42, 0x1e, 0xcf, 0x76, 0x12, 0x32, 0x4e, 0x33, 0xe3,
	0x69, 0xf9, 0x59, 0x50, 0x9a, 0x4f, 0x02, 0x91, 0xb9, 0xfc, 0x2c, 0xc8, 0xb7, 0xb0, 0x23, 0xe3,
	0x5b, 0x11, 0x18, 0x0f, 0x5f, 0x74, 0x24, 0x7d, 0xfa, 0xac, 0x5f, 0x19, 0x34, 0x58, 0x27, 0x87,
	0xed, 0x3d, 0x9b, 0x4a, 0x9f, 0x8c, 0xc1, 0x59, 0x6f, 0x02, 0x25, 0x7e, 0xcf, 0x84, 0x36, 0xb4,
	0xd7, 0xaf, 0x3c, 0xba, 0x95, 0xe5, 0x95, 0x7a, 0x8f, 0xbb, 0x82, 0xed, 0x94, 0x16, 0x96, 0x3b,
	0x48, 0x0f, 0x1a, 0xa9, 0x92, 0x89, 0x92, 0xe6, 0x8e, 0x7e, 0xd1, 0xaf, 0x0c, 0x3a, 0x6c, 0x1d,
	0x93, 0x9f, 0xa1, 0x89, 0x2d, 0x98, 0xbb, 0x54, 0xd0, 0x2f, 0x71, 0x09, 0x9c, 0xfc, 0xf7, 0x12,
	0x58, 0xdc, 0xa5, 0x82, 0x35, 0x4c, 0x71, 0x22, 0x3f, 0x41, 0x7b, 0x73, 0x21, 0xd1, 0xe7, 0xd8,
	0x5e, 0xef, 0x51, 0x8e, 0x0b, 0x2b, 0x61, 0xa8, 0x60, 0xad, 0xe5, 0x7d, 0x40, 0xfa, 0xd0, 0x36,
	0x8a, 0x07, 0xc2, 0x2b, 0x76, 0xc9, 0x8b, 0xfc, 0x52, 0x20, 0xf6, 0x0b, 0x2e, 0x94, 0x13, 0xe8,
	0xe4, 0x0a, 0xec, 0x53, 0x86, 0xf4, 0x18, 0x25, 0x2d, 0x04, 0x6d, 0xee, 0x49, 0x48, 0x06, 0x50,
	0xc7, 0xed, 0xaa, 0x69, 0x1f, 0xdf, 0xac, 0xb3, 0x51, 0x7e, 0x6a, 0x09, 0x56, 0xf0, 0x64, 0x08,
	0x3b, 0x0f, 0xf7, 0xb0, 0xa6, 0x2f, 0xd1, 0x42, 0x1f, 0x5b, 0xe6, 0x85, 0x80, 0x75, 0xa3, 0xcd,
	0x50, 0x9f, 0xfc, 0x55, 0x81, 0xe6, 0x7a, 0x1b, 0x92, 0x16, 0x6c, 0x9f, 0x33, 0x77, 0xb8, 0x70,
	0xc7, 0xce, 0x13, 0xd2, 0x86, 0xc6, 0x68, 0x7a, 0x75, 0xfe, 0x76, 0x32, 0xbb, 0x74, 0x2a, 0x36,
	0x62, 0xd7, 0xb3, 0xd9, 0x70, 0x34, 0x75, 0x9d, 0xa7, 0x36, 0x1a, 0xce, 0xe7, 0x93, 0xcb, 0x99,
	0x3b, 0x76, 0xaa, 0xd6, 0x66, 0x39, 0x2b, 0xdc, 0x22, 0x1d, 0x68, 0x9e, 0x5f, 0xfd, 0xfa, 0x6e,
	0xea, 0xda, 0x2c, 0x35, 0x02, 0x50, 0xbf, 0x18, 0x4e, 0xa6, 0xee, 0xd8, 0xa9, 0x5b, 0xdd, 0x70,
	0x74, 0xc5, 0x2c, 0xb1, 0x6d, 0x75, 0x63, 0x77, 0xea, 0x5e, 0x62, 0xb5, 0x86, 0xe5, 0xae, 0x67,
	0x6f, 0x67, 0x57, 0xbf, 0xcd, 0x9c, 0xe6, 0xc9, 0x6b, 0x68, 0x94, 0x5f, 0x87, 0x34, 0xa1, 0x36,
	0x7f, 0xe3, 0xba, 0xef, 0x9c, 0x27, 0x36, 0x17, 0x1b, 0x8e, 0x46, 0x93, 0x85, 0x53, 0xb1, 0xf0,
	0xd8, 0x7d, 0x3f, 0x99, 0x3a, 0x4f, 0x2d, 0xbc, 0xb8, 0x66, 0x8b, 0xa9, 0xeb, 0x54, 0xfd, 0x3a,
	0xfe, 0x43, 0x5e, 0xfd, 0x13, 0x00, 0x00, 0xff, 0xff, 0x58, 0x88, 0x83, 0x90, 0xca, 0x06, 0x00,
	0x00,
}