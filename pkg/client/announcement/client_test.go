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
	readstatemgrpb "github.com/NpoolPlatform/message/npool/notif/mgr/v1/announcement/readstate"

	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	valuedef "github.com/NpoolPlatform/message/npool"

	channelpb "github.com/NpoolPlatform/message/npool/notif/mgr/v1/channel"
	npool "github.com/NpoolPlatform/message/npool/notif/mw/v1/announcement"
	announcementcrud "github.com/NpoolPlatform/notif-manager/pkg/crud/announcement"
	readstatecrud "github.com/NpoolPlatform/notif-manager/pkg/crud/announcement/readstate"

	userstatemgrpb "github.com/NpoolPlatform/message/npool/notif/mgr/v1/announcement/user"
	userstatecrud "github.com/NpoolPlatform/notif-manager/pkg/crud/announcement/user"

	mgrpb "github.com/NpoolPlatform/message/npool/notif/mgr/v1/announcement"

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
	endAt  = uint32(time.Now().Add(1 * time.Hour).Unix())
	userID = uuid.NewString()
	aType  = mgrpb.AnnouncementType_AllUsers
	data   = npool.Announcement{
		AnnouncementID:      uuid.NewString(),
		AppID:               uuid.NewString(),
		LangID:              uuid.NewString(),
		Title:               uuid.NewString(),
		Content:             uuid.NewString(),
		AlreadyRead:         true,
		EndAt:               endAt,
		AnnouncementTypeStr: aType.String(),
		AnnouncementType:    aType,
		ReadUserID:          userID,
	}
)

func getAnnouncementStates(t *testing.T) {
	_, err := announcementcrud.Create(context.Background(), &mgrpb.AnnouncementReq{
		ID:               &data.AnnouncementID,
		AppID:            &data.AppID,
		LangID:           &data.LangID,
		Title:            &data.Title,
		Content:          &data.Content,
		Channels:         []channelpb.NotifChannel{channelpb.NotifChannel_ChannelEmail},
		EndAt:            &endAt,
		AnnouncementType: &aType,
	})
	assert.Nil(t, err)

	userID1 := uuid.NewString()
	aType1 := mgrpb.AnnouncementType_AppointUsers
	_, err = announcementcrud.Create(context.Background(), &mgrpb.AnnouncementReq{
		ID:               &userID,
		AppID:            &data.AppID,
		LangID:           &data.LangID,
		Title:            &data.Title,
		Content:          &data.Content,
		Channels:         []channelpb.NotifChannel{channelpb.NotifChannel_ChannelEmail},
		EndAt:            &endAt,
		AnnouncementType: &aType1,
	})
	assert.Nil(t, err)

	_, err = readstatecrud.Create(context.Background(), &readstatemgrpb.ReadStateReq{
		AppID:          &data.AppID,
		UserID:         &userID,
		AnnouncementID: &data.AnnouncementID,
	})
	assert.Nil(t, err)

	_, err = readstatecrud.Create(context.Background(), &readstatemgrpb.ReadStateReq{
		AppID:          &data.AppID,
		UserID:         &userID,
		AnnouncementID: &data.AnnouncementID,
	})
	assert.Nil(t, err)

	_, err = userstatecrud.Create(context.Background(), &userstatemgrpb.UserReq{
		AppID:          &data.AppID,
		UserID:         &userID1,
		AnnouncementID: &data.AnnouncementID,
	})
	assert.Nil(t, err)

	infos, total, err := GetAnnouncementStates(context.Background(), &npool.Conds{
		AppID: &valuedef.StringVal{
			Op:    cruder.EQ,
			Value: data.AppID,
		},
		UserID: &valuedef.StringVal{
			Op:    cruder.EQ,
			Value: userID,
		},
		LangID: &valuedef.StringVal{
			Op:    cruder.EQ,
			Value: data.LangID,
		},
	}, 0, 10)
	if assert.Nil(t, err) {
		data.CreatedAt = infos[0].CreatedAt
		data.UpdatedAt = infos[0].UpdatedAt
		assert.NotEqual(t, total, 1)
		assert.Equal(t, &data, infos[0])
	}

	_, total, err = GetAnnouncementStates(context.Background(), &npool.Conds{
		AppID: &valuedef.StringVal{
			Op:    cruder.EQ,
			Value: data.AppID,
		},
		UserID: &valuedef.StringVal{
			Op:    cruder.EQ,
			Value: userID,
		},
		LangID: &valuedef.StringVal{
			Op:    cruder.EQ,
			Value: data.LangID,
		},
	}, 0, 10)
	if assert.Nil(t, err) {
		assert.NotEqual(t, total, 2)
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
	t.Run("getAnnouncementStates", getAnnouncementStates)
}
