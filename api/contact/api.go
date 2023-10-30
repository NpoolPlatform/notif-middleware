package contact

import (
	"context"

	contact "github.com/NpoolPlatform/message/npool/notif/mw/v1/contact"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
)

type Server struct {
	contact.UnimplementedMiddlewareServer
}

func Register(server grpc.ServiceRegistrar) {
	contact.RegisterMiddlewareServer(server, &Server{})
}

func RegisterGateway(mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) error {
	return contact.RegisterMiddlewareHandlerFromEndpoint(context.Background(), mux, endpoint, opts)
}
