package readstate

import (
	"context"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/notif/mw/v1/announcement/readstate"
	"github.com/NpoolPlatform/notif-middleware/pkg/mw/announcement/handler"
	announcement1 "github.com/NpoolPlatform/notif-middleware/pkg/mw/announcement/readstate"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) CreateReadState(ctx context.Context, in *npool.CreateReadStateRequest) (*npool.CreateReadStateResponse, error) {
	req := in.GetInfo()
	handler, err := announcement1.NewHandler(
		ctx,
		handler.WithAppID(req.AppID),
		handler.WithUserID(req.AppID, req.UserID),
		handler.WithAnnouncementID(req.AppID, req.AnnouncementID),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"CreateReadState",
			"Req", in,
			"Error", err,
		)
		return &npool.CreateReadStateResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	info, err := handler.CreateReadState(ctx)
	if err != nil {
		return &npool.CreateReadStateResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.CreateReadStateResponse{
		Info: info,
	}, nil
}
