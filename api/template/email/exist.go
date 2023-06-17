//nolint:nolintlint,dupl
package email

import (
	"context"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/notif/mw/v1/template/email"
	emailtemplate1 "github.com/NpoolPlatform/notif-middleware/pkg/mw/template/email"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) ExistEmailTemplate(ctx context.Context, in *npool.ExistEmailTemplateRequest) (*npool.ExistEmailTemplateResponse, error) {
	handler, err := emailtemplate1.NewHandler(
		ctx,
		emailtemplate1.WithID(&in.ID),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"ExistEmailTemplate",
			"In", in,
			"Error", err,
		)
		return &npool.ExistEmailTemplateResponse{}, status.Error(codes.Aborted, err.Error())
	}

	exist, err := handler.ExistEmailTemplate(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"ExistEmailTemplate",
			"In", in,
			"Error", err,
		)
		return &npool.ExistEmailTemplateResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.ExistEmailTemplateResponse{
		Info: exist,
	}, nil
}

func (s *Server) ExistEmailTemplateConds(
	ctx context.Context,
	in *npool.ExistEmailTemplateCondsRequest,
) (
	*npool.ExistEmailTemplateCondsResponse,
	error,
) {
	handler, err := emailtemplate1.NewHandler(ctx,
		emailtemplate1.WithConds(in.GetConds()),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"ExistEmailTemplate",
			"Req", in,
			"Error", err,
		)
		return &npool.ExistEmailTemplateCondsResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	info, err := handler.ExistEmailTemplateConds(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"ExistEmailTemplate",
			"Req", in,
			"Error", err,
		)
		return &npool.ExistEmailTemplateCondsResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.ExistEmailTemplateCondsResponse{
		Info: info,
	}, nil
}
