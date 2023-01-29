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
	}, nil
}

func GetNotif(ctx context.Context, id string) (*npool.Notif, error) {
	info, err := mgrcli.GetNotif(ctx, id)
	if err != nil {
		return nil, err
	}
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
	}, nil
}

func GetNotifs(ctx context.Context, conds *mgrpb.Conds, offset, limit int32) ([]*npool.Notif, uint32, error) {
	infos, total, err := mgrcli.GetNotifs(ctx, conds, offset, limit)
	if err != nil {
		return nil, 0, err
	}
	infos1 := []*npool.Notif{}
	for _, info := range infos {
		infos1 = append(infos1, &npool.Notif{
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
		})
	}
	return infos1, total, nil
}

func GetNotifOnly(ctx context.Context, conds *mgrpb.Conds) (*npool.Notif, error) {
	info, err := mgrcli.GetNotifOnly(ctx, conds)
	if err != nil {
		return nil, err
	}
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
	}, nil
}
