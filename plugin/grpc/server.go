package grpc

import (
	"context"

	pinpoint "github.com/pinpoint-apm/pinpoint-go-agent"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

const serviceTypeGrpcServer = 1130

type serverStream struct {
	grpc.ServerStream
	context context.Context
}

func (s *serverStream) Context() context.Context {
	return s.context
}

type DistributedTracingContextReaderMD struct {
	md metadata.MD
}

func (m DistributedTracingContextReaderMD) Get(key string) string {
	v := m.md.Get(key)
	if len(v) == 0 {
		return ""
	}
	return v[0]
}

func startSpan(ctx context.Context, rpcName string) pinpoint.Tracer {
	md, _ := metadata.FromIncomingContext(ctx) // nil is ok
	reader := &DistributedTracingContextReaderMD{md}
	tracer := pinpoint.GetAgent().NewSpanTracerWithReader("gRPC Server", rpcName, reader)
	tracer.Span().SetServiceType(serviceTypeGrpcServer)

	return tracer
}

func UnaryServerInterceptor() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		if !pinpoint.GetAgent().Enable() {
			return handler(ctx, req)
		}

		tracer := startSpan(ctx, info.FullMethod)
		defer tracer.EndSpan()
		defer tracer.NewSpanEvent(info.FullMethod).EndSpanEvent()

		ctx = pinpoint.NewContext(ctx, tracer)
		resp, err := handler(ctx, req)
		if err != nil {
			tracer.Span().SetError(err)
		}
		return resp, err
	}
}

func StreamServerInterceptor() grpc.StreamServerInterceptor {
	return func(srv interface{}, stream grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) error {
		if !pinpoint.GetAgent().Enable() {
			return handler(srv, stream)
		}

		tracer := startSpan(stream.Context(), info.FullMethod)
		defer tracer.EndSpan()
		defer tracer.NewSpanEvent(info.FullMethod).EndSpanEvent()

		ctx := pinpoint.NewContext(stream.Context(), tracer)
		wrappedStream := &serverStream{stream, ctx}
		err := handler(srv, wrappedStream)
		if err != nil {
			tracer.Span().SetError(err)
		}
		return err
	}
}
