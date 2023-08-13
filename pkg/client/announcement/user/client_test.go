package user

import (
	"context"
	"fmt"

	"os"
	"strconv"
	"testing"
	"time"

	"github.com/google/uuid"

	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"

	"github.com/NpoolPlatform/go-service-framework/pkg/config"

	"bou.ke/monkey"
	grpc2 "github.com/NpoolPlatform/go-service-framework/pkg/grpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"github.com/NpoolPlatform/notif-middleware/pkg/testinit"

	basetypes "github.com/NpoolPlatform/message/npool/basetypes/v1"
	announcementpb "github.com/NpoolPlatform/message/npool/notif/mw/v1/announcement"
	npool "github.com/NpoolPlatform/message/npool/notif/mw/v1/announcement/user"
	announcement1 "github.com/NpoolPlatform/notif-middleware/pkg/mw/announcement"
	"github.com/stretchr/testify/assert"
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
		announcement1.WithTitle(&announcement.Title),
		announcement1.WithContent(&announcement.Content),
		announcement1.WithAppID(&announcement.AppID),
		announcement1.WithLangID(&announcement.LangID),
		announcement1.WithChannel(&announcement.Channel),
		announcement1.WithAnnouncementType(&announcement.AnnouncementType),
		announcement1.WithStartAt(&announcement.StartAt),
		announcement1.WithEndAt(&announcement.EndAt),
	)
	assert.Nil(t, err)

	_announcement, err := handler.CreateAnnouncement(context.Background())
	assert.Nil(t, err)
	assert.NotNil(t, _announcement)

	ret.AnnouncementID = _announcement.ID

	_id, err := uuid.Parse(_announcement.ID)
	assert.Nil(t, err)
	handler.ID = &_id

	return func(*testing.T) {
		_, _ = handler.DeleteAnnouncement(context.Background())
	}
}

func createAnnouncementUser(t *testing.T) {
	info, err := CreateAnnouncementUser(context.Background(), &npool.AnnouncementUserReq{
		AppID:          &ret.AppID,
		UserID:         &ret.UserID,
		AnnouncementID: &ret.AnnouncementID,
	})
	if assert.Nil(t, err) {
		ret.CreatedAt = info.CreatedAt
		ret.UpdatedAt = info.UpdatedAt
		ret.ID = info.ID
		assert.Equal(t, info, &ret)
	}
}

func getAnnouncementUser(t *testing.T) {
	info, err := GetAnnouncementUser(context.Background(), ret.AppID, ret.ID)
	assert.Nil(t, err)
	assert.NotNil(t, info)
}

func getAnnouncementUsers(t *testing.T) {
	infos, _, err := GetAnnouncementUsers(context.Background(), &npool.Conds{
		ID: &basetypes.StringVal{
			Op:    cruder.EQ,
			Value: ret.ID,
		},
	}, 0, 1)
	if assert.Nil(t, err) {
		assert.NotEqual(t, len(infos), 0)
	}
}

func deleteAnnouncementUser(t *testing.T) {
	info, err := DeleteAnnouncementUser(context.Background(), ret.ID)
	if assert.Nil(t, err) {
		assert.Equal(t, info, &ret)
	}
	info, err = GetAnnouncementUser(context.Background(), info.AppID, info.ID)
	assert.Nil(t, err)
	assert.Nil(t, info)
}

func TestClient(t *testing.T) {
	if runByGithubAction, err := strconv.ParseBool(os.Getenv("RUN_BY_GITHUB_ACTION")); err == nil && runByGithubAction {
		return
	}

	gport := config.GetIntValueWithNameSpace("", config.KeyGRPCPort)

	teardown := setupAnnouncementUser(t)
	defer teardown(t)

	patch := monkey.Patch(grpc2.GetGRPCConn, func(service string, tags ...string) (*grpc.ClientConn, error) {
		return grpc.Dial(fmt.Sprintf("localhost:%v", gport), grpc.WithTransportCredentials(insecure.NewCredentials()))
	})
	monkey.Patch(grpc2.GetGRPCConnV1, func(service string, recvMsgBytes int, tags ...string) (*grpc.ClientConn, error) {
		return grpc.Dial(fmt.Sprintf("localhost:%v", gport), grpc.WithTransportCredentials(insecure.NewCredentials()))
	})

	t.Run("createAnnouncementUser", createAnnouncementUser)
	t.Run("getAnnouncementUser", getAnnouncementUser)
	t.Run("getAnnouncementUsers", getAnnouncementUsers)
	t.Run("deleteAnnouncementUser", deleteAnnouncementUser)

	patch.Unpatch()
}
