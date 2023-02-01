package sendstate

import (
	sendstate "github.com/NpoolPlatform/message/npool/notif/mw/v1/announcement/sendstate"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
)

type Server struct {
	sendstate.UnimplementedMiddlewareServer
}

func Register(server grpc.ServiceRegistrar) {
	sendstate.RegisterMiddlewareServer(server, &Server{})
}

func RegisterGateway(mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) error {
	return nil
}
