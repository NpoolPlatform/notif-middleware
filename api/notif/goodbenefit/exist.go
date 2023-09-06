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

func (s *Server) ExistGoodBenefitConds(ctx context.Context, in *npool.ExistGoodBenefitCondsRequest) (*npool.ExistGoodBenefitCondsResponse, error) {
	handler, err := goodbenefit1.NewHandler(
		ctx,
		goodbenefit1.WithConds(in.GetConds()),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"ExistGoodBenefitConds",
			"In", in,
			"Error", err,
		)
		return &npool.ExistGoodBenefitCondsResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	exist, err := handler.ExistGoodBenefitConds(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"ExistGoodBenefitConds",
			"In", in,
			"Error", err,
		)
		return &npool.ExistGoodBenefitCondsResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.ExistGoodBenefitCondsResponse{
		Info: exist,
	}, nil
}
