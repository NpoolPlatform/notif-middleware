package sendstate

import (
	"context"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/notif/mw/v1/announcement/sendstate"
	"github.com/NpoolPlatform/notif-middleware/pkg/mw/announcement/handler"
	sendamt "github.com/NpoolPlatform/notif-middleware/pkg/mw/announcement/sendstate"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) CreateSendState(ctx context.Context, in *npool.CreateSendStateRequest) (*npool.CreateSendStateResponse, error) {
	req := in.GetInfo()
	handler, err := sendamt.NewHandler(
		ctx,
		handler.WithAppID(req.AppID),
		handler.WithUserID(req.AppID, req.UserID),
		handler.WithAnnouncementID(req.AppID, req.AnnouncementID),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"CreateSendState",
			"Req", in,
			"Error", err,
		)
		return &npool.CreateSendStateResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	info, err := handler.CreateSendState(ctx)
	if err != nil {
		return &npool.CreateSendStateResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.CreateSendStateResponse{
		Info: info,
	}, nil
}
