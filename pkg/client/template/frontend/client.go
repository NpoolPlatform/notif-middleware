//nolint:nolintlint,dupl
package frontend

import (
	"context"
	"fmt"
	"time"

	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"

	grpc2 "github.com/NpoolPlatform/go-service-framework/pkg/grpc"

	npool "github.com/NpoolPlatform/message/npool/notif/mw/v1/template/frontend"

	servicename "github.com/NpoolPlatform/notif-middleware/pkg/servicename"
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

func CreateFrontendTemplate(ctx context.Context, req *npool.FrontendTemplateReq) (*npool.FrontendTemplate, error) {
	info, err := do(ctx, func(_ctx context.Context, cli npool.MiddlewareClient) (cruder.Any, error) {
		resp, err := cli.CreateFrontendTemplate(ctx, &npool.CreateFrontendTemplateRequest{
			Info: req,
		})
		if err != nil {
			return nil, err
		}
		return resp.Info, nil
	})
	if err != nil {
		return nil, err
	}
	return info.(*npool.FrontendTemplate), nil
}

func CreateFrontendTemplates(ctx context.Context, reqs []*npool.FrontendTemplateReq) ([]*npool.FrontendTemplate, error) {
	infos, err := do(ctx, func(_ctx context.Context, cli npool.MiddlewareClient) (cruder.Any, error) {
		resp, err := cli.CreateFrontendTemplates(ctx, &npool.CreateFrontendTemplatesRequest{
			Infos: reqs,
		})
		if err != nil {
			return nil, err
		}
		return resp.Infos, nil
	})
	if err != nil {
		return nil, err
	}
	return infos.([]*npool.FrontendTemplate), nil
}

func UpdateFrontendTemplate(ctx context.Context, req *npool.FrontendTemplateReq) (*npool.FrontendTemplate, error) {
	info, err := do(ctx, func(_ctx context.Context, cli npool.MiddlewareClient) (cruder.Any, error) {
		resp, err := cli.UpdateFrontendTemplate(ctx, &npool.UpdateFrontendTemplateRequest{
			Info: req,
		})
		if err != nil {
			return nil, err
		}
		return resp.Info, nil
	})
	if err != nil {
		return nil, err
	}
	return info.(*npool.FrontendTemplate), nil
}

func GetFrontendTemplate(ctx context.Context, id string) (*npool.FrontendTemplate, error) {
	info, err := do(ctx, func(_ctx context.Context, cli npool.MiddlewareClient) (cruder.Any, error) {
		resp, err := cli.GetFrontendTemplate(ctx, &npool.GetFrontendTemplateRequest{
			ID: id,
		})
		if err != nil {
			return nil, err
		}
		return resp.Info, nil
	})
	if err != nil {
		return nil, err
	}
	return info.(*npool.FrontendTemplate), nil
}

func GetFrontendTemplates(ctx context.Context, conds *npool.Conds, offset, limit int32) ([]*npool.FrontendTemplate, uint32, error) {
	var total uint32
	infos, err := do(ctx, func(_ctx context.Context, cli npool.MiddlewareClient) (cruder.Any, error) {
		resp, err := cli.GetFrontendTemplates(ctx, &npool.GetFrontendTemplatesRequest{
			Conds:  conds,
			Offset: offset,
			Limit:  limit,
		})
		if err != nil {
			return nil, err
		}
		total = resp.GetTotal()
		return resp.Infos, nil
	})
	if err != nil {
		return nil, 0, err
	}
	return infos.([]*npool.FrontendTemplate), total, nil
}

func GetFrontendTemplateOnly(ctx context.Context, conds *npool.Conds) (*npool.FrontendTemplate, error) {
	infos, err := do(ctx, func(_ctx context.Context, cli npool.MiddlewareClient) (cruder.Any, error) {
		resp, err := cli.GetFrontendTemplates(ctx, &npool.GetFrontendTemplatesRequest{
			Conds:  conds,
			Offset: 0,
			Limit:  2, //nolint
		})
		if err != nil {
			return nil, err
		}
		return resp.Infos, nil
	})
	if err != nil {
		return nil, err
	}
	if len(infos.([]*npool.FrontendTemplate)) == 0 {
		return nil, nil
	}
	if len(infos.([]*npool.FrontendTemplate)) > 1 {
		return nil, fmt.Errorf("too many record")
	}
	return infos.([]*npool.FrontendTemplate)[0], nil
}

func DeleteFrontendTemplate(ctx context.Context, req *npool.FrontendTemplateReq) (*npool.FrontendTemplate, error) {
	info, err := do(ctx, func(_ctx context.Context, cli npool.MiddlewareClient) (cruder.Any, error) {
		resp, err := cli.DeleteFrontendTemplate(ctx, &npool.DeleteFrontendTemplateRequest{
			Info: req,
		})
		if err != nil {
			return nil, err
		}
		return resp.Info, nil
	})
	if err != nil {
		return nil, err
	}
	return info.(*npool.FrontendTemplate), nil
}

func ExistFrontendTemplateConds(ctx context.Context, conds *npool.Conds) (bool, error) {
	infos, err := do(ctx, func(_ctx context.Context, cli npool.MiddlewareClient) (cruder.Any, error) {
		resp, err := cli.ExistFrontendTemplateConds(ctx, &npool.ExistFrontendTemplateCondsRequest{
			Conds: conds,
		})
		if err != nil {
			return nil, fmt.Errorf("fail get frontendtemplate: %v", err)
		}
		return resp.GetInfo(), nil
	})
	if err != nil {
		return false, fmt.Errorf("fail get frontendtemplate: %v", err)
	}
	return infos.(bool), nil
}
