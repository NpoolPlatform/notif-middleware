package tx

import (
	"context"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/notif/mw/v1/notif/tx"
	tx1 "github.com/NpoolPlatform/notif-middleware/pkg/mw/notif/tx"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) DeleteTx(ctx context.Context, in *npool.DeleteTxRequest) (*npool.DeleteTxResponse, error) {
	id := in.GetInfo().GetID()
	handler, err := tx1.NewHandler(
		ctx,
		tx1.WithID(&id),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"DeleteTx",
			"In", in,
			"Error", err,
		)
		return &npool.DeleteTxResponse{}, status.Error(codes.Aborted, err.Error())
	}
	info, err := handler.DeleteTx(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"DeleteTx",
			"In", in,
			"Error", err,
		)
		return &npool.DeleteTxResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.DeleteTxResponse{
		Info: info,
	}, nil
}
