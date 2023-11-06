//nolint:nolintlint,dupl
package frontend

import (
	"context"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/notif/mw/v1/template/frontend"
	frontendtemplate1 "github.com/NpoolPlatform/notif-middleware/pkg/mw/template/frontend"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) ExistFrontendTemplate(
	ctx context.Context,
	in *npool.ExistFrontendTemplateRequest,
) (
	*npool.ExistFrontendTemplateResponse,
	error,
) {
	handler, err := frontendtemplate1.NewHandler(
		ctx,
		frontendtemplate1.WithEntID(&in.EntID, true),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"ExistFrontendTemplate",
			"In", in,
			"Error", err,
		)
		return &npool.ExistFrontendTemplateResponse{}, status.Error(codes.Aborted, err.Error())
	}

	exist, err := handler.ExistFrontendTemplate(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"ExistFrontendTemplate",
			"In", in,
			"Error", err,
		)
		return &npool.ExistFrontendTemplateResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.ExistFrontendTemplateResponse{
		Info: exist,
	}, nil
}

func (s *Server) ExistFrontendTemplateConds(
	ctx context.Context,
	in *npool.ExistFrontendTemplateCondsRequest,
) (
	*npool.ExistFrontendTemplateCondsResponse,
	error,
) {
	handler, err := frontendtemplate1.NewHandler(
		ctx,
		frontendtemplate1.WithConds(in.GetConds()),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"ExistFrontendTemplate",
			"Req", in,
			"Error", err,
		)
		return &npool.ExistFrontendTemplateCondsResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	info, err := handler.ExistFrontendTemplateConds(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"ExistFrontendTemplate",
			"Req", in,
			"Error", err,
		)
		return &npool.ExistFrontendTemplateCondsResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.ExistFrontendTemplateCondsResponse{
		Info: info,
	}, nil
}
