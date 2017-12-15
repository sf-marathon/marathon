package main

import (
	"context"
	"flag"
	ca "marathon/cargo-assistant"
	dao "marathon/cargo-assistant/dao"
	"os"
	"time"

	kitlog "github.com/go-kit/kit/log"
)

func main() {
	var (
		mysqlUrl      = flag.String("mysqlUrl", "10.2.4.113", "")
		mysqlPort     = flag.String("mysqlPort", "3306", "")
		mysqlUsername = flag.String("mysqlUsername", "root", "")
		mysqlPassword = flag.String("mysqlPassword", "sfai", "")
		mysqlDBName   = flag.String("mysqlDBName", "marathon", "")
		checkinterval = flag.Duration("checkinterval", time.Second, "")
	)
	flag.Parse()
	var logger kitlog.Logger
	logger = kitlog.NewJSONLogger(os.Stderr)
	var groupDao dao.IGroupDao
	var proMktBaseDao dao.IProMarketBaseDao
	var joinDao dao.IJoinDao
	errs := make(chan error)
	var err error
	//init DB
	ca.NewMysqlManager(*mysqlUrl, *mysqlPort, *mysqlDBName, *mysqlUsername, *mysqlPassword, logger)
	groupDao, err = dao.NewGroupDao(logger)
	if err != nil {
		errs <- err
	}
	proMktBaseDao, err = dao.NewProMarketBaseDao(logger)
	if err != nil {
		errs <- err
	}
	joinDao, err = dao.NewJoinDao(logger)
	if err != nil {
		errs <- err
	}

	grpMgr := NewGroupManager(groupDao, joinDao, proMktBaseDao, *checkinterval, logger)

	grpMgr.Run(context.Background())

}
