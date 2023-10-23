package readstate

import (
	"context"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/notif/mw/v1/announcement/readstate"
	handler1 "github.com/NpoolPlatform/notif-middleware/pkg/mw/announcement/handler"
	amtread1 "github.com/NpoolPlatform/notif-middleware/pkg/mw/announcement/readstate"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) CreateReadState(ctx context.Context, in *npool.CreateReadStateRequest) (*npool.CreateReadStateResponse, error) {
	req := in.GetInfo()
	handler, err := amtread1.NewHandler(
		ctx,
		handler1.WithAppID(req.AppID, true),
		handler1.WithUserID(req.UserID, true),
		handler1.WithAnnouncementID(req.AppID, req.AnnouncementID, true),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"CreateReadState",
			"Req", in,
			"Error", err,
		)
		return &npool.CreateReadStateResponse{}, status.Error(codes.Aborted, err.Error())
	}

	info, err := handler.CreateReadState(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"CreateReadState",
			"Req", in,
			"Error", err,
		)
		return &npool.CreateReadStateResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.CreateReadStateResponse{
		Info: info,
	}, nil
}
