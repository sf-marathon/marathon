package main

import (
	"flag"
	"fmt"
	ca "marathon/cargo-assistant"
	"marathon/cargo-assistant/dao"
	svc "marathon/cargo-assistant/service"
	tp "marathon/cargo-assistant/transport"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"syscall"

	kitlog "github.com/go-kit/kit/log"
	"github.com/golang/glog"

	"marathon/redispool"

	"github.com/gorilla/mux"
)

func main() {
	var (
		mysqlUrl      = flag.String("mysqlUrl", "10.2.4.113", "")
		RedisRawUrl   = flag.String("redisrawurl", "redis://:@10.2.4.113:6379/0", "redis raw url")
		mysqlPort     = flag.String("mysqlPort", "3306", "")
		mysqlUsername = flag.String("mysqlUsername", "root", "")
		mysqlPassword = flag.String("mysqlPassword", "sfai", "")
		mysqlDBName   = flag.String("mysqlDBName", "marathon", "")
		httpAddr      = flag.String("addr", ":8081", "The address of listen and serve")
	)
	flag.Parse()
	var logger kitlog.Logger
	logger = kitlog.NewJSONLogger(os.Stderr)
	var groupDao dao.IGroupDao
	var proMktBaseDao dao.IProMarketBaseDao
	var joinDao dao.IJoinDao
	var addDao dao.IAddressDao
	var groupService svc.IGroupService
	var joinService svc.IJoinService
	var addService svc.IAddressService
	errs := make(chan error)
	var err error

	//init redispool
	redisPool := redispool.NewRedispool(*RedisRawUrl)
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
	addDao, err = dao.NewAddressDao(logger, redisPool)
	if err != nil {
		errs <- err
	}

	groupService = svc.NewGroupService(groupDao, proMktBaseDao)
	joinService = svc.NewJoinService(joinDao, groupDao, proMktBaseDao)
	addService = svc.NewAddressService(addDao)

	groupService = ca.NewLoggingMiddleware(logger, groupService)

	route := mux.NewRouter()
	route = route.PathPrefix("/ca").Subrouter()
	tp.MakeHttpHandler(groupService, route, logger)
	tp.MakeJoinHttpHandler(joinService, route, logger)
	tp.MakeAddressHttpHandler(addService, route, logger)
	go func() {
		c := make(chan os.Signal)
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
		errs <- fmt.Errorf("%s", <-c)
	}()
	go func() {
		logger.Log("transport", "HTTP", "addr", *httpAddr)
		errs <- http.ListenAndServe(*httpAddr, allowCORS(route))
	}()
	logger.Log("exit:", <-errs)
}

// allowCORS allows Cross Origin Resoruce Sharing from any origin.
// Don't do this without consideration in production systems.
func allowCORS(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if origin := r.Header.Get("Origin"); origin != "" {
			w.Header().Set("Access-Control-Allow-Origin", origin)
			if r.Method == "OPTIONS" && r.Header.Get("Access-Control-Request-Method") != "" {
				preflightHandler(w, r)
				return
			}
		}
		h.ServeHTTP(w, r)
	})
}

func preflightHandler(w http.ResponseWriter, r *http.Request) {
	headers := []string{"Content-Type", "Accept"}
	w.Header().Set("Access-Control-Allow-Headers", strings.Join(headers, ","))
	methods := []string{"GET", "HEAD", "POST", "PUT", "DELETE"}
	w.Header().Set("Access-Control-Allow-Methods", strings.Join(methods, ","))
	glog.Infof("preflight request for %s", r.URL.Path)
	return
}
