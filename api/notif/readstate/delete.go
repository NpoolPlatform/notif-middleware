package readstate

import (
	"context"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/notif/mw/v1/notif/readstate"
	readstate1 "github.com/NpoolPlatform/notif-middleware/pkg/mw/notif/readstate"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) DeleteReadState(ctx context.Context, in *npool.DeleteReadStateRequest) (*npool.DeleteReadStateResponse, error) {
	req := in.GetInfo()
	handler, err := readstate1.NewHandler(
		ctx,
		readstate1.WithID(req.ID),
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
