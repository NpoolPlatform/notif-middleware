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

	grpc2 "github.com/NpoolPlatform/go-service-framework/pkg/grpc"
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
	ret = npool.Channel{
		AppID:        uuid.NewString(),
		EventType:    basetypes.UsedFor_KYCApproved,
		EventTypeStr: basetypes.UsedFor_KYCApproved.String(),
		Channel:      basetypes.NotifChannel_ChannelEmail,
		ChannelStr:   basetypes.NotifChannel_ChannelEmail.String(),
	}
)

func setupChannel(t *testing.T) func(*testing.T) {
	return func(*testing.T) {}
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
		ret.EntID = info.EntID
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
	info, err := GetChannel(context.Background(), ret.AppID, ret.EntID)
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
	info, err = GetChannel(context.Background(), info.AppID, info.EntID)
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
	monkey.Patch(grpc2.GetGRPCConnV1, func(service string, recvMsgBytes int, tags ...string) (*grpc.ClientConn, error) {
		return grpc.Dial(fmt.Sprintf("localhost:%v", gport), grpc.WithTransportCredentials(insecure.NewCredentials()))
	})

	t.Run("createChannels", createChannel)
	t.Run("getChannel", getChannel)
	t.Run("existChannelConds", existChannelConds)
	t.Run("getChannels", getChannels)
	t.Run("deleteChannel", deleteChannel)

	patch.Unpatch()
}
