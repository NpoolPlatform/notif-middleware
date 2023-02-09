package sendstate

import (
	"context"

	mgrpb "github.com/NpoolPlatform/message/npool/notif/mgr/v1/announcement"
	mgrcli "github.com/NpoolPlatform/notif-manager/pkg/client/announcement"

	entreadannouncement "github.com/NpoolPlatform/notif-manager/pkg/db/ent/readannouncement"

	"entgo.io/ent/dialect/sql"
	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	"github.com/NpoolPlatform/notif-manager/pkg/db"
	"github.com/google/uuid"

	npool "github.com/NpoolPlatform/message/npool/notif/mw/v1/announcement"
	"github.com/NpoolPlatform/notif-manager/pkg/db/ent"
	entannouncement "github.com/NpoolPlatform/notif-manager/pkg/db/ent/announcement"
	entuserannouncement "github.com/NpoolPlatform/notif-manager/pkg/db/ent/userannouncement"
)

func GetAnnouncements(
	ctx context.Context,
	conds *mgrpb.Conds,
	offset, limit int32,
) (
	[]*npool.Announcement,
	uint32,
	error,
) {
	rows, total, err := mgrcli.GetAnnouncements(ctx, conds, offset, limit)
	if err != nil {
		return nil, 0, err
	}

	infos := []*npool.Announcement{}
	for _, val := range rows {
		infos = append(infos, &npool.Announcement{
			AnnouncementID:   val.ID,
			AppID:            val.AppID,
			Title:            val.Title,
			Content:          val.Content,
			EndAt:            val.EndAt,
			AnnouncementType: val.AnnouncementType,
		})
	}
	return infos, total, err
}

func GetAnnouncementStates(
	ctx context.Context,
	appID, userID, langID string,
	offset, limit int32,
) (
	[]*npool.Announcement,
	uint32,
	error,
) {
	var infos []*npool.Announcement
	var total uint32

	err := db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		stm := cli.
			Announcement.
			Query().
			Where(
				entannouncement.AppID(uuid.MustParse(appID)),
				entannouncement.LangID(uuid.MustParse(langID)),
			)
		stm.Select().Modify(func(s *sql.Selector) {
			t1 := sql.Table(entreadannouncement.Table)
			t2 := sql.Table(entuserannouncement.Table)
			s.
				LeftJoin(t1).
				On(
					s.C(entannouncement.FieldID),
					t1.C(entreadannouncement.FieldAnnouncementID),
				).
				OnP(
					sql.EQ(t1.C(entreadannouncement.FieldUserID), userID),
				)
			s.
				LeftJoin(t2).
				On(
					s.C(entannouncement.FieldID),
					t2.C(entuserannouncement.FieldAnnouncementID),
				).
				OnP(
					sql.EQ(t2.C(entuserannouncement.FieldUserID), userID),
				)
			s.Select(
				sql.As(s.C(entannouncement.FieldID), "announcement_id"),
				s.C(entannouncement.FieldAppID),
				t2.C(entuserannouncement.FieldUserID),
				s.C(entannouncement.FieldLangID),
				s.C(entannouncement.FieldTitle),
				s.C(entannouncement.FieldContent),
				s.C(entannouncement.FieldCreatedAt),
				s.C(entannouncement.FieldUpdatedAt),
				s.C(entannouncement.FieldEndAt),
				s.C(entannouncement.FieldType),
				sql.As(t1.C(entreadannouncement.FieldUserID), "read_user_id"),
			)
		})
		_total, err := stm.Count(ctx)
		if err != nil {
			return err
		}
		total = uint32(_total)
		return stm.
			Order(ent.Desc(entannouncement.FieldCreatedAt)).
			Offset(int(offset)).
			Limit(int(limit)).
			Modify(func(s *sql.Selector) {
			}).
			Scan(ctx, &infos)
	})
	if err != nil {
		logger.Sugar().Errorw("GetSendState", "err", err)
		return nil, 0, err
	}
	if len(infos) == 0 {
		return nil, 0, nil
	}

	infos = expand(infos)
	if err != nil {
		logger.Sugar().Errorw("GetSendState", "err", err)
		return nil, 0, err
	}

	return infos, total, nil
}

func expand(infos []*npool.Announcement) []*npool.Announcement {
	retInfos := []*npool.Announcement{}
	for _, info := range infos {
		if info.ReadUserID != "" {
			info.AlreadyRead = true
		}

		info.AnnouncementType = mgrpb.AnnouncementType(mgrpb.AnnouncementType_value[info.AnnouncementTypeStr])

		if info.AnnouncementType == mgrpb.AnnouncementType_AllUsers {
			retInfos = append(retInfos, info)
		}
		if info.AnnouncementType == mgrpb.AnnouncementType_AppointUsers && info.UserID != "" {
			retInfos = append(retInfos, info)
		}
	}
	return retInfos
}
