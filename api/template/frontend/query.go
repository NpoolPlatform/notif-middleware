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

func (s *Server) GetFrontendTemplate(
	ctx context.Context,
	in *npool.GetFrontendTemplateRequest,
) (
	*npool.GetFrontendTemplateResponse,
	error,
) {
	handler, err := frontendtemplate1.NewHandler(
		ctx,
		frontendtemplate1.WithID(&in.ID),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"GetFrontendTemplate",
			"In", in,
			"Error", err,
		)
		return &npool.GetFrontendTemplateResponse{}, status.Error(codes.Aborted, err.Error())
	}
	info, err := handler.GetFrontendTemplate(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"GetFrontendTemplate",
			"In", in,
			"Error", err,
		)
		return &npool.GetFrontendTemplateResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.GetFrontendTemplateResponse{
		Info: info,
	}, nil
}

func (s *Server) GetFrontendTemplateOnly(
	ctx context.Context,
	in *npool.GetFrontendTemplateOnlyRequest,
) (
	*npool.GetFrontendTemplateOnlyResponse,
	error,
) {
	handler, err := frontendtemplate1.NewHandler(
		ctx,
		frontendtemplate1.WithConds(in.Conds),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"GetFrontendTemplateOnly",
			"In", in,
			"Error", err,
		)
		return &npool.GetFrontendTemplateOnlyResponse{}, status.Error(codes.Aborted, err.Error())
	}
	info, err := handler.GetFrontendTemplateOnly(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"GetFrontendTemplateOnly",
			"In", in,
			"Error", err,
		)
		return &npool.GetFrontendTemplateOnlyResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.GetFrontendTemplateOnlyResponse{
		Info: info,
	}, nil
}

func (s *Server) GetFrontendTemplates(
	ctx context.Context,
	in *npool.GetFrontendTemplatesRequest,
) (
	*npool.GetFrontendTemplatesResponse,
	error,
) {
	handler, err := frontendtemplate1.NewHandler(
		ctx,
		frontendtemplate1.WithConds(in.GetConds()),
		frontendtemplate1.WithOffset(in.GetOffset()),
		frontendtemplate1.WithLimit(in.GetLimit()),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"GetFrontendTemplates",
			"In", in,
			"Error", err,
		)
		return &npool.GetFrontendTemplatesResponse{}, status.Error(codes.Aborted, err.Error())
	}
	infos, total, err := handler.GetFrontendTemplates(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"GetFrontendTemplates",
			"In", in,
			"Error", err,
		)
		return &npool.GetFrontendTemplatesResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.GetFrontendTemplatesResponse{
		Infos: infos,
		Total: total,
	}, nil
}
