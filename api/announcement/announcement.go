//nolint:nolintlint,dupl
package announcement

import (
	"context"
	"fmt"

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

	fmt.Println(rows)
	return &npool.GetAnnouncementsResponse{
		Infos: rows,
		Total: total,
	}, nil
}
