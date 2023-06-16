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

func (s *Server) GetAnnouncementUsers(ctx context.Context, in *npool.GetAnnouncementUsersRequest) (*npool.GetAnnouncementUsersResponse, error) {
	handler, err := amtuser1.NewHandler(
		ctx,
		amtuser1.WithConds(in.GetConds()),
		handler.WithOffset(in.Offset),
		handler.WithLimit(in.Limit),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"GetAnnouncementUsers",
			"In", in,
			"Error", err,
		)
		return &npool.GetAnnouncementUsersResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	infos, total, err := handler.GetAnnouncementUsers(ctx)
	if err != nil {
		return &npool.GetAnnouncementUsersResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.GetAnnouncementUsersResponse{
		Infos: infos,
		Total: total,
	}, nil
}

func (s *Server) GetAnnouncementUser(ctx context.Context, in *npool.GetAnnouncementUserRequest) (*npool.GetAnnouncementUserResponse, error) {
	handler, err := amtuser1.NewHandler(
		ctx,
		handler.WithID(&in.ID),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"GetAnnouncementUser",
			"In", in,
			"error", err,
		)
		return &npool.GetAnnouncementUserResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	info, err := handler.GetAnnouncementUser(ctx)
	if err != nil {
		return &npool.GetAnnouncementUserResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.GetAnnouncementUserResponse{
		Info: info,
	}, nil
}
