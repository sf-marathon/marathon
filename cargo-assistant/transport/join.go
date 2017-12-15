package transport

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/go-kit/kit/circuitbreaker"
	"github.com/go-kit/kit/log"
//	kitprometheus "github.com/go-kit/kit/metrics/prometheus"
	"github.com/go-kit/kit/ratelimit"
	kithttp "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
//	stdprometheus "github.com/prometheus/client_golang/prometheus"
	"github.com/sony/gobreaker"
	"golang.org/x/time/rate"
	ca "marathon/cargo-assistant"
	svc "marathon/cargo-assistant/service"
	"net/http"
	"strconv"
	"time"
//	"marathon/cargo-assistant/dao"
)

func MakeJoinHttpHandler(s svc.IJoinService,router *mux.Router, logger log.Logger) http.Handler {

	options := []kithttp.ServerOption{
		kithttp.ServerErrorLogger(logger),
	}
	//todo: prometheus

	getJoinEndpoint := ca.MakeGetJoinEndpoint(s)
	getJoinEndpoint = ratelimit.NewErroringLimiter(rate.NewLimiter(rate.Every(time.Nanosecond), 1))(getJoinEndpoint)
	getJoinEndpoint = circuitbreaker.Gobreaker(gobreaker.NewCircuitBreaker(gobreaker.Settings{}))(getJoinEndpoint)
	joinEndpoint := ca.MakeJoinEndpoint(s)
	joinEndpoint = ratelimit.NewErroringLimiter(rate.NewLimiter(rate.Every(time.Nanosecond), 1))(joinEndpoint)
	joinEndpoint = circuitbreaker.Gobreaker(gobreaker.NewCircuitBreaker(gobreaker.Settings{}))(joinEndpoint)

	router.Methods("GET").
		Path("/join/{group_id}").
		Handler(kithttp.NewServer(
			getJoinEndpoint,
			decodeGetJoinRequest,
			encodeResponse,
			options...,
		))
	router.Methods("POST").
		Path("/join").
		Handler(kithttp.NewServer(
			joinEndpoint,
			decodeJoinRequest,
			encodeResponse,
			options...,
		))

	return router
}

func decodeJoinRequest(_ context.Context, r *http.Request) (request interface{}, err error) {
	var joinReq ca.JoinRequest
	err = json.NewDecoder(r.Body).Decode(&joinReq)
	if err != nil {
		return nil, errors.New("param json err")
	}
	return joinReq, nil
}

func decodeGetJoinRequest(_ context.Context, r *http.Request) (request interface{}, err error) {
	vars := mux.Vars(r)
	id, ok := vars["group_id"]
	if !ok {
		return nil, errors.New("param err")
	}
	intId, err := strconv.Atoi(id)
	if err != nil {
		return nil, errors.New("param err")
	}
	return ca.GetJoinRequest{GroupId: intId}, nil
}
