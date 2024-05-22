package interceptor

import (
	"context"
	"strings"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"

	"github.com/virzz/vlog"
)

type authInterceptor struct {
	grpc.ServerStream
	ctx     context.Context
	token   string
	logging bool
}

const (
	headerAuthorize = "Authorization"
	expectedScheme  = "Bearer"
)

func (i *authInterceptor) Context() context.Context { return i.ctx }
func (i *authInterceptor) RecvMsg(m any) error      { return i.ServerStream.RecvMsg(m) }
func (i *authInterceptor) SendMsg(m any) error      { return i.ServerStream.SendMsg(m) }
func (i *authInterceptor) Warp(s grpc.ServerStream) *authInterceptor {
	if existing, ok := s.(*authInterceptor); ok {
		return existing
	}
	i.ServerStream, i.ctx = s, s.Context()
	return i
}

func (i *authInterceptor) Stream(srv any, ss grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) (err error) {
	if i.logging {
		vlog.Info(info.FullMethod)
	}
	vals := metadata.ValueFromIncomingContext(ss.Context(), headerAuthorize)
	if vals == nil || len(vals) < 2 || !strings.EqualFold(vals[0], expectedScheme) {
		err = status.Error(codes.Unauthenticated, "Request unauthenticated with "+expectedScheme)
	} else if !strings.EqualFold(vals[1], i.token) {
		err = status.Error(codes.PermissionDenied, "Request bad token")
	}
	if err != nil {
		vlog.Error(info.FullMethod, "err", err.Error())
		return err
	}
	err = handler(srv, i.Warp(ss))
	if err != nil {
		vlog.Error(info.FullMethod, "err", err.Error())
		return err
	}
	return nil
}

func (i *authInterceptor) Unary(ctx context.Context, req any, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (rsp any, err error) {
	if i.logging {
		vlog.Info(info.FullMethod)
	}
	vals := metadata.ValueFromIncomingContext(ctx, headerAuthorize)
	if vals == nil || len(vals) < 2 || !strings.EqualFold(vals[0], expectedScheme) {
		err = status.Error(codes.Unauthenticated, "Request unauthenticated with "+expectedScheme)
	} else if !strings.EqualFold(vals[1], i.token) {
		err = status.Error(codes.PermissionDenied, "Request bad token")
	}
	if err != nil {
		vlog.Error(info.FullMethod, "err", err.Error())
		return nil, err
	}
	rsp, err = handler(ctx, req)
	if err != nil {
		vlog.Error(info.FullMethod, "err", err.Error())
		return nil, err
	}
	return
}

func (i *authInterceptor) WithLogging() { i.logging = !i.logging }

func NewAuthInterceptor(token string) *authInterceptor {
	return &authInterceptor{token: token}
}
