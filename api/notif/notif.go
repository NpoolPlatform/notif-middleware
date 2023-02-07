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
	if in == nil {
		logger.Sugar().Errorw("validate", "in", in, "error", "invalid info")
		return fmt.Errorf("invalid info")
	}
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
	case mgrpb.EventType_WithdrawalRequest:
	case mgrpb.EventType_WithdrawalCompleted:
	case mgrpb.EventType_DepositReceived:
	case mgrpb.EventType_KYCApproved:
	case mgrpb.EventType_KYCRejected:
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
		case mgrpb.EventType_WithdrawalRequest:
		case mgrpb.EventType_WithdrawalCompleted:
		case mgrpb.EventType_DepositReceived:
		case mgrpb.EventType_KYCApproved:
		case mgrpb.EventType_KYCRejected:
		default:
			logger.Sugar().Errorw("validate", "error", "EventType is invalid")
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
	if len(in.GetInfo().GetChannels()) == 0 {
		logger.Sugar().Errorw("validate", "Channels", in.GetInfo().GetChannels())
		return &npool.UpdateNotifResponse{}, status.Error(codes.InvalidArgument, "channels is invalid")
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

	for _, id := range in.GetIDs() {
		if _, err := uuid.Parse(id); err != nil {
			logger.Sugar().Errorw("validate", "ID", id, "error", err)
			return &npool.UpdateNotifsResponse{}, status.Error(codes.Internal, err.Error())
		}
	}

	infos, _, err := notif1.UpdateNotifs(ctx, in.GetIDs(), in.EmailSend, in.AlreadyRead)
	if err != nil {
		logger.Sugar().Errorw("UpdatesNotif", "error", err)
		return &npool.UpdateNotifsResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.UpdateNotifsResponse{
		Infos: infos,
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
		case uint32(mgrpb.EventType_WithdrawalRequest):
		case uint32(mgrpb.EventType_WithdrawalCompleted):
		case uint32(mgrpb.EventType_DepositReceived):
		case uint32(mgrpb.EventType_KYCApproved):
		case uint32(mgrpb.EventType_KYCRejected):
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
