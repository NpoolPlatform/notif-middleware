package email

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"testing"

	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"

	npool "github.com/NpoolPlatform/message/npool/notif/mw/v1/template/email"
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
	ret = npool.EmailTemplate{
		ID:                uuid.NewString(),
		AppID:             uuid.NewString(),
		LangID:            uuid.NewString(),
		UsedFor:           basetypes.UsedFor_AffiliatePurchase,
		Sender:            uuid.NewString(),
		ReplyTos:          []string{uuid.NewString(), uuid.NewString()},
		CCTos:             []string{uuid.NewString(), uuid.NewString()},
		Subject:           "subject " + uuid.NewString(),
		Body:              "body " + uuid.NewString(),
		DefaultToUsername: "DefaultToUsername " + uuid.NewString(),
	}
)

func creatEmailTemplate(t *testing.T) {
	handler, err := NewHandler(
		context.Background(),
		WithID(&ret.ID),
		WithAppID(&ret.AppID),
		WithLangID(&ret.LangID),
		WithUsedFor(&ret.UsedFor),
		WithSender(&ret.Sender),
		WithReplyTos(&ret.ReplyTos),
		WithSubject(&ret.Subject),
		WithBody(&ret.Body),
		WithDefaultToUsername(&ret.DefaultToUsername),
	)
	assert.Nil(t, err)

	info, err := handler.CreateEmailTemplate(context.Background())
	if assert.Nil(t, err) {
		assert.Equal(t, info, &ret)
	}
}

func updateEmailTemplate(t *testing.T) {
	ret.UsedFor = basetypes.UsedFor_Announcement
	ret.Sender = uuid.NewString()
	ret.ReplyTos = []string{"change1 " + uuid.NewString(), "change2 " + uuid.NewString()}
	ret.CCTos = []string{"change1 " + uuid.NewString(), "change2 " + uuid.NewString()}
	ret.Subject = "change Subject " + uuid.NewString()
	ret.Body = "change Body " + uuid.NewString()
	ret.DefaultToUsername = "change DefaultToUsername " + uuid.NewString()
	handler, err := NewHandler(
		context.Background(),
		WithID(&ret.ID),
		WithAppID(&ret.AppID),
		WithLangID(&ret.LangID),
		WithUsedFor(&ret.UsedFor),
		WithSender(&ret.Sender),
		WithReplyTos(&ret.ReplyTos),
		WithSubject(&ret.Subject),
		WithBody(&ret.Body),
		WithDefaultToUsername(&ret.DefaultToUsername),
	)
	assert.Nil(t, err)

	info, err := handler.UpdateEmailTemplate(context.Background())
	if assert.Nil(t, err) {
		assert.Equal(t, info, &ret)
	}
}

func getEmailTemplate(t *testing.T) {
	handler, err := NewHandler(
		context.Background(),
		WithID(&ret.ID),
	)
	assert.Nil(t, err)

	info, err := handler.GetEmailTemplate(context.Background())
	if assert.Nil(t, err) {
		assert.Equal(t, info, &ret)
	}
}

func getEmailTemplates(t *testing.T) {
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

	infos, _, err := handler.GetEmailTemplates(context.Background())
	if !assert.Nil(t, err) {
		assert.NotEqual(t, len(infos), 0)
	}
}

func deleteEmailTemplate(t *testing.T) {
	handler, err := NewHandler(
		context.Background(),
		WithID(&ret.ID),
	)
	assert.Nil(t, err)

	info, err := handler.DeleteEmailTemplate(context.Background())
	if assert.Nil(t, err) {
		assert.Equal(t, info, &ret)
	}

	info, err = handler.GetEmailTemplate(context.Background())
	assert.Nil(t, err)
	assert.Nil(t, info)
}

func TestEmailTemplate(t *testing.T) {
	if runByGithubAction, err := strconv.ParseBool(os.Getenv("RUN_BY_GITHUB_ACTION")); err == nil && runByGithubAction {
		return
	}

	t.Run("creatEmailTemplate", creatEmailTemplate)
	t.Run("updateEmailTemplate", updateEmailTemplate)
	t.Run("getEmailTemplate", getEmailTemplate)
	t.Run("getEmailTemplates", getEmailTemplates)
	t.Run("deleteEmailTemplate", deleteEmailTemplate)
}
