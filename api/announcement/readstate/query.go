package readstate

import (
	"context"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"

	npool "github.com/NpoolPlatform/message/npool/notif/mw/v1/announcement/readstate"
	handler1 "github.com/NpoolPlatform/notif-middleware/pkg/mw/announcement/handler"
	amtread1 "github.com/NpoolPlatform/notif-middleware/pkg/mw/announcement/readstate"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) GetReadStates(ctx context.Context, in *npool.GetReadStatesRequest) (*npool.GetReadStatesResponse, error) {
	handler, err := amtread1.NewHandler(
		ctx,
		amtread1.WithConds(in.GetConds()),
		handler1.WithOffset(in.Offset),
		handler1.WithLimit(in.Limit),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"GetReadStates",
			"In", in,
			"Error", err,
		)
		return &npool.GetReadStatesResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	infos, total, err := handler.GetReadStates(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"GetReadStates",
			"In", in,
			"Error", err,
		)
		return &npool.GetReadStatesResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.GetReadStatesResponse{
		Infos: infos,
		Total: total,
	}, nil
}

func (s *Server) GetReadState(ctx context.Context, in *npool.GetReadStateRequest) (*npool.GetReadStateResponse, error) {
	handler, err := amtread1.NewHandler(
		ctx,
		handler1.WithEntID(&in.EntID, true),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"GetReadState",
			"In", in,
			"error", err,
		)
		return &npool.GetReadStateResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	info, err := handler.GetReadState(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"GetReadState",
			"In", in,
			"error", err,
		)
		return &npool.GetReadStateResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.GetReadStateResponse{
		Info: info,
	}, nil
}
