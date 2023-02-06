package sendstate

import (
	"context"
	"fmt"
	"time"

	"github.com/NpoolPlatform/message/npool/notif/mgr/v1/channel"

	grpc2 "github.com/NpoolPlatform/go-service-framework/pkg/grpc"

	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"

	npool "github.com/NpoolPlatform/message/npool/notif/mw/v1/announcement/sendstate"

	constant "github.com/NpoolPlatform/notif-middleware/pkg/message/const"
)

var timeout = 10 * time.Second

type handler func(context.Context, npool.MiddlewareClient) (cruder.Any, error)

func withCRUD(ctx context.Context, handler handler) (cruder.Any, error) {
	_ctx, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()

	conn, err := grpc2.GetGRPCConn(constant.ServiceName, grpc2.GRPCTAG)
	if err != nil {
		return nil, fmt.Errorf("fail get sendstate connection: %v", err)
	}

	defer conn.Close()

	cli := npool.NewMiddlewareClient(conn)

	return handler(_ctx, cli)
}

func CreateSendState(ctx context.Context, appID, userID, announcementID string, c channel.NotifChannel) error {
	_, err := withCRUD(ctx, func(_ctx context.Context, cli npool.MiddlewareClient) (cruder.Any, error) {
		_, err := cli.CreateSendState(ctx, &npool.CreateSendStateRequest{
			AppID:          appID,
			UserID:         userID,
			AnnouncementID: announcementID,
			Channel:        c,
		})
		if err != nil {
			return nil, fmt.Errorf("fail get sendstates: %v", err)
		}
		return nil, nil
	})
	if err != nil {
		return fmt.Errorf("fail get sendstates: %v", err)
	}
	return nil
}

func GetSendStates(ctx context.Context, conds *npool.Conds, offset, limit int32) ([]*npool.SendState, uint32, error) {
	var total uint32
	infos, err := withCRUD(ctx, func(_ctx context.Context, cli npool.MiddlewareClient) (cruder.Any, error) {
		resp, err := cli.GetSendStates(ctx, &npool.GetSendStatesRequest{
			Conds:  conds,
			Limit:  limit,
			Offset: offset,
		})
		if err != nil {
			return nil, fmt.Errorf("fail get sendstates: %v", err)
		}
		total = resp.GetTotal()
		return resp.GetInfos(), nil
	})
	if err != nil {
		return nil, 0, fmt.Errorf("fail get sendstates: %v", err)
	}
	return infos.([]*npool.SendState), total, nil
}
