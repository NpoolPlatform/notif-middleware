//nolint:dupl
package tx

import (
	"context"
	"fmt"
	"time"

	grpc2 "github.com/NpoolPlatform/go-service-framework/pkg/grpc"
	"github.com/NpoolPlatform/notif-middleware/pkg/servicename"

	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	npool "github.com/NpoolPlatform/message/npool/notif/mw/v1/notif/tx"
)

func do(ctx context.Context, fn func(_ctx context.Context, cli npool.MiddlewareClient) (cruder.Any, error)) (cruder.Any, error) {
	_ctx, cancel := context.WithTimeout(ctx, 10*time.Second) //nolint
	defer cancel()

	conn, err := grpc2.GetGRPCConn(servicename.ServiceDomain, grpc2.GRPCTAG)
	if err != nil {
		return nil, err
	}
	defer conn.Close()

	cli := npool.NewMiddlewareClient(conn)

	return fn(_ctx, cli)
}

func CreateTx(ctx context.Context, in *npool.TxReq) (*npool.Tx, error) {
	info, err := do(ctx, func(_ctx context.Context, cli npool.MiddlewareClient) (cruder.Any, error) {
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
	return info.(*npool.Tx), nil
}

func UpdateTx(ctx context.Context, in *npool.TxReq) (*npool.Tx, error) {
	info, err := do(ctx, func(_ctx context.Context, cli npool.MiddlewareClient) (cruder.Any, error) {
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
	return info.(*npool.Tx), nil
}

func GetTxOnly(ctx context.Context, conds *npool.Conds) (*npool.Tx, error) {
	info, err := do(ctx, func(_ctx context.Context, cli npool.MiddlewareClient) (cruder.Any, error) {
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
	return info.(*npool.Tx), nil
}

func GetTxs(ctx context.Context, conds *npool.Conds, offset, limit int32) ([]*npool.Tx, uint32, error) {
	var total uint32
	infos, err := do(ctx, func(_ctx context.Context, cli npool.MiddlewareClient) (cruder.Any, error) {
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
	return infos.([]*npool.Tx), total, nil
}

func ExistTx(ctx context.Context, id string) (bool, error) {
	info, err := do(ctx, func(_ctx context.Context, cli npool.MiddlewareClient) (cruder.Any, error) {
		resp, err := cli.ExistTx(ctx, &npool.ExistTxRequest{
			ID: id,
		})
		if err != nil {
			return nil, fmt.Errorf("fail exist tx: %v", err)
		}
		return resp.Info, nil
	})
	if err != nil {
		return false, err
	}
	return info.(bool), nil
}

func ExistTxConds(ctx context.Context, conds *npool.Conds) (bool, error) {
	info, err := do(ctx, func(_ctx context.Context, cli npool.MiddlewareClient) (cruder.Any, error) {
		resp, err := cli.ExistTxConds(ctx, &npool.ExistTxCondsRequest{
			Conds: conds,
		})
		if err != nil {
			return nil, fmt.Errorf("fail exist tx conds: %v", err)
		}
		return resp.Info, nil
	})
	if err != nil {
		return false, err
	}
	return info.(bool), nil
}
