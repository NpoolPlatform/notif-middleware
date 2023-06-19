package notif

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"testing"

	basetypes "github.com/NpoolPlatform/message/npool/basetypes/v1"

	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"

	"github.com/NpoolPlatform/go-service-framework/pkg/config"

	"bou.ke/monkey"
	grpc2 "github.com/NpoolPlatform/go-service-framework/pkg/grpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"github.com/NpoolPlatform/notif-middleware/pkg/testinit"

	npool "github.com/NpoolPlatform/message/npool/notif/mw/v1/notif"
	"github.com/stretchr/testify/assert"

	"github.com/google/uuid"
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
	ret = &npool.Notif{
		ID:           uuid.NewString(),
		AppID:        uuid.NewString(),
		UserID:       uuid.NewString(),
		Notified:     true,
		LangID:       uuid.NewString(),
		EventID:      uuid.NewString(),
		EventType:    basetypes.UsedFor_KYCApproved,
		EventTypeStr: basetypes.UsedFor_KYCApproved.String(),
		UseTemplate:  true,
		Title:        "Title " + uuid.NewString(),
		Content:      "Content " + uuid.NewString(),
		Channel:      basetypes.NotifChannel_ChannelSMS,
		ChannelStr:   basetypes.NotifChannel_ChannelSMS.String(),
		NotifType:    npool.NotifType_Multicast,
		NotifTypeStr: npool.NotifType_Multicast.String(),
	}

	retReq = &npool.NotifReq{
		ID:          &ret.ID,
		AppID:       &ret.AppID,
		UserID:      &ret.UserID,
		Notified:    &ret.Notified,
		LangID:      &ret.LangID,
		EventID:     &ret.EventID,
		EventType:   &ret.EventType,
		UseTemplate: &ret.UseTemplate,
		Title:       &ret.Title,
		Content:     &ret.Content,
		Channel:     &ret.Channel,
		NotifType:   &ret.NotifType,
	}
)

func createNotif(t *testing.T) {
	info, err := CreateNotif(context.Background(), retReq)
	if assert.Nil(t, err) {
		ret.CreatedAt = info.CreatedAt
		ret.UpdatedAt = info.UpdatedAt
		ret.Notified = info.Notified
		assert.Equal(t, ret, info)
	}
}

func updateNotif(t *testing.T) {
	ret.Notified = true
	info, err := UpdateNotif(context.Background(), retReq)
	if assert.Nil(t, err) {
		ret.CreatedAt = info.CreatedAt
		ret.UpdatedAt = info.UpdatedAt
		ret.Notified = info.Notified
		assert.Equal(t, ret, info)
	}
}

func updateNotifs(t *testing.T) {
	b := true
	infos, err := UpdateNotifs(context.Background(), []string{ret.ID}, b)
	if assert.Nil(t, err) {
		ret.CreatedAt = infos[0].CreatedAt
		ret.UpdatedAt = infos[0].UpdatedAt
		assert.Equal(t, infos[0].String(), ret.String())
	}
}

func getNotif(t *testing.T) {
	info, err := GetNotif(context.Background(), ret.ID)
	if assert.Nil(t, err) {
		assert.Equal(t, ret.String(), info.String())
	}
}

func getNotifs(t *testing.T) {
	infos, total, err := GetNotifs(context.Background(), &npool.Conds{
		ID: &basetypes.StringVal{
			Op:    cruder.EQ,
			Value: ret.ID,
		},
	}, 0, 1)
	if assert.Nil(t, err) {
		assert.Equal(t, total, uint32(1))
		assert.Equal(t, infos[0].String(), ret.String())
	}
}

func getNotifOnly(t *testing.T) {
	info, err := GetNotifOnly(context.Background(), &npool.Conds{
		ID: &basetypes.StringVal{
			Op:    cruder.EQ,
			Value: ret.ID,
		},
	})
	if assert.Nil(t, err) {
		assert.Equal(t, info.String(), ret.String())
	}
}
func TestClient(t *testing.T) {
	if runByGithubAction, err := strconv.ParseBool(os.Getenv("RUN_BY_GITHUB_ACTION")); err == nil && runByGithubAction {
		return
	}

	gport := config.GetIntValueWithNameSpace("", config.KeyGRPCPort)

	monkey.Patch(grpc2.GetGRPCConn, func(service string, tags ...string) (*grpc.ClientConn, error) {
		return grpc.Dial(fmt.Sprintf("localhost:%v", gport), grpc.WithTransportCredentials(insecure.NewCredentials()))
	})

	t.Run("createNotif", createNotif)
	t.Run("updateNotif", updateNotif)
	t.Run("updateNotifs", updateNotifs)
	t.Run("getNotif", getNotif)
	t.Run("getNotifs", getNotifs)
	t.Run("getNotifOnly", getNotifOnly)
}
