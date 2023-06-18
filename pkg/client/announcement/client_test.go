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

	appmwcli "github.com/NpoolPlatform/appuser-middleware/pkg/client/app"
	appusercli "github.com/NpoolPlatform/appuser-middleware/pkg/client/user"
	appmwpb "github.com/NpoolPlatform/message/npool/appuser/mw/v1/app"
	appuserpb "github.com/NpoolPlatform/message/npool/appuser/mw/v1/user"
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
	appID = uuid.NewString()
	ret   = appuserpb.User{
		ID:                uuid.NewString(),
		AppID:             appID,
		EmailAddress:      "aaa@aaa.aaa",
		PhoneNO:           "+8613612203133",
		ImportedFromAppID: uuid.NewString(),
		Username:          uuid.NewString(),
	}
	amt = npool.Announcement{
		AppID:               "",
		LangID:              uuid.NewString(),
		Title:               uuid.NewString(),
		Content:             uuid.NewString(),
		Channel:             basetypes.NotifChannel_ChannelEmail,
		ChannelStr:          basetypes.NotifChannel_ChannelEmail.String(),
		AnnouncementType:    npool.AnnouncementType_Multicast,
		AnnouncementTypeStr: npool.AnnouncementType_Multicast.String(),
		EndAt:               uint32(time.Now().Add(1 * time.Hour).Unix()),
	}
)

func setupAnnouncement(t *testing.T) func(*testing.T) {
	app1, err := appmwcli.CreateApp(
		context.Background(),
		&appmwpb.AppReq{
			ID:        &ret.AppID,
			CreatedBy: &ret.ID,
			Name:      &ret.AppID,
		},
	)
	assert.Nil(t, err)
	assert.NotNil(t, app1)

	amt.AppID = app1.ID

	var (
		id           = ret.ID
		appID        = ret.AppID
		passwordHash = uuid.NewString()
		req          = appuserpb.UserReq{
			ID:           &id,
			AppID:        &appID,
			EmailAddress: &ret.EmailAddress,
			PasswordHash: &passwordHash,
		}
	)

	info, err := appusercli.CreateUser(context.Background(), &req)
	assert.Nil(t, err)
	assert.NotNil(t, info)

	ret.ID = info.ID

	return func(*testing.T) {
		_, _ = appmwcli.DeleteApp(context.Background(), ret.AppID)
		_, _ = appusercli.DeleteUser(context.Background(), info.AppID, info.ID)
	}
}

func createAnnouncement(t *testing.T) {
	info, err := CreateAnnouncement(context.Background(), &npool.AnnouncementReq{
		AppID:            &amt.AppID,
		LangID:           &amt.LangID,
		Title:            &amt.Title,
		Content:          &amt.Content,
		Channel:          &amt.Channel,
		AnnouncementType: &amt.AnnouncementType,
		EndAt:            &amt.EndAt,
	})
	if assert.Nil(t, err) {
		amt.CreatedAt = info.CreatedAt
		amt.UpdatedAt = info.UpdatedAt
		amt.ID = info.ID
		assert.Equal(t, info, &amt)
	}
}

func updateAnnouncement(t *testing.T) {
	amt.Title = uuid.NewString()
	amt.Content = "-----" + uuid.NewString()
	amt.AnnouncementType = npool.AnnouncementType_Broadcast
	amt.AnnouncementTypeStr = npool.AnnouncementType_Broadcast.String()
	amt.EndAt = uint32(time.Now().Add(3 * time.Hour).Unix())
	info, err := UpdateAnnouncement(context.Background(), &npool.AnnouncementReq{
		ID:               &amt.ID,
		Title:            &amt.Title,
		Content:          &amt.Content,
		AnnouncementType: &amt.AnnouncementType,
		EndAt:            &amt.EndAt,
	})
	if assert.Nil(t, err) {
		amt.UpdatedAt = info.UpdatedAt
		assert.Equal(t, info, &amt)
	}
}

func getAnnouncement(t *testing.T) {
	info, err := GetAnnouncement(context.Background(), amt.ID)
	assert.Nil(t, err)
	assert.NotNil(t, info)
}

func getAnnouncements(t *testing.T) {
	infos, _, err := GetAnnouncements(context.Background(), &npool.Conds{
		ID: &basetypes.StringVal{
			Op:    cruder.EQ,
			Value: amt.ID,
		},
		Channel: &basetypes.Uint32Val{
			Op:    cruder.EQ,
			Value: uint32(basetypes.NotifChannel_value[amt.Channel.String()]),
		},
		EndAt: &basetypes.Uint32Val{
			Op:    cruder.GTE,
			Value: amt.EndAt,
		},
	}, 0, 1)
	if assert.Nil(t, err) {
		assert.NotEqual(t, len(infos), 0)
	}
}

func existAnnouncement(t *testing.T) {
	exist, err := ExistAnnouncement(context.Background(), amt.ID, amt.AppID)
	assert.Nil(t, err)
	assert.True(t, exist)
}

func deleteAnnouncement(t *testing.T) {
	info, err := DeleteAnnouncement(context.Background(), amt.ID)
	if assert.Nil(t, err) {
		assert.Equal(t, info, &amt)
	}
	info, err = GetAnnouncement(context.Background(), info.ID)
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
	t.Run("createAnnouncement", createAnnouncement)
	t.Run("updateAnnouncement", updateAnnouncement)
	t.Run("getAnnouncement", getAnnouncement)
	t.Run("getAnnouncements", getAnnouncements)
	t.Run("existAnnouncement", existAnnouncement)
	t.Run("deleteAnnouncement", deleteAnnouncement)

	patch.Unpatch()
}

