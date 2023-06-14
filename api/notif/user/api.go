package user

import (
	"github.com/NpoolPlatform/message/npool/notif/mw/v1/notif/user"
	"google.golang.org/grpc"
)

type Server struct {
	user.UnimplementedMiddlewareServer
}

func Register(server grpc.ServiceRegistrar) {
	user.RegisterMiddlewareServer(server, &Server{})
}
