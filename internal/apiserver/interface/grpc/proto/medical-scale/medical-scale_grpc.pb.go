// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.29.3
// source: medical-scale.proto

package medical_scale

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	MedicalScaleService_GetMedicalScaleByCode_FullMethodName              = "/medical_scale.MedicalScaleService/GetMedicalScaleByCode"
	MedicalScaleService_GetMedicalScaleByQuestionnaireCode_FullMethodName = "/medical_scale.MedicalScaleService/GetMedicalScaleByQuestionnaireCode"
)

// MedicalScaleServiceClient is the client API for MedicalScaleService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
//
// MedicalScaleService 医学量表服务
type MedicalScaleServiceClient interface {
	// GetMedicalScaleByCode 根据医学量表代码获取医学量表详情
	GetMedicalScaleByCode(ctx context.Context, in *GetMedicalScaleByCodeRequest, opts ...grpc.CallOption) (*GetMedicalScaleByCodeResponse, error)
	// GetMedicalScaleByQuestionnaireCode 根据问卷代码获取医学量表详情
	GetMedicalScaleByQuestionnaireCode(ctx context.Context, in *GetMedicalScaleByQuestionnaireCodeRequest, opts ...grpc.CallOption) (*GetMedicalScaleByQuestionnaireCodeResponse, error)
}

type medicalScaleServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewMedicalScaleServiceClient(cc grpc.ClientConnInterface) MedicalScaleServiceClient {
	return &medicalScaleServiceClient{cc}
}

func (c *medicalScaleServiceClient) GetMedicalScaleByCode(ctx context.Context, in *GetMedicalScaleByCodeRequest, opts ...grpc.CallOption) (*GetMedicalScaleByCodeResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetMedicalScaleByCodeResponse)
	err := c.cc.Invoke(ctx, MedicalScaleService_GetMedicalScaleByCode_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *medicalScaleServiceClient) GetMedicalScaleByQuestionnaireCode(ctx context.Context, in *GetMedicalScaleByQuestionnaireCodeRequest, opts ...grpc.CallOption) (*GetMedicalScaleByQuestionnaireCodeResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetMedicalScaleByQuestionnaireCodeResponse)
	err := c.cc.Invoke(ctx, MedicalScaleService_GetMedicalScaleByQuestionnaireCode_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MedicalScaleServiceServer is the server API for MedicalScaleService service.
// All implementations must embed UnimplementedMedicalScaleServiceServer
// for forward compatibility.
//
// MedicalScaleService 医学量表服务
type MedicalScaleServiceServer interface {
	// GetMedicalScaleByCode 根据医学量表代码获取医学量表详情
	GetMedicalScaleByCode(context.Context, *GetMedicalScaleByCodeRequest) (*GetMedicalScaleByCodeResponse, error)
	// GetMedicalScaleByQuestionnaireCode 根据问卷代码获取医学量表详情
	GetMedicalScaleByQuestionnaireCode(context.Context, *GetMedicalScaleByQuestionnaireCodeRequest) (*GetMedicalScaleByQuestionnaireCodeResponse, error)
	mustEmbedUnimplementedMedicalScaleServiceServer()
}

// UnimplementedMedicalScaleServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedMedicalScaleServiceServer struct{}

func (UnimplementedMedicalScaleServiceServer) GetMedicalScaleByCode(context.Context, *GetMedicalScaleByCodeRequest) (*GetMedicalScaleByCodeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetMedicalScaleByCode not implemented")
}
func (UnimplementedMedicalScaleServiceServer) GetMedicalScaleByQuestionnaireCode(context.Context, *GetMedicalScaleByQuestionnaireCodeRequest) (*GetMedicalScaleByQuestionnaireCodeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetMedicalScaleByQuestionnaireCode not implemented")
}
func (UnimplementedMedicalScaleServiceServer) mustEmbedUnimplementedMedicalScaleServiceServer() {}
func (UnimplementedMedicalScaleServiceServer) testEmbeddedByValue()                             {}

// UnsafeMedicalScaleServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MedicalScaleServiceServer will
// result in compilation errors.
type UnsafeMedicalScaleServiceServer interface {
	mustEmbedUnimplementedMedicalScaleServiceServer()
}

func RegisterMedicalScaleServiceServer(s grpc.ServiceRegistrar, srv MedicalScaleServiceServer) {
	// If the following call pancis, it indicates UnimplementedMedicalScaleServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&MedicalScaleService_ServiceDesc, srv)
}

func _MedicalScaleService_GetMedicalScaleByCode_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetMedicalScaleByCodeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MedicalScaleServiceServer).GetMedicalScaleByCode(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MedicalScaleService_GetMedicalScaleByCode_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MedicalScaleServiceServer).GetMedicalScaleByCode(ctx, req.(*GetMedicalScaleByCodeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MedicalScaleService_GetMedicalScaleByQuestionnaireCode_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetMedicalScaleByQuestionnaireCodeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MedicalScaleServiceServer).GetMedicalScaleByQuestionnaireCode(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MedicalScaleService_GetMedicalScaleByQuestionnaireCode_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MedicalScaleServiceServer).GetMedicalScaleByQuestionnaireCode(ctx, req.(*GetMedicalScaleByQuestionnaireCodeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// MedicalScaleService_ServiceDesc is the grpc.ServiceDesc for MedicalScaleService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var MedicalScaleService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "medical_scale.MedicalScaleService",
	HandlerType: (*MedicalScaleServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetMedicalScaleByCode",
			Handler:    _MedicalScaleService_GetMedicalScaleByCode_Handler,
		},
		{
			MethodName: "GetMedicalScaleByQuestionnaireCode",
			Handler:    _MedicalScaleService_GetMedicalScaleByQuestionnaireCode_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "medical-scale.proto",
}
