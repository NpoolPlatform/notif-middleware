package notif

import (
	mgrpb "github.com/NpoolPlatform/message/npool/notif/mgr/v1/notif"
	npool "github.com/NpoolPlatform/message/npool/notif/mw/v1/notif"
)

func expand(info *mgrpb.Notif) *npool.Notif {
	return &npool.Notif{
		ID:          info.ID,
		AppID:       info.AppID,
		UserID:      info.UserID,
		Notified:    info.Notified,
		LangID:      info.LangID,
		EventType:   info.EventType,
		UseTemplate: info.UseTemplate,
		Title:       info.Title,
		Content:     info.Content,
		Channel:     info.Channel,
		CreatedAt:   info.CreatedAt,
		UpdatedAt:   info.UpdatedAt,
	}
}
