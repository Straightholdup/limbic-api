// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.20.0
// source: protos/speech_emotion_recognition.proto

package speech_emotion_recognition

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

// SpeechEmotionRecognitionClient is the client API for SpeechEmotionRecognition service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SpeechEmotionRecognitionClient interface {
	LoadData(ctx context.Context, opts ...grpc.CallOption) (SpeechEmotionRecognition_LoadDataClient, error)
}

type speechEmotionRecognitionClient struct {
	cc grpc.ClientConnInterface
}

func NewSpeechEmotionRecognitionClient(cc grpc.ClientConnInterface) SpeechEmotionRecognitionClient {
	return &speechEmotionRecognitionClient{cc}
}

func (c *speechEmotionRecognitionClient) LoadData(ctx context.Context, opts ...grpc.CallOption) (SpeechEmotionRecognition_LoadDataClient, error) {
	stream, err := c.cc.NewStream(ctx, &SpeechEmotionRecognition_ServiceDesc.Streams[0], "/limbic.protos.speech_emotion_recognition.SpeechEmotionRecognition/LoadData", opts...)
	if err != nil {
		return nil, err
	}
	x := &speechEmotionRecognitionLoadDataClient{stream}
	return x, nil
}

type SpeechEmotionRecognition_LoadDataClient interface {
	Send(*Chunk) error
	CloseAndRecv() (*LoadDataReply, error)
	grpc.ClientStream
}

type speechEmotionRecognitionLoadDataClient struct {
	grpc.ClientStream
}

func (x *speechEmotionRecognitionLoadDataClient) Send(m *Chunk) error {
	return x.ClientStream.SendMsg(m)
}

func (x *speechEmotionRecognitionLoadDataClient) CloseAndRecv() (*LoadDataReply, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(LoadDataReply)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// SpeechEmotionRecognitionServer is the server API for SpeechEmotionRecognition service.
// All implementations must embed UnimplementedSpeechEmotionRecognitionServer
// for forward compatibility
type SpeechEmotionRecognitionServer interface {
	LoadData(SpeechEmotionRecognition_LoadDataServer) error
	mustEmbedUnimplementedSpeechEmotionRecognitionServer()
}

// UnimplementedSpeechEmotionRecognitionServer must be embedded to have forward compatible implementations.
type UnimplementedSpeechEmotionRecognitionServer struct {
}

func (UnimplementedSpeechEmotionRecognitionServer) LoadData(SpeechEmotionRecognition_LoadDataServer) error {
	return status.Errorf(codes.Unimplemented, "method LoadData not implemented")
}
func (UnimplementedSpeechEmotionRecognitionServer) mustEmbedUnimplementedSpeechEmotionRecognitionServer() {
}

// UnsafeSpeechEmotionRecognitionServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SpeechEmotionRecognitionServer will
// result in compilation errors.
type UnsafeSpeechEmotionRecognitionServer interface {
	mustEmbedUnimplementedSpeechEmotionRecognitionServer()
}

func RegisterSpeechEmotionRecognitionServer(s grpc.ServiceRegistrar, srv SpeechEmotionRecognitionServer) {
	s.RegisterService(&SpeechEmotionRecognition_ServiceDesc, srv)
}

func _SpeechEmotionRecognition_LoadData_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(SpeechEmotionRecognitionServer).LoadData(&speechEmotionRecognitionLoadDataServer{stream})
}

type SpeechEmotionRecognition_LoadDataServer interface {
	SendAndClose(*LoadDataReply) error
	Recv() (*Chunk, error)
	grpc.ServerStream
}

type speechEmotionRecognitionLoadDataServer struct {
	grpc.ServerStream
}

func (x *speechEmotionRecognitionLoadDataServer) SendAndClose(m *LoadDataReply) error {
	return x.ServerStream.SendMsg(m)
}

func (x *speechEmotionRecognitionLoadDataServer) Recv() (*Chunk, error) {
	m := new(Chunk)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// SpeechEmotionRecognition_ServiceDesc is the grpc.ServiceDesc for SpeechEmotionRecognition service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var SpeechEmotionRecognition_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "limbic.protos.speech_emotion_recognition.SpeechEmotionRecognition",
	HandlerType: (*SpeechEmotionRecognitionServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "LoadData",
			Handler:       _SpeechEmotionRecognition_LoadData_Handler,
			ClientStreams: true,
		},
	},
	Metadata: "protos/speech_emotion_recognition.proto",
}
