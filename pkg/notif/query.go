package notif

import (
	"context"

	mgrpb "github.com/NpoolPlatform/message/npool/notif/mgr/v1/notif"
	mgrcli "github.com/NpoolPlatform/notif-manager/pkg/client/notif"

	npool "github.com/NpoolPlatform/message/npool/notif/mw/v1/notif"
)

func GetNotif(ctx context.Context, id string) (*npool.Notif, error) {
	info, err := mgrcli.GetNotif(ctx, id)
	if err != nil {
		return nil, err
	}
	return expand(info), nil
}

func GetNotifs(ctx context.Context, conds *mgrpb.Conds, offset, limit int32) ([]*npool.Notif, uint32, error) {
	rows, total, err := mgrcli.GetNotifs(ctx, conds, offset, limit)
	if err != nil {
		return nil, 0, err
	}
	infos := []*npool.Notif{}
	for _, info := range rows {
		infos = append(infos, expand(info))
	}
	return infos, total, nil
}

func GetNotifOnly(ctx context.Context, conds *mgrpb.Conds) (*npool.Notif, error) {
	info, err := mgrcli.GetNotifOnly(ctx, conds)
	if err != nil {
		return nil, err
	}
	if info == nil {
		return nil, nil
	}
	return expand(info), nil
}
