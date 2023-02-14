package notifchannel

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"testing"

	"github.com/NpoolPlatform/message/npool/notif/mgr/v1/channel"
	"github.com/NpoolPlatform/message/npool/third/mgr/v1/usedfor"
	crud "github.com/NpoolPlatform/notif-manager/pkg/crud/notif/notifchannel"

	"github.com/NpoolPlatform/go-service-framework/pkg/config"
	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	valuedef "github.com/NpoolPlatform/message/npool"

	"bou.ke/monkey"
	grpc2 "github.com/NpoolPlatform/go-service-framework/pkg/grpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"github.com/NpoolPlatform/notif-middleware/pkg/testinit"

	mgrpb "github.com/NpoolPlatform/message/npool/notif/mgr/v1/notif/notifchannel"
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
	eventType = usedfor.UsedFor_KYCApproved
	channel1  = channel.NotifChannel_ChannelFrontend
	data      = &mgrpb.NotifChannel{
		ID:        uuid.NewString(),
		AppID:     uuid.NewString(),
		EventType: eventType,
		Channel:   channel1,
	}
)

var dataReq = &mgrpb.NotifChannelReq{
	ID:        &data.ID,
	AppID:     &data.AppID,
	EventType: &data.EventType,
	Channel:   &data.Channel,
}

func getNotifChannels(t *testing.T) {
	_, err := crud.Create(context.Background(), dataReq)
	assert.Nil(t, err)

	infos, total, err := GetNotifChannels(context.Background(), &mgrpb.Conds{
		ID: &valuedef.StringVal{
			Op:    cruder.EQ,
			Value: data.ID,
		},
	}, 0, 1)
	if assert.Nil(t, err) {
		data.CreatedAt = infos[0].CreatedAt
		data.UpdatedAt = infos[0].UpdatedAt
		assert.Equal(t, total, uint32(1))
		assert.Equal(t, infos[0].String(), data.String())
	}
}

func getNotifChannelOnly(t *testing.T) {
	info, err := GetNotifChannelOnly(context.Background(), &mgrpb.Conds{
		ID: &valuedef.StringVal{
			Op:    cruder.EQ,
			Value: data.ID,
		},
	})
	if assert.Nil(t, err) {
		data.CreatedAt = info.CreatedAt
		data.UpdatedAt = info.UpdatedAt
		assert.Equal(t, info.String(), data.String())
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

	t.Run("getNotifChannels", getNotifChannels)
	t.Run("getNotifChannelOnly", getNotifChannelOnly)
}