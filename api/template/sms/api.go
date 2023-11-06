package sms

import (
	"context"

	"github.com/NpoolPlatform/message/npool/notif/mw/v1/template/sms"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"

	"google.golang.org/grpc"
)

type Server struct {
	sms.UnimplementedMiddlewareServer
}

func Register(server grpc.ServiceRegistrar) {
	sms.RegisterMiddlewareServer(server, &Server{})
}

func RegisterGateway(mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) error {
	return sms.RegisterMiddlewareHandlerFromEndpoint(context.Background(), mux, endpoint, opts)
}
