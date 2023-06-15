package channel

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"testing"

	basetypes "github.com/NpoolPlatform/message/npool/basetypes/v1"

	"github.com/NpoolPlatform/go-service-framework/pkg/config"
	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"

	"bou.ke/monkey"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"github.com/NpoolPlatform/notif-middleware/pkg/testinit"

	npool "github.com/NpoolPlatform/message/npool/notif/mw/v1/notif/channel"
	"github.com/stretchr/testify/assert"

	appmwcli "github.com/NpoolPlatform/appuser-middleware/pkg/client/app"
	grpc2 "github.com/NpoolPlatform/go-service-framework/pkg/grpc"
	appmwpb "github.com/NpoolPlatform/message/npool/appuser/mw/v1/app"
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
	appID = uuid.NewString()
	ret   = npool.Channel{
		AppID:        appID,
		EventType:    basetypes.UsedFor_KYCApproved,
		EventTypeStr: basetypes.UsedFor_KYCApproved.String(),
		Channel:      basetypes.NotifChannel_ChannelEmail,
		ChannelStr:   basetypes.NotifChannel_ChannelEmail.String(),
	}
)

func setupChannel(t *testing.T) func(*testing.T) {
	app1, err := appmwcli.CreateApp(
		context.Background(),
		&appmwpb.AppReq{
			ID:        &appID,
			CreatedBy: &appID,
			Name:      &appID,
		},
	)
	assert.Nil(t, err)
	assert.NotNil(t, app1)

	ret.AppID = app1.ID

	return func(*testing.T) {
		_, _ = appmwcli.DeleteApp(context.Background(), ret.AppID)
	}
}

func createChannel(t *testing.T) {
	info, err := CreateChannel(context.Background(), &npool.ChannelReq{
		AppID:     &ret.AppID,
		EventType: &ret.EventType,
		Channel:   &ret.Channel,
	})
	if assert.Nil(t, err) {
		ret.CreatedAt = info.CreatedAt
		ret.UpdatedAt = info.UpdatedAt
		ret.ID = info.ID
		assert.Equal(t, info, &ret)
	}
}

func existChannelConds(t *testing.T) {
	exist, err := ExistChannelConds(context.Background(), &npool.ExistChannelCondsRequest{
		Conds: &npool.Conds{
			AppID: &basetypes.StringVal{
				Op:    cruder.EQ,
				Value: ret.AppID,
			},
			EventType: &basetypes.Uint32Val{
				Op:    cruder.EQ,
				Value: uint32(ret.EventType),
			},
			Channel: &basetypes.Uint32Val{
				Op:    cruder.EQ,
				Value: uint32(ret.Channel),
			},
		},
	})
	if assert.Nil(t, err) {
		assert.True(t, exist)
	}
}

func getChannel(t *testing.T) {
	info, err := GetChannel(context.Background(), ret.ID)
	assert.Nil(t, err)
	assert.NotNil(t, info)
}

func getChannels(t *testing.T) {
	infos, _, err := GetChannels(context.Background(), &npool.Conds{
		AppID: &basetypes.StringVal{
			Op:    cruder.EQ,
			Value: ret.AppID,
		},
		EventType: &basetypes.Uint32Val{
			Op:    cruder.EQ,
			Value: uint32(ret.EventType),
		},
		Channel: &basetypes.Uint32Val{
			Op:    cruder.EQ,
			Value: uint32(ret.Channel),
		},
	}, 0, 1)
	if assert.Nil(t, err) {
		assert.NotEqual(t, len(infos), 0)
	}
}

func deleteChannel(t *testing.T) {
	info, err := DeleteChannel(context.Background(), ret.ID)
	if assert.Nil(t, err) {
		assert.Equal(t, info, &ret)
	}
	info, err = GetChannel(context.Background(), info.ID)
	assert.Nil(t, err)
	assert.Nil(t, info)
}

func TestClient(t *testing.T) {
	if runByGithubAction, err := strconv.ParseBool(os.Getenv("RUN_BY_GITHUB_ACTION")); err == nil && runByGithubAction {
		return
	}

	gport := config.GetIntValueWithNameSpace("", config.KeyGRPCPort)

	teardown := setupChannel(t)
	defer teardown(t)

	patch := monkey.Patch(grpc2.GetGRPCConn, func(service string, tags ...string) (*grpc.ClientConn, error) {
		return grpc.Dial(fmt.Sprintf("localhost:%v", gport), grpc.WithTransportCredentials(insecure.NewCredentials()))
	})

	t.Run("getChannels", createChannel)
	t.Run("existChannelConds", existChannelConds)
	t.Run("getChannel", getChannel)
	t.Run("getChannels", getChannels)
	t.Run("deleteChannel", deleteChannel)

	patch.Unpatch()
}
