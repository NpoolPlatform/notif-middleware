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
		frontendtemplate1.WithID(&in.ID),
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
