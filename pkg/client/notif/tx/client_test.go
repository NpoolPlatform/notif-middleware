package tx

// import (
// 	"context"
// 	"fmt"
// 	"os"
// 	"strconv"
// 	"testing"

// 	"github.com/NpoolPlatform/go-service-framework/pkg/config"
// 	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"

// 	"bou.ke/monkey"
// 	grpc2 "github.com/NpoolPlatform/go-service-framework/pkg/grpc"
// 	"google.golang.org/grpc"
// 	"google.golang.org/grpc/credentials/insecure"

// 	"github.com/NpoolPlatform/notif-middleware/pkg/testinit"

// 	basetypes "github.com/NpoolPlatform/message/npool/basetypes/v1"
// 	npool "github.com/NpoolPlatform/message/npool/notif/mw/v1/notif/tx"
// 	"github.com/stretchr/testify/assert"

// 	"github.com/google/uuid"
// )

// func init() {
// 	if runByGithubAction, err := strconv.ParseBool(os.Getenv("RUN_BY_GITHUB_ACTION")); err == nil && runByGithubAction {
// 		return
// 	}
// 	if err := testinit.Init(); err != nil {
// 		fmt.Printf("cannot init test stub: %v\n", err)
// 	}
// }

// var (
// 	ret = npool.Tx{
// 		ID:            uuid.NewString(),
// 		TxID:          uuid.NewString(),
// 		NotifState:    npool.TxState_WaitNotified,
// 		NotifStateStr: npool.TxState_WaitNotified.String(),
// 		TxType:        basetypes.TxType_TxWithdraw,
// 		TxTypeStr:     basetypes.TxType_TxWithdraw.String(),
// 	}
// )

// func createTx(t *testing.T) {
// 	info, err := CreateTx(context.Background(), &npool.TxReq{
// 		TxID:       &ret.TxID,
// 		NotifState: &ret.NotifState,
// 		TxType:     &ret.TxType,
// 	})
// 	if assert.Nil(t, err) {
// 		ret.CreatedAt = info.CreatedAt
// 		ret.UpdatedAt = info.UpdatedAt
// 		ret.ID = info.ID
// 		assert.Equal(t, info, &ret)
// 	}
// }

// func updateTx(t *testing.T) {
// 	ret.NotifState = npool.TxState_Notified
// 	ret.NotifStateStr = npool.TxState_Notified.String()
// 	info, err := UpdateTx(context.Background(), &npool.TxReq{
// 		ID:         &ret.ID,
// 		NotifState: &ret.NotifState,
// 	})
// 	if assert.Nil(t, err) {
// 		ret.UpdatedAt = info.UpdatedAt
// 		assert.Equal(t, info, &ret)
// 	}
// }

// func getTxs(t *testing.T) {
// 	infos, _, err := GetTxs(context.Background(), &npool.Conds{
// 		ID: &basetypes.StringVal{
// 			Op:    cruder.EQ,
// 			Value: ret.ID,
// 		},
// 		TxID: &basetypes.StringVal{
// 			Op:    cruder.EQ,
// 			Value: ret.TxID,
// 		},
// 		NotifState: &basetypes.Uint32Val{
// 			Op:    cruder.EQ,
// 			Value: uint32(ret.NotifState),
// 		},
// 		TxType: &basetypes.Uint32Val{
// 			Op:    cruder.EQ,
// 			Value: uint32(ret.TxType),
// 		},
// 	}, 0, 1)
// 	if assert.Nil(t, err) {
// 		assert.NotEqual(t, len(infos), 0)
// 	}
// }

// func getTxOnly(t *testing.T) {
// 	info, err := GetTxOnly(context.Background(), &npool.Conds{
// 		ID: &basetypes.StringVal{
// 			Op:    cruder.EQ,
// 			Value: ret.ID,
// 		},
// 		NotifState: &basetypes.Uint32Val{
// 			Op:    cruder.EQ,
// 			Value: uint32(ret.NotifState),
// 		},
// 	})
// 	if assert.Nil(t, err) {
// 		assert.Equal(t, info, &ret)
// 	}
// }

// func TestClient(t *testing.T) {
// 	if runByGithubAction, err := strconv.ParseBool(os.Getenv("RUN_BY_GITHUB_ACTION")); err == nil && runByGithubAction {
// 		return
// 	}

// 	gport := config.GetIntValueWithNameSpace("", config.KeyGRPCPort)

// 	monkey.Patch(grpc2.GetGRPCConn, func(service string, tags ...string) (*grpc.ClientConn, error) {
// 		return grpc.Dial(fmt.Sprintf("localhost:%v", gport), grpc.WithTransportCredentials(insecure.NewCredentials()))
// 	})

// 	t.Run("createTx", createTx)
// 	t.Run("updateTx", updateTx)
// 	t.Run("getTxs", getTxs)
// 	t.Run("getTxOnly", getTxOnly)
// }
