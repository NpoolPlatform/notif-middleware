//nolint:nolintlint,dupl
package notifchannel

import (
	"context"

	mgrcli "github.com/NpoolPlatform/notif-manager/pkg/client/notif/notifchannel"

	constant "github.com/NpoolPlatform/notif-middleware/pkg/message/const"

	"go.opentelemetry.io/otel"
	scodes "go.opentelemetry.io/otel/codes"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/notif/mw/v1/notif/notifchannel"
)

func (s *Server) GetNotifChannelOnly(
	ctx context.Context,
	in *npool.GetNotifChannelOnlyRequest,
) (
	*npool.GetNotifChannelOnlyResponse,
	error,
) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "GetUsers")
	defer span.End()

	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()

	row, err := mgrcli.GetNotifChannelOnly(ctx, in.GetConds())
	if err != nil {
		logger.Sugar().Errorw("GetUsers", "error", err)
		return &npool.GetNotifChannelOnlyResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.GetNotifChannelOnlyResponse{
		Info: row,
	}, nil
}

func (s *Server) GetNotifChannels(
	ctx context.Context,
	in *npool.GetNotifChannelsRequest,
) (
	*npool.GetNotifChannelsResponse,
	error,
) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "GetUsers")
	defer span.End()

	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()

	rows, total, err := mgrcli.GetNotifChannels(ctx, in.GetConds(), in.GetOffset(), in.GetLimit())
	if err != nil {
		logger.Sugar().Errorw("GetUsers", "error", err)
		return &npool.GetNotifChannelsResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.GetNotifChannelsResponse{
		Infos: rows,
		Total: total,
	}, nil
}
