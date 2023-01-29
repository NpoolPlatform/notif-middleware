//nolint:nolintlint,dupl
package notif

import (
	"context"
	"fmt"
	"github.com/NpoolPlatform/go-service-framework/pkg/logger"

	mgrpb "github.com/NpoolPlatform/message/npool/notif/mgr/v1/notif"
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

func validate(in *mgrpb.NotifReq) error {
	if in.ID != nil {
		if _, err := uuid.Parse(in.GetID()); err != nil {
			logger.Sugar().Errorw("validate", "ID", in.GetID(), "error", err)
			return err
		}
	}
	if _, err := uuid.Parse(in.GetAppID()); err != nil {
		logger.Sugar().Errorw("validate", "AppID", in.GetAppID(), "error", err)
		return err
	}
	if _, err := uuid.Parse(in.GetUserID()); err != nil {
		logger.Sugar().Errorw("validate", "UserID", in.GetUserID(), "error", err)
		return err
	}
	if _, err := uuid.Parse(in.GetLangID()); err != nil {
		logger.Sugar().Errorw("validate", "LangID", in.GetLangID(), "error", err)
		return err
	}

	switch in.GetEventType() {
	case mgrpb.EventType_KycReviewApproved:
	case mgrpb.EventType_KycReviewRejected:
	case mgrpb.EventType_WithdrawReviewApproved:
	case mgrpb.EventType_WithdrawReviewRejected:
	case mgrpb.EventType_WithdrawAddressReviewApproved:
	case mgrpb.EventType_WithdrawAddressReviewRejected:
	default:
		return fmt.Errorf("EventType is invalid")
	}

	if in.GetTitle() == "" {
		logger.Sugar().Errorw("validate", "Title", in.GetTitle())
		return fmt.Errorf("title is invalid")
	}
	if in.GetContent() == "" {
		logger.Sugar().Errorw("validate", "Content", in.GetContent())
		return fmt.Errorf("title is invalid")
	}
	if len(in.GetChannels()) == 0 {
		logger.Sugar().Errorw("validate", "Channels", in.GetChannels())
		return fmt.Errorf("channels is invalid")
	}
	return nil
}

func (s *Server) CreateNotif(ctx context.Context, in *npool.CreateNotifRequest) (*npool.CreateNotifResponse, error) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "CreateNotif")
	defer span.End()

	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()

	span = commontracer.TraceInvoker(span, "notif", "crud", "Create")

	if err = validate(in.GetInfo()); err != nil {
		logger.Sugar().Errorw("CreateNotif", "error", err)
		return &npool.CreateNotifResponse{}, status.Error(codes.Internal, err.Error())
	}

	info, err := notif1.CreateNotif(ctx, in.GetInfo())
	if err != nil {
		logger.Sugar().Errorw("CreateNotif", "error", err)
		return &npool.CreateNotifResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.CreateNotifResponse{
		Info: info,
	}, nil
}

func (s *Server) GetNotif(ctx context.Context, in *npool.GetNotifRequest) (*npool.GetNotifResponse, error) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "GetNotif")
	defer span.End()

	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()

	span = commontracer.TraceID(span, in.GetID())

	_, err = uuid.Parse(in.GetID())
	if err != nil {
		logger.Sugar().Errorw("GetNotif", "ID", in.GetID(), "error", err)
		return &npool.GetNotifResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	span = commontracer.TraceInvoker(span, "notif", "crud", "Row")

	info, err := notif1.GetNotif(ctx, in.GetID())
	if err != nil {
		logger.Sugar().Errorw("GetNotif", "error", err)
		return &npool.GetNotifResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.GetNotifResponse{
		Info: info,
	}, nil
}

//nolint:gocyclo
func validateConds(in *mgrpb.Conds) error {
	if in.ID != nil {
		if _, err := uuid.Parse(in.GetID().GetValue()); err != nil {
			logger.Sugar().Errorw("validateConds", "ID", in.GetID().GetValue(), "error", err)
			return err
		}
	}
	if in.AppID != nil {
		if _, err := uuid.Parse(in.GetAppID().GetValue()); err != nil {
			logger.Sugar().Errorw("validateConds", "AppID", in.GetAppID().GetValue(), "error", err)
			return err
		}
	}
	if in.UserID != nil {
		if _, err := uuid.Parse(in.GetUserID().GetValue()); err != nil {
			logger.Sugar().Errorw("validateConds", "UserID", in.GetUserID().GetValue(), "error", err)
			return err
		}
	}
	if in.LangID != nil {
		if _, err := uuid.Parse(in.GetLangID().GetValue()); err != nil {
			logger.Sugar().Errorw("validateConds", "LangID", in.GetLangID().GetValue(), "error", err)
			return err
		}
	}
	if in.EventType != nil {
		switch in.GetEventType().GetValue() {
		case uint32(mgrpb.EventType_KycReviewApproved):
		case uint32(mgrpb.EventType_KycReviewRejected):
		case uint32(mgrpb.EventType_WithdrawReviewApproved):
		case uint32(mgrpb.EventType_WithdrawReviewRejected):
		case uint32(mgrpb.EventType_WithdrawAddressReviewApproved):
		case uint32(mgrpb.EventType_WithdrawAddressReviewRejected):
		default:
			return fmt.Errorf("EventType is invalid")
		}
	}
	if in.Channels != nil {
		if len(in.GetChannels().GetValue()) == 0 {
			return fmt.Errorf("channels is invalid")
		}
	}
	return nil
}

func (s *Server) GetNotifOnly(ctx context.Context, in *npool.GetNotifOnlyRequest) (*npool.GetNotifOnlyResponse, error) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "GetNotifOnly")
	defer span.End()

	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()

	if err := validateConds(in.GetConds()); err != nil {
		return &npool.GetNotifOnlyResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	span = commontracer.TraceInvoker(span, "notif", "crud", "RowOnly")

	info, err := notif1.GetNotifOnly(ctx, in.GetConds())
	if err != nil {
		logger.Sugar().Errorw("GetNotifOnly", "error", err)
		return &npool.GetNotifOnlyResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.GetNotifOnlyResponse{
		Info: info,
	}, nil
}

func (s *Server) GetNotifs(ctx context.Context, in *npool.GetNotifsRequest) (*npool.GetNotifsResponse, error) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "GetNotifs")
	defer span.End()

	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()

	span = commontracer.TraceOffsetLimit(span, int(in.GetOffset()), int(in.GetLimit()))

	if err := validateConds(in.GetConds()); err != nil {
		return &npool.GetNotifsResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	span = commontracer.TraceInvoker(span, "notif", "crud", "Rows")

	rows, total, err := notif1.GetNotifs(ctx, in.GetConds(), in.GetOffset(), in.GetLimit())
	if err != nil {
		logger.Sugar().Errorw("GetNotifs", "error", err)
		return &npool.GetNotifsResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.GetNotifsResponse{
		Infos: rows,
		Total: total,
	}, nil
}
