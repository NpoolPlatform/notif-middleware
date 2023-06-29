package frontend

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"testing"

	"github.com/google/uuid"

	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"

	"github.com/NpoolPlatform/go-service-framework/pkg/config"

	"bou.ke/monkey"
	grpc2 "github.com/NpoolPlatform/go-service-framework/pkg/grpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"github.com/NpoolPlatform/notif-middleware/pkg/testinit"

	basetypes "github.com/NpoolPlatform/message/npool/basetypes/v1"

	"github.com/stretchr/testify/assert"

	npool "github.com/NpoolPlatform/message/npool/notif/mw/v1/template/frontend"
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
	ret = npool.FrontendTemplate{
		ID:         uuid.NewString(),
		AppID:      uuid.NewString(),
		LangID:     uuid.NewString(),
		UsedFor:    basetypes.UsedFor_KYCApproved,
		UsedForStr: basetypes.UsedFor_KYCApproved.String(),
		Title:      "Title " + uuid.NewString(),
		Content:    "Content " + uuid.NewString(),
	}

	appInfo = npool.FrontendTemplateReq{
		ID:      &ret.ID,
		AppID:   &ret.AppID,
		LangID:  &ret.LangID,
		UsedFor: &ret.UsedFor,
		Title:   &ret.Title,
		Content: &ret.Content,
	}
)

var info *npool.FrontendTemplate

func createFrontendTemplate(t *testing.T) {
	var err error
	info, err = CreateFrontendTemplate(context.Background(), &appInfo)
	if assert.Nil(t, err) {
		ret.UpdatedAt = info.UpdatedAt
		ret.CreatedAt = info.CreatedAt
		assert.Equal(t, info, &ret)
	}
}

func createFrontendTemplates(t *testing.T) {
	rets := []npool.FrontendTemplate{
		{
			ID:         uuid.NewString(),
			AppID:      ret.AppID,
			LangID:     uuid.NewString(),
			UsedFor:    basetypes.UsedFor_KYCRejected,
			UsedForStr: basetypes.UsedFor_KYCRejected.String(),
			Title:      "Title1 " + uuid.NewString(),
			Content:    "Content1 " + uuid.NewString(),
		},
		{
			ID:         uuid.NewString(),
			AppID:      ret.AppID,
			LangID:     uuid.NewString(),
			UsedFor:    basetypes.UsedFor_WithdrawalCompleted,
			UsedForStr: basetypes.UsedFor_WithdrawalCompleted.String(),
			Title:      "Title1 " + uuid.NewString(),
			Content:    "Content1 " + uuid.NewString(),
		},
	}

	apps := []*npool.FrontendTemplateReq{}
	for key := range rets {
		apps = append(apps, &npool.FrontendTemplateReq{
			ID:      &rets[key].ID,
			AppID:   &rets[key].AppID,
			LangID:  &rets[key].LangID,
			UsedFor: &rets[key].UsedFor,
			Title:   &rets[key].Title,
			Content: &rets[key].Content,
		})
	}

	infos, err := CreateFrontendTemplates(context.Background(), apps)
	if assert.Nil(t, err) {
		assert.Equal(t, len(infos), 2)
	}
}

func updateFrontendTemplate(t *testing.T) {
	var err error
	info, err = UpdateFrontendTemplate(context.Background(), &appInfo)
	if assert.Nil(t, err) {
		ret.UpdatedAt = info.UpdatedAt
		assert.Equal(t, info, &ret)
	}
}

func getFrontendTemplate(t *testing.T) {
	var err error
	info, err = GetFrontendTemplate(context.Background(), info.ID)
	if assert.Nil(t, err) {
		assert.Equal(t, info, &ret)
	}
}

func getFrontendTemplates(t *testing.T) {
	infos, total, err := GetFrontendTemplates(context.Background(),
		&npool.Conds{
			ID: &basetypes.StringVal{
				Value: info.ID,
				Op:    cruder.EQ,
			},
		}, 0, 1)
	if assert.Nil(t, err) {
		assert.Equal(t, total, uint32(1))
		assert.Equal(t, infos[0], &ret)
	}
}

func getFrontendTemplateOnly(t *testing.T) {
	var err error
	info, err = GetFrontendTemplateOnly(context.Background(),
		&npool.Conds{
			ID: &basetypes.StringVal{
				Value: info.ID,
				Op:    cruder.EQ,
			},
		})
	if assert.Nil(t, err) {
		ret.UpdatedAt = info.UpdatedAt
		ret.CreatedAt = info.CreatedAt
		assert.Equal(t, info, &ret)
	}
}

func existFrontendTemplate(t *testing.T) {
	exist, err := ExistFrontendTemplate(context.Background(), info.ID)
	if assert.Nil(t, err) {
		assert.Equal(t, exist, true)
	}
}

func existFrontendTemplateConds(t *testing.T) {
	exist, err := ExistFrontendTemplateConds(context.Background(),
		&npool.Conds{
			ID: &basetypes.StringVal{
				Value: info.ID,
				Op:    cruder.EQ,
			},
		},
	)
	if assert.Nil(t, err) {
		assert.Equal(t, exist, true)
	}
}

func deleteFrontendTemplate(t *testing.T) {
	info, err := DeleteFrontendTemplate(context.Background(), &npool.FrontendTemplateReq{
		ID: &ret.ID,
	})
	if assert.Nil(t, err) {
		assert.Equal(t, info, &ret)
	}
}

func TestMainOrder(t *testing.T) {
	if runByGithubAction, err := strconv.ParseBool(os.Getenv("RUN_BY_GITHUB_ACTION")); err == nil && runByGithubAction {
		return
	}

	gport := config.GetIntValueWithNameSpace("", config.KeyGRPCPort)

	monkey.Patch(grpc2.GetGRPCConn, func(service string, tags ...string) (*grpc.ClientConn, error) {
		return grpc.Dial(fmt.Sprintf("localhost:%v", gport), grpc.WithTransportCredentials(insecure.NewCredentials()))
	})

	t.Run("createFrontendTemplate", createFrontendTemplate)
	t.Run("createFrontendTemplates", createFrontendTemplates)
	t.Run("getFrontendTemplate", getFrontendTemplate)
	t.Run("getFrontendTemplates", getFrontendTemplates)
	t.Run("getFrontendTemplateOnly", getFrontendTemplateOnly)
	t.Run("updateFrontendTemplate", updateFrontendTemplate)
	t.Run("existFrontendTemplate", existFrontendTemplate)
	t.Run("existFrontendTemplateConds", existFrontendTemplateConds)
	t.Run("delete", deleteFrontendTemplate)
}
