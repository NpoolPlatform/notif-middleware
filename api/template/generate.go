package template

import (
	"context"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"

	npool "github.com/NpoolPlatform/message/npool/notif/mw/v1/template"

	template1 "github.com/NpoolPlatform/notif-middleware/pkg/mw/template"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) GenerateText(ctx context.Context, in *npool.GenerateTextRequest) (*npool.GenerateTextResponse, error) {
	handler, err := template1.NewHandler(
		ctx,
		template1.WithAppID(&in.AppID, true),
		template1.WithLangID(&in.LangID, true),
		template1.WithUsedFor(&in.EventType, true),
		template1.WithChannel(&in.Channel, true),
		template1.WithVars(in.Vars, false),
	)
	if err != nil {
		return nil, err
	}

	info, err := handler.GenerateText(ctx)
	if err != nil {
		logger.Sugar().Errorw("GenerateText", "Error", err)
		return &npool.GenerateTextResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.GenerateTextResponse{
		Info: info,
	}, nil
}
