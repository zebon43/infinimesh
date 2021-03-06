//--------------------------------------------------------------------------
// Copyright 2018 Infinite Devices GmbH
// www.infinimesh.io
//
//   Licensed under the Apache License, Version 2.0 (the "License");
//   you may not use this file except in compliance with the License.
//   You may obtain a copy of the License at
//
//       http://www.apache.org/licenses/LICENSE-2.0
//
//   Unless required by applicable law or agreed to in writing, software
//   distributed under the License is distributed on an "AS IS" BASIS,
//   WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//   See the License for the specific language governing permissions and
//   limitations under the License.
//--------------------------------------------------------------------------

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.23.0
// 	protoc        v3.12.3
// source: pkg/shadow/shadowpb/shadow.proto

package shadowpb

import (
	context "context"
	proto "github.com/golang/protobuf/proto"
	_struct "github.com/golang/protobuf/ptypes/struct"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type VersionedValue struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Version   uint64               `protobuf:"varint,1,opt,name=version,proto3" json:"version,omitempty"`
	Data      *_struct.Value       `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
	Timestamp *timestamp.Timestamp `protobuf:"bytes,3,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
}

func (x *VersionedValue) Reset() {
	*x = VersionedValue{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_shadow_shadowpb_shadow_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *VersionedValue) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VersionedValue) ProtoMessage() {}

func (x *VersionedValue) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_shadow_shadowpb_shadow_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VersionedValue.ProtoReflect.Descriptor instead.
func (*VersionedValue) Descriptor() ([]byte, []int) {
	return file_pkg_shadow_shadowpb_shadow_proto_rawDescGZIP(), []int{0}
}

func (x *VersionedValue) GetVersion() uint64 {
	if x != nil {
		return x.Version
	}
	return 0
}

func (x *VersionedValue) GetData() *_struct.Value {
	if x != nil {
		return x.Data
	}
	return nil
}

func (x *VersionedValue) GetTimestamp() *timestamp.Timestamp {
	if x != nil {
		return x.Timestamp
	}
	return nil
}

type GetRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *GetRequest) Reset() {
	*x = GetRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_shadow_shadowpb_shadow_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetRequest) ProtoMessage() {}

func (x *GetRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_shadow_shadowpb_shadow_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetRequest.ProtoReflect.Descriptor instead.
func (*GetRequest) Descriptor() ([]byte, []int) {
	return file_pkg_shadow_shadowpb_shadow_proto_rawDescGZIP(), []int{1}
}

func (x *GetRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type GetResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Shadow *Shadow `protobuf:"bytes,1,opt,name=shadow,proto3" json:"shadow,omitempty"`
}

func (x *GetResponse) Reset() {
	*x = GetResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_shadow_shadowpb_shadow_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetResponse) ProtoMessage() {}

func (x *GetResponse) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_shadow_shadowpb_shadow_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetResponse.ProtoReflect.Descriptor instead.
func (*GetResponse) Descriptor() ([]byte, []int) {
	return file_pkg_shadow_shadowpb_shadow_proto_rawDescGZIP(), []int{2}
}

func (x *GetResponse) GetShadow() *Shadow {
	if x != nil {
		return x.Shadow
	}
	return nil
}

type Shadow struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Config   *VersionedValue `protobuf:"bytes,2,opt,name=config,proto3" json:"config,omitempty"`
	Reported *VersionedValue `protobuf:"bytes,3,opt,name=reported,proto3" json:"reported,omitempty"`
	Desired  *VersionedValue `protobuf:"bytes,4,opt,name=desired,proto3" json:"desired,omitempty"`
}

func (x *Shadow) Reset() {
	*x = Shadow{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_shadow_shadowpb_shadow_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Shadow) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Shadow) ProtoMessage() {}

func (x *Shadow) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_shadow_shadowpb_shadow_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Shadow.ProtoReflect.Descriptor instead.
func (*Shadow) Descriptor() ([]byte, []int) {
	return file_pkg_shadow_shadowpb_shadow_proto_rawDescGZIP(), []int{3}
}

func (x *Shadow) GetConfig() *VersionedValue {
	if x != nil {
		return x.Config
	}
	return nil
}

func (x *Shadow) GetReported() *VersionedValue {
	if x != nil {
		return x.Reported
	}
	return nil
}

func (x *Shadow) GetDesired() *VersionedValue {
	if x != nil {
		return x.Desired
	}
	return nil
}

type PatchDesiredStateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id   string         `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Data *_struct.Value `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *PatchDesiredStateRequest) Reset() {
	*x = PatchDesiredStateRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_shadow_shadowpb_shadow_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PatchDesiredStateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PatchDesiredStateRequest) ProtoMessage() {}

func (x *PatchDesiredStateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_shadow_shadowpb_shadow_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PatchDesiredStateRequest.ProtoReflect.Descriptor instead.
func (*PatchDesiredStateRequest) Descriptor() ([]byte, []int) {
	return file_pkg_shadow_shadowpb_shadow_proto_rawDescGZIP(), []int{4}
}

func (x *PatchDesiredStateRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *PatchDesiredStateRequest) GetData() *_struct.Value {
	if x != nil {
		return x.Data
	}
	return nil
}

type PatchDesiredStateResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *PatchDesiredStateResponse) Reset() {
	*x = PatchDesiredStateResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_shadow_shadowpb_shadow_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PatchDesiredStateResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PatchDesiredStateResponse) ProtoMessage() {}

func (x *PatchDesiredStateResponse) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_shadow_shadowpb_shadow_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PatchDesiredStateResponse.ProtoReflect.Descriptor instead.
func (*PatchDesiredStateResponse) Descriptor() ([]byte, []int) {
	return file_pkg_shadow_shadowpb_shadow_proto_rawDescGZIP(), []int{5}
}

type StreamReportedStateChangesRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	OnlyDelta bool   `protobuf:"varint,2,opt,name=only_delta,json=onlyDelta,proto3" json:"only_delta,omitempty"`
}

func (x *StreamReportedStateChangesRequest) Reset() {
	*x = StreamReportedStateChangesRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_shadow_shadowpb_shadow_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StreamReportedStateChangesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StreamReportedStateChangesRequest) ProtoMessage() {}

func (x *StreamReportedStateChangesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_shadow_shadowpb_shadow_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StreamReportedStateChangesRequest.ProtoReflect.Descriptor instead.
func (*StreamReportedStateChangesRequest) Descriptor() ([]byte, []int) {
	return file_pkg_shadow_shadowpb_shadow_proto_rawDescGZIP(), []int{6}
}

func (x *StreamReportedStateChangesRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *StreamReportedStateChangesRequest) GetOnlyDelta() bool {
	if x != nil {
		return x.OnlyDelta
	}
	return false
}

type StreamReportedStateChangesResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ReportedState *VersionedValue `protobuf:"bytes,2,opt,name=reportedState,proto3" json:"reportedState,omitempty"`
	DesiredState  *VersionedValue `protobuf:"bytes,3,opt,name=desiredState,proto3" json:"desiredState,omitempty"`
}

func (x *StreamReportedStateChangesResponse) Reset() {
	*x = StreamReportedStateChangesResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_shadow_shadowpb_shadow_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StreamReportedStateChangesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StreamReportedStateChangesResponse) ProtoMessage() {}

func (x *StreamReportedStateChangesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_shadow_shadowpb_shadow_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StreamReportedStateChangesResponse.ProtoReflect.Descriptor instead.
func (*StreamReportedStateChangesResponse) Descriptor() ([]byte, []int) {
	return file_pkg_shadow_shadowpb_shadow_proto_rawDescGZIP(), []int{7}
}

func (x *StreamReportedStateChangesResponse) GetReportedState() *VersionedValue {
	if x != nil {
		return x.ReportedState
	}
	return nil
}

func (x *StreamReportedStateChangesResponse) GetDesiredState() *VersionedValue {
	if x != nil {
		return x.DesiredState
	}
	return nil
}

var File_pkg_shadow_shadowpb_shadow_proto protoreflect.FileDescriptor

var file_pkg_shadow_shadowpb_shadow_proto_rawDesc = []byte{
	0x0a, 0x20, 0x70, 0x6b, 0x67, 0x2f, 0x73, 0x68, 0x61, 0x64, 0x6f, 0x77, 0x2f, 0x73, 0x68, 0x61,
	0x64, 0x6f, 0x77, 0x70, 0x62, 0x2f, 0x73, 0x68, 0x61, 0x64, 0x6f, 0x77, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x11, 0x69, 0x6e, 0x66, 0x69, 0x6e, 0x69, 0x6d, 0x65, 0x73, 0x68, 0x2e, 0x73,
	0x68, 0x61, 0x64, 0x6f, 0x77, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0x90, 0x01, 0x0a, 0x0e, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e,
	0x65, 0x64, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69,
	0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f,
	0x6e, 0x12, 0x2a, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x12, 0x38, 0x0a,
	0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x74, 0x69,
	0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x22, 0x1c, 0x0a, 0x0a, 0x47, 0x65, 0x74, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x40, 0x0a, 0x0b, 0x47, 0x65, 0x74, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x31, 0x0a, 0x06, 0x73, 0x68, 0x61, 0x64, 0x6f, 0x77, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x69, 0x6e, 0x66, 0x69, 0x6e, 0x69, 0x6d, 0x65, 0x73,
	0x68, 0x2e, 0x73, 0x68, 0x61, 0x64, 0x6f, 0x77, 0x2e, 0x53, 0x68, 0x61, 0x64, 0x6f, 0x77, 0x52,
	0x06, 0x73, 0x68, 0x61, 0x64, 0x6f, 0x77, 0x22, 0xbf, 0x01, 0x0a, 0x06, 0x53, 0x68, 0x61, 0x64,
	0x6f, 0x77, 0x12, 0x39, 0x0a, 0x06, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x21, 0x2e, 0x69, 0x6e, 0x66, 0x69, 0x6e, 0x69, 0x6d, 0x65, 0x73, 0x68, 0x2e,
	0x73, 0x68, 0x61, 0x64, 0x6f, 0x77, 0x2e, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x65, 0x64,
	0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x06, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x3d, 0x0a,
	0x08, 0x72, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x65, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x21, 0x2e, 0x69, 0x6e, 0x66, 0x69, 0x6e, 0x69, 0x6d, 0x65, 0x73, 0x68, 0x2e, 0x73, 0x68, 0x61,
	0x64, 0x6f, 0x77, 0x2e, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x65, 0x64, 0x56, 0x61, 0x6c,
	0x75, 0x65, 0x52, 0x08, 0x72, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x65, 0x64, 0x12, 0x3b, 0x0a, 0x07,
	0x64, 0x65, 0x73, 0x69, 0x72, 0x65, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x21, 0x2e,
	0x69, 0x6e, 0x66, 0x69, 0x6e, 0x69, 0x6d, 0x65, 0x73, 0x68, 0x2e, 0x73, 0x68, 0x61, 0x64, 0x6f,
	0x77, 0x2e, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x65, 0x64, 0x56, 0x61, 0x6c, 0x75, 0x65,
	0x52, 0x07, 0x64, 0x65, 0x73, 0x69, 0x72, 0x65, 0x64, 0x22, 0x56, 0x0a, 0x18, 0x50, 0x61, 0x74,
	0x63, 0x68, 0x44, 0x65, 0x73, 0x69, 0x72, 0x65, 0x64, 0x53, 0x74, 0x61, 0x74, 0x65, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x2a, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x04, 0x64, 0x61, 0x74,
	0x61, 0x22, 0x1b, 0x0a, 0x19, 0x50, 0x61, 0x74, 0x63, 0x68, 0x44, 0x65, 0x73, 0x69, 0x72, 0x65,
	0x64, 0x53, 0x74, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x52,
	0x0a, 0x21, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x65, 0x64,
	0x53, 0x74, 0x61, 0x74, 0x65, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x02, 0x69, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x6f, 0x6e, 0x6c, 0x79, 0x5f, 0x64, 0x65, 0x6c, 0x74,
	0x61, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x09, 0x6f, 0x6e, 0x6c, 0x79, 0x44, 0x65, 0x6c,
	0x74, 0x61, 0x22, 0xb4, 0x01, 0x0a, 0x22, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x52, 0x65, 0x70,
	0x6f, 0x72, 0x74, 0x65, 0x64, 0x53, 0x74, 0x61, 0x74, 0x65, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65,
	0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x47, 0x0a, 0x0d, 0x72, 0x65, 0x70,
	0x6f, 0x72, 0x74, 0x65, 0x64, 0x53, 0x74, 0x61, 0x74, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x21, 0x2e, 0x69, 0x6e, 0x66, 0x69, 0x6e, 0x69, 0x6d, 0x65, 0x73, 0x68, 0x2e, 0x73, 0x68,
	0x61, 0x64, 0x6f, 0x77, 0x2e, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x65, 0x64, 0x56, 0x61,
	0x6c, 0x75, 0x65, 0x52, 0x0d, 0x72, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x65, 0x64, 0x53, 0x74, 0x61,
	0x74, 0x65, 0x12, 0x45, 0x0a, 0x0c, 0x64, 0x65, 0x73, 0x69, 0x72, 0x65, 0x64, 0x53, 0x74, 0x61,
	0x74, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x21, 0x2e, 0x69, 0x6e, 0x66, 0x69, 0x6e,
	0x69, 0x6d, 0x65, 0x73, 0x68, 0x2e, 0x73, 0x68, 0x61, 0x64, 0x6f, 0x77, 0x2e, 0x56, 0x65, 0x72,
	0x73, 0x69, 0x6f, 0x6e, 0x65, 0x64, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x0c, 0x64, 0x65, 0x73,
	0x69, 0x72, 0x65, 0x64, 0x53, 0x74, 0x61, 0x74, 0x65, 0x32, 0xcd, 0x02, 0x0a, 0x07, 0x53, 0x68,
	0x61, 0x64, 0x6f, 0x77, 0x73, 0x12, 0x44, 0x0a, 0x03, 0x47, 0x65, 0x74, 0x12, 0x1d, 0x2e, 0x69,
	0x6e, 0x66, 0x69, 0x6e, 0x69, 0x6d, 0x65, 0x73, 0x68, 0x2e, 0x73, 0x68, 0x61, 0x64, 0x6f, 0x77,
	0x2e, 0x47, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1e, 0x2e, 0x69, 0x6e,
	0x66, 0x69, 0x6e, 0x69, 0x6d, 0x65, 0x73, 0x68, 0x2e, 0x73, 0x68, 0x61, 0x64, 0x6f, 0x77, 0x2e,
	0x47, 0x65, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x6e, 0x0a, 0x11, 0x50,
	0x61, 0x74, 0x63, 0x68, 0x44, 0x65, 0x73, 0x69, 0x72, 0x65, 0x64, 0x53, 0x74, 0x61, 0x74, 0x65,
	0x12, 0x2b, 0x2e, 0x69, 0x6e, 0x66, 0x69, 0x6e, 0x69, 0x6d, 0x65, 0x73, 0x68, 0x2e, 0x73, 0x68,
	0x61, 0x64, 0x6f, 0x77, 0x2e, 0x50, 0x61, 0x74, 0x63, 0x68, 0x44, 0x65, 0x73, 0x69, 0x72, 0x65,
	0x64, 0x53, 0x74, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2c, 0x2e,
	0x69, 0x6e, 0x66, 0x69, 0x6e, 0x69, 0x6d, 0x65, 0x73, 0x68, 0x2e, 0x73, 0x68, 0x61, 0x64, 0x6f,
	0x77, 0x2e, 0x50, 0x61, 0x74, 0x63, 0x68, 0x44, 0x65, 0x73, 0x69, 0x72, 0x65, 0x64, 0x53, 0x74,
	0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x8b, 0x01, 0x0a, 0x1a,
	0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x65, 0x64, 0x53, 0x74,
	0x61, 0x74, 0x65, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x73, 0x12, 0x34, 0x2e, 0x69, 0x6e, 0x66,
	0x69, 0x6e, 0x69, 0x6d, 0x65, 0x73, 0x68, 0x2e, 0x73, 0x68, 0x61, 0x64, 0x6f, 0x77, 0x2e, 0x53,
	0x74, 0x72, 0x65, 0x61, 0x6d, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x65, 0x64, 0x53, 0x74, 0x61,
	0x74, 0x65, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x35, 0x2e, 0x69, 0x6e, 0x66, 0x69, 0x6e, 0x69, 0x6d, 0x65, 0x73, 0x68, 0x2e, 0x73, 0x68,
	0x61, 0x64, 0x6f, 0x77, 0x2e, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x52, 0x65, 0x70, 0x6f, 0x72,
	0x74, 0x65, 0x64, 0x53, 0x74, 0x61, 0x74, 0x65, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x73, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x30, 0x01, 0x42, 0x0a, 0x5a, 0x08, 0x73, 0x68, 0x61,
	0x64, 0x6f, 0x77, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_pkg_shadow_shadowpb_shadow_proto_rawDescOnce sync.Once
	file_pkg_shadow_shadowpb_shadow_proto_rawDescData = file_pkg_shadow_shadowpb_shadow_proto_rawDesc
)

func file_pkg_shadow_shadowpb_shadow_proto_rawDescGZIP() []byte {
	file_pkg_shadow_shadowpb_shadow_proto_rawDescOnce.Do(func() {
		file_pkg_shadow_shadowpb_shadow_proto_rawDescData = protoimpl.X.CompressGZIP(file_pkg_shadow_shadowpb_shadow_proto_rawDescData)
	})
	return file_pkg_shadow_shadowpb_shadow_proto_rawDescData
}

var file_pkg_shadow_shadowpb_shadow_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_pkg_shadow_shadowpb_shadow_proto_goTypes = []interface{}{
	(*VersionedValue)(nil),                     // 0: infinimesh.shadow.VersionedValue
	(*GetRequest)(nil),                         // 1: infinimesh.shadow.GetRequest
	(*GetResponse)(nil),                        // 2: infinimesh.shadow.GetResponse
	(*Shadow)(nil),                             // 3: infinimesh.shadow.Shadow
	(*PatchDesiredStateRequest)(nil),           // 4: infinimesh.shadow.PatchDesiredStateRequest
	(*PatchDesiredStateResponse)(nil),          // 5: infinimesh.shadow.PatchDesiredStateResponse
	(*StreamReportedStateChangesRequest)(nil),  // 6: infinimesh.shadow.StreamReportedStateChangesRequest
	(*StreamReportedStateChangesResponse)(nil), // 7: infinimesh.shadow.StreamReportedStateChangesResponse
	(*_struct.Value)(nil),                      // 8: google.protobuf.Value
	(*timestamp.Timestamp)(nil),                // 9: google.protobuf.Timestamp
}
var file_pkg_shadow_shadowpb_shadow_proto_depIdxs = []int32{
	8,  // 0: infinimesh.shadow.VersionedValue.data:type_name -> google.protobuf.Value
	9,  // 1: infinimesh.shadow.VersionedValue.timestamp:type_name -> google.protobuf.Timestamp
	3,  // 2: infinimesh.shadow.GetResponse.shadow:type_name -> infinimesh.shadow.Shadow
	0,  // 3: infinimesh.shadow.Shadow.config:type_name -> infinimesh.shadow.VersionedValue
	0,  // 4: infinimesh.shadow.Shadow.reported:type_name -> infinimesh.shadow.VersionedValue
	0,  // 5: infinimesh.shadow.Shadow.desired:type_name -> infinimesh.shadow.VersionedValue
	8,  // 6: infinimesh.shadow.PatchDesiredStateRequest.data:type_name -> google.protobuf.Value
	0,  // 7: infinimesh.shadow.StreamReportedStateChangesResponse.reportedState:type_name -> infinimesh.shadow.VersionedValue
	0,  // 8: infinimesh.shadow.StreamReportedStateChangesResponse.desiredState:type_name -> infinimesh.shadow.VersionedValue
	1,  // 9: infinimesh.shadow.Shadows.Get:input_type -> infinimesh.shadow.GetRequest
	4,  // 10: infinimesh.shadow.Shadows.PatchDesiredState:input_type -> infinimesh.shadow.PatchDesiredStateRequest
	6,  // 11: infinimesh.shadow.Shadows.StreamReportedStateChanges:input_type -> infinimesh.shadow.StreamReportedStateChangesRequest
	2,  // 12: infinimesh.shadow.Shadows.Get:output_type -> infinimesh.shadow.GetResponse
	5,  // 13: infinimesh.shadow.Shadows.PatchDesiredState:output_type -> infinimesh.shadow.PatchDesiredStateResponse
	7,  // 14: infinimesh.shadow.Shadows.StreamReportedStateChanges:output_type -> infinimesh.shadow.StreamReportedStateChangesResponse
	12, // [12:15] is the sub-list for method output_type
	9,  // [9:12] is the sub-list for method input_type
	9,  // [9:9] is the sub-list for extension type_name
	9,  // [9:9] is the sub-list for extension extendee
	0,  // [0:9] is the sub-list for field type_name
}

func init() { file_pkg_shadow_shadowpb_shadow_proto_init() }
func file_pkg_shadow_shadowpb_shadow_proto_init() {
	if File_pkg_shadow_shadowpb_shadow_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_pkg_shadow_shadowpb_shadow_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*VersionedValue); i {
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
		file_pkg_shadow_shadowpb_shadow_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetRequest); i {
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
		file_pkg_shadow_shadowpb_shadow_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetResponse); i {
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
		file_pkg_shadow_shadowpb_shadow_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Shadow); i {
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
		file_pkg_shadow_shadowpb_shadow_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PatchDesiredStateRequest); i {
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
		file_pkg_shadow_shadowpb_shadow_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PatchDesiredStateResponse); i {
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
		file_pkg_shadow_shadowpb_shadow_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StreamReportedStateChangesRequest); i {
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
		file_pkg_shadow_shadowpb_shadow_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StreamReportedStateChangesResponse); i {
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
			RawDescriptor: file_pkg_shadow_shadowpb_shadow_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_pkg_shadow_shadowpb_shadow_proto_goTypes,
		DependencyIndexes: file_pkg_shadow_shadowpb_shadow_proto_depIdxs,
		MessageInfos:      file_pkg_shadow_shadowpb_shadow_proto_msgTypes,
	}.Build()
	File_pkg_shadow_shadowpb_shadow_proto = out.File
	file_pkg_shadow_shadowpb_shadow_proto_rawDesc = nil
	file_pkg_shadow_shadowpb_shadow_proto_goTypes = nil
	file_pkg_shadow_shadowpb_shadow_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// ShadowsClient is the client API for Shadows service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ShadowsClient interface {
	Get(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*GetResponse, error)
	PatchDesiredState(ctx context.Context, in *PatchDesiredStateRequest, opts ...grpc.CallOption) (*PatchDesiredStateResponse, error)
	StreamReportedStateChanges(ctx context.Context, in *StreamReportedStateChangesRequest, opts ...grpc.CallOption) (Shadows_StreamReportedStateChangesClient, error)
}

type shadowsClient struct {
	cc grpc.ClientConnInterface
}

func NewShadowsClient(cc grpc.ClientConnInterface) ShadowsClient {
	return &shadowsClient{cc}
}

func (c *shadowsClient) Get(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*GetResponse, error) {
	out := new(GetResponse)
	err := c.cc.Invoke(ctx, "/infinimesh.shadow.Shadows/Get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *shadowsClient) PatchDesiredState(ctx context.Context, in *PatchDesiredStateRequest, opts ...grpc.CallOption) (*PatchDesiredStateResponse, error) {
	out := new(PatchDesiredStateResponse)
	err := c.cc.Invoke(ctx, "/infinimesh.shadow.Shadows/PatchDesiredState", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *shadowsClient) StreamReportedStateChanges(ctx context.Context, in *StreamReportedStateChangesRequest, opts ...grpc.CallOption) (Shadows_StreamReportedStateChangesClient, error) {
	stream, err := c.cc.NewStream(ctx, &_Shadows_serviceDesc.Streams[0], "/infinimesh.shadow.Shadows/StreamReportedStateChanges", opts...)
	if err != nil {
		return nil, err
	}
	x := &shadowsStreamReportedStateChangesClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Shadows_StreamReportedStateChangesClient interface {
	Recv() (*StreamReportedStateChangesResponse, error)
	grpc.ClientStream
}

type shadowsStreamReportedStateChangesClient struct {
	grpc.ClientStream
}

func (x *shadowsStreamReportedStateChangesClient) Recv() (*StreamReportedStateChangesResponse, error) {
	m := new(StreamReportedStateChangesResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// ShadowsServer is the server API for Shadows service.
type ShadowsServer interface {
	Get(context.Context, *GetRequest) (*GetResponse, error)
	PatchDesiredState(context.Context, *PatchDesiredStateRequest) (*PatchDesiredStateResponse, error)
	StreamReportedStateChanges(*StreamReportedStateChangesRequest, Shadows_StreamReportedStateChangesServer) error
}

// UnimplementedShadowsServer can be embedded to have forward compatible implementations.
type UnimplementedShadowsServer struct {
}

func (*UnimplementedShadowsServer) Get(context.Context, *GetRequest) (*GetResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (*UnimplementedShadowsServer) PatchDesiredState(context.Context, *PatchDesiredStateRequest) (*PatchDesiredStateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PatchDesiredState not implemented")
}
func (*UnimplementedShadowsServer) StreamReportedStateChanges(*StreamReportedStateChangesRequest, Shadows_StreamReportedStateChangesServer) error {
	return status.Errorf(codes.Unimplemented, "method StreamReportedStateChanges not implemented")
}

func RegisterShadowsServer(s *grpc.Server, srv ShadowsServer) {
	s.RegisterService(&_Shadows_serviceDesc, srv)
}

func _Shadows_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ShadowsServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/infinimesh.shadow.Shadows/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ShadowsServer).Get(ctx, req.(*GetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Shadows_PatchDesiredState_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PatchDesiredStateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ShadowsServer).PatchDesiredState(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/infinimesh.shadow.Shadows/PatchDesiredState",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ShadowsServer).PatchDesiredState(ctx, req.(*PatchDesiredStateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Shadows_StreamReportedStateChanges_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(StreamReportedStateChangesRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(ShadowsServer).StreamReportedStateChanges(m, &shadowsStreamReportedStateChangesServer{stream})
}

type Shadows_StreamReportedStateChangesServer interface {
	Send(*StreamReportedStateChangesResponse) error
	grpc.ServerStream
}

type shadowsStreamReportedStateChangesServer struct {
	grpc.ServerStream
}

func (x *shadowsStreamReportedStateChangesServer) Send(m *StreamReportedStateChangesResponse) error {
	return x.ServerStream.SendMsg(m)
}

var _Shadows_serviceDesc = grpc.ServiceDesc{
	ServiceName: "infinimesh.shadow.Shadows",
	HandlerType: (*ShadowsServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Get",
			Handler:    _Shadows_Get_Handler,
		},
		{
			MethodName: "PatchDesiredState",
			Handler:    _Shadows_PatchDesiredState_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "StreamReportedStateChanges",
			Handler:       _Shadows_StreamReportedStateChanges_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "pkg/shadow/shadowpb/shadow.proto",
}
