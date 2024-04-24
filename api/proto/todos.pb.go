// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.12.4
// source: proto/todos.proto

package todos

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

type GetTodoRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uuid string `protobuf:"bytes,1,opt,name=uuid,proto3" json:"uuid,omitempty"`
}

func (x *GetTodoRequest) Reset() {
	*x = GetTodoRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_todos_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetTodoRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetTodoRequest) ProtoMessage() {}

func (x *GetTodoRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_todos_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetTodoRequest.ProtoReflect.Descriptor instead.
func (*GetTodoRequest) Descriptor() ([]byte, []int) {
	return file_proto_todos_proto_rawDescGZIP(), []int{0}
}

func (x *GetTodoRequest) GetUuid() string {
	if x != nil {
		return x.Uuid
	}
	return ""
}

type CreateTodoRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Todo *CreateTodoRequest_Todo `protobuf:"bytes,1,opt,name=todo,proto3" json:"todo,omitempty"`
}

func (x *CreateTodoRequest) Reset() {
	*x = CreateTodoRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_todos_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateTodoRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateTodoRequest) ProtoMessage() {}

func (x *CreateTodoRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_todos_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateTodoRequest.ProtoReflect.Descriptor instead.
func (*CreateTodoRequest) Descriptor() ([]byte, []int) {
	return file_proto_todos_proto_rawDescGZIP(), []int{1}
}

func (x *CreateTodoRequest) GetTodo() *CreateTodoRequest_Todo {
	if x != nil {
		return x.Todo
	}
	return nil
}

type TodoCreateResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uuid string `protobuf:"bytes,1,opt,name=uuid,proto3" json:"uuid,omitempty"`
}

func (x *TodoCreateResponse) Reset() {
	*x = TodoCreateResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_todos_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TodoCreateResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TodoCreateResponse) ProtoMessage() {}

func (x *TodoCreateResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_todos_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TodoCreateResponse.ProtoReflect.Descriptor instead.
func (*TodoCreateResponse) Descriptor() ([]byte, []int) {
	return file_proto_todos_proto_rawDescGZIP(), []int{2}
}

func (x *TodoCreateResponse) GetUuid() string {
	if x != nil {
		return x.Uuid
	}
	return ""
}

type TodoResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Todo *Todo `protobuf:"bytes,1,opt,name=todo,proto3" json:"todo,omitempty"`
}

func (x *TodoResponse) Reset() {
	*x = TodoResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_todos_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TodoResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TodoResponse) ProtoMessage() {}

func (x *TodoResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_todos_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TodoResponse.ProtoReflect.Descriptor instead.
func (*TodoResponse) Descriptor() ([]byte, []int) {
	return file_proto_todos_proto_rawDescGZIP(), []int{3}
}

func (x *TodoResponse) GetTodo() *Todo {
	if x != nil {
		return x.Todo
	}
	return nil
}

type Todo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uuid   string `protobuf:"bytes,1,opt,name=uuid,proto3" json:"uuid,omitempty"`
	Text   string `protobuf:"bytes,2,opt,name=text,proto3" json:"text,omitempty"`
	Done   bool   `protobuf:"varint,3,opt,name=done,proto3" json:"done,omitempty"`
	UserId string `protobuf:"bytes,4,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
}

func (x *Todo) Reset() {
	*x = Todo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_todos_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Todo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Todo) ProtoMessage() {}

func (x *Todo) ProtoReflect() protoreflect.Message {
	mi := &file_proto_todos_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Todo.ProtoReflect.Descriptor instead.
func (*Todo) Descriptor() ([]byte, []int) {
	return file_proto_todos_proto_rawDescGZIP(), []int{4}
}

func (x *Todo) GetUuid() string {
	if x != nil {
		return x.Uuid
	}
	return ""
}

func (x *Todo) GetText() string {
	if x != nil {
		return x.Text
	}
	return ""
}

func (x *Todo) GetDone() bool {
	if x != nil {
		return x.Done
	}
	return false
}

func (x *Todo) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

type CreateTodoRequest_Todo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Text string `protobuf:"bytes,1,opt,name=text,proto3" json:"text,omitempty"`
}

func (x *CreateTodoRequest_Todo) Reset() {
	*x = CreateTodoRequest_Todo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_todos_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateTodoRequest_Todo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateTodoRequest_Todo) ProtoMessage() {}

func (x *CreateTodoRequest_Todo) ProtoReflect() protoreflect.Message {
	mi := &file_proto_todos_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateTodoRequest_Todo.ProtoReflect.Descriptor instead.
func (*CreateTodoRequest_Todo) Descriptor() ([]byte, []int) {
	return file_proto_todos_proto_rawDescGZIP(), []int{1, 0}
}

func (x *CreateTodoRequest_Todo) GetText() string {
	if x != nil {
		return x.Text
	}
	return ""
}

var File_proto_todos_proto protoreflect.FileDescriptor

var file_proto_todos_proto_rawDesc = []byte{
	0x0a, 0x11, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x74, 0x6f, 0x64, 0x6f, 0x73, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x05, 0x74, 0x6f, 0x64, 0x6f, 0x73, 0x22, 0x24, 0x0a, 0x0e, 0x47, 0x65,
	0x74, 0x54, 0x6f, 0x64, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04,
	0x75, 0x75, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x75, 0x75, 0x69, 0x64,
	0x22, 0x62, 0x0a, 0x11, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x6f, 0x64, 0x6f, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x31, 0x0a, 0x04, 0x74, 0x6f, 0x64, 0x6f, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x74, 0x6f, 0x64, 0x6f, 0x73, 0x2e, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x54, 0x6f, 0x64, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x54, 0x6f,
	0x64, 0x6f, 0x52, 0x04, 0x74, 0x6f, 0x64, 0x6f, 0x1a, 0x1a, 0x0a, 0x04, 0x54, 0x6f, 0x64, 0x6f,
	0x12, 0x12, 0x0a, 0x04, 0x74, 0x65, 0x78, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x74, 0x65, 0x78, 0x74, 0x22, 0x28, 0x0a, 0x12, 0x54, 0x6f, 0x64, 0x6f, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x75, 0x75,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x75, 0x75, 0x69, 0x64, 0x22, 0x2f,
	0x0a, 0x0c, 0x54, 0x6f, 0x64, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1f,
	0x0a, 0x04, 0x74, 0x6f, 0x64, 0x6f, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x74,
	0x6f, 0x64, 0x6f, 0x73, 0x2e, 0x54, 0x6f, 0x64, 0x6f, 0x52, 0x04, 0x74, 0x6f, 0x64, 0x6f, 0x22,
	0x5b, 0x0a, 0x04, 0x54, 0x6f, 0x64, 0x6f, 0x12, 0x12, 0x0a, 0x04, 0x75, 0x75, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x75, 0x75, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x74,
	0x65, 0x78, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x65, 0x78, 0x74, 0x12,
	0x12, 0x0a, 0x04, 0x64, 0x6f, 0x6e, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x04, 0x64,
	0x6f, 0x6e, 0x65, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x32, 0x81, 0x01, 0x0a,
	0x05, 0x54, 0x6f, 0x64, 0x6f, 0x73, 0x12, 0x35, 0x0a, 0x07, 0x47, 0x65, 0x74, 0x54, 0x6f, 0x64,
	0x6f, 0x12, 0x15, 0x2e, 0x74, 0x6f, 0x64, 0x6f, 0x73, 0x2e, 0x47, 0x65, 0x74, 0x54, 0x6f, 0x64,
	0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x13, 0x2e, 0x74, 0x6f, 0x64, 0x6f, 0x73,
	0x2e, 0x54, 0x6f, 0x64, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x41, 0x0a,
	0x0a, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x6f, 0x64, 0x6f, 0x12, 0x18, 0x2e, 0x74, 0x6f,
	0x64, 0x6f, 0x73, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x6f, 0x64, 0x6f, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x19, 0x2e, 0x74, 0x6f, 0x64, 0x6f, 0x73, 0x2e, 0x54, 0x6f,
	0x64, 0x6f, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x42, 0x32, 0x5a, 0x30, 0x67, 0x69, 0x74, 0x6c, 0x61, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x62,
	0x65, 0x6e, 0x73, 0x63, 0x68, 0x75, 0x64, 0x64, 0x65, 0x62, 0x6f, 0x6f, 0x6d, 0x2f, 0x68, 0x65,
	0x78, 0x61, 0x72, 0x63, 0x68, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2f, 0x74,
	0x6f, 0x64, 0x6f, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_todos_proto_rawDescOnce sync.Once
	file_proto_todos_proto_rawDescData = file_proto_todos_proto_rawDesc
)

func file_proto_todos_proto_rawDescGZIP() []byte {
	file_proto_todos_proto_rawDescOnce.Do(func() {
		file_proto_todos_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_todos_proto_rawDescData)
	})
	return file_proto_todos_proto_rawDescData
}

var file_proto_todos_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_proto_todos_proto_goTypes = []interface{}{
	(*GetTodoRequest)(nil),         // 0: todos.GetTodoRequest
	(*CreateTodoRequest)(nil),      // 1: todos.CreateTodoRequest
	(*TodoCreateResponse)(nil),     // 2: todos.TodoCreateResponse
	(*TodoResponse)(nil),           // 3: todos.TodoResponse
	(*Todo)(nil),                   // 4: todos.Todo
	(*CreateTodoRequest_Todo)(nil), // 5: todos.CreateTodoRequest.Todo
}
var file_proto_todos_proto_depIdxs = []int32{
	5, // 0: todos.CreateTodoRequest.todo:type_name -> todos.CreateTodoRequest.Todo
	4, // 1: todos.TodoResponse.todo:type_name -> todos.Todo
	0, // 2: todos.Todos.GetTodo:input_type -> todos.GetTodoRequest
	1, // 3: todos.Todos.CreateTodo:input_type -> todos.CreateTodoRequest
	3, // 4: todos.Todos.GetTodo:output_type -> todos.TodoResponse
	2, // 5: todos.Todos.CreateTodo:output_type -> todos.TodoCreateResponse
	4, // [4:6] is the sub-list for method output_type
	2, // [2:4] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_proto_todos_proto_init() }
func file_proto_todos_proto_init() {
	if File_proto_todos_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_todos_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetTodoRequest); i {
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
		file_proto_todos_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateTodoRequest); i {
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
		file_proto_todos_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TodoCreateResponse); i {
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
		file_proto_todos_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TodoResponse); i {
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
		file_proto_todos_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Todo); i {
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
		file_proto_todos_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateTodoRequest_Todo); i {
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
			RawDescriptor: file_proto_todos_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_todos_proto_goTypes,
		DependencyIndexes: file_proto_todos_proto_depIdxs,
		MessageInfos:      file_proto_todos_proto_msgTypes,
	}.Build()
	File_proto_todos_proto = out.File
	file_proto_todos_proto_rawDesc = nil
	file_proto_todos_proto_goTypes = nil
	file_proto_todos_proto_depIdxs = nil
}