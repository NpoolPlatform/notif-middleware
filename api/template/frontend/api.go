package frontend

import (
	"context"

	"github.com/NpoolPlatform/message/npool/notif/mw/v1/template/frontend"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"

	"google.golang.org/grpc"
)

type Server struct {
	frontend.UnimplementedMiddlewareServer
}

func Register(server grpc.ServiceRegistrar) {
	frontend.RegisterMiddlewareServer(server, &Server{})
}

func RegisterGateway(mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) error {
	return frontend.RegisterMiddlewareHandlerFromEndpoint(context.Background(), mux, endpoint, opts)
}
