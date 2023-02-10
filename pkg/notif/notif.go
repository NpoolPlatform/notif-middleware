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

func UpdateNotifs(ctx context.Context, ids []string, emailSend, alreadyRead *bool) ([]*npool.Notif, uint32, error) {
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
		if emailSend != nil {
			stm = stm.SetEmailSend(*emailSend)
		}
		if alreadyRead != nil {
			stm = stm.SetAlreadyRead(*alreadyRead)
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
	return expand(info), nil
}

func expand(info *mgrpb.Notif) *npool.Notif {
	return &npool.Notif{
		ID:          info.ID,
		AppID:       info.AppID,
		UserID:      info.UserID,
		AlreadyRead: info.AlreadyRead,
		LangID:      info.LangID,
		EventType:   info.EventType,
		UseTemplate: info.UseTemplate,
		Title:       info.Title,
		Content:     info.Content,
		Channels:    info.Channels,
		EmailSend:   info.EmailSend,
		CreatedAt:   info.CreatedAt,
		UpdatedAt:   info.UpdatedAt,
	}
}
