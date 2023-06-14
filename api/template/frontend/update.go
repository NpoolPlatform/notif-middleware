package frontend

import (
	"context"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/notif/mw/v1/template/frontend"
	frontendtemplate1 "github.com/NpoolPlatform/notif-middleware/pkg/mw/template/frontend"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) UpdateFrontendTemplate(ctx context.Context, in *npool.UpdateFrontendTemplateRequest) (*npool.UpdateFrontendTemplateResponse, error) {
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
			"UpdateFrontendTemplate",
			"In", in,
			"Error", err,
		)
		return &npool.UpdateFrontendTemplateResponse{}, status.Error(codes.Aborted, err.Error())
	}
	info, err := handler.UpdateFrontendTemplate(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"UpdateFrontendTemplate",
			"In", in,
			"Error", err,
		)
		return &npool.UpdateFrontendTemplateResponse{}, status.Error(codes.Aborted, err.Error())
	}
	return &npool.UpdateFrontendTemplateResponse{
		Info: info,
	}, nil
}
