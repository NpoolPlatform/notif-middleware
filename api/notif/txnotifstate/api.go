package txnotifstate

import (
	txnotifstate "github.com/NpoolPlatform/message/npool/notif/mw/v1/notif/txnotifstate"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
)

type Server struct {
	txnotifstate.UnimplementedMiddlewareServer
}

func Register(server grpc.ServiceRegistrar) {
	txnotifstate.RegisterMiddlewareServer(server, &Server{})
}

func RegisterGateway(mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) error {
	return nil
}
