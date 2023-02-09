//nolint:nolintlint,dupl
package announcement

import (
	"context"

	"github.com/google/uuid"

	announcement1 "github.com/NpoolPlatform/notif-middleware/pkg/announcement"

	commontracer "github.com/NpoolPlatform/notif-middleware/pkg/tracer"

	constant "github.com/NpoolPlatform/notif-middleware/pkg/message/const"

	"go.opentelemetry.io/otel"
	scodes "go.opentelemetry.io/otel/codes"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/notif/mw/v1/announcement"
)

func (s *Server) GetAnnouncementStates(
	ctx context.Context,
	in *npool.GetAnnouncementStatesRequest,
) (
	*npool.GetAnnouncementStatesResponse,
	error,
) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "GetAnnouncements")
	defer span.End()

	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()

	span = commontracer.TraceOffsetLimit(span, int(in.GetOffset()), int(in.GetLimit()))

	span = commontracer.TraceInvoker(span, "announcement/announcement", "crud", "Rows")

	if in.Conds == nil {
		logger.Sugar().Errorw("validateConds", "Conds", in.GetConds(), "error", err)
		return &npool.GetAnnouncementStatesResponse{}, status.Error(codes.InvalidArgument, "Conds is empty")
	}
	if in.GetConds().AppID == nil {
		logger.Sugar().Errorw("validateConds", "AppID", in.GetConds().GetAppID(), "error", err)
		return &npool.GetAnnouncementStatesResponse{}, status.Error(codes.InvalidArgument, "Conds AppID is empty")
	}
	if in.GetConds().UserID == nil {
		logger.Sugar().Errorw("validateConds", "UserID", in.GetConds().GetUserID(), "error", err)
		return &npool.GetAnnouncementStatesResponse{}, status.Error(codes.InvalidArgument, "Conds UserID is empty")
	}
	if in.GetConds().LangID == nil {
		logger.Sugar().Errorw("validateConds", "LangID", in.GetConds().GetLangID(), "error", err)
		return &npool.GetAnnouncementStatesResponse{}, status.Error(codes.InvalidArgument, "Conds LangID is empty")
	}

	if _, err := uuid.Parse(in.GetConds().GetAppID().GetValue()); err != nil {
		logger.Sugar().Errorw("validateConds", "AppID", in.GetConds().GetAppID().GetValue(), "error", err)
		return &npool.GetAnnouncementStatesResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	if _, err := uuid.Parse(in.GetConds().GetUserID().GetValue()); err != nil {
		logger.Sugar().Errorw("validateConds", "UserID", in.GetConds().GetAppID().GetValue(), "error", err)
		return &npool.GetAnnouncementStatesResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	if _, err := uuid.Parse(in.GetConds().GetLangID().GetValue()); err != nil {
		logger.Sugar().Errorw("validateConds", "UserID", in.GetConds().GetAppID().GetValue(), "error", err)
		return &npool.GetAnnouncementStatesResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	rows, total, err := announcement1.GetAnnouncementStates(
		ctx,
		in.GetConds().GetAppID().GetValue(),
		in.GetConds().GetUserID().GetValue(),
		in.GetConds().GetLangID().GetValue(),
		in.GetOffset(),
		in.GetLimit(),
	)
	if err != nil {
		logger.Sugar().Errorw("GetAnnouncements", "error", err)
		return &npool.GetAnnouncementStatesResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.GetAnnouncementStatesResponse{
		Infos: rows,
		Total: total,
	}, nil
}

func (s *Server) GetAnnouncements(ctx context.Context, in *npool.GetAnnouncementsRequest) (*npool.GetAnnouncementsResponse, error) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "GetAnnouncements")
	defer span.End()

	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()

	span = commontracer.TraceOffsetLimit(span, int(in.GetOffset()), int(in.GetLimit()))

	span = commontracer.TraceInvoker(span, "announcement/announcement", "crud", "Rows")

	rows, total, err := announcement1.GetAnnouncements(ctx, in.GetConds(), in.GetOffset(), in.GetLimit())
	if err != nil {
		logger.Sugar().Errorw("GetAnnouncements", "error", err)
		return &npool.GetAnnouncementsResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.GetAnnouncementsResponse{
		Infos: rows,
		Total: total,
	}, nil
}
