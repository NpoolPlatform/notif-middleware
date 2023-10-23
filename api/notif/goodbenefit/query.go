//nolint:nolintlint,dupl
package goodbenefit

import (
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"

	npool "github.com/NpoolPlatform/message/npool/notif/mw/v1/notif/goodbenefit"

	goodbenefit1 "github.com/NpoolPlatform/notif-middleware/pkg/mw/notif/goodbenefit"
)

func (s *Server) GetGoodBenefits(ctx context.Context, in *npool.GetGoodBenefitsRequest) (*npool.GetGoodBenefitsResponse, error) {
	handler, err := goodbenefit1.NewHandler(
		ctx,
		goodbenefit1.WithConds(in.GetConds()),
		goodbenefit1.WithOffset(in.Offset),
		goodbenefit1.WithLimit(in.Limit),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"GetGoodBenefits",
			"In", in,
			"Error", err,
		)
		return &npool.GetGoodBenefitsResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	infos, total, err := handler.GetGoodBenefits(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"GetGoodBenefits",
			"In", in,
			"Error", err,
		)
		return &npool.GetGoodBenefitsResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.GetGoodBenefitsResponse{
		Infos: infos,
		Total: total,
	}, nil
}

func (s *Server) GetGoodBenefit(ctx context.Context, in *npool.GetGoodBenefitRequest) (*npool.GetGoodBenefitResponse, error) {
	handler, err := goodbenefit1.NewHandler(
		ctx,
		goodbenefit1.WithEntID(&in.EntID, true),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"GetGoodBenefit",
			"In", in,
			"error", err,
		)
		return &npool.GetGoodBenefitResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	info, err := handler.GetGoodBenefit(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"GetGoodBenefits",
			"In", in,
			"Error", err,
		)
		return &npool.GetGoodBenefitResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.GetGoodBenefitResponse{
		Info: info,
	}, nil
}
