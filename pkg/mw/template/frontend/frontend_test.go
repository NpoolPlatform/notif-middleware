package frontend

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"testing"

	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"

	npool "github.com/NpoolPlatform/message/npool/notif/mw/v1/template/frontend"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"

	"github.com/NpoolPlatform/g11n-middleware/pkg/testinit"
	basetypes "github.com/NpoolPlatform/message/npool/basetypes/v1"
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
		ID:      uuid.NewString(),
		AppID:   uuid.NewString(),
		LangID:  uuid.NewString(),
		UsedFor: basetypes.UsedFor_AffiliatePurchase,
		Title:   "title " + uuid.NewString(),
		Content: "content " + uuid.NewString(),
	}
)

func creatFrontendTemplate(t *testing.T) {
	handler, err := NewHandler(
		context.Background(),
		WithID(&ret.ID),
		WithAppID(&ret.AppID),
		WithLangID(&ret.LangID),
		WithUsedFor(&ret.UsedFor),
		WithTitle(&ret.Title),
		WithContent(&ret.Content),
	)
	assert.Nil(t, err)

	info, err := handler.CreateFrontendTemplate(context.Background())
	if assert.Nil(t, err) {
		assert.Equal(t, info, &ret)
	}
}

func updateFrontendTemplate(t *testing.T) {
	ret.UsedFor = basetypes.UsedFor_Announcement
	ret.Title = "change Title " + uuid.NewString()
	ret.Content = "change Content " + uuid.NewString()
	handler, err := NewHandler(
		context.Background(),
		WithID(&ret.ID),
		WithAppID(&ret.AppID),
		WithLangID(&ret.LangID),
		WithUsedFor(&ret.UsedFor),
		WithTitle(&ret.Title),
		WithContent(&ret.Content),
	)
	assert.Nil(t, err)

	info, err := handler.UpdateFrontendTemplate(context.Background())
	if assert.Nil(t, err) {
		assert.Equal(t, info, &ret)
	}
}

func getFrontendTemplate(t *testing.T) {
	handler, err := NewHandler(
		context.Background(),
		WithID(&ret.ID),
	)
	assert.Nil(t, err)

	info, err := handler.GetFrontendTemplate(context.Background())
	if assert.Nil(t, err) {
		assert.Equal(t, info, &ret)
	}
}

func getFrontendTemplates(t *testing.T) {
	conds := &npool.Conds{
		ID: &basetypes.StringVal{Op: cruder.EQ, Value: ret.ID},
	}

	handler, err := NewHandler(
		context.Background(),
		WithConds(conds),
		WithOffset(0),
		WithLimit(0),
	)
	assert.Nil(t, err)

	infos, _, err := handler.GetFrontendTemplates(context.Background())
	if !assert.Nil(t, err) {
		assert.NotEqual(t, len(infos), 0)
	}
}

func deleteFrontendTemplate(t *testing.T) {
	handler, err := NewHandler(
		context.Background(),
		WithID(&ret.ID),
	)
	assert.Nil(t, err)

	info, err := handler.DeleteFrontendTemplate(context.Background())
	if assert.Nil(t, err) {
		assert.Equal(t, info, &ret)
	}

	info, err = handler.GetFrontendTemplate(context.Background())
	assert.Nil(t, err)
	assert.Nil(t, info)
}

func TestFrontendTemplate(t *testing.T) {
	if runByGithubAction, err := strconv.ParseBool(os.Getenv("RUN_BY_GITHUB_ACTION")); err == nil && runByGithubAction {
		return
	}

	t.Run("creatFrontendTemplate", creatFrontendTemplate)
	t.Run("updateFrontendTemplate", updateFrontendTemplate)
	t.Run("getFrontendTemplate", getFrontendTemplate)
	t.Run("getFrontendTemplates", getFrontendTemplates)
	t.Run("deleteFrontendTemplate", deleteFrontendTemplate)
}
