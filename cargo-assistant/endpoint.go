package cargo_assistant

import (
	"context"
	"errors"
	"github.com/go-kit/kit/endpoint"
	"time"
	svc "marathon/cargo-assistant/service"

	"fmt"
	"marathon/cargo-assistant/dao"
)

var (
	ERROR_TYPE_ASSERTION = errors.New("type assertion error!")
)


type GetGroupRequest struct {

}

type GetJoinRequest struct {
	GroupId           int     `json:"group_id"`
}

type JoinRequest struct {
	GroupId           int     `json:"group_id"`
	Phone             string  `json:"phone"`
	Address           string  `json:"address"`
	ExpectDailyAmount int     `json:"expect_daily_amount"`
	TotalAmount       int     `json:"total_amount"`
	AverageWeight     float64 `json:"average_weight"`
}


type GetStartRequest struct {
	Id string `json:"uid"`
}
type CargoInfo struct {
	Name string `json:"name"`
	WeightRequire string `json:"weight_require"`
	MinAmount int `json:"min_amount"`
	BasePrice float32 `json:"base_price"`
	BaseWeight float32 `json:"base_weight"`
	PictureUrl string `json:"picture_url"`
	Percentage float32 `json:"percentage"`
	Lack int `json:"lack"`
	Joined int `json:"joined"`
	Deadline string `json:"deadline"`
	Duration string `json:"duration"`
	UseRequire string `json:"use_require"`
}



type CommonResponse struct {
	RequestId    string      `json:"requestId"`
	Success      bool        `json:"success"`
	Business     string      `json:"business"`
	ErrorCode    string      `json:"errorCode"`
	ErrorMessage string      `json:"errorMessage"`
	Date         string      `json:"date"`
	Version      string      `json:"version"`
	Obj          interface{} `json:"obj"`
}


func MakeGetGroupEndpoint(s svc.IGroupService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		if _, ok := request.(GetGroupRequest); ok {
			route, err := s.GetGroup(ctx)
			if err != nil {
				return CommonResponse{
					Success:      false,
					Date:         time.Now().Format("2006-01-02 15:04:05"),
					ErrorCode:    "",
					ErrorMessage: err.Error(),
					Version:      "v1",
				}, err
			}
			return CommonResponse{
				Success: true,
				Date:    time.Now().Format("2006-01-02 15:04:05"),
				Version: "v1",
				Obj:     route,
			}, nil
		} else {
			return CommonResponse{
				Success:      false,
				Date:         time.Now().Format("2006-01-02 15:04:05"),
				ErrorCode:    "",
				ErrorMessage: ERROR_TYPE_ASSERTION.Error(),
				Version:      "v1",
			}, ERROR_TYPE_ASSERTION
		}
	}
}

func MakeJoinEndpoint(s svc.IJoinService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		if req, ok := request.(JoinRequest); ok {
			join := &dao.Join{
				GroupId:           req.GroupId,
				Phone:             req.Phone,
				Address:           req.Address,
				ExpectDailyAmount: req.ExpectDailyAmount,
				TotalAmount:       req.TotalAmount,
				AverageWeight:     req.AverageWeight,
			}
			err := s.Join(ctx, join)
			if err != nil {
				return CommonResponse{
					Success:      false,
					Date:         time.Now().Format("2006-01-02 15:04:05"),
					ErrorCode:    "",
					ErrorMessage: err.Error(),
					Version:      "v1",
				}, err
			}
			return CommonResponse{
				Success: true,
				Date:    time.Now().Format("2006-01-02 15:04:05"),
				Version: "v1",
				Obj:     nil,
			}, nil
		} else {
			return CommonResponse{
				Success:      false,
				Date:         time.Now().Format("2006-01-02 15:04:05"),
				ErrorCode:    "",
				ErrorMessage: ERROR_TYPE_ASSERTION.Error(),
				Version:      "v1",
			}, ERROR_TYPE_ASSERTION
		}
	}
}

func MakeGetJoinEndpoint(s svc.IJoinService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		if req, ok := request.(GetJoinRequest); ok {
			route, err := s.GetJoin(ctx, fmt.Sprintf("%d",req.GroupId))
			if err != nil {
				return CommonResponse{
					Success:      false,
					Date:         time.Now().Format("2006-01-02 15:04:05"),
					ErrorCode:    "",
					ErrorMessage: err.Error(),
					Version:      "v1",
				}, err
			}
			return CommonResponse{
				Success: true,
				Date:    time.Now().Format("2006-01-02 15:04:05"),
				Version: "v1",
				Obj:     route,
			}, nil
		} else {
			return CommonResponse{
				Success:      false,
				Date:         time.Now().Format("2006-01-02 15:04:05"),
				ErrorCode:    "",
				ErrorMessage: ERROR_TYPE_ASSERTION.Error(),
				Version:      "v1",
			}, ERROR_TYPE_ASSERTION
		}
	}
}
//func MakeAddOrderEndpoint(s IOrderService) endpoint.Endpoint {
//	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
//		if req, ok := request.(*Order); ok {
//			err := s.Order(ctx, req)
//			if err != nil {
//				return CommonResponse{
//					Success:      false,
//					Date:         time.Now().Format("2006-01-02 15:04:05"),
//					ErrorCode:    "",
//					ErrorMessage: err.Error(),
//					Version:      "v1",
//				}, err
//			}
//			return CommonResponse{
//				Success: true,
//				Date:    time.Now().Format("2006-01-02 15:04:05"),
//				Version: "v1",
//			}, nil
//		} else {
//			return CommonResponse{
//				Success:      false,
//				Date:         time.Now().Format("2006-01-02 15:04:05"),
//				ErrorCode:    "",
//				ErrorMessage: ERROR_TYPE_ASSERTION.Error(),
//				Version:      "v1",
//			}, ERROR_TYPE_ASSERTION
//		}
//	}
//}

