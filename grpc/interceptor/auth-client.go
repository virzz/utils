package interceptor

import (
	"context"

	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

type AuthClient struct {
	grpc.ClientStream

	md metadata.MD
}

func (i *AuthClient) RecvMsg(m interface{}) error { return i.ClientStream.RecvMsg(m) }
func (i *AuthClient) SendMsg(m interface{}) error { return i.ClientStream.SendMsg(m) }
func (i *AuthClient) Wrap(ctx context.Context, desc *grpc.StreamDesc, cc *grpc.ClientConn, method string, streamer grpc.Streamer, opts ...grpc.CallOption) (grpc.ClientStream, error) {
	s, err := streamer(ctx, desc, cc, method, opts...)
	if err != nil {
		return nil, err
	}
	i.ClientStream = s
	return i, nil
}

func (i *AuthClient) Unary(ctx context.Context, method string, req, reply any, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
	ctx = metadata.NewOutgoingContext(ctx, i.md)
	return invoker(ctx, method, req, reply, cc, opts...)
}

func (i *AuthClient) Stream(ctx context.Context, desc *grpc.StreamDesc, cc *grpc.ClientConn, method string, streamer grpc.Streamer, opts ...grpc.CallOption) (grpc.ClientStream, error) {
	ctx = metadata.NewOutgoingContext(ctx, i.md)
	return i.Wrap(ctx, desc, cc, method, streamer, opts...)
}

func NewAuthClient(token string) *AuthClient {
	return &AuthClient{md: metadata.New(map[string]string{
		HeaderAuthorize: TokenScheme + " " + token,
	})}
}
