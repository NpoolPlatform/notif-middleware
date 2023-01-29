//nolint:nolintlint,dupl
package readstate

import (
	"context"

	mgrpb "github.com/NpoolPlatform/message/npool/notif/mgr/v1/announcement/readstate"

	readstate1 "github.com/NpoolPlatform/notif-middleware/pkg/announcement/readstate"
	commontracer "github.com/NpoolPlatform/notif-middleware/pkg/tracer"

	constant "github.com/NpoolPlatform/notif-middleware/pkg/message/const"

	"go.opentelemetry.io/otel"
	scodes "go.opentelemetry.io/otel/codes"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/notif/mw/v1/announcement/readstate"

	"github.com/google/uuid"
)

func (s *Server) GetReadState(ctx context.Context, in *npool.GetReadStateRequest) (*npool.GetReadStateResponse, error) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "GetReadState")
	defer span.End()

	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()

	span = commontracer.TraceID(span, in.GetAnnouncementID())

	_, err = uuid.Parse(in.GetAnnouncementID())
	if err != nil {
		logger.Sugar().Errorw("GetReadState", "AnnouncementID", in.GetAnnouncementID(), "error", err)
		return &npool.GetReadStateResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}
	_, err = uuid.Parse(in.GetUserID())
	if err != nil {
		logger.Sugar().Errorw("GetReadState", "UserID", in.GetUserID(), "error", err)
		return &npool.GetReadStateResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	span = commontracer.TraceInvoker(span, "announcement/readstate", "crud", "Row")

	info, err := readstate1.GetReadState(ctx, in.GetAnnouncementID(), in.GetUserID())
	if err != nil {
		logger.Sugar().Errorw("GetReadState", "error", err)
		return &npool.GetReadStateResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.GetReadStateResponse{
		Info: info,
	}, nil
}

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
	if in.AnnouncementID != nil {
		if _, err := uuid.Parse(in.GetAnnouncementID().GetValue()); err != nil {
			logger.Sugar().Errorw("validateConds", "AnnouncementID", in.GetAnnouncementID().GetValue(), "error", err)
			return err
		}
	}
	return nil
}

func (s *Server) GetReadStates(ctx context.Context, in *npool.GetReadStatesRequest) (*npool.GetReadStatesResponse, error) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "GetReadStates")
	defer span.End()

	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()

	span = commontracer.TraceOffsetLimit(span, int(in.GetOffset()), int(in.GetLimit()))

	if err = validateConds(in.GetConds()); err != nil {
		logger.Sugar().Errorw("validateConds", "Conds", in.GetConds(), "error", err)
		return &npool.GetReadStatesResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	span = commontracer.TraceInvoker(span, "announcement/readstate", "crud", "Rows")

	rows, total, err := readstate1.GetReadStates(ctx, in.GetConds(), in.GetOffset(), in.GetLimit())
	if err != nil {
		logger.Sugar().Errorw("GetReadStates", "error", err)
		return &npool.GetReadStatesResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.GetReadStatesResponse{
		Infos: rows,
		Total: total,
	}, nil
}
