package goodbenefit

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"testing"
	"time"

	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"

	npool "github.com/NpoolPlatform/message/npool/notif/mw/v1/notif/goodbenefit"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"

	basetypes "github.com/NpoolPlatform/message/npool/basetypes/v1"
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
	yesterday = uint32(time.Now().AddDate(0, 0, -1).Unix())
	ret       = npool.GoodBenefit{
		GoodID:      uuid.NewString(),
		GoodName:    uuid.NewString(),
		Amount:      "100",
		State:       basetypes.Result_Success,
		StateStr:    basetypes.Result_Success.String(),
		Message:     uuid.NewString(),
		BenefitDate: yesterday,
		TxID:        uuid.NewString(),
		Generated:   false,
	}

	ret2 = npool.GoodBenefit{
		GoodID:      uuid.NewString(),
		GoodName:    uuid.NewString(),
		Amount:      "10",
		State:       basetypes.Result_Success,
		StateStr:    basetypes.Result_Success.String(),
		Message:     uuid.NewString(),
		BenefitDate: uint32(time.Now().Add(-3 * time.Minute).Unix()),
		TxID:        uuid.NewString(),
		Generated:   false,
	}
)

func createGoodBenefit(t *testing.T) {
	handler, err := NewHandler(
		context.Background(),
		WithGoodID(&ret.GoodID),
		WithGoodName(&ret.GoodName),
		WithAmount(&ret.Amount),
		WithState(&ret.State),
		WithMessage(&ret.Message),
		WithBenefitDate(&ret.BenefitDate),
		WithTxID(&ret.TxID),
		WithGenerated(&ret.Generated),
	)
	assert.Nil(t, err)

	info, err := handler.CreateGoodBenefit(context.Background())
	if assert.Nil(t, err) {
		ret.CreatedAt = info.CreatedAt
		ret.UpdatedAt = info.UpdatedAt
		ret.ID = info.ID
		assert.Equal(t, info, &ret)
	}

	handler2, err := NewHandler(
		context.Background(),
		WithGoodID(&ret2.GoodID),
		WithGoodName(&ret2.GoodName),
		WithAmount(&ret2.Amount),
		WithState(&ret2.State),
		WithMessage(&ret2.Message),
		WithBenefitDate(&ret2.BenefitDate),
		WithTxID(&ret2.TxID),
		WithGenerated(&ret2.Generated),
	)
	assert.Nil(t, err)

	_info, err := handler2.CreateGoodBenefit(context.Background())
	if assert.Nil(t, err) {
		ret2.CreatedAt = _info.CreatedAt
		ret2.UpdatedAt = _info.UpdatedAt
		ret2.ID = _info.ID
		assert.Equal(t, _info, &ret2)
	}
}

func updateGoodBenefit(t *testing.T) {
	ret.Generated = true
	handler, err := NewHandler(
		context.Background(),
		WithID(&ret.ID),
		WithGenerated(&ret.Generated),
	)
	assert.Nil(t, err)

	info, err := handler.UpdateGoodBenefit(context.Background())
	if assert.Nil(t, err) {
		ret.UpdatedAt = info.UpdatedAt
		assert.Equal(t, info, &ret)
	}
}

func getGoodBenefit(t *testing.T) {
	handler, err := NewHandler(
		context.Background(),
		WithID(&ret.ID),
	)
	assert.Nil(t, err)

	info, err := handler.GetGoodBenefit(context.Background())
	if assert.Nil(t, err) {
		assert.Equal(t, info, &ret)
	}
}

func getGoodBenefits(t *testing.T) {
	conds := &npool.Conds{
		Generated:        &basetypes.BoolVal{Op: cruder.EQ, Value: ret.Generated},
		BenefitDateStart: &basetypes.Uint32Val{Op: cruder.LTE, Value: yesterday},
		BenefitDateEnd:   &basetypes.Uint32Val{Op: cruder.GTE, Value: uint32(time.Now().Unix())},
	}

	handler, err := NewHandler(
		context.Background(),
		WithConds(conds),
		WithOffset(0),
		WithLimit(0),
	)
	assert.Nil(t, err)

	infos, _, err := handler.GetGoodBenefits(context.Background())
	if !assert.Nil(t, err) {
		assert.NotEqual(t, len(infos), 0)
	}
}

func deleteGoodBenefit(t *testing.T) {
	handler, err := NewHandler(
		context.Background(),
		WithID(&ret.ID),
	)
	assert.Nil(t, err)

	info, err := handler.DeleteGoodBenefit(context.Background())
	if assert.Nil(t, err) {
		ret.UpdatedAt = info.UpdatedAt
		assert.Equal(t, info, &ret)
	}

	info, err = handler.GetGoodBenefit(context.Background())
	assert.Nil(t, err)
	assert.Nil(t, info)

	handler2, err := NewHandler(
		context.Background(),
		WithID(&ret2.ID),
	)
	assert.Nil(t, err)

	_info, err := handler2.DeleteGoodBenefit(context.Background())
	if assert.Nil(t, err) {
		ret2.UpdatedAt = _info.UpdatedAt
		assert.Equal(t, _info, &ret2)
	}

	_info, err = handler2.GetGoodBenefit(context.Background())
	assert.Nil(t, err)
	assert.Nil(t, _info)
}

func TestGoodBenefit(t *testing.T) {
	if runByGithubAction, err := strconv.ParseBool(os.Getenv("RUN_BY_GITHUB_ACTION")); err == nil && runByGithubAction {
		return
	}

	t.Run("createGoodBenefit", createGoodBenefit)
	t.Run("updateGoodBenefit", updateGoodBenefit)
	t.Run("getGoodBenefit", getGoodBenefit)
	t.Run("getGoodBenefits", getGoodBenefits)
	t.Run("deleteGoodBenefit", deleteGoodBenefit)
}
