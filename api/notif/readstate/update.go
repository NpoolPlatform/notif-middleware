package readstate

import (
	"context"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/notif/mw/v1/notif/readstate"
	readstate1 "github.com/NpoolPlatform/notif-middleware/pkg/mw/notif/readstate"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) UpdateReadState(ctx context.Context, in *npool.UpdateReadStateRequest) (*npool.UpdateReadStateResponse, error) {
	req := in.GetInfo()
	handler, err := readstate1.NewHandler(
		ctx,
		readstate1.WithID(req.ID),
		readstate1.WithAppID(req.AppID),
		readstate1.WithUserID(req.UserID),
		readstate1.WithNotifID(req.NotifID),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"UpdateReadState",
			"In", in,
			"Error", err,
		)
		return &npool.UpdateReadStateResponse{}, status.Error(codes.Aborted, err.Error())
	}
	info, err := handler.UpdateReadState(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"UpdateReadState",
			"In", in,
			"Error", err,
		)
		return &npool.UpdateReadStateResponse{}, status.Error(codes.Aborted, err.Error())
	}
	return &npool.UpdateReadStateResponse{
		Info: info,
	}, nil
}
