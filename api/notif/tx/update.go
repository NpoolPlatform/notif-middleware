//nolint:nolintlint,dupl
package tx

import (
	"context"

	mgrcli "github.com/NpoolPlatform/notif-manager/pkg/client/notif/tx"

	constant "github.com/NpoolPlatform/notif-middleware/pkg/message/const"

	"go.opentelemetry.io/otel"
	scodes "go.opentelemetry.io/otel/codes"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/notif/mw/v1/notif/tx"
)

func (s *Server) UpdateTx(
	ctx context.Context,
	in *npool.UpdateTxRequest,
) (
	*npool.UpdateTxResponse,
	error,
) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "UpdateTx")
	defer span.End()

	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()

	row, err := mgrcli.UpdateTx(ctx, in.GetInfo())
	if err != nil {
		logger.Sugar().Errorw("UpdateTx", "error", err)
		return &npool.UpdateTxResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.UpdateTxResponse{
		Info: row,
	}, nil
}
