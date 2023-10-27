package contact

import (
	"context"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/notif/mw/v1/contact"
	contact1 "github.com/NpoolPlatform/notif-middleware/pkg/mw/contact"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) CreateContact(ctx context.Context, in *npool.CreateContactRequest) (*npool.CreateContactResponse, error) {
	req := in.GetInfo()
	if req == nil {
		logger.Sugar().Errorw(
			"CreateContact",
			"In", in,
		)
		return &npool.CreateContactResponse{}, status.Error(codes.InvalidArgument, "Info is empty")
	}
	handler, err := contact1.NewHandler(
		ctx,
		contact1.WithAppID(req.AppID, true),
		contact1.WithAccount(req.Account, true),
		contact1.WithAccountType(req.AccountType, true),
		contact1.WithUsedFor(req.UsedFor, true),
		contact1.WithSender(req.Sender, true),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"CreateContact",
			"Req", in,
			"Error", err,
		)
		return &npool.CreateContactResponse{}, status.Error(codes.Aborted, err.Error())
	}

	info, err := handler.CreateContact(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"CreateContact",
			"Req", in,
			"Error", err,
		)
		return &npool.CreateContactResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.CreateContactResponse{
		Info: info,
	}, nil
}
