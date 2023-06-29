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
	handler, err := contact1.NewHandler(
		ctx,
		contact1.WithAppID(req.AppID),
		contact1.WithAccount(req.Account),
		contact1.WithAccountType(req.AccountType),
		contact1.WithUsedFor(req.UsedFor),
		contact1.WithSender(req.Sender),
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
