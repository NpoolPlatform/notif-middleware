package notif

import (
	"context"

	mgrpb "github.com/NpoolPlatform/message/npool/notif/mgr/v1/notif"
	mgrcli "github.com/NpoolPlatform/notif-manager/pkg/client/notif"

	npool "github.com/NpoolPlatform/message/npool/notif/mw/v1/notif"
)

func CreateNotif(ctx context.Context, req *mgrpb.NotifReq) (*npool.Notif, error) {
	info, err := mgrcli.CreateNotif(ctx, req)
	if err != nil {
		return nil, err
	}
	return expand(info), nil
}

func CreateNotifs(ctx context.Context, req []*mgrpb.NotifReq) ([]*npool.Notif, error) {
	rows, err := mgrcli.CreateNotifs(ctx, req)
	if err != nil {
		return nil, err
	}
	infos := []*npool.Notif{}
	for _, val := range rows {
		infos = append(infos, expand(val))
	}
	return infos, nil
}
