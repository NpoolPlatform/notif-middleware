//nolint:nolintlint,dupl
package contact

import (
	"context"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"

	basetypes "github.com/NpoolPlatform/message/npool/basetypes/v1"
	npool "github.com/NpoolPlatform/message/npool/notif/mw/v1/contact"

	constant "github.com/NpoolPlatform/notif-middleware/pkg/message/const"
	commontracer "github.com/NpoolPlatform/notif-middleware/pkg/tracer"

	"go.opentelemetry.io/otel"
	scodes "go.opentelemetry.io/otel/codes"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	contact1 "github.com/NpoolPlatform/notif-middleware/pkg/contact"

	"github.com/google/uuid"
)

//nolint:funlen,gocyclo
func (s *Server) GenerateContact(ctx context.Context, in *npool.GenerateContactRequest) (*npool.GenerateContactResponse, error) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "GenerateContact")
	defer span.End()

	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()

	span = commontracer.TraceInvoker(span, "contact", "crud", "Update")

	if _, err := uuid.Parse(in.GetAppID()); err != nil {
		logger.Sugar().Errorw("GenerateContact", "AppID", in.GetAppID(), "error", err)
		return &npool.GenerateContactResponse{}, status.Error(codes.Internal, err.Error())
	}
	switch in.GetUsedFor() {
	case basetypes.UsedFor_Contact:
	default:
		logger.Sugar().Errorw("GenerateContact", "UsedFor", in.GetUsedFor())
		return &npool.GenerateContactResponse{}, status.Error(codes.InvalidArgument, "UsedFor is invalid")
	}
	if in.GetSender() == "" {
		logger.Sugar().Errorw("GenerateContact", "Sender", in.GetSender())
		return &npool.GenerateContactResponse{}, status.Error(codes.InvalidArgument, "Sender is invalid")
	}
	if in.GetSubject() == "" {
		logger.Sugar().Errorw("GenerateContact", "Subject", in.GetSubject())
		return &npool.GenerateContactResponse{}, status.Error(codes.InvalidArgument, "Subject is invalid")
	}
	if in.GetBody() == "" {
		logger.Sugar().Errorw("GenerateContact", "Body", in.GetBody())
		return &npool.GenerateContactResponse{}, status.Error(codes.InvalidArgument, "Body is invalid")
	}

	info, err := contact1.GenerateContact(
		ctx,
		in.GetSubject(),
		in.GetBody(),
		in.GetAppID(),
		in.GetSender(),
		in.GetSenderName(),
		in.GetUsedFor())
	if err != nil {
		logger.Sugar().Errorw("GenerateContact", "error", err)
		return &npool.GenerateContactResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.GenerateContactResponse{
		Info: info,
	}, nil
}
