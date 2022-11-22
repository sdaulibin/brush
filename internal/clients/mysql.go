package clients

import (
	"binginx.com/brush/config"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"time"
)

var WriteDBCli *gorm.DB
var ReadDBCli *gorm.DB
var err error

func InitDBClients() {
	WriteDBCli, err = GetSqlDriver(config.GlobalConfig.WriteDB)
	if err != nil {
		panic(err)
	}
	ReadDBCli, err = GetSqlDriver(config.GlobalConfig.ReadDB)
	if err != nil {
		panic(err)
	}
	if config.GlobalConfig.DebugMode {
		WriteDBCli = WriteDBCli.Debug()
		ReadDBCli = ReadDBCli.Debug()
	}
}

func GetSqlDriver(dbConf config.DBConfig) (*gorm.DB, error) {
	var dbDialector = getDbDialector(dbConf)
	gormDb, err := gorm.Open(dbDialector, &gorm.Config{
		SkipDefaultTransaction: true,
		PrepareStmt:            true,
	})
	if err != nil {
		return nil, err
	}
	if rawDb, err := gormDb.DB(); err != nil {
		return nil, err
	} else {
		rawDb.SetConnMaxIdleTime(time.Second * 30)
		rawDb.SetMaxIdleConns(dbConf.MaxIdleConns)
		rawDb.SetMaxOpenConns(dbConf.MaxOpenConns)
		return gormDb, nil
	}
}

func getDbDialector(conf config.DBConfig) gorm.Dialector {
	var dbDialector gorm.Dialector
	dsn := getDsn(conf)
	dbDialector = mysql.Open(dsn)
	return dbDialector
}

func getDsn(dbConf config.DBConfig) string {
	Host := dbConf.Host
	DataBase := dbConf.Name
	Port := dbConf.Port
	User := dbConf.User
	Pass := dbConf.Password
	Charset := "utf8mb4"
	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=True&loc=Local", User, Pass, Host, Port, DataBase, Charset)
}
