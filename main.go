package main

import (
	"flag"
	"fmt"
	kitlog "github.com/go-kit/kit/log"
	"net/http"
	"os"
	"os/signal"
	order "marathon/cargo-assistant"
	"syscall"
)

func main() {
	var (
		mysqlUrl      = flag.String("mysqlUrl", "localhost", "")
		mysqlPort     = flag.String("mysqlPort", "3306", "")
		mysqlUsername = flag.String("mysqlUsername", "root", "")
		mysqlPassword = flag.String("mysqlPassword", "123456", "")
		mysqlDBName   = flag.String("mysqlDBName", "marathon", "")
		httpAddr      = flag.String("addr", ":8088", "The address of listen and serve")
	)
	flag.Parse()
	var logger kitlog.Logger
	logger = kitlog.NewJSONLogger(os.Stderr)
	var orderDao order.IOrderDao
	var orderService order.IOrderService
	errs := make(chan error)
	var err error
	//init DB
	order.NewMysqlManager(*mysqlUrl, *mysqlPort, *mysqlDBName, *mysqlUsername, *mysqlPassword,logger)
	//orderDao, err = order.NewOrderDao(logger)
	if err != nil {
		errs <- err
	}
	orderService = order.NewOrderService(orderDao)
	orderService=order.NewLoggingMiddleware(logger,orderService)
	httpHandler := order.MakeHttpHandler(orderService, logger)
	go func() {
		c := make(chan os.Signal)
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
		errs <- fmt.Errorf("%s", <-c)
	}()
	go func() {
		logger.Log("transport", "HTTP", "addr", *httpAddr)
		errs <- http.ListenAndServe(*httpAddr, httpHandler)
	}()
	logger.Log("exit:", <-errs)
}
