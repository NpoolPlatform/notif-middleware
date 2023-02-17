package sms

import (
	"github.com/NpoolPlatform/message/npool/notif/mw/v1/template/sms"

	"google.golang.org/grpc"
)

type Server struct {
	sms.UnimplementedMiddlewareServer
}

func Register(server grpc.ServiceRegistrar) {
	sms.RegisterMiddlewareServer(server, &Server{})
}
