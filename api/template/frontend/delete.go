package frontend

import (
	"context"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/notif/mw/v1/template/frontend"
	frontendtemplate1 "github.com/NpoolPlatform/notif-middleware/pkg/mw/template/frontend"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) DeleteFrontendTemplate(ctx context.Context, in *npool.DeleteFrontendTemplateRequest) (*npool.DeleteFrontendTemplateResponse, error) {
	req := in.GetInfo()
	handler, err := frontendtemplate1.NewHandler(
		ctx,
		frontendtemplate1.WithID(req.ID),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"DeleteFrontendTemplate",
			"In", in,
			"Error", err,
		)
		return &npool.DeleteFrontendTemplateResponse{}, status.Error(codes.Aborted, err.Error())
	}
	info, err := handler.DeleteFrontendTemplate(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"DeleteFrontendTemplate",
			"In", in,
			"Error", err,
		)
		return &npool.DeleteFrontendTemplateResponse{}, status.Error(codes.Aborted, err.Error())
	}
	return &npool.DeleteFrontendTemplateResponse{
		Info: info,
	}, nil
}
