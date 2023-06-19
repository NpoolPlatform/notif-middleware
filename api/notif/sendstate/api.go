package sendstate

import (
	"github.com/NpoolPlatform/message/npool/notif/mw/v1/notif/sendstate"
	"google.golang.org/grpc"
)

type Server struct {
	sendstate.UnimplementedMiddlewareServer
}

func Register(server grpc.ServiceRegistrar) {
	sendstate.RegisterMiddlewareServer(server, &Server{})
}
