//nolint:nolintlint,dupl
package sms

import (
	"context"
	"fmt"
	"time"

	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"

	grpc2 "github.com/NpoolPlatform/go-service-framework/pkg/grpc"

	npool "github.com/NpoolPlatform/message/npool/notif/mw/v1/template/sms"

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

func CreateSMSTemplate(ctx context.Context, req *npool.SMSTemplateReq) (*npool.SMSTemplate, error) {
	info, err := do(ctx, func(_ctx context.Context, cli npool.MiddlewareClient) (cruder.Any, error) {
		resp, err := cli.CreateSMSTemplate(ctx, &npool.CreateSMSTemplateRequest{
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
	return info.(*npool.SMSTemplate), nil
}

func CreateSMSTemplates(ctx context.Context, reqs []*npool.SMSTemplateReq) ([]*npool.SMSTemplate, error) {
	infos, err := do(ctx, func(_ctx context.Context, cli npool.MiddlewareClient) (cruder.Any, error) {
		resp, err := cli.CreateSMSTemplates(ctx, &npool.CreateSMSTemplatesRequest{
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
	return infos.([]*npool.SMSTemplate), nil
}

func UpdateSMSTemplate(ctx context.Context, req *npool.SMSTemplateReq) (*npool.SMSTemplate, error) {
	info, err := do(ctx, func(_ctx context.Context, cli npool.MiddlewareClient) (cruder.Any, error) {
		resp, err := cli.UpdateSMSTemplate(ctx, &npool.UpdateSMSTemplateRequest{
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
	return info.(*npool.SMSTemplate), nil
}

func GetSMSTemplate(ctx context.Context, id string) (*npool.SMSTemplate, error) {
	info, err := do(ctx, func(_ctx context.Context, cli npool.MiddlewareClient) (cruder.Any, error) {
		resp, err := cli.GetSMSTemplate(ctx, &npool.GetSMSTemplateRequest{
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
	return info.(*npool.SMSTemplate), nil
}

func GetSMSTemplates(ctx context.Context, conds *npool.Conds, offset, limit int32) ([]*npool.SMSTemplate, uint32, error) {
	var total uint32
	infos, err := do(ctx, func(_ctx context.Context, cli npool.MiddlewareClient) (cruder.Any, error) {
		resp, err := cli.GetSMSTemplates(ctx, &npool.GetSMSTemplatesRequest{
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
	return infos.([]*npool.SMSTemplate), total, nil
}

func GetSMSTemplateOnly(ctx context.Context, conds *npool.Conds) (*npool.SMSTemplate, error) {
	infos, err := do(ctx, func(_ctx context.Context, cli npool.MiddlewareClient) (cruder.Any, error) {
		resp, err := cli.GetSMSTemplates(ctx, &npool.GetSMSTemplatesRequest{
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
	if len(infos.([]*npool.SMSTemplate)) == 0 {
		return nil, nil
	}
	if len(infos.([]*npool.SMSTemplate)) > 1 {
		return nil, fmt.Errorf("too many record")
	}
	return infos.([]*npool.SMSTemplate)[0], nil
}

func DeleteSMSTemplate(ctx context.Context, req *npool.SMSTemplateReq) (*npool.SMSTemplate, error) {
	info, err := do(ctx, func(_ctx context.Context, cli npool.MiddlewareClient) (cruder.Any, error) {
		resp, err := cli.DeleteSMSTemplate(ctx, &npool.DeleteSMSTemplateRequest{
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
	return info.(*npool.SMSTemplate), nil
}

func ExistSMSTemplate(ctx context.Context, id string) (bool, error) {
	info, err := do(ctx, func(_ctx context.Context, cli npool.MiddlewareClient) (cruder.Any, error) {
		resp, err := cli.ExistSMSTemplate(ctx, &npool.ExistSMSTemplateRequest{
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

func ExistSMSTemplateConds(ctx context.Context, conds *npool.Conds) (bool, error) {
	infos, err := do(ctx, func(_ctx context.Context, cli npool.MiddlewareClient) (cruder.Any, error) {
		resp, err := cli.ExistSMSTemplateConds(ctx, &npool.ExistSMSTemplateCondsRequest{
			Conds: conds,
		})
		if err != nil {
			return nil, fmt.Errorf("fail get smstemplate: %v", err)
		}
		return resp.GetInfo(), nil
	})
	if err != nil {
		return false, fmt.Errorf("fail get smstemplate: %v", err)
	}
	return infos.(bool), nil
}
