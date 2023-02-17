//nolint:nolintlint,dupl
package notif

import (
	"context"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"

	basetypes "github.com/NpoolPlatform/message/npool/basetypes/v1"
	npool "github.com/NpoolPlatform/message/npool/notif/mw/v1/notif"

	constant "github.com/NpoolPlatform/notif-middleware/pkg/message/const"
	commontracer "github.com/NpoolPlatform/notif-middleware/pkg/tracer"

	"go.opentelemetry.io/otel"
	scodes "go.opentelemetry.io/otel/codes"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	notif1 "github.com/NpoolPlatform/notif-middleware/pkg/notif"

	"github.com/google/uuid"
)

//nolint:funlen,gocyclo
func (s *Server) GenerateNotifs(ctx context.Context, in *npool.GenerateNotifsRequest) (*npool.GenerateNotifsResponse, error) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "GenerateNotifs")
	defer span.End()

	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()

	span = commontracer.TraceInvoker(span, "notif", "crud", "Update")

	if _, err := uuid.Parse(in.GetAppID()); err != nil {
		logger.Sugar().Errorw("GenerateNotifs", "AppID", in.GetAppID(), "error", err)
		return &npool.GenerateNotifsResponse{}, status.Error(codes.Internal, err.Error())
	}
	if _, err := uuid.Parse(in.GetUserID()); err != nil {
		logger.Sugar().Errorw("GenerateNotifs", "UserID", in.GetUserID(), "error", err)
		return &npool.GenerateNotifsResponse{}, status.Error(codes.Internal, err.Error())
	}

	switch in.GetEventType() {
	case basetypes.UsedFor_WithdrawalRequest:
	case basetypes.UsedFor_WithdrawalCompleted:
	case basetypes.UsedFor_DepositReceived:
	case basetypes.UsedFor_KYCApproved:
	case basetypes.UsedFor_KYCRejected:
	default:
		logger.Sugar().Errorw("GenerateNotifs", "EventType", in.GetEventType())
		return &npool.GenerateNotifsResponse{}, status.Error(codes.InvalidArgument, "EventType is invalid")
	}

	infos, err := notif1.GenerateNotifs(
		ctx,
		in.GetAppID(),
		in.GetUserID(),
		in.GetEventType(),
		in.Vars,
		in.Extra,
	)
	if err != nil {
		logger.Sugar().Errorw("GenerateNotifs", "error", err)
		return &npool.GenerateNotifsResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.GenerateNotifsResponse{
		Infos: infos,
	}, nil
}
