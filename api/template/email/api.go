package email

import (
	"context"

	"github.com/NpoolPlatform/message/npool/notif/mw/v1/template/email"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
)

type Server struct {
	email.UnimplementedMiddlewareServer
}

func Register(server grpc.ServiceRegistrar) {
	email.RegisterMiddlewareServer(server, &Server{})
}

func RegisterGateway(mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) error {
	return email.RegisterMiddlewareHandlerFromEndpoint(context.Background(), mux, endpoint, opts)
}
