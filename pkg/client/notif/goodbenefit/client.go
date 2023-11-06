//nolint:dupl
package goodbenefit

import (
	"context"
	"fmt"
	"time"

	grpc2 "github.com/NpoolPlatform/go-service-framework/pkg/grpc"
	"github.com/NpoolPlatform/notif-middleware/pkg/servicename"

	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	npool "github.com/NpoolPlatform/message/npool/notif/mw/v1/notif/goodbenefit"
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

func GetGoodBenefit(ctx context.Context, id string) (*npool.GoodBenefit, error) {
	info, err := do(ctx, func(_ctx context.Context, cli npool.MiddlewareClient) (cruder.Any, error) {
		resp, err := cli.GetGoodBenefit(ctx, &npool.GetGoodBenefitRequest{
			EntID: id,
		})
		if err != nil {
			return nil, fmt.Errorf("fail get good benefit: %v", err)
		}
		return resp.Info, nil
	})
	if err != nil {
		return nil, fmt.Errorf("fail get good benefit: %v", err)
	}
	return info.(*npool.GoodBenefit), nil
}

func CreateGoodBenefit(ctx context.Context, in *npool.GoodBenefitReq) (*npool.GoodBenefit, error) {
	info, err := do(ctx, func(_ctx context.Context, cli npool.MiddlewareClient) (cruder.Any, error) {
		resp, err := cli.CreateGoodBenefit(ctx, &npool.CreateGoodBenefitRequest{
			Info: in,
		})
		if err != nil {
			return nil, fmt.Errorf("fail create good benefit: %v", err)
		}
		return resp.Info, nil
	})
	if err != nil {
		return nil, fmt.Errorf("fail create good benefit: %v", err)
	}
	return info.(*npool.GoodBenefit), nil
}

func UpdateGoodBenefit(ctx context.Context, in *npool.GoodBenefitReq) (*npool.GoodBenefit, error) {
	info, err := do(ctx, func(_ctx context.Context, cli npool.MiddlewareClient) (cruder.Any, error) {
		resp, err := cli.UpdateGoodBenefit(ctx, &npool.UpdateGoodBenefitRequest{
			Info: in,
		})
		if err != nil {
			return nil, fmt.Errorf("fail create good benefit: %v", err)
		}
		return resp.Info, nil
	})
	if err != nil {
		return nil, fmt.Errorf("fail create good benefit: %v", err)
	}
	return info.(*npool.GoodBenefit), nil
}

func GetGoodBenefits(ctx context.Context, conds *npool.Conds, offset, limit int32) ([]*npool.GoodBenefit, uint32, error) {
	var total uint32
	infos, err := do(ctx, func(_ctx context.Context, cli npool.MiddlewareClient) (cruder.Any, error) {
		resp, err := cli.GetGoodBenefits(ctx, &npool.GetGoodBenefitsRequest{
			Conds:  conds,
			Limit:  limit,
			Offset: offset,
		})
		if err != nil {
			return nil, fmt.Errorf("fail get good benefits: %v", err)
		}
		total = resp.GetTotal()
		return resp.Infos, nil
	})
	if err != nil {
		return nil, 0, fmt.Errorf("fail get good benefits: %v", err)
	}
	return infos.([]*npool.GoodBenefit), total, nil
}

func ExistGoodBenefitConds(ctx context.Context, conds *npool.Conds) (bool, error) {
	info, err := do(ctx, func(_ctx context.Context, cli npool.MiddlewareClient) (cruder.Any, error) {
		resp, err := cli.ExistGoodBenefitConds(ctx, &npool.ExistGoodBenefitCondsRequest{
			Conds: conds,
		})
		if err != nil {
			return nil, fmt.Errorf("fail exist good benefit conds: %v", err)
		}
		return resp.Info, nil
	})
	if err != nil {
		return false, err
	}
	return info.(bool), nil
}
