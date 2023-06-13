package readstate

import (
	readamt "github.com/NpoolPlatform/message/npool/notif/mw/v1/announcement/read"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
)

type Server struct {
	readamt.UnimplementedMiddlewareServer
}

func Register(server grpc.ServiceRegistrar) {
	readamt.RegisterMiddlewareServer(server, &Server{})
}

func RegisterGateway(mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) error {
	return nil
}
