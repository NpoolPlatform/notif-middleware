//nolint:nolintlint,dupl
package user

import (
	"context"

	mgrpb "github.com/NpoolPlatform/message/npool/notif/mgr/v1/announcement/user"

	user1 "github.com/NpoolPlatform/notif-middleware/pkg/announcement/user"
	commontracer "github.com/NpoolPlatform/notif-middleware/pkg/tracer"

	constant "github.com/NpoolPlatform/notif-middleware/pkg/message/const"

	"go.opentelemetry.io/otel"
	scodes "go.opentelemetry.io/otel/codes"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/notif/mw/v1/announcement/user"

	"github.com/google/uuid"
)

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

func (s *Server) GetUsers(ctx context.Context, in *npool.GetUsersRequest) (*npool.GetUsersResponse, error) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "GetUsers")
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
		return &npool.GetUsersResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	span = commontracer.TraceInvoker(span, "announcement/user", "crud", "Rows")

	rows, total, err := user1.GetUsers(ctx, in.GetConds(), in.GetOffset(), in.GetLimit())
	if err != nil {
		logger.Sugar().Errorw("GetUsers", "error", err)
		return &npool.GetUsersResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.GetUsersResponse{
		Infos: rows,
		Total: total,
	}, nil
}
