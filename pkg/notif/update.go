package notif

import (
	"context"

	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"

	npool2 "github.com/NpoolPlatform/message/npool"

	"github.com/NpoolPlatform/notif-manager/pkg/db"
	"github.com/NpoolPlatform/notif-manager/pkg/db/ent"
	"github.com/google/uuid"

	mgrpb "github.com/NpoolPlatform/message/npool/notif/mgr/v1/notif"
	mgrcli "github.com/NpoolPlatform/notif-manager/pkg/client/notif"

	npool "github.com/NpoolPlatform/message/npool/notif/mw/v1/notif"
	entnotif "github.com/NpoolPlatform/notif-manager/pkg/db/ent/notif"
)

func UpdateNotifs(ctx context.Context, ids []string, notified *bool) ([]*npool.Notif, uint32, error) {
	uIDs := []uuid.UUID{}
	for _, val := range ids {
		id, err := uuid.Parse(val)
		if err != nil {
			return nil, 0, err
		}
		uIDs = append(uIDs, id)
	}
	err := db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		stm := cli.
			Notif.
			Update().
			Where(
				entnotif.IDIn(uIDs...),
			)

		if notified != nil {
			stm = stm.SetNotified(*notified)
		}
		_, err := stm.Save(ctx)
		return err
	})
	if err != nil {
		return nil, 0, err
	}
	return GetNotifs(ctx, &mgrpb.Conds{
		IDs: &npool2.StringSliceVal{
			Op:    cruder.IN,
			Value: ids,
		},
	}, 0, int32(len(ids)))
}

func UpdateNotif(ctx context.Context, req *mgrpb.NotifReq) (*npool.Notif, error) {
	info, err := mgrcli.UpdateNotif(ctx, req)
	if err != nil {
		return nil, err
	}
	return expand(info), nil
}
