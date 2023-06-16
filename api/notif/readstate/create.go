//nolint:nolintlint,dupl
package readstate

import (
	"context"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/notif/mw/v1/notif/readstate"
	readstate1 "github.com/NpoolPlatform/notif-middleware/pkg/mw/notif/readstate"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) CreateReadState(ctx context.Context, in *npool.CreateReadStateRequest) (*npool.CreateReadStateResponse, error) {
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
			"CreateReadState",
			"In", in,
			"Error", err,
		)
		return &npool.CreateReadStateResponse{}, status.Error(codes.Aborted, err.Error())
	}
	info, err := handler.CreateReadState(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"CreateReadState",
			"In", in,
			"Error", err,
		)
		return &npool.CreateReadStateResponse{}, status.Error(codes.Aborted, err.Error())
	}
	return &npool.CreateReadStateResponse{
		Info: info,
	}, nil
}

func (s *Server) CreateReadStates(ctx context.Context, in *npool.CreateReadStatesRequest) (*npool.CreateReadStatesResponse, error) {
	handler, err := readstate1.NewHandler(
		ctx,
		readstate1.WithReqs(in.GetInfos()),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"CreateReadStates",
			"In", in,
			"Error", err,
		)
		return &npool.CreateReadStatesResponse{}, status.Error(codes.Aborted, err.Error())
	}
	infos, err := handler.CreateReadStates(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"CreateReadStates",
			"In", in,
			"Error", err,
		)
		return &npool.CreateReadStatesResponse{}, status.Error(codes.Aborted, err.Error())
	}
	return &npool.CreateReadStatesResponse{
		Infos: infos,
	}, nil
}
