package announcement

import (
	announcement "github.com/NpoolPlatform/message/npool/notif/mw/v1/announcement"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
)

type Server struct {
	announcement.UnimplementedMiddlewareServer
}

func Register(server grpc.ServiceRegistrar) {
	announcement.RegisterMiddlewareServer(server, &Server{})
}

func RegisterGateway(mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) error {
	return nil
}
