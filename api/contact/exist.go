package contact

import (
	"context"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/notif/mw/v1/contact"
	contact1 "github.com/NpoolPlatform/notif-middleware/pkg/mw/contact"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) ExistContactConds(ctx context.Context, in *npool.ExistContactCondsRequest) (*npool.ExistContactCondsResponse, error) {
	handler, err := contact1.NewHandler(
		ctx,
		contact1.WithConds(in.GetConds()),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"ExistContactConds",
			"Req", in,
			"Error", err,
		)
		return &npool.ExistContactCondsResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	info, err := handler.ExistContactConds(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"ExistContactConds",
			"Req", in,
			"Error", err,
		)
		return &npool.ExistContactCondsResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.ExistContactCondsResponse{
		Info: info,
	}, nil
}
