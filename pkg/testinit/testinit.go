package testinit

import (
	"context"
	"fmt"
	"path"
	"runtime"

	"github.com/NpoolPlatform/go-service-framework/pkg/app"

	"github.com/NpoolPlatform/notif-middleware/pkg/db"

	migrator "github.com/NpoolPlatform/notif-middleware/pkg/migrator"
	servicename "github.com/NpoolPlatform/notif-middleware/pkg/servicename"

	mysqlconst "github.com/NpoolPlatform/go-service-framework/pkg/mysql/const"
	rabbitmqconst "github.com/NpoolPlatform/go-service-framework/pkg/rabbitmq/const"
	redisconst "github.com/NpoolPlatform/go-service-framework/pkg/redis/const"
)

func Init() error {
	_, myPath, _, ok := runtime.Caller(0)
	if !ok {
		return fmt.Errorf("cannot get source file path")
	}

	appName := path.Base(path.Dir(path.Dir(path.Dir(myPath))))
	configPath := fmt.Sprintf("%s/../../cmd/%v", path.Dir(myPath), appName)

	err := app.Init(
		servicename.ServiceName,
		"",
		"",
		"",
		configPath,
		nil,
		nil,
		mysqlconst.MysqlServiceName,
		rabbitmqconst.RabbitMQServiceName,
		redisconst.RedisServiceName,
	)
	if err != nil {
		return fmt.Errorf("cannot init app stub: %v", err)
	}
	if err := migrator.Migrate(context.Background()); err != nil {
		panic(fmt.Errorf("fail migrate db: %v", err))
	}
	err = db.Init()
	if err != nil {
		return fmt.Errorf("cannot init database: %v", err)
	}

	return nil
}
