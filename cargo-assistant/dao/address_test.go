package dao

import (
	"testing"
	"flag"
	"os"
	kitlog "github.com/go-kit/kit/log"
	_ "github.com/go-sql-driver/mysql"
	"github.com/astaxie/beego/orm"
	"fmt"
)

/**
*Author:hanyajun
*Time:17-12-15 下午2:54
*Discription:
**/
const (
	DEFAULT_MAX_IDLE_CONNS = 10
	DEFAULT_MAX_OPEN_CONNS = 100
)

func TestAddressDao_SelectAll(t *testing.T) {
	var (
		mysqlUrl      = flag.String("mysqlUrl", "10.2.4.113", "")
		mysqlPort     = flag.String("mysqlPort", "3306", "")
		mysqlUsername = flag.String("mysqlUsername", "root", "")
		mysqlPassword = flag.String("mysqlPassword", "sfai", "")
		mysqlDBName   = flag.String("mysqlDBName", "marathon", "")
	)
	flag.Parse()
	var logger kitlog.Logger
	logger = kitlog.NewJSONLogger(os.Stderr)
	//ca.NewMysqlManager(*mysqlUrl, *mysqlPort, *mysqlDBName, *mysqlUsername, *mysqlPassword, logger)
	orm.RegisterDriver("mysql", orm.DRMySQL)
	ds := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8", *mysqlUsername,*mysqlPassword, *mysqlUrl, *mysqlPort, *mysqlDBName)
	err := orm.RegisterDataBase("default", "mysql", ds, DEFAULT_MAX_IDLE_CONNS, DEFAULT_MAX_OPEN_CONNS)
	if err != nil {
		panic(err)
	}
	ad,_:=NewAddressDao(logger)
	addresses,_:=ad.SelectAll()
	logger.Log("address",addresses)
}
