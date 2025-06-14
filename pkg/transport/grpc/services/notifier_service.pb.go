// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v5.28.0
// source: services/notifier/proto/notifier_service.proto

package pb

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type GreetRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Email string `protobuf:"bytes,1,opt,name=email,proto3" json:"email,omitempty"`
	Name  string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *GreetRequest) Reset() {
	*x = GreetRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_services_notifier_proto_notifier_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GreetRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GreetRequest) ProtoMessage() {}

func (x *GreetRequest) ProtoReflect() protoreflect.Message {
	mi := &file_services_notifier_proto_notifier_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GreetRequest.ProtoReflect.Descriptor instead.
func (*GreetRequest) Descriptor() ([]byte, []int) {
	return file_services_notifier_proto_notifier_service_proto_rawDescGZIP(), []int{0}
}

func (x *GreetRequest) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *GreetRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type GreetResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success bool `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
}

func (x *GreetResponse) Reset() {
	*x = GreetResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_services_notifier_proto_notifier_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GreetResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GreetResponse) ProtoMessage() {}

func (x *GreetResponse) ProtoReflect() protoreflect.Message {
	mi := &file_services_notifier_proto_notifier_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GreetResponse.ProtoReflect.Descriptor instead.
func (*GreetResponse) Descriptor() ([]byte, []int) {
	return file_services_notifier_proto_notifier_service_proto_rawDescGZIP(), []int{1}
}

func (x *GreetResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

type NotifyTaskAssignmentRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AssigneeName  string `protobuf:"bytes,1,opt,name=assignee_name,json=assigneeName,proto3" json:"assignee_name,omitempty"`
	AssigneeEmail string `protobuf:"bytes,2,opt,name=assignee_email,json=assigneeEmail,proto3" json:"assignee_email,omitempty"`
	TaskName      string `protobuf:"bytes,3,opt,name=task_name,json=taskName,proto3" json:"task_name,omitempty"`
	TaskId        string `protobuf:"bytes,4,opt,name=task_id,json=taskId,proto3" json:"task_id,omitempty"`
	ProjectName   string `protobuf:"bytes,5,opt,name=project_name,json=projectName,proto3" json:"project_name,omitempty"`
}

func (x *NotifyTaskAssignmentRequest) Reset() {
	*x = NotifyTaskAssignmentRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_services_notifier_proto_notifier_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NotifyTaskAssignmentRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NotifyTaskAssignmentRequest) ProtoMessage() {}

func (x *NotifyTaskAssignmentRequest) ProtoReflect() protoreflect.Message {
	mi := &file_services_notifier_proto_notifier_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NotifyTaskAssignmentRequest.ProtoReflect.Descriptor instead.
func (*NotifyTaskAssignmentRequest) Descriptor() ([]byte, []int) {
	return file_services_notifier_proto_notifier_service_proto_rawDescGZIP(), []int{2}
}

func (x *NotifyTaskAssignmentRequest) GetAssigneeName() string {
	if x != nil {
		return x.AssigneeName
	}
	return ""
}

func (x *NotifyTaskAssignmentRequest) GetAssigneeEmail() string {
	if x != nil {
		return x.AssigneeEmail
	}
	return ""
}

func (x *NotifyTaskAssignmentRequest) GetTaskName() string {
	if x != nil {
		return x.TaskName
	}
	return ""
}

func (x *NotifyTaskAssignmentRequest) GetTaskId() string {
	if x != nil {
		return x.TaskId
	}
	return ""
}

func (x *NotifyTaskAssignmentRequest) GetProjectName() string {
	if x != nil {
		return x.ProjectName
	}
	return ""
}

type NotifyTaskAssignmentResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success bool `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
}

func (x *NotifyTaskAssignmentResponse) Reset() {
	*x = NotifyTaskAssignmentResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_services_notifier_proto_notifier_service_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NotifyTaskAssignmentResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NotifyTaskAssignmentResponse) ProtoMessage() {}

func (x *NotifyTaskAssignmentResponse) ProtoReflect() protoreflect.Message {
	mi := &file_services_notifier_proto_notifier_service_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NotifyTaskAssignmentResponse.ProtoReflect.Descriptor instead.
func (*NotifyTaskAssignmentResponse) Descriptor() ([]byte, []int) {
	return file_services_notifier_proto_notifier_service_proto_rawDescGZIP(), []int{3}
}

func (x *NotifyTaskAssignmentResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

var File_services_notifier_proto_notifier_service_proto protoreflect.FileDescriptor

var file_services_notifier_proto_notifier_service_proto_rawDesc = []byte{
	0x0a, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2f, 0x6e, 0x6f, 0x74, 0x69, 0x66,
	0x69, 0x65, 0x72, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69,
	0x65, 0x72, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x08, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x65, 0x72, 0x22, 0x38, 0x0a, 0x0c, 0x47, 0x72,
	0x65, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d,
	0x61, 0x69, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c,
	0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x22, 0x29, 0x0a, 0x0d, 0x47, 0x72, 0x65, 0x65, 0x74, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x22,
	0xc2, 0x01, 0x0a, 0x1b, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x54, 0x61, 0x73, 0x6b, 0x41, 0x73,
	0x73, 0x69, 0x67, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x23, 0x0a, 0x0d, 0x61, 0x73, 0x73, 0x69, 0x67, 0x6e, 0x65, 0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x61, 0x73, 0x73, 0x69, 0x67, 0x6e, 0x65, 0x65,
	0x4e, 0x61, 0x6d, 0x65, 0x12, 0x25, 0x0a, 0x0e, 0x61, 0x73, 0x73, 0x69, 0x67, 0x6e, 0x65, 0x65,
	0x5f, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x61, 0x73,
	0x73, 0x69, 0x67, 0x6e, 0x65, 0x65, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x1b, 0x0a, 0x09, 0x74,
	0x61, 0x73, 0x6b, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x74, 0x61, 0x73, 0x6b, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x17, 0x0a, 0x07, 0x74, 0x61, 0x73, 0x6b,
	0x5f, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x74, 0x61, 0x73, 0x6b, 0x49,
	0x64, 0x12, 0x21, 0x0a, 0x0c, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x5f, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74,
	0x4e, 0x61, 0x6d, 0x65, 0x22, 0x38, 0x0a, 0x1c, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x54, 0x61,
	0x73, 0x6b, 0x41, 0x73, 0x73, 0x69, 0x67, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x32, 0xab,
	0x01, 0x0a, 0x08, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x65, 0x72, 0x12, 0x38, 0x0a, 0x05, 0x47,
	0x72, 0x65, 0x65, 0x74, 0x12, 0x16, 0x2e, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x65, 0x72, 0x2e,
	0x47, 0x72, 0x65, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x17, 0x2e, 0x6e,
	0x6f, 0x74, 0x69, 0x66, 0x69, 0x65, 0x72, 0x2e, 0x47, 0x72, 0x65, 0x65, 0x74, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x65, 0x0a, 0x14, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x54,
	0x61, 0x73, 0x6b, 0x41, 0x73, 0x73, 0x69, 0x67, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x25, 0x2e,
	0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x65, 0x72, 0x2e, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x54,
	0x61, 0x73, 0x6b, 0x41, 0x73, 0x73, 0x69, 0x67, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x26, 0x2e, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x65, 0x72, 0x2e,
	0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x54, 0x61, 0x73, 0x6b, 0x41, 0x73, 0x73, 0x69, 0x67, 0x6e,
	0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x17, 0x5a, 0x15,
	0x70, 0x6d, 0x73, 0x2e, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x65, 0x72, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x3b, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_services_notifier_proto_notifier_service_proto_rawDescOnce sync.Once
	file_services_notifier_proto_notifier_service_proto_rawDescData = file_services_notifier_proto_notifier_service_proto_rawDesc
)

func file_services_notifier_proto_notifier_service_proto_rawDescGZIP() []byte {
	file_services_notifier_proto_notifier_service_proto_rawDescOnce.Do(func() {
		file_services_notifier_proto_notifier_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_services_notifier_proto_notifier_service_proto_rawDescData)
	})
	return file_services_notifier_proto_notifier_service_proto_rawDescData
}

var file_services_notifier_proto_notifier_service_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_services_notifier_proto_notifier_service_proto_goTypes = []any{
	(*GreetRequest)(nil),                 // 0: notifier.GreetRequest
	(*GreetResponse)(nil),                // 1: notifier.GreetResponse
	(*NotifyTaskAssignmentRequest)(nil),  // 2: notifier.NotifyTaskAssignmentRequest
	(*NotifyTaskAssignmentResponse)(nil), // 3: notifier.NotifyTaskAssignmentResponse
}
var file_services_notifier_proto_notifier_service_proto_depIdxs = []int32{
	0, // 0: notifier.Notifier.Greet:input_type -> notifier.GreetRequest
	2, // 1: notifier.Notifier.NotifyTaskAssignment:input_type -> notifier.NotifyTaskAssignmentRequest
	1, // 2: notifier.Notifier.Greet:output_type -> notifier.GreetResponse
	3, // 3: notifier.Notifier.NotifyTaskAssignment:output_type -> notifier.NotifyTaskAssignmentResponse
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_services_notifier_proto_notifier_service_proto_init() }
func file_services_notifier_proto_notifier_service_proto_init() {
	if File_services_notifier_proto_notifier_service_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_services_notifier_proto_notifier_service_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*GreetRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_services_notifier_proto_notifier_service_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*GreetResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_services_notifier_proto_notifier_service_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*NotifyTaskAssignmentRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_services_notifier_proto_notifier_service_proto_msgTypes[3].Exporter = func(v any, i int) any {
			switch v := v.(*NotifyTaskAssignmentResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_services_notifier_proto_notifier_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_services_notifier_proto_notifier_service_proto_goTypes,
		DependencyIndexes: file_services_notifier_proto_notifier_service_proto_depIdxs,
		MessageInfos:      file_services_notifier_proto_notifier_service_proto_msgTypes,
	}.Build()
	File_services_notifier_proto_notifier_service_proto = out.File
	file_services_notifier_proto_notifier_service_proto_rawDesc = nil
	file_services_notifier_proto_notifier_service_proto_goTypes = nil
	file_services_notifier_proto_notifier_service_proto_depIdxs = nil
}
