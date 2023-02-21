//nolint:dupl
package tx

import (
	"context"
	"fmt"
	"time"

	grpc2 "github.com/NpoolPlatform/go-service-framework/pkg/grpc"

	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	npool "github.com/NpoolPlatform/message/npool/notif/mw/v1/notif/tx"

	constant "github.com/NpoolPlatform/notif-middleware/pkg/message/const"

	mgrpb "github.com/NpoolPlatform/message/npool/notif/mgr/v1/notif/tx"
)

var timeout = 10 * time.Second

type handler func(context.Context, npool.MiddlewareClient) (cruder.Any, error)

func withCRUD(ctx context.Context, handler handler) (cruder.Any, error) {
	_ctx, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()

	conn, err := grpc2.GetGRPCConn(constant.ServiceName, grpc2.GRPCTAG)
	if err != nil {
		return nil, fmt.Errorf("fail get tx connection: %v", err)
	}

	defer conn.Close()

	cli := npool.NewMiddlewareClient(conn)

	return handler(_ctx, cli)
}

func CreateTx(ctx context.Context, in *mgrpb.TxReq) (*mgrpb.Tx, error) {
	info, err := withCRUD(ctx, func(_ctx context.Context, cli npool.MiddlewareClient) (cruder.Any, error) {
		resp, err := cli.CreateTx(ctx, &npool.CreateTxRequest{
			Info: in,
		})
		if err != nil {
			return nil, fmt.Errorf("fail create tx: %v", err)
		}
		return resp.Info, nil
	})
	if err != nil {
		return nil, fmt.Errorf("fail create tx: %v", err)
	}
	return info.(*mgrpb.Tx), nil
}

func UpdateTx(ctx context.Context, in *mgrpb.TxReq) (*mgrpb.Tx, error) {
	info, err := withCRUD(ctx, func(_ctx context.Context, cli npool.MiddlewareClient) (cruder.Any, error) {
		resp, err := cli.UpdateTx(ctx, &npool.UpdateTxRequest{
			Info: in,
		})
		if err != nil {
			return nil, fmt.Errorf("fail create tx: %v", err)
		}
		return resp.Info, nil
	})
	if err != nil {
		return nil, fmt.Errorf("fail create tx: %v", err)
	}
	return info.(*mgrpb.Tx), nil
}

func GetTxOnly(ctx context.Context, conds *mgrpb.Conds) (*mgrpb.Tx, error) {
	info, err := withCRUD(ctx, func(_ctx context.Context, cli npool.MiddlewareClient) (cruder.Any, error) {
		resp, err := cli.GetTxOnly(ctx, &npool.GetTxOnlyRequest{
			Conds: conds,
		})
		if err != nil {
			return nil, fmt.Errorf("fail get tx: %v", err)
		}
		return resp.Info, nil
	})
	if err != nil {
		return nil, fmt.Errorf("fail get tx: %v", err)
	}
	return info.(*mgrpb.Tx), nil
}

func GetTxs(ctx context.Context, conds *mgrpb.Conds, offset, limit int32) ([]*mgrpb.Tx, uint32, error) {
	var total uint32
	infos, err := withCRUD(ctx, func(_ctx context.Context, cli npool.MiddlewareClient) (cruder.Any, error) {
		resp, err := cli.GetTxs(ctx, &npool.GetTxsRequest{
			Conds:  conds,
			Limit:  limit,
			Offset: offset,
		})
		if err != nil {
			return nil, fmt.Errorf("fail get txs: %v", err)
		}
		total = resp.GetTotal()
		return resp.Infos, nil
	})
	if err != nil {
		return nil, 0, fmt.Errorf("fail get txs: %v", err)
	}
	return infos.([]*mgrpb.Tx), total, nil
}
