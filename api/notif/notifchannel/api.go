package notifchannel

import (
	notifchannel "github.com/NpoolPlatform/message/npool/notif/mw/v1/notif/notifchannel"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
)

type Server struct {
	notifchannel.UnimplementedMiddlewareServer
}

func Register(server grpc.ServiceRegistrar) {
	notifchannel.RegisterMiddlewareServer(server, &Server{})
}

func RegisterGateway(mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) error {
	return nil
}
