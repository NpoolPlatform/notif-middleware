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

func (s *Server) CreateTx(
	ctx context.Context,
	in *npool.CreateTxRequest,
) (
	*npool.CreateTxResponse,
	error,
) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "CreateTx")
	defer span.End()

	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()

	row, err := mgrcli.CreateTx(ctx, in.GetInfo())
	if err != nil {
		logger.Sugar().Errorw("CreateTx", "error", err)
		return &npool.CreateTxResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.CreateTxResponse{
		Info: row,
	}, nil
}
