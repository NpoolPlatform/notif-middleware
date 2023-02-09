package user

import (
	"context"
	"fmt"

	"os"
	"strconv"
	"testing"
	"time"

	announcementcrud "github.com/NpoolPlatform/notif-manager/pkg/crud/announcement"
	userstatecrud "github.com/NpoolPlatform/notif-manager/pkg/crud/announcement/user"
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
	usermgrpb "github.com/NpoolPlatform/message/npool/notif/mgr/v1/announcement/user"
	channelpb "github.com/NpoolPlatform/message/npool/notif/mgr/v1/channel"
	npool "github.com/NpoolPlatform/message/npool/notif/mw/v1/announcement/user"
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
	aType = announcementmgrpb.AnnouncementType_AppointUsers
	data  = npool.User{
		ID:                  uuid.NewString(),
		AnnouncementID:      uuid.NewString(),
		AppID:               uuid.NewString(),
		UserID:              uuid.NewString(),
		Title:               uuid.NewString(),
		Content:             uuid.NewString(),
		ChannelsStr:         `["ChannelEmail", "ChannelSMS"]`,
		Channels:            []channelpb.NotifChannel{channelpb.NotifChannel_ChannelEmail, channelpb.NotifChannel_ChannelSMS},
		AnnouncementTypeStr: aType.String(),
		AnnouncementType:    aType,
	}
)

func getUsers(t *testing.T) {
	endAt := uint32(time.Now().Add(1 * time.Hour).Unix())
	_, err := announcementcrud.Create(context.Background(), &announcementmgrpb.AnnouncementReq{
		ID:               &data.AnnouncementID,
		AppID:            &data.AppID,
		Title:            &data.Title,
		Content:          &data.Content,
		Channels:         data.Channels,
		EndAt:            &endAt,
		AnnouncementType: &aType,
	})
	assert.Nil(t, err)

	_, err = userstatecrud.Create(context.Background(), &usermgrpb.UserReq{
		ID:             &data.ID,
		AppID:          &data.AppID,
		UserID:         &data.UserID,
		AnnouncementID: &data.AnnouncementID,
	})
	assert.Nil(t, err)

	infos, total, err := GetUsers(context.Background(), &usermgrpb.Conds{
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
	t.Run("getUsers", getUsers)
}
