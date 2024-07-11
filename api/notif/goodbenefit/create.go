//nolint:nolintlint,dupl
package goodbenefit

import (
	"context"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/notif/mw/v1/notif/goodbenefit"
	goodbenefit1 "github.com/NpoolPlatform/notif-middleware/pkg/mw/notif/goodbenefit"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) CreateGoodBenefit(
	ctx context.Context,
	in *npool.CreateGoodBenefitRequest,
) (
	*npool.CreateGoodBenefitResponse,
	error,
) {
	req := in.GetInfo()
	if req == nil {
		logger.Sugar().Errorw(
			"CreateGoodBenefit",
			"In", in,
		)
		return &npool.CreateGoodBenefitResponse{}, status.Error(codes.InvalidArgument, "Info is empty")
	}
	handler, err := goodbenefit1.NewHandler(
		ctx,
		goodbenefit1.WithEntID(req.EntID, false),
		goodbenefit1.WithGoodID(req.GoodID, true),
		goodbenefit1.WithGoodType(req.GoodType, true),
		goodbenefit1.WithGoodName(req.GoodName, true),
		goodbenefit1.WithCoinTypeID(req.CoinTypeID, true),
		goodbenefit1.WithAmount(req.Amount, false),
		goodbenefit1.WithState(req.State, true),
		goodbenefit1.WithMessage(req.Message, false),
		goodbenefit1.WithBenefitDate(req.BenefitDate, true),
		goodbenefit1.WithTxID(req.TxID, false),
		goodbenefit1.WithGenerated(req.Generated, false),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"CreateGoodBenefit",
			"Req", in,
			"Error", err,
		)
		return &npool.CreateGoodBenefitResponse{}, status.Error(codes.Aborted, err.Error())
	}

	info, err := handler.CreateGoodBenefit(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"CreateGoodBenefit",
			"Req", in,
			"Error", err,
		)
		return &npool.CreateGoodBenefitResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.CreateGoodBenefitResponse{
		Info: info,
	}, nil
}
