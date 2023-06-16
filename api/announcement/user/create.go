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

func (s *Server) CreateAnnouncementUser(ctx context.Context, in *npool.CreateAnnouncementUserRequest) (*npool.CreateAnnouncementUserResponse, error) { // nolint
	req := in.GetInfo()
	handler1, err := amtuser1.NewHandler(
		ctx,
		handler.WithAppID(req.AppID),
		handler.WithUserID(req.AppID, req.UserID),
		handler.WithAnnouncementID(req.AppID, req.AnnouncementID),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"CreateAnnouncementUser",
			"Req", in,
			"Error", err,
		)
		return &npool.CreateAnnouncementUserResponse{}, status.Error(codes.InvalidArgument, err.Error())
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

// nolint
func (s *Server) CreateAnnouncementUsers(ctx context.Context, in *npool.CreateAnnouncementUsersRequest) (*npool.CreateAnnouncementUsersResponse, error) {
	handler1, err := amtuser1.NewHandler(
		ctx,
		amtuser1.WithReqs(in.GetInfos()),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"CreateAnnouncementUsers",
			"Req", in,
			"Error", err,
		)
		return &npool.CreateAnnouncementUsersResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	info, err := handler1.CreateAnnouncementUsers(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"CreateAnnouncementUsers",
			"Req", in,
			"Error", err,
		)
		return &npool.CreateAnnouncementUsersResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.CreateAnnouncementUsersResponse{
		Infos: info,
	}, nil
}
