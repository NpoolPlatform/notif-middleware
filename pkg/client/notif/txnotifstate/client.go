//nolint:dupl
package txnotifstate

import (
	"context"
	"fmt"
	"time"

	grpc2 "github.com/NpoolPlatform/go-service-framework/pkg/grpc"

	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	npool "github.com/NpoolPlatform/message/npool/notif/mw/v1/notif/txnotifstate"

	constant "github.com/NpoolPlatform/notif-middleware/pkg/message/const"

	mgrpb "github.com/NpoolPlatform/message/npool/notif/mgr/v1/notif/txnotifstate"
)

var timeout = 10 * time.Second

type handler func(context.Context, npool.MiddlewareClient) (cruder.Any, error)

func withCRUD(ctx context.Context, handler handler) (cruder.Any, error) {
	_ctx, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()

	conn, err := grpc2.GetGRPCConn(constant.ServiceName, grpc2.GRPCTAG)
	if err != nil {
		return nil, fmt.Errorf("fail get txnotifstate connection: %v", err)
	}

	defer conn.Close()

	cli := npool.NewMiddlewareClient(conn)

	return handler(_ctx, cli)
}

func CreateTxNotifState(ctx context.Context, in *mgrpb.TxNotifStateReq) (*mgrpb.TxNotifState, error) {
	info, err := withCRUD(ctx, func(_ctx context.Context, cli npool.MiddlewareClient) (cruder.Any, error) {
		resp, err := cli.CreateTxNotifState(ctx, &npool.CreateTxNotifStateRequest{
			Info: in,
		})
		if err != nil {
			return nil, fmt.Errorf("fail create txnotifstate: %v", err)
		}
		return resp.Info, nil
	})
	if err != nil {
		return nil, fmt.Errorf("fail create txnotifstate: %v", err)
	}
	return info.(*mgrpb.TxNotifState), nil
}

func UpdateTxNotifState(ctx context.Context, in *mgrpb.TxNotifStateReq) (*mgrpb.TxNotifState, error) {
	info, err := withCRUD(ctx, func(_ctx context.Context, cli npool.MiddlewareClient) (cruder.Any, error) {
		resp, err := cli.UpdateTxNotifState(ctx, &npool.UpdateTxNotifStateRequest{
			Info: in,
		})
		if err != nil {
			return nil, fmt.Errorf("fail create txnotifstate: %v", err)
		}
		return resp.Info, nil
	})
	if err != nil {
		return nil, fmt.Errorf("fail create txnotifstate: %v", err)
	}
	return info.(*mgrpb.TxNotifState), nil
}

func GetTxNotifStateOnly(ctx context.Context, conds *mgrpb.Conds) (*mgrpb.TxNotifState, error) {
	info, err := withCRUD(ctx, func(_ctx context.Context, cli npool.MiddlewareClient) (cruder.Any, error) {
		resp, err := cli.GetTxNotifStateOnly(ctx, &npool.GetTxNotifStateOnlyRequest{
			Conds: conds,
		})
		if err != nil {
			return nil, fmt.Errorf("fail get txnotifstate: %v", err)
		}
		return resp.Info, nil
	})
	if err != nil {
		return nil, fmt.Errorf("fail get txnotifstate: %v", err)
	}
	return info.(*mgrpb.TxNotifState), nil
}

func GetTxNotifStates(ctx context.Context, conds *mgrpb.Conds, offset, limit int32) ([]*mgrpb.TxNotifState, uint32, error) {
	var total uint32
	infos, err := withCRUD(ctx, func(_ctx context.Context, cli npool.MiddlewareClient) (cruder.Any, error) {
		resp, err := cli.GetTxNotifStates(ctx, &npool.GetTxNotifStatesRequest{
			Conds:  conds,
			Limit:  limit,
			Offset: offset,
		})
		if err != nil {
			return nil, fmt.Errorf("fail get txnotifstates: %v", err)
		}
		total = resp.GetTotal()
		return resp.Infos, nil
	})
	if err != nil {
		return nil, 0, fmt.Errorf("fail get txnotifstates: %v", err)
	}
	return infos.([]*mgrpb.TxNotifState), total, nil
}
