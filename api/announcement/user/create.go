package user

import (
	"context"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/notif/mw/v1/announcement/user"
	"github.com/NpoolPlatform/notif-middleware/pkg/mw/announcement/handler"
	amtuser1 "github.com/NpoolPlatform/notif-middleware/pkg/mw/announcement/user"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) CreateAnnouncementUser(ctx context.Context, in *npool.CreateAnnouncementUserRequest) (*npool.CreateAnnouncementUserResponse, error) {
	req := in.GetInfo()
	if req == nil {
		logger.Sugar().Errorw(
			"CreateAnnouncementUser",
			"In", in,
		)
		return &npool.CreateAnnouncementUserResponse{}, status.Error(codes.InvalidArgument, "Info is empty")
	}
	handler1, err := amtuser1.NewHandler(
		ctx,
		handler.WithEntID(req.EntID, false),
		handler.WithAppID(req.AppID, true),
		handler.WithUserID(req.UserID, true),
		handler.WithAnnouncementID(req.AppID, req.AnnouncementID, true),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"CreateAnnouncementUser",
			"Req", in,
			"Error", err,
		)
		return &npool.CreateAnnouncementUserResponse{}, status.Error(codes.Aborted, err.Error())
	}

	info, err := handler1.CreateAnnouncementUser(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"CreateAnnouncementUser",
			"Req", in,
			"Error", err,
		)
		return &npool.CreateAnnouncementUserResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.CreateAnnouncementUserResponse{
		Info: info,
	}, nil
}
