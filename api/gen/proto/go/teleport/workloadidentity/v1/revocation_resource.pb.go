// Copyright 2025 Gravitational, Inc
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        (unknown)
// source: teleport/workloadidentity/v1/revocation_resource.proto

package workloadidentityv1

import (
	v1 "github.com/gravitational/teleport/api/gen/proto/go/teleport/header/v1"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// WorkloadIdentityX509Revocation represents the revocation of a single X509
// workload identity credential. Creating or deleting these resources triggers
// the regeneration of the trust domain CRL.
//
// The name of a WorkloadIdentityX509Revocation must be the base16, lower case,
// encoded serial number of the revoked X509 certificate. Therefore, only a
// single revocation entry can exist for a given certificate.
type WorkloadIdentityX509Revocation struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The kind of resource represented. For WorkloadIdentityX509Revocation
	// resources, this is always `workload_identity_x509_revocation`.
	Kind string `protobuf:"bytes,1,opt,name=kind,proto3" json:"kind,omitempty"`
	// Differentiates variations of the same kind. All resources should
	// contain one, even if it is never populated.
	SubKind string `protobuf:"bytes,2,opt,name=sub_kind,json=subKind,proto3" json:"sub_kind,omitempty"`
	// The version of the resource being represented.
	Version string `protobuf:"bytes,3,opt,name=version,proto3" json:"version,omitempty"`
	// Common metadata that all resources share.
	Metadata *v1.Metadata `protobuf:"bytes,4,opt,name=metadata,proto3" json:"metadata,omitempty"`
	// The configured properties of the WorkloadIdentityX509Revocation
	Spec *WorkloadIdentityX509RevocationSpec `protobuf:"bytes,5,opt,name=spec,proto3" json:"spec,omitempty"`
}

func (x *WorkloadIdentityX509Revocation) Reset() {
	*x = WorkloadIdentityX509Revocation{}
	if protoimpl.UnsafeEnabled {
		mi := &file_teleport_workloadidentity_v1_revocation_resource_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WorkloadIdentityX509Revocation) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WorkloadIdentityX509Revocation) ProtoMessage() {}

func (x *WorkloadIdentityX509Revocation) ProtoReflect() protoreflect.Message {
	mi := &file_teleport_workloadidentity_v1_revocation_resource_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WorkloadIdentityX509Revocation.ProtoReflect.Descriptor instead.
func (*WorkloadIdentityX509Revocation) Descriptor() ([]byte, []int) {
	return file_teleport_workloadidentity_v1_revocation_resource_proto_rawDescGZIP(), []int{0}
}

func (x *WorkloadIdentityX509Revocation) GetKind() string {
	if x != nil {
		return x.Kind
	}
	return ""
}

func (x *WorkloadIdentityX509Revocation) GetSubKind() string {
	if x != nil {
		return x.SubKind
	}
	return ""
}

func (x *WorkloadIdentityX509Revocation) GetVersion() string {
	if x != nil {
		return x.Version
	}
	return ""
}

func (x *WorkloadIdentityX509Revocation) GetMetadata() *v1.Metadata {
	if x != nil {
		return x.Metadata
	}
	return nil
}

func (x *WorkloadIdentityX509Revocation) GetSpec() *WorkloadIdentityX509RevocationSpec {
	if x != nil {
		return x.Spec
	}
	return nil
}

// Configuration specific to WorkloadIdentityX509Revocation.
type WorkloadIdentityX509RevocationSpec struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Describes why this revocation entry was created.
	// Required.
	Reason string `protobuf:"bytes,1,opt,name=reason,proto3" json:"reason,omitempty"`
	// The time at which the revocation entry was created.
	// Required.
	RevokedAt *timestamppb.Timestamp `protobuf:"bytes,2,opt,name=revoked_at,json=revokedAt,proto3" json:"revoked_at,omitempty"`
}

func (x *WorkloadIdentityX509RevocationSpec) Reset() {
	*x = WorkloadIdentityX509RevocationSpec{}
	if protoimpl.UnsafeEnabled {
		mi := &file_teleport_workloadidentity_v1_revocation_resource_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WorkloadIdentityX509RevocationSpec) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WorkloadIdentityX509RevocationSpec) ProtoMessage() {}

func (x *WorkloadIdentityX509RevocationSpec) ProtoReflect() protoreflect.Message {
	mi := &file_teleport_workloadidentity_v1_revocation_resource_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WorkloadIdentityX509RevocationSpec.ProtoReflect.Descriptor instead.
func (*WorkloadIdentityX509RevocationSpec) Descriptor() ([]byte, []int) {
	return file_teleport_workloadidentity_v1_revocation_resource_proto_rawDescGZIP(), []int{1}
}

func (x *WorkloadIdentityX509RevocationSpec) GetReason() string {
	if x != nil {
		return x.Reason
	}
	return ""
}

func (x *WorkloadIdentityX509RevocationSpec) GetRevokedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.RevokedAt
	}
	return nil
}

var File_teleport_workloadidentity_v1_revocation_resource_proto protoreflect.FileDescriptor

var file_teleport_workloadidentity_v1_revocation_resource_proto_rawDesc = []byte{
	0x0a, 0x36, 0x74, 0x65, 0x6c, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x2f, 0x77, 0x6f, 0x72, 0x6b, 0x6c,
	0x6f, 0x61, 0x64, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x2f, 0x76, 0x31, 0x2f, 0x72,
	0x65, 0x76, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1c, 0x74, 0x65, 0x6c, 0x65, 0x70, 0x6f,
	0x72, 0x74, 0x2e, 0x77, 0x6f, 0x72, 0x6b, 0x6c, 0x6f, 0x61, 0x64, 0x69, 0x64, 0x65, 0x6e, 0x74,
	0x69, 0x74, 0x79, 0x2e, 0x76, 0x31, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d,
	0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x21, 0x74, 0x65, 0x6c, 0x65, 0x70, 0x6f, 0x72,
	0x74, 0x2f, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x2f, 0x76, 0x31, 0x2f, 0x6d, 0x65, 0x74, 0x61,
	0x64, 0x61, 0x74, 0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xf9, 0x01, 0x0a, 0x1e, 0x57,
	0x6f, 0x72, 0x6b, 0x6c, 0x6f, 0x61, 0x64, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x58,
	0x35, 0x30, 0x39, 0x52, 0x65, 0x76, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x12, 0x0a,
	0x04, 0x6b, 0x69, 0x6e, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6b, 0x69, 0x6e,
	0x64, 0x12, 0x19, 0x0a, 0x08, 0x73, 0x75, 0x62, 0x5f, 0x6b, 0x69, 0x6e, 0x64, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x07, 0x73, 0x75, 0x62, 0x4b, 0x69, 0x6e, 0x64, 0x12, 0x18, 0x0a, 0x07,
	0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x76,
	0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x38, 0x0a, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61,
	0x74, 0x61, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x74, 0x65, 0x6c, 0x65, 0x70,
	0x6f, 0x72, 0x74, 0x2e, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x4d, 0x65,
	0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x52, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61,
	0x12, 0x54, 0x0a, 0x04, 0x73, 0x70, 0x65, 0x63, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x40,
	0x2e, 0x74, 0x65, 0x6c, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x2e, 0x77, 0x6f, 0x72, 0x6b, 0x6c, 0x6f,
	0x61, 0x64, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x2e, 0x76, 0x31, 0x2e, 0x57, 0x6f,
	0x72, 0x6b, 0x6c, 0x6f, 0x61, 0x64, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x58, 0x35,
	0x30, 0x39, 0x52, 0x65, 0x76, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x70, 0x65, 0x63,
	0x52, 0x04, 0x73, 0x70, 0x65, 0x63, 0x22, 0x77, 0x0a, 0x22, 0x57, 0x6f, 0x72, 0x6b, 0x6c, 0x6f,
	0x61, 0x64, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x58, 0x35, 0x30, 0x39, 0x52, 0x65,
	0x76, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x70, 0x65, 0x63, 0x12, 0x16, 0x0a, 0x06,
	0x72, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x72, 0x65,
	0x61, 0x73, 0x6f, 0x6e, 0x12, 0x39, 0x0a, 0x0a, 0x72, 0x65, 0x76, 0x6f, 0x6b, 0x65, 0x64, 0x5f,
	0x61, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73,
	0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x72, 0x65, 0x76, 0x6f, 0x6b, 0x65, 0x64, 0x41, 0x74, 0x42,
	0x64, 0x5a, 0x62, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x67, 0x72,
	0x61, 0x76, 0x69, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x61, 0x6c, 0x2f, 0x74, 0x65, 0x6c, 0x65,
	0x70, 0x6f, 0x72, 0x74, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2f, 0x67, 0x6f, 0x2f, 0x74, 0x65, 0x6c, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x2f, 0x77,
	0x6f, 0x72, 0x6b, 0x6c, 0x6f, 0x61, 0x64, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x2f,
	0x76, 0x31, 0x3b, 0x77, 0x6f, 0x72, 0x6b, 0x6c, 0x6f, 0x61, 0x64, 0x69, 0x64, 0x65, 0x6e, 0x74,
	0x69, 0x74, 0x79, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_teleport_workloadidentity_v1_revocation_resource_proto_rawDescOnce sync.Once
	file_teleport_workloadidentity_v1_revocation_resource_proto_rawDescData = file_teleport_workloadidentity_v1_revocation_resource_proto_rawDesc
)

func file_teleport_workloadidentity_v1_revocation_resource_proto_rawDescGZIP() []byte {
	file_teleport_workloadidentity_v1_revocation_resource_proto_rawDescOnce.Do(func() {
		file_teleport_workloadidentity_v1_revocation_resource_proto_rawDescData = protoimpl.X.CompressGZIP(file_teleport_workloadidentity_v1_revocation_resource_proto_rawDescData)
	})
	return file_teleport_workloadidentity_v1_revocation_resource_proto_rawDescData
}

var file_teleport_workloadidentity_v1_revocation_resource_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_teleport_workloadidentity_v1_revocation_resource_proto_goTypes = []any{
	(*WorkloadIdentityX509Revocation)(nil),     // 0: teleport.workloadidentity.v1.WorkloadIdentityX509Revocation
	(*WorkloadIdentityX509RevocationSpec)(nil), // 1: teleport.workloadidentity.v1.WorkloadIdentityX509RevocationSpec
	(*v1.Metadata)(nil),                        // 2: teleport.header.v1.Metadata
	(*timestamppb.Timestamp)(nil),              // 3: google.protobuf.Timestamp
}
var file_teleport_workloadidentity_v1_revocation_resource_proto_depIdxs = []int32{
	2, // 0: teleport.workloadidentity.v1.WorkloadIdentityX509Revocation.metadata:type_name -> teleport.header.v1.Metadata
	1, // 1: teleport.workloadidentity.v1.WorkloadIdentityX509Revocation.spec:type_name -> teleport.workloadidentity.v1.WorkloadIdentityX509RevocationSpec
	3, // 2: teleport.workloadidentity.v1.WorkloadIdentityX509RevocationSpec.revoked_at:type_name -> google.protobuf.Timestamp
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_teleport_workloadidentity_v1_revocation_resource_proto_init() }
func file_teleport_workloadidentity_v1_revocation_resource_proto_init() {
	if File_teleport_workloadidentity_v1_revocation_resource_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_teleport_workloadidentity_v1_revocation_resource_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*WorkloadIdentityX509Revocation); i {
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
		file_teleport_workloadidentity_v1_revocation_resource_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*WorkloadIdentityX509RevocationSpec); i {
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
			RawDescriptor: file_teleport_workloadidentity_v1_revocation_resource_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_teleport_workloadidentity_v1_revocation_resource_proto_goTypes,
		DependencyIndexes: file_teleport_workloadidentity_v1_revocation_resource_proto_depIdxs,
		MessageInfos:      file_teleport_workloadidentity_v1_revocation_resource_proto_msgTypes,
	}.Build()
	File_teleport_workloadidentity_v1_revocation_resource_proto = out.File
	file_teleport_workloadidentity_v1_revocation_resource_proto_rawDesc = nil
	file_teleport_workloadidentity_v1_revocation_resource_proto_goTypes = nil
	file_teleport_workloadidentity_v1_revocation_resource_proto_depIdxs = nil
}
