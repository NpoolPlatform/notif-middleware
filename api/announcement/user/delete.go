package user

import (
	"context"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/notif/mw/v1/announcement/user"
	handler1 "github.com/NpoolPlatform/notif-middleware/pkg/mw/announcement/handler"
	amtuser1 "github.com/NpoolPlatform/notif-middleware/pkg/mw/announcement/user"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) DeleteAnnouncementUser(ctx context.Context, in *npool.DeleteAnnouncementUserRequest) (*npool.DeleteAnnouncementUserResponse, error) {
	req := in.GetInfo()
	if req == nil {
		logger.Sugar().Errorw(
			"DeleteAnnouncementUser",
			"In", in,
		)
		return &npool.DeleteAnnouncementUserResponse{}, status.Error(codes.InvalidArgument, "Info is empty")
	}
	handler, err := amtuser1.NewHandler(
		ctx,
		handler1.WithID(req.ID, true),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"DeleteAnnouncementUser",
			"In", in,
			"Error", err,
		)
		return &npool.DeleteAnnouncementUserResponse{}, status.Error(codes.Aborted, err.Error())
	}
	info, err := handler.DeleteAnnouncementUser(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"DeleteAnnouncementUser",
			"In", in,
			"Error", err,
		)
		return &npool.DeleteAnnouncementUserResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.DeleteAnnouncementUserResponse{
		Info: info,
	}, nil
}
