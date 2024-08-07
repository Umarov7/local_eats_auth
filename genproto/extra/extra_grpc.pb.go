// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.12.4
// source: extra.proto

package extra

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

// ExtraClient is the client API for Extra service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ExtraClient interface {
	GetStatistics(ctx context.Context, in *Period, opts ...grpc.CallOption) (*Statistics, error)
	TrackActivity(ctx context.Context, in *Period, opts ...grpc.CallOption) (*Activity, error)
	SetWorkingHours(ctx context.Context, in *WorkingHours, opts ...grpc.CallOption) (*WorkingHoursResp, error)
	GetNutrition(ctx context.Context, in *ID, opts ...grpc.CallOption) (*NutritionalInfo, error)
}

type extraClient struct {
	cc grpc.ClientConnInterface
}

func NewExtraClient(cc grpc.ClientConnInterface) ExtraClient {
	return &extraClient{cc}
}

func (c *extraClient) GetStatistics(ctx context.Context, in *Period, opts ...grpc.CallOption) (*Statistics, error) {
	out := new(Statistics)
	err := c.cc.Invoke(ctx, "/extra.Extra/GetStatistics", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *extraClient) TrackActivity(ctx context.Context, in *Period, opts ...grpc.CallOption) (*Activity, error) {
	out := new(Activity)
	err := c.cc.Invoke(ctx, "/extra.Extra/TrackActivity", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *extraClient) SetWorkingHours(ctx context.Context, in *WorkingHours, opts ...grpc.CallOption) (*WorkingHoursResp, error) {
	out := new(WorkingHoursResp)
	err := c.cc.Invoke(ctx, "/extra.Extra/SetWorkingHours", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *extraClient) GetNutrition(ctx context.Context, in *ID, opts ...grpc.CallOption) (*NutritionalInfo, error) {
	out := new(NutritionalInfo)
	err := c.cc.Invoke(ctx, "/extra.Extra/GetNutrition", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ExtraServer is the server API for Extra service.
// All implementations must embed UnimplementedExtraServer
// for forward compatibility
type ExtraServer interface {
	GetStatistics(context.Context, *Period) (*Statistics, error)
	TrackActivity(context.Context, *Period) (*Activity, error)
	SetWorkingHours(context.Context, *WorkingHours) (*WorkingHoursResp, error)
	GetNutrition(context.Context, *ID) (*NutritionalInfo, error)
	mustEmbedUnimplementedExtraServer()
}

// UnimplementedExtraServer must be embedded to have forward compatible implementations.
type UnimplementedExtraServer struct {
}

func (UnimplementedExtraServer) GetStatistics(context.Context, *Period) (*Statistics, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetStatistics not implemented")
}
func (UnimplementedExtraServer) TrackActivity(context.Context, *Period) (*Activity, error) {
	return nil, status.Errorf(codes.Unimplemented, "method TrackActivity not implemented")
}
func (UnimplementedExtraServer) SetWorkingHours(context.Context, *WorkingHours) (*WorkingHoursResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetWorkingHours not implemented")
}
func (UnimplementedExtraServer) GetNutrition(context.Context, *ID) (*NutritionalInfo, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetNutrition not implemented")
}
func (UnimplementedExtraServer) mustEmbedUnimplementedExtraServer() {}

// UnsafeExtraServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ExtraServer will
// result in compilation errors.
type UnsafeExtraServer interface {
	mustEmbedUnimplementedExtraServer()
}

func RegisterExtraServer(s grpc.ServiceRegistrar, srv ExtraServer) {
	s.RegisterService(&Extra_ServiceDesc, srv)
}

func _Extra_GetStatistics_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Period)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ExtraServer).GetStatistics(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/extra.Extra/GetStatistics",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ExtraServer).GetStatistics(ctx, req.(*Period))
	}
	return interceptor(ctx, in, info, handler)
}

func _Extra_TrackActivity_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Period)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ExtraServer).TrackActivity(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/extra.Extra/TrackActivity",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ExtraServer).TrackActivity(ctx, req.(*Period))
	}
	return interceptor(ctx, in, info, handler)
}

func _Extra_SetWorkingHours_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(WorkingHours)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ExtraServer).SetWorkingHours(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/extra.Extra/SetWorkingHours",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ExtraServer).SetWorkingHours(ctx, req.(*WorkingHours))
	}
	return interceptor(ctx, in, info, handler)
}

func _Extra_GetNutrition_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ExtraServer).GetNutrition(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/extra.Extra/GetNutrition",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ExtraServer).GetNutrition(ctx, req.(*ID))
	}
	return interceptor(ctx, in, info, handler)
}

// Extra_ServiceDesc is the grpc.ServiceDesc for Extra service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Extra_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "extra.Extra",
	HandlerType: (*ExtraServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetStatistics",
			Handler:    _Extra_GetStatistics_Handler,
		},
		{
			MethodName: "TrackActivity",
			Handler:    _Extra_TrackActivity_Handler,
		},
		{
			MethodName: "SetWorkingHours",
			Handler:    _Extra_SetWorkingHours_Handler,
		},
		{
			MethodName: "GetNutrition",
			Handler:    _Extra_GetNutrition_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "extra.proto",
}
