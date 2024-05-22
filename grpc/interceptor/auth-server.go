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

const (
	HeaderAuthorize = "Authorization"
	TokenScheme     = "Bearer"
)

type AuthServer struct {
	grpc.ServerStream
	ctx     context.Context
	token   string
	logging bool
}

func (i *AuthServer) WithLogging()             { i.logging = !i.logging }
func (i *AuthServer) Context() context.Context { return i.ctx }
func (i *AuthServer) RecvMsg(m any) error      { return i.ServerStream.RecvMsg(m) }
func (i *AuthServer) SendMsg(m any) error      { return i.ServerStream.SendMsg(m) }
func (i *AuthServer) Warp(s grpc.ServerStream) *AuthServer {
	if existing, ok := s.(*AuthServer); ok {
		return existing
	}
	i.ServerStream, i.ctx = s, s.Context()
	return i
}

func (i *AuthServer) Stream(srv any, ss grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) (err error) {
	if i.logging {
		vlog.Info(info.FullMethod)
	}
	vals := metadata.ValueFromIncomingContext(ss.Context(), HeaderAuthorize)
	if vals == nil || len(vals) < 2 || !strings.EqualFold(vals[0], TokenScheme) {
		err = status.Error(codes.Unauthenticated, "Request unauthenticated with "+TokenScheme)
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

func (i *AuthServer) Unary(ctx context.Context, req any, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (rsp any, err error) {
	if i.logging {
		vlog.Info(info.FullMethod)
	}
	vals := metadata.ValueFromIncomingContext(ctx, HeaderAuthorize)
	if vals == nil || len(vals) < 2 || !strings.EqualFold(vals[0], TokenScheme) {
		err = status.Error(codes.Unauthenticated, "Request unauthenticated with "+TokenScheme)
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

func NewAuthServer(token string) *AuthServer { return &AuthServer{token: token} }
