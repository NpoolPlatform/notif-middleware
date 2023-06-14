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

func (s *Server) DeleteReadState(ctx context.Context, in *npool.DeleteReadStateRequest) (*npool.DeleteReadStateResponse, error) {
	id := in.GetInfo().GetID()
	handler, err := announcement1.NewHandler(
		ctx,
		handler.WithID(&id),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"DeleteReadState",
			"In", in,
			"Error", err,
		)
		return &npool.DeleteReadStateResponse{}, status.Error(codes.Aborted, err.Error())
	}
	info, err := handler.DeleteReadState(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"DeleteReadState",
			"In", in,
			"Error", err,
		)
		return &npool.DeleteReadStateResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.DeleteReadStateResponse{
		Info: info,
	}, nil
}
