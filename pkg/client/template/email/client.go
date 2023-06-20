//nolint:nolintlint,dupl
package email

import (
	"context"
	"fmt"
	"time"

	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"

	grpc2 "github.com/NpoolPlatform/go-service-framework/pkg/grpc"

	npool "github.com/NpoolPlatform/message/npool/notif/mw/v1/template/email"

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

func CreateEmailTemplate(ctx context.Context, req *npool.EmailTemplateReq) (*npool.EmailTemplate, error) {
	info, err := do(ctx, func(_ctx context.Context, cli npool.MiddlewareClient) (cruder.Any, error) {
		resp, err := cli.CreateEmailTemplate(ctx, &npool.CreateEmailTemplateRequest{
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
	return info.(*npool.EmailTemplate), nil
}

func CreateEmailTemplates(ctx context.Context, reqs []*npool.EmailTemplateReq) ([]*npool.EmailTemplate, error) {
	infos, err := do(ctx, func(_ctx context.Context, cli npool.MiddlewareClient) (cruder.Any, error) {
		resp, err := cli.CreateEmailTemplates(ctx, &npool.CreateEmailTemplatesRequest{
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
	return infos.([]*npool.EmailTemplate), nil
}

func UpdateEmailTemplate(ctx context.Context, req *npool.EmailTemplateReq) (*npool.EmailTemplate, error) {
	info, err := do(ctx, func(_ctx context.Context, cli npool.MiddlewareClient) (cruder.Any, error) {
		resp, err := cli.UpdateEmailTemplate(ctx, &npool.UpdateEmailTemplateRequest{
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
	return info.(*npool.EmailTemplate), nil
}

func GetEmailTemplate(ctx context.Context, id string) (*npool.EmailTemplate, error) {
	info, err := do(ctx, func(_ctx context.Context, cli npool.MiddlewareClient) (cruder.Any, error) {
		resp, err := cli.GetEmailTemplate(ctx, &npool.GetEmailTemplateRequest{
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
	return info.(*npool.EmailTemplate), nil
}

func GetEmailTemplates(ctx context.Context, conds *npool.Conds, offset, limit int32) ([]*npool.EmailTemplate, uint32, error) {
	var total uint32
	infos, err := do(ctx, func(_ctx context.Context, cli npool.MiddlewareClient) (cruder.Any, error) {
		resp, err := cli.GetEmailTemplates(ctx, &npool.GetEmailTemplatesRequest{
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
	return infos.([]*npool.EmailTemplate), total, nil
}

func GetEmailTemplateOnly(ctx context.Context, conds *npool.Conds) (*npool.EmailTemplate, error) {
	infos, err := do(ctx, func(_ctx context.Context, cli npool.MiddlewareClient) (cruder.Any, error) {
		const singleRowLimit = 2
		resp, err := cli.GetEmailTemplates(ctx, &npool.GetEmailTemplatesRequest{
			Conds:  conds,
			Offset: 0,
			Limit:  singleRowLimit,
		})
		if err != nil {
			return nil, err
		}
		return resp.Infos, nil
	})
	if err != nil {
		return nil, err
	}
	if len(infos.([]*npool.EmailTemplate)) == 0 {
		return nil, nil
	}
	if len(infos.([]*npool.EmailTemplate)) > 1 {
		return nil, fmt.Errorf("too many record")
	}
	return infos.([]*npool.EmailTemplate)[0], nil
}

func DeleteEmailTemplate(ctx context.Context, req *npool.EmailTemplateReq) (*npool.EmailTemplate, error) {
	info, err := do(ctx, func(_ctx context.Context, cli npool.MiddlewareClient) (cruder.Any, error) {
		resp, err := cli.DeleteEmailTemplate(ctx, &npool.DeleteEmailTemplateRequest{
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
	return info.(*npool.EmailTemplate), nil
}

func ExistEmailTemplate(ctx context.Context, id string) (bool, error) {
	info, err := do(ctx, func(_ctx context.Context, cli npool.MiddlewareClient) (cruder.Any, error) {
		resp, err := cli.ExistEmailTemplate(ctx, &npool.ExistEmailTemplateRequest{
			ID: id,
		})
		if err != nil {
			return nil, err
		}
		return resp.Info, nil
	})
	if err != nil {
		return false, err
	}
	return info.(bool), nil
}

func ExistEmailTemplateConds(ctx context.Context, conds *npool.Conds) (bool, error) {
	infos, err := do(ctx, func(_ctx context.Context, cli npool.MiddlewareClient) (cruder.Any, error) {
		resp, err := cli.ExistEmailTemplateConds(ctx, &npool.ExistEmailTemplateCondsRequest{
			Conds: conds,
		})
		if err != nil {
			return nil, fmt.Errorf("fail get emailtemplate: %v", err)
		}
		return resp.GetInfo(), nil
	})
	if err != nil {
		return false, fmt.Errorf("fail get emailtemplate: %v", err)
	}
	return infos.(bool), nil
}
