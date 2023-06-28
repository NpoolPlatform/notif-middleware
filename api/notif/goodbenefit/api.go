package goodbenefit

import (
	goodbenefit "github.com/NpoolPlatform/message/npool/notif/mw/v1/notif/goodbenefit"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
)

type Server struct {
	goodbenefit.UnimplementedMiddlewareServer
}

func Register(server grpc.ServiceRegistrar) {
	goodbenefit.RegisterMiddlewareServer(server, &Server{})
}

func RegisterGateway(mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) error {
	return nil
}
