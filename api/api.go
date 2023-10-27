package api

import (
	"context"

	"github.com/NpoolPlatform/notif-middleware/api/announcement"
	"github.com/NpoolPlatform/notif-middleware/api/announcement/readstate"
	"github.com/NpoolPlatform/notif-middleware/api/announcement/sendstate"
	"github.com/NpoolPlatform/notif-middleware/api/announcement/user"
	"github.com/NpoolPlatform/notif-middleware/api/contact"
	"github.com/NpoolPlatform/notif-middleware/api/notif"
	"github.com/NpoolPlatform/notif-middleware/api/notif/channel"
	goodbenefit "github.com/NpoolPlatform/notif-middleware/api/notif/goodbenefit"
	notifuser "github.com/NpoolPlatform/notif-middleware/api/notif/user"
	"github.com/NpoolPlatform/notif-middleware/api/template"
	"github.com/NpoolPlatform/notif-middleware/api/template/email"
	"github.com/NpoolPlatform/notif-middleware/api/template/frontend"
	"github.com/NpoolPlatform/notif-middleware/api/template/sms"

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
	contact.Register(server)
	template.Register(server)
	email.Register(server)
	sms.Register(server)
	frontend.Register(server)
	readstate.Register(server)
	sendstate.Register(server)
	announcement.Register(server)
	user.Register(server)
	channel.Register(server)
	notifuser.Register(server)
	goodbenefit.Register(server)
}

func RegisterGateway(mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) error {
	if err := v1.RegisterMiddlewareHandlerFromEndpoint(context.Background(), mux, endpoint, opts); err != nil {
		return err
	}
	if err := notif.RegisterGateway(mux, endpoint, opts); err != nil {
		return err
	}
	if err := email.RegisterGateway(mux, endpoint, opts); err != nil {
		return err
	}
	if err := frontend.RegisterGateway(mux, endpoint, opts); err != nil {
		return err
	}
	if err := sms.RegisterGateway(mux, endpoint, opts); err != nil {
		return err
	}
	return nil
}
