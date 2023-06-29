package user

import (
	user "github.com/NpoolPlatform/message/npool/notif/mw/v1/notif/user"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
)

type Server struct {
	user.UnimplementedMiddlewareServer
}

func Register(server grpc.ServiceRegistrar) {
	user.RegisterMiddlewareServer(server, &Server{})
}

func RegisterGateway(mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) error {
	return nil
}
