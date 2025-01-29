// Copyright 2024 Gravitational, Inc
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

// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             (unknown)
// source: teleport/decision/v1alpha1/decision_service.proto

package decisionpb

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	DecisionService_EvaluateSSHAccess_FullMethodName      = "/teleport.decision.v1alpha1.DecisionService/EvaluateSSHAccess"
	DecisionService_EvaluateDatabaseAccess_FullMethodName = "/teleport.decision.v1alpha1.DecisionService/EvaluateDatabaseAccess"
)

// DecisionServiceClient is the client API for DecisionService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type DecisionServiceClient interface {
	// EvaluateSSHAccess evaluates an SSH access attempt.
	EvaluateSSHAccess(ctx context.Context, in *EvaluateSSHAccessRequest, opts ...grpc.CallOption) (*EvaluateSSHAccessResponse, error)
	// EvaluateDatabaseAccess evaluate a database access attempt.
	EvaluateDatabaseAccess(ctx context.Context, in *EvaluateDatabaseAccessRequest, opts ...grpc.CallOption) (*EvaluateDatabaseAccessResponse, error)
}

type decisionServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewDecisionServiceClient(cc grpc.ClientConnInterface) DecisionServiceClient {
	return &decisionServiceClient{cc}
}

func (c *decisionServiceClient) EvaluateSSHAccess(ctx context.Context, in *EvaluateSSHAccessRequest, opts ...grpc.CallOption) (*EvaluateSSHAccessResponse, error) {
	out := new(EvaluateSSHAccessResponse)
	err := c.cc.Invoke(ctx, DecisionService_EvaluateSSHAccess_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *decisionServiceClient) EvaluateDatabaseAccess(ctx context.Context, in *EvaluateDatabaseAccessRequest, opts ...grpc.CallOption) (*EvaluateDatabaseAccessResponse, error) {
	out := new(EvaluateDatabaseAccessResponse)
	err := c.cc.Invoke(ctx, DecisionService_EvaluateDatabaseAccess_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DecisionServiceServer is the server API for DecisionService service.
// All implementations must embed UnimplementedDecisionServiceServer
// for forward compatibility
type DecisionServiceServer interface {
	// EvaluateSSHAccess evaluates an SSH access attempt.
	EvaluateSSHAccess(context.Context, *EvaluateSSHAccessRequest) (*EvaluateSSHAccessResponse, error)
	// EvaluateDatabaseAccess evaluate a database access attempt.
	EvaluateDatabaseAccess(context.Context, *EvaluateDatabaseAccessRequest) (*EvaluateDatabaseAccessResponse, error)
	mustEmbedUnimplementedDecisionServiceServer()
}

// UnimplementedDecisionServiceServer must be embedded to have forward compatible implementations.
type UnimplementedDecisionServiceServer struct {
}

func (UnimplementedDecisionServiceServer) EvaluateSSHAccess(context.Context, *EvaluateSSHAccessRequest) (*EvaluateSSHAccessResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method EvaluateSSHAccess not implemented")
}
func (UnimplementedDecisionServiceServer) EvaluateDatabaseAccess(context.Context, *EvaluateDatabaseAccessRequest) (*EvaluateDatabaseAccessResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method EvaluateDatabaseAccess not implemented")
}
func (UnimplementedDecisionServiceServer) mustEmbedUnimplementedDecisionServiceServer() {}

// UnsafeDecisionServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to DecisionServiceServer will
// result in compilation errors.
type UnsafeDecisionServiceServer interface {
	mustEmbedUnimplementedDecisionServiceServer()
}

func RegisterDecisionServiceServer(s grpc.ServiceRegistrar, srv DecisionServiceServer) {
	s.RegisterService(&DecisionService_ServiceDesc, srv)
}

func _DecisionService_EvaluateSSHAccess_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EvaluateSSHAccessRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DecisionServiceServer).EvaluateSSHAccess(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DecisionService_EvaluateSSHAccess_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DecisionServiceServer).EvaluateSSHAccess(ctx, req.(*EvaluateSSHAccessRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DecisionService_EvaluateDatabaseAccess_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EvaluateDatabaseAccessRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DecisionServiceServer).EvaluateDatabaseAccess(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DecisionService_EvaluateDatabaseAccess_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DecisionServiceServer).EvaluateDatabaseAccess(ctx, req.(*EvaluateDatabaseAccessRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// DecisionService_ServiceDesc is the grpc.ServiceDesc for DecisionService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var DecisionService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "teleport.decision.v1alpha1.DecisionService",
	HandlerType: (*DecisionServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "EvaluateSSHAccess",
			Handler:    _DecisionService_EvaluateSSHAccess_Handler,
		},
		{
			MethodName: "EvaluateDatabaseAccess",
			Handler:    _DecisionService_EvaluateDatabaseAccess_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "teleport/decision/v1alpha1/decision_service.proto",
}
