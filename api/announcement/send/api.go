package send

import (
	sendamt "github.com/NpoolPlatform/message/npool/notif/mw/v1/announcement/send"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
)

type Server struct {
	sendamt.UnimplementedMiddlewareServer
}

func Register(server grpc.ServiceRegistrar) {
	sendamt.RegisterMiddlewareServer(server, &Server{})
}

func RegisterGateway(mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) error {
	return nil
}
