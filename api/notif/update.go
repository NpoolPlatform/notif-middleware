//nolint:nolintlint,dupl
package notif

import (
	"context"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"

	basetypes "github.com/NpoolPlatform/message/npool/basetypes/v1"
	npool "github.com/NpoolPlatform/message/npool/notif/mw/v1/notif"

	constant "github.com/NpoolPlatform/notif-middleware/pkg/message/const"
	commontracer "github.com/NpoolPlatform/notif-middleware/pkg/tracer"

	"github.com/google/uuid"
	"go.opentelemetry.io/otel"
	scodes "go.opentelemetry.io/otel/codes"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	notif1 "github.com/NpoolPlatform/notif-middleware/pkg/notif"
)

//nolint:funlen,gocyclo
func (s *Server) UpdateNotif(ctx context.Context, in *npool.UpdateNotifRequest) (*npool.UpdateNotifResponse, error) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "UpdateNotif")
	defer span.End()

	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()

	span = commontracer.TraceInvoker(span, "notif", "crud", "Update")

	if in.Info == nil {
		logger.Sugar().Errorw("validate", "in", in, "error", "invalid info")
		return &npool.UpdateNotifResponse{}, status.Error(codes.Internal, err.Error())
	}

	if _, err := uuid.Parse(in.GetInfo().GetID()); err != nil {
		logger.Sugar().Errorw("validate", "ID", in.GetInfo().GetID(), "error", err)
		return &npool.UpdateNotifResponse{}, status.Error(codes.Internal, err.Error())
	}

	if in.Info.AppID != nil {
		if _, err := uuid.Parse(in.GetInfo().GetAppID()); err != nil {
			logger.Sugar().Errorw("validate", "AppID", in.GetInfo().GetAppID(), "error", err)
			return &npool.UpdateNotifResponse{}, status.Error(codes.Internal, err.Error())
		}
	}
	if in.Info.UserID != nil {
		if _, err := uuid.Parse(in.GetInfo().GetUserID()); err != nil {
			logger.Sugar().Errorw("validate", "UserID", in.GetInfo().GetUserID(), "error", err)
			return &npool.UpdateNotifResponse{}, status.Error(codes.Internal, err.Error())
		}
	}
	if in.Info.LangID != nil {
		if _, err := uuid.Parse(in.GetInfo().GetLangID()); err != nil {
			logger.Sugar().Errorw("validate", "LangID", in.GetInfo().GetLangID(), "error", err)
			return &npool.UpdateNotifResponse{}, status.Error(codes.Internal, err.Error())
		}
	}

	if in.Info.EventType != nil {
		switch in.GetInfo().GetEventType() {
		case basetypes.UsedFor_WithdrawalRequest:
		case basetypes.UsedFor_WithdrawalCompleted:
		case basetypes.UsedFor_DepositReceived:
		case basetypes.UsedFor_KYCApproved:
		case basetypes.UsedFor_KYCRejected:
		case basetypes.UsedFor_Announcement:
		default:
			logger.Sugar().Errorw("validate", "EventType", in.GetInfo().GetEventType())
			return &npool.UpdateNotifResponse{}, status.Error(codes.InvalidArgument, "EventType is invalid")
		}
	}

	if in.GetInfo().GetTitle() == "" && in.Info.Title != nil {
		logger.Sugar().Errorw("validate", "Title", in.GetInfo().GetTitle())
		return &npool.UpdateNotifResponse{}, status.Error(codes.InvalidArgument, "title is invalid")
	}
	if in.GetInfo().GetContent() == "" && in.Info.Content != nil {
		logger.Sugar().Errorw("validate", "Content", in.GetInfo().GetContent())
		return &npool.UpdateNotifResponse{}, status.Error(codes.InvalidArgument, "content is invalid")
	}

	info, err := notif1.UpdateNotif(ctx, in.GetInfo())
	if err != nil {
		logger.Sugar().Errorw("UpdateNotif", "error", err)
		return &npool.UpdateNotifResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.UpdateNotifResponse{
		Info: info,
	}, nil
}

//nolint:funlen,gocyclo
func (s *Server) UpdateNotifs(ctx context.Context, in *npool.UpdateNotifsRequest) (*npool.UpdateNotifsResponse, error) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "UpdatesNotif")
	defer span.End()

	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()

	span = commontracer.TraceInvoker(span, "notif", "crud", "Updates")

	if len(in.GetIDs()) == 0 {
		return &npool.UpdateNotifsResponse{}, status.Error(codes.Internal, "IDs is invalid")
	}

	for _, id := range in.GetIDs() {
		if _, err := uuid.Parse(id); err != nil {
			logger.Sugar().Errorw("validate", "ID", id, "error", err)
			return &npool.UpdateNotifsResponse{}, status.Error(codes.Internal, err.Error())
		}
	}

	infos, _, err := notif1.UpdateNotifs(ctx, in.GetIDs(), in.Notified)
	if err != nil {
		logger.Sugar().Errorw("UpdatesNotif", "error", err)
		return &npool.UpdateNotifsResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.UpdateNotifsResponse{
		Infos: infos,
	}, nil
}
