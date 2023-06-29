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

func (s *Server) CreateFrontendTemplate(
	ctx context.Context,
	in *npool.CreateFrontendTemplateRequest,
) (
	*npool.CreateFrontendTemplateResponse,
	error,
) {
	req := in.GetInfo()
	handler, err := frontendtemplate1.NewHandler(
		ctx,
		frontendtemplate1.WithID(req.ID),
		frontendtemplate1.WithAppID(req.AppID),
		frontendtemplate1.WithLangID(req.LangID),
		frontendtemplate1.WithUsedFor(req.UsedFor),
		frontendtemplate1.WithTitle(req.Title),
		frontendtemplate1.WithContent(req.Content),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"CreateFrontendTemplate",
			"In", in,
			"Error", err,
		)
		return &npool.CreateFrontendTemplateResponse{}, status.Error(codes.Aborted, err.Error())
	}
	info, err := handler.CreateFrontendTemplate(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"CreateFrontendTemplate",
			"In", in,
			"Error", err,
		)
		return &npool.CreateFrontendTemplateResponse{}, status.Error(codes.Aborted, err.Error())
	}
	return &npool.CreateFrontendTemplateResponse{
		Info: info,
	}, nil
}

func (s *Server) CreateFrontendTemplates(
	ctx context.Context,
	in *npool.CreateFrontendTemplatesRequest,
) (
	*npool.CreateFrontendTemplatesResponse,
	error,
) {
	handler, err := frontendtemplate1.NewHandler(
		ctx,
		frontendtemplate1.WithReqs(in.GetInfos()),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"CreateFrontendTemplates",
			"In", in,
			"Error", err,
		)
		return &npool.CreateFrontendTemplatesResponse{}, status.Error(codes.Aborted, err.Error())
	}
	infos, err := handler.CreateFrontendTemplates(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"CreateFrontendTemplates",
			"In", in,
			"Error", err,
		)
		return &npool.CreateFrontendTemplatesResponse{}, status.Error(codes.Aborted, err.Error())
	}
	return &npool.CreateFrontendTemplatesResponse{
		Infos: infos,
	}, nil
}
