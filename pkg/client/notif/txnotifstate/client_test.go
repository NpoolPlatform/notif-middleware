package txnotifstate

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"testing"

	"github.com/NpoolPlatform/go-service-framework/pkg/config"
	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	valuedef "github.com/NpoolPlatform/message/npool"

	"bou.ke/monkey"
	grpc2 "github.com/NpoolPlatform/go-service-framework/pkg/grpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"github.com/NpoolPlatform/notif-middleware/pkg/testinit"

	mgrpb "github.com/NpoolPlatform/message/npool/notif/mgr/v1/notif/txnotifstate"
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
	txState = mgrpb.TxState_WaitSend
	txType  = mgrpb.TxType_Withdraw
	data    = &mgrpb.TxNotifState{
		ID:         uuid.NewString(),
		TxID:       uuid.NewString(),
		NotifState: txState,
		NotifType:  txType,
	}
)

var dataReq = &mgrpb.TxNotifStateReq{
	ID:         &data.ID,
	TxID:       &data.TxID,
	NotifState: &data.NotifState,
	NotifType:  &data.NotifType,
}

func createTxNotifState(t *testing.T) {
	info, err := CreateTxNotifState(context.Background(), dataReq)
	if assert.Nil(t, err) {
		data.CreatedAt = info.CreatedAt
		data.UpdatedAt = info.UpdatedAt
		assert.Equal(t, data, info)
	}
}

func updateTxNotifState(t *testing.T) {
	info, err := UpdateTxNotifState(context.Background(), dataReq)
	if assert.Nil(t, err) {
		data.CreatedAt = info.CreatedAt
		data.UpdatedAt = info.UpdatedAt
		assert.Equal(t, data, info)
	}
}

func getTxNotifStates(t *testing.T) {
	infos, total, err := GetTxNotifStates(context.Background(), &mgrpb.Conds{
		ID: &valuedef.StringVal{
			Op:    cruder.EQ,
			Value: data.ID,
		},
	}, 0, 1)
	if assert.Nil(t, err) {
		assert.Equal(t, total, uint32(1))
		assert.Equal(t, infos[0].String(), data.String())
	}
}

func getTxNotifStateOnly(t *testing.T) {
	info, err := GetTxNotifStateOnly(context.Background(), &mgrpb.Conds{
		ID: &valuedef.StringVal{
			Op:    cruder.EQ,
			Value: data.ID,
		},
	})
	if assert.Nil(t, err) {
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

	t.Run("createTxNotifState", createTxNotifState)
	t.Run("updateTxNotifState", updateTxNotifState)
	t.Run("getTxNotifStates", getTxNotifStates)
	t.Run("getTxNotifStateOnly", getTxNotifStateOnly)
}
