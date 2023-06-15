//nolint:nolintlint,dupl
package tx

import (
	"context"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/notif/mw/v1/notif/tx"
	tx1 "github.com/NpoolPlatform/notif-middleware/pkg/mw/notif/tx"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) CreateTx(ctx context.Context, in *npool.CreateTxRequest) (*npool.CreateTxResponse, error) {
	req := in.GetInfo()
	handler, err := tx1.NewHandler(
		ctx,
		tx1.WithTxID(req.TxID),
		tx1.WithTxType(req.TxType),
		tx1.WithNotifState(req.NotifState),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"CreateTx",
			"Req", in,
			"Error", err,
		)
		return &npool.CreateTxResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	info, err := handler.CreateTx(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"CreateTx",
			"Req", in,
			"Error", err,
		)
		return &npool.CreateTxResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.CreateTxResponse{
		Info: info,
	}, nil
}
