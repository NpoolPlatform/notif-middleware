package email

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

	npool "github.com/NpoolPlatform/message/npool/notif/mw/v1/template/email"

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
	langID1 = uuid.NewString()
	langID2 = uuid.NewString()
	ret     = npool.EmailTemplate{
		ID:                uuid.NewString(),
		AppID:             uuid.NewString(),
		LangID:            uuid.NewString(),
		UsedFor:           basetypes.UsedFor_KYCApproved,
		UsedForStr:        basetypes.UsedFor_KYCApproved.String(),
		Sender:            "Sender " + uuid.NewString(),
		ReplyTos:          []string{uuid.NewString(), uuid.NewString()},
		ReplyTosStr:       "",
		CCTos:             []string{uuid.NewString(), uuid.NewString()},
		CCTosStr:          "",
		Subject:           "Subject " + uuid.NewString(),
		Body:              "Body " + uuid.NewString(),
		DefaultToUsername: "DefaultToUsername " + uuid.NewString(),
	}

	appInfo = npool.EmailTemplateReq{
		ID:                &ret.ID,
		AppID:             &ret.AppID,
		LangID:            &ret.LangID,
		UsedFor:           &ret.UsedFor,
		Sender:            &ret.Sender,
		ReplyTos:          ret.ReplyTos,
		CCTos:             ret.CCTos,
		Subject:           &ret.Subject,
		Body:              &ret.Body,
		DefaultToUsername: &ret.DefaultToUsername,
	}

	info *npool.EmailTemplate
)

func createEmailTemplate(t *testing.T) {
	var err error
	info, err = CreateEmailTemplate(context.Background(), &appInfo)
	if assert.Nil(t, err) {
		info.ReplyTosStr = ret.ReplyTosStr
		info.CCTosStr = ret.CCTosStr
		ret.CreatedAt = info.CreatedAt
		ret.UpdatedAt = info.UpdatedAt
		assert.Equal(t, info, &ret)
	}
}

func createEmailTemplates(t *testing.T) {
	rets := []npool.EmailTemplate{
		{
			ID:                uuid.NewString(),
			AppID:             ret.AppID,
			LangID:            langID1,
			UsedFor:           basetypes.UsedFor_KYCApproved,
			UsedForStr:        basetypes.UsedFor_KYCApproved.String(),
			Sender:            "Sender1 " + uuid.NewString(),
			ReplyTos:          []string{uuid.NewString()},
			CCTos:             []string{uuid.NewString()},
			Subject:           "Subject1 " + uuid.NewString(),
			Body:              "Body1 " + uuid.NewString(),
			DefaultToUsername: "DefaultToUsername1 " + uuid.NewString(),
		},
		{
			ID:                uuid.NewString(),
			AppID:             ret.AppID,
			LangID:            langID2,
			UsedFor:           basetypes.UsedFor_Signin,
			UsedForStr:        basetypes.UsedFor_Signin.String(),
			Sender:            "Sender2 " + uuid.NewString(),
			ReplyTos:          []string{uuid.NewString()},
			ReplyTosStr:       "",
			CCTos:             []string{uuid.NewString()},
			CCTosStr:          "",
			Subject:           "Subject2 " + uuid.NewString(),
			Body:              "Body2 " + uuid.NewString(),
			DefaultToUsername: "DefaultToUsername2 " + uuid.NewString(),
		},
	}

	apps := []*npool.EmailTemplateReq{}
	for key := range rets {
		apps = append(apps, &npool.EmailTemplateReq{
			ID:                &rets[key].ID,
			AppID:             &rets[key].AppID,
			LangID:            &rets[key].LangID,
			UsedFor:           &rets[key].UsedFor,
			Sender:            &rets[key].Sender,
			ReplyTos:          rets[key].ReplyTos,
			CCTos:             rets[key].CCTos,
			Subject:           &rets[key].Subject,
			Body:              &rets[key].Body,
			DefaultToUsername: &rets[key].DefaultToUsername,
		})
	}

	infos, err := CreateEmailTemplates(context.Background(), apps)
	if assert.Nil(t, err) {
		assert.Equal(t, len(infos), 2)
	}
}

func updateEmailTemplate(t *testing.T) {
	var err error
	info, err = UpdateEmailTemplate(context.Background(), &appInfo)
	if assert.Nil(t, err) {
		info.ReplyTosStr = ret.ReplyTosStr
		info.CCTosStr = ret.CCTosStr
		assert.Equal(t, info, &ret)
	}
}

func getEmailTemplate(t *testing.T) {
	var err error
	info, err = GetEmailTemplate(context.Background(), info.ID)
	if assert.Nil(t, err) {
		info.ReplyTosStr = ret.ReplyTosStr
		info.CCTosStr = ret.CCTosStr
		assert.Equal(t, info, &ret)
	}
}

func getEmailTemplates(t *testing.T) {
	infos, total, err := GetEmailTemplates(context.Background(),
		&npool.Conds{
			ID: &basetypes.StringVal{
				Value: info.ID,
				Op:    cruder.EQ,
			},
		}, 0, 1)
	if assert.Nil(t, err) {
		assert.Equal(t, total, uint32(1))
		infos[0].ReplyTosStr = ret.ReplyTosStr
		infos[0].CCTosStr = ret.CCTosStr
		assert.Equal(t, infos[0], &ret)
	}
}

func getEmailTemplateOnly(t *testing.T) {
	var err error
	info, err = GetEmailTemplateOnly(context.Background(),
		&npool.Conds{
			ID: &basetypes.StringVal{
				Value: info.ID,
				Op:    cruder.EQ,
			},
		})
	if assert.Nil(t, err) {
		info.ReplyTosStr = ret.ReplyTosStr
		info.CCTosStr = ret.CCTosStr
		assert.Equal(t, info, &ret)
	}
}

func existEmailTemplate(t *testing.T) {
	exist, err := ExistEmailTemplate(context.Background(), info.ID)
	if assert.Nil(t, err) {
		assert.Equal(t, exist, true)
	}
}

func existEmailTemplateConds(t *testing.T) {
	exist, err := ExistEmailTemplateConds(context.Background(),
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

func deleteEmailTemplate(t *testing.T) {
	info, err := DeleteEmailTemplate(context.Background(), &npool.EmailTemplateReq{
		ID: &ret.ID,
	})
	if assert.Nil(t, err) {
		info.ReplyTosStr = ret.ReplyTosStr
		info.CCTosStr = ret.CCTosStr
		assert.Equal(t, info, &ret)
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
	t.Run("createEmailTemplate", createEmailTemplate)
	t.Run("createEmailTemplates", createEmailTemplates)
	t.Run("getEmailTemplate", getEmailTemplate)
	t.Run("getEmailTemplates", getEmailTemplates)
	t.Run("getEmailTemplateOnly", getEmailTemplateOnly)
	t.Run("updateEmailTemplate", updateEmailTemplate)
	t.Run("existEmailTemplate", existEmailTemplate)
	t.Run("existEmailTemplateConds", existEmailTemplateConds)
	t.Run("delete", deleteEmailTemplate)
}
