//nolint:nolintlint,dupl
package contact

import (
	"context"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"

	npool "github.com/NpoolPlatform/message/npool/notif/mw/v1/contact"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	contact "github.com/NpoolPlatform/notif-middleware/pkg/mw/contact"
	contact1 "github.com/NpoolPlatform/notif-middleware/pkg/mw/contact/generate"
)

//nolint:funlen,gocyclo
func (s *Server) GenerateContact(ctx context.Context, in *npool.GenerateContactRequest) (*npool.GenerateContactResponse, error) {
	handler, err := contact1.NewHandler(
		ctx,
		contact1.WithSubject(&in.Subject),
		contact1.WithBody(&in.Body),
		contact.WithUsedFor(&in.UsedFor),
		contact.WithAppID(&in.AppID),
		contact.WithSender(&in.Sender),
		contact1.WithSenderName(&in.SenderName),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"GenerateContact",
			"Req", in,
			"Error", err,
		)
		return &npool.GenerateContactResponse{}, status.Error(codes.Aborted, err.Error())
	}

	info, err := handler.GenerateContact(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"GenerateContact",
			"Req", in,
			"Error", err,
		)
		return &npool.GenerateContactResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.GenerateContactResponse{
		Info: info,
	}, nil
}
