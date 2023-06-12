package sendstate

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"testing"
	"time"

	"github.com/google/uuid"

	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	valuedef "github.com/NpoolPlatform/message/npool"

	"github.com/NpoolPlatform/go-service-framework/pkg/config"

	"bou.ke/monkey"
	grpc2 "github.com/NpoolPlatform/go-service-framework/pkg/grpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"github.com/NpoolPlatform/notif-middleware/pkg/testinit"

	announcementmgrpb "github.com/NpoolPlatform/message/npool/notif/mgr/v1/announcement"
	channelpb "github.com/NpoolPlatform/message/npool/notif/mgr/v1/channel"
	npool "github.com/NpoolPlatform/message/npool/notif/mw/v1/announcement/send"
	"github.com/stretchr/testify/assert"

	announcementcrud "github.com/NpoolPlatform/notif-manager/pkg/crud/announcement"
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
	aType = announcementmgrpb.AnnouncementType_Multicast
	data  = npool.SendState{
		AnnouncementID:      uuid.NewString(),
		AppID:               uuid.NewString(),
		UserID:              uuid.NewString(),
		Title:               uuid.NewString(),
		Content:             uuid.NewString(),
		ChannelStr:          channelpb.NotifChannel_ChannelEmail.String(),
		Channel:             channelpb.NotifChannel_ChannelEmail,
		AnnouncementTypeStr: aType.String(),
		AnnouncementType:    aType,
	}
)

func getSendStates(t *testing.T) {
	endAt := uint32(time.Now().Add(1 * time.Hour).Unix())
	channel1 := channelpb.NotifChannel_ChannelEmail

	_, err := announcementcrud.Create(context.Background(), &announcementmgrpb.AnnouncementReq{
		ID:               &data.AnnouncementID,
		AppID:            &data.AppID,
		Title:            &data.Title,
		Content:          &data.Content,
		Channel:          &channel1,
		EndAt:            &endAt,
		AnnouncementType: &aType,
	})
	assert.Nil(t, err)

	err = CreateSendState(context.Background(), data.AppID, data.UserID, data.AnnouncementID, data.Channel)
	assert.Nil(t, err)

	infos, total, err := GetSendStates(context.Background(), &npool.Conds{
		AppID: &valuedef.StringVal{
			Op:    cruder.EQ,
			Value: data.AppID,
		},
		UserID: &valuedef.StringVal{
			Op:    cruder.EQ,
			Value: data.UserID,
		},
		AnnouncementID: &valuedef.StringVal{
			Op:    cruder.EQ,
			Value: data.AnnouncementID,
		},
		Channel: &valuedef.Uint32Val{
			Op:    cruder.EQ,
			Value: uint32(channelpb.NotifChannel_ChannelEmail),
		},
	}, 0, 1)
	if assert.Nil(t, err) {
		assert.Equal(t, total, uint32(1))
		data.CreatedAt = infos[0].CreatedAt
		data.UpdatedAt = infos[0].UpdatedAt
		assert.Equal(t, infos[0], &data)
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
	t.Run("getSendStates", getSendStates)
}
