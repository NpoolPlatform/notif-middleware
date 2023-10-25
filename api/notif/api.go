package notif

import (
	"context"

	notif "github.com/NpoolPlatform/message/npool/notif/mw/v1/notif"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
)

type Server struct {
	notif.UnimplementedMiddlewareServer
}

func Register(server grpc.ServiceRegistrar) {
	notif.RegisterMiddlewareServer(server, &Server{})
}

func RegisterGateway(mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) error {
	return notif.RegisterMiddlewareHandlerFromEndpoint(context.Background(), mux, endpoint, opts)
}
