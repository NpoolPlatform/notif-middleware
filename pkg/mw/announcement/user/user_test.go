package user

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"testing"
	"time"

	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"

	announcementpb "github.com/NpoolPlatform/message/npool/notif/mw/v1/announcement"
	npool "github.com/NpoolPlatform/message/npool/notif/mw/v1/announcement/user"
	announcement1 "github.com/NpoolPlatform/notif-middleware/pkg/mw/announcement"
	handler1 "github.com/NpoolPlatform/notif-middleware/pkg/mw/announcement/handler"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"

	basetypes "github.com/NpoolPlatform/message/npool/basetypes/v1"
	"github.com/NpoolPlatform/notif-middleware/pkg/testinit"
)

func init() {
	if runByGithubAction, err := strconv.ParseBool(os.Getenv("RUN_BY_GITHUB_ACTION")); err == nil && runByGithubAction {
		return
	}
	if err := testinit.Init(); err != nil {
		fmt.Printf("cannot init test stub: %v\n", err)
	}
}

var (
	announcement = announcementpb.Announcement{
		AppID:               uuid.NewString(),
		LangID:              uuid.NewString(),
		Title:               uuid.NewString(),
		Content:             uuid.NewString(),
		Channel:             basetypes.NotifChannel_ChannelEmail,
		ChannelStr:          basetypes.NotifChannel_ChannelEmail.String(),
		AnnouncementType:    basetypes.NotifType_NotifMulticast,
		AnnouncementTypeStr: basetypes.NotifType_NotifMulticast.String(),
		StartAt:             uint32(time.Now().Add(10 * time.Minute).Unix()),
		EndAt:               uint32(time.Now().Add(1 * time.Hour).Unix()),
	}

	ret = npool.AnnouncementUser{
		AppID:            announcement.AppID,
		UserID:           uuid.NewString(),
		AnnouncementID:   "",
		LangID:           announcement.LangID,
		Title:            announcement.Title,
		Content:          announcement.Content,
		Channel:          announcement.ChannelStr,
		AnnouncementType: announcement.AnnouncementTypeStr,
		EndAt:            announcement.EndAt,
	}
)

func setupAnnouncementUser(t *testing.T) func(*testing.T) {
	handler, err := announcement1.NewHandler(
		context.Background(),
		announcement1.WithTitle(&announcement.Title, true),
		announcement1.WithContent(&announcement.Content, true),
		announcement1.WithAppID(&announcement.AppID, true),
		announcement1.WithLangID(&announcement.LangID, true),
		announcement1.WithChannel(&announcement.Channel, true),
		announcement1.WithAnnouncementType(&announcement.AnnouncementType, true),
		announcement1.WithStartAt(&announcement.StartAt, true),
		announcement1.WithEndAt(&announcement.EndAt, true),
	)
	assert.Nil(t, err)

	_announcement, err := handler.CreateAnnouncement(context.Background())
	assert.Nil(t, err)
	assert.NotNil(t, _announcement)

	ret.AnnouncementID = _announcement.EntID

	handler.ID = &_announcement.ID

	return func(*testing.T) {
		_, _ = handler.DeleteAnnouncement(context.Background())
	}
}

func createAnnouncementUser(t *testing.T) {
	handler, err := NewHandler(
		context.Background(),
		handler1.WithAppID(&ret.AppID, true),
		handler1.WithUserID(&ret.UserID, true),
		handler1.WithAnnouncementID(&ret.AppID, &ret.AnnouncementID, true),
	)
	assert.Nil(t, err)

	info, err := handler.CreateAnnouncementUser(context.Background())
	if assert.Nil(t, err) {
		ret.CreatedAt = info.CreatedAt
		ret.UpdatedAt = info.UpdatedAt
		ret.ID = info.ID
		ret.EntID = info.EntID
		assert.Equal(t, info, &ret)
	}
}

func getAnnouncementUser(t *testing.T) {
	handler, err := NewHandler(
		context.Background(),
		handler1.WithEntID(&ret.EntID, true),
	)
	assert.Nil(t, err)

	info, err := handler.GetAnnouncementUser(context.Background())
	if assert.Nil(t, err) {
		assert.Equal(t, info, &ret)
	}
}

func getAnnouncementUsers(t *testing.T) {
	conds := &npool.Conds{
		AnnouncementID: &basetypes.StringVal{Op: cruder.EQ, Value: ret.AnnouncementID},
	}

	handler, err := NewHandler(
		context.Background(),
		WithConds(conds),
		handler1.WithOffset(0),
		handler1.WithLimit(0),
	)
	assert.Nil(t, err)

	infos, _, err := handler.GetAnnouncementUsers(context.Background())
	if !assert.Nil(t, err) {
		assert.NotEqual(t, len(infos), 0)
	}
}

func deleteAnnouncementUser(t *testing.T) {
	handler, err := NewHandler(
		context.Background(),
		handler1.WithID(&ret.ID, true),
	)
	assert.Nil(t, err)

	info, err := handler.DeleteAnnouncementUser(context.Background())
	if assert.Nil(t, err) {
		ret.UpdatedAt = info.UpdatedAt
		assert.Equal(t, info, &ret)
	}

	info, err = handler.GetAnnouncementUser(context.Background())
	assert.Nil(t, err)
	assert.Nil(t, info)
}

func TestAnnouncementUser(t *testing.T) {
	if runByGithubAction, err := strconv.ParseBool(os.Getenv("RUN_BY_GITHUB_ACTION")); err == nil && runByGithubAction {
		return
	}

	teardown := setupAnnouncementUser(t)
	defer teardown(t)

	t.Run("createAnnouncementUser", createAnnouncementUser)
	t.Run("getAnnouncementUser", getAnnouncementUser)
	t.Run("getAnnouncementUsers", getAnnouncementUsers)
	t.Run("deleteAnnouncementUser", deleteAnnouncementUser)
}
