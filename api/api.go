package api

import (
	"context"

	"github.com/NpoolPlatform/notif-middleware/api/announcement"
	"github.com/NpoolPlatform/notif-middleware/api/announcement/sendstate"

	"github.com/NpoolPlatform/notif-middleware/api/announcement/readstate"
	"github.com/NpoolPlatform/notif-middleware/api/notif"

	v1 "github.com/NpoolPlatform/message/npool/notif/mw/v1"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
)

type Server struct {
	v1.UnimplementedMiddlewareServer
}

func Register(server grpc.ServiceRegistrar) {
	v1.RegisterMiddlewareServer(server, &Server{})
	notif.Register(server)
	readstate.Register(server)
	sendstate.Register(server)
	announcement.Register(server)
}

func RegisterGateway(mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) error {
	if err := v1.RegisterMiddlewareHandlerFromEndpoint(context.Background(), mux, endpoint, opts); err != nil {
		return err
	}
	return nil
}
