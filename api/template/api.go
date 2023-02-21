package template

import (
	template "github.com/NpoolPlatform/message/npool/notif/mw/v1/template"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
)

type Server struct {
	template.UnimplementedMiddlewareServer
}

func Register(server grpc.ServiceRegistrar) {
	template.RegisterMiddlewareServer(server, &Server{})
}

func RegisterGateway(mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) error {
	return nil
}
