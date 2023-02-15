package tx

import (
	tx "github.com/NpoolPlatform/message/npool/notif/mw/v1/notif/tx"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
)

type Server struct {
	tx.UnimplementedMiddlewareServer
}

func Register(server grpc.ServiceRegistrar) {
	tx.RegisterMiddlewareServer(server, &Server{})
}

func RegisterGateway(mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) error {
	return nil
}
