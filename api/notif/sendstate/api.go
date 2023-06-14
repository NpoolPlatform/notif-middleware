package sendstate

import (
	"github.com/NpoolPlatform/message/npool/notif/mw/v1/notif/readstate"
	"google.golang.org/grpc"
)

type Server struct {
	readstate.UnimplementedMiddlewareServer
}

func Register(server grpc.ServiceRegistrar) {
	readstate.RegisterMiddlewareServer(server, &Server{})
}
