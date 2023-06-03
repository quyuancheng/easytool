package db

import (
	"os"
	"time"
	"tool/cmd/tool/config"
	"tool/pkg/models"

	"github.com/sirupsen/logrus"
)

// Inst db instance
var Inst *Instance

// GetDbInst returns the db instance.
func GetDbInst() *Instance {
	if Inst != nil {
		return Inst
	}

	// create db instance
	Inst, err := NewDBInstance(&InstanceConfig{
		DBType:               MysqlDB,
		User:                 config.C.DB.User,
		Passwd:               config.C.DB.Pass,
		Net:                  "tcp",
		Addr:                 config.C.DB.Addr,
		DBName:               config.C.DB.Name,
		AllowNativePasswords: true,
		ParseTime:            true,
		Loc:                  time.Local,
		Models:               models.GetModels(),
		TablePrefix:          config.C.DB.TablePrefix,
	})
	if err != nil {
		logrus.Errorf("create db instance failed: %s", err.Error())
		os.Exit(1)
	}
	return Inst
}
