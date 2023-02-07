package announcements

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
	valuedef "github.com/NpoolPlatform/message/npool"
	announcementmgrpb "github.com/NpoolPlatform/message/npool/notif/mgr/v1/announcement"
	readstatemgrpb "github.com/NpoolPlatform/message/npool/notif/mgr/v1/announcement/readstate"
	sendstatemgrpb "github.com/NpoolPlatform/message/npool/notif/mgr/v1/announcement/sendstate"
	channelpb "github.com/NpoolPlatform/message/npool/notif/mgr/v1/channel"
	npool "github.com/NpoolPlatform/message/npool/notif/mw/v1/announcement"
	announcementcrud "github.com/NpoolPlatform/notif-manager/pkg/crud/announcement"
	readstatecrud "github.com/NpoolPlatform/notif-manager/pkg/crud/announcement/readstate"
	sendstatecrud "github.com/NpoolPlatform/notif-manager/pkg/crud/announcement/sendstate"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

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
	userID = uuid.NewString()
	data   = npool.Announcement{
		AnnouncementID: uuid.NewString(),
		AppID:          uuid.NewString(),
		UserID:         userID,
		Title:          uuid.NewString(),
		Content:        uuid.NewString(),
		ChannelStr:     channelpb.NotifChannel_ChannelEmail.String(),
		Channel:        channelpb.NotifChannel_ChannelEmail,
		AlreadySend:    true,
		ReadUserID:     userID,
		AlreadyRead:    true,
	}
)

func getAnnouncements(t *testing.T) {
	endAt := uint32(time.Now().Add(1 * time.Hour).Unix())
	_, err := announcementcrud.Create(context.Background(), &announcementmgrpb.AnnouncementReq{
		ID:       &data.AnnouncementID,
		AppID:    &data.AppID,
		Title:    &data.Title,
		Content:  &data.Content,
		Channels: []channelpb.NotifChannel{channelpb.NotifChannel_ChannelEmail},
		EndAt:    &endAt,
	})
	assert.Nil(t, err)

	_, err = readstatecrud.Create(context.Background(), &readstatemgrpb.ReadStateReq{
		AppID:          &data.AppID,
		UserID:         &data.UserID,
		AnnouncementID: &data.AnnouncementID,
	})
	assert.Nil(t, err)

	_, err = sendstatecrud.Create(context.Background(), &sendstatemgrpb.SendStateReq{
		AppID:          &data.AppID,
		UserID:         &data.UserID,
		AnnouncementID: &data.AnnouncementID,
		Channel:        &data.Channel,
	})
	assert.Nil(t, err)

	infos, total, err := GetAnnouncements(context.Background(), &npool.Conds{
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
		data.CreatedAt = infos[0].CreatedAt
		data.UpdatedAt = infos[0].UpdatedAt
		assert.NotEqual(t, total, 0)
		assert.Equal(t, &data, infos[0])
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
	t.Run("getAnnouncements", getAnnouncements)
}
