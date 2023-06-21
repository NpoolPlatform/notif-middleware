package sendstate

import (
	"context"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/notif/mw/v1/announcement/sendstate"
	handler1 "github.com/NpoolPlatform/notif-middleware/pkg/mw/announcement/handler"
	amtsend1 "github.com/NpoolPlatform/notif-middleware/pkg/mw/announcement/sendstate"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) DeleteSendState(ctx context.Context, in *npool.DeleteSendStateRequest) (*npool.DeleteSendStateResponse, error) {
	id := in.GetInfo().GetID()
	handler, err := amtsend1.NewHandler(
		ctx,
		handler1.WithID(&id),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"DeleteSendState",
			"In", in,
			"Error", err,
		)
		return &npool.DeleteSendStateResponse{}, status.Error(codes.Aborted, err.Error())
	}

	info, err := handler.DeleteSendState(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"DeleteSendState",
			"In", in,
			"Error", err,
		)
		return &npool.DeleteSendStateResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.DeleteSendStateResponse{
		Info: info,
	}, nil
}
