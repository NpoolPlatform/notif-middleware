package readstate

import (
	readstate "github.com/NpoolPlatform/message/npool/notif/mw/v1/announcement/readstate"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
)

type Server struct {
	readstate.UnimplementedMiddlewareServer
}

func Register(server grpc.ServiceRegistrar) {
	readstate.RegisterMiddlewareServer(server, &Server{})
}

func RegisterGateway(mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) error {
	return nil
}
