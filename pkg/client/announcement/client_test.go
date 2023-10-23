package announcement

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"testing"
	"time"

	"bou.ke/monkey"
	"github.com/NpoolPlatform/go-service-framework/pkg/config"
	grpc2 "github.com/NpoolPlatform/go-service-framework/pkg/grpc"

	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"

	npool "github.com/NpoolPlatform/message/npool/notif/mw/v1/announcement"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

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
	ret = npool.Announcement{
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
)

func setupAnnouncement(t *testing.T) func(*testing.T) {
	return func(*testing.T) {}
}

func createAnnouncement(t *testing.T) {
	info, err := CreateAnnouncement(context.Background(), &npool.AnnouncementReq{
		AppID:            &ret.AppID,
		LangID:           &ret.LangID,
		Title:            &ret.Title,
		Content:          &ret.Content,
		Channel:          &ret.Channel,
		AnnouncementType: &ret.AnnouncementType,
		StartAt:          &ret.StartAt,
		EndAt:            &ret.EndAt,
	})
	if assert.Nil(t, err) {
		ret.CreatedAt = info.CreatedAt
		ret.UpdatedAt = info.UpdatedAt
		ret.ID = info.ID
		ret.EntID = info.EntID
		assert.Equal(t, info, &ret)
	}
}

func updateAnnouncement(t *testing.T) {
	ret.Title = uuid.NewString()
	ret.Content = "-----" + uuid.NewString()
	ret.AnnouncementType = basetypes.NotifType_NotifBroadcast
	ret.AnnouncementTypeStr = basetypes.NotifType_NotifBroadcast.String()
	ret.StartAt = uint32(time.Now().Add(3 * time.Hour).Unix())
	ret.EndAt = uint32(time.Now().Add(10 * time.Hour).Unix())
	info, err := UpdateAnnouncement(context.Background(), &npool.AnnouncementReq{
		ID:               &ret.ID,
		Title:            &ret.Title,
		Content:          &ret.Content,
		AnnouncementType: &ret.AnnouncementType,
		StartAt:          &ret.StartAt,
		EndAt:            &ret.EndAt,
	})
	if assert.Nil(t, err) {
		ret.UpdatedAt = info.UpdatedAt
		assert.Equal(t, info, &ret)
	}
}

func getAnnouncement(t *testing.T) {
	info, err := GetAnnouncement(context.Background(), ret.EntID)
	assert.Nil(t, err)
	assert.NotNil(t, info)
}

func getAnnouncements(t *testing.T) {
	infos, _, err := GetAnnouncements(context.Background(), &npool.Conds{
		EntID:   &basetypes.StringVal{Op: cruder.EQ, Value: ret.EntID},
		Channel: &basetypes.Uint32Val{Op: cruder.EQ, Value: uint32(basetypes.NotifChannel_value[ret.Channel.String()])},
		StartAt: &basetypes.Uint32Val{Op: cruder.LTE, Value: ret.StartAt},
		EndAt:   &basetypes.Uint32Val{Op: cruder.GTE, Value: ret.StartAt},
	}, 0, 1)
	if assert.Nil(t, err) {
		assert.NotEqual(t, len(infos), 0)
	}
}

func existAnnouncement(t *testing.T) {
	exist, err := ExistAnnouncement(context.Background(), ret.EntID)
	assert.Nil(t, err)
	assert.True(t, exist)
}

func deleteAnnouncement(t *testing.T) {
	info, err := DeleteAnnouncement(context.Background(), ret.ID)
	if assert.Nil(t, err) {
		assert.Equal(t, info, &ret)
	}
	info, err = GetAnnouncement(context.Background(), info.EntID)
	assert.Nil(t, err)
	assert.Nil(t, info)
}

func TestClient(t *testing.T) {
	if runByGithubAction, err := strconv.ParseBool(os.Getenv("RUN_BY_GITHUB_ACTION")); err == nil && runByGithubAction {
		return
	}

	gport := config.GetIntValueWithNameSpace("", config.KeyGRPCPort)

	teardown := setupAnnouncement(t)
	defer teardown(t)

	patch := monkey.Patch(grpc2.GetGRPCConn, func(service string, tags ...string) (*grpc.ClientConn, error) {
		return grpc.Dial(fmt.Sprintf("localhost:%v", gport), grpc.WithTransportCredentials(insecure.NewCredentials()))
	})
	monkey.Patch(grpc2.GetGRPCConnV1, func(service string, recvMsgBytes int, tags ...string) (*grpc.ClientConn, error) {
		return grpc.Dial(fmt.Sprintf("localhost:%v", gport), grpc.WithTransportCredentials(insecure.NewCredentials()))
	})

	t.Run("createAnnouncement", createAnnouncement)
	t.Run("updateAnnouncement", updateAnnouncement)
	t.Run("getAnnouncement", getAnnouncement)
	t.Run("getAnnouncements", getAnnouncements)
	t.Run("existAnnouncement", existAnnouncement)
	t.Run("deleteAnnouncement", deleteAnnouncement)

	patch.Unpatch()
}
